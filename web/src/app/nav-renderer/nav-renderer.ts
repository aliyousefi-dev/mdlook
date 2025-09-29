import { Component, inject, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MarkdownService } from '../services/markdown.service';
import { marked } from 'marked'; // Import marked
import { Pipe, PipeTransform } from '@angular/core';
import { DomSanitizer } from '@angular/platform-browser';

@Pipe({ name: 'safeHtml' })
export class SafeHtmlPipe implements PipeTransform {
  constructor(private sanitized: DomSanitizer) {}
  transform(value: string) {
    return this.sanitized.bypassSecurityTrustHtml(value);
  }
}

@Component({
  selector: 'app-nav-renderer',
  standalone: true,
  imports: [CommonModule, RouterModule, SafeHtmlPipe],
  templateUrl: './nav-renderer.html',
})
export class NavRenderer implements OnInit {
  private markdownService = inject(MarkdownService);
  constructor(private router: Router) {}

  indexContent: string = '';
  htmlContent: string = ''; // This will hold the processed HTML content

  ngOnInit() {
    this.markdownService
      .getMarkdownContent(`/nav.md`)
      .subscribe(async (content) => {
        this.indexContent = content;
        this.htmlContent = await this.convertMarkdownToHtml(this.indexContent);
      });

    // Delegate click events for router navigation
    setTimeout(() => {
      const container = document.querySelector('#nav-renderer');
      if (container) {
        container.addEventListener('click', (event: any) => {
          const target = event.target.closest('a[routerLink]');
          if (target) {
            event.preventDefault();
            const link = target.getAttribute('routerLink');
            if (link) {
              this.router.navigateByUrl(link);
            }
          }
        });
      }
    }, 500);
  }

  private convertHrefToRouterLink(content: string): string {
    // Replace <a href="..."> with <a routerLink="..." href="..."> and remove .md extension
    return content.replace(
      /<a\s+href="([^"]+)"([^>]*)>(.*?)<\/a>/g,
      (match, href, attrs, text) => {
        // Remove .md extension from href
        const cleanHref = href.replace(/\.md$/, '');
        return `<a routerLink="${cleanHref}" href="${cleanHref}"${attrs}>${text}</a>`;
      }
    );
  }

  async convertMarkdownToHtml(markdown: string): Promise<string> {
    // Convert markdown to HTML
    const html = marked(markdown).toString();
    // Add 'text-xl' class to the first <h1>
    const htmlWithClass = html.replace(
      /<h1(.*?)>/,
      '<h1$1 class="text-xl mb-5">'
    );
    // Replace href with routerLink for internal links
    return this.convertHrefToRouterLink(htmlWithClass);
  }
}
