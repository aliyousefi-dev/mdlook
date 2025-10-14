import { Component, inject, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MarkdownService } from '../../services/markdown.service';
import { Marked } from 'marked'; // Import marked
import { SafeHtmlPipe } from '../../services/safeHtmlPipe';
import { UrlService } from '../../services/url-service';
import { Subscription } from 'rxjs';
import customNavRenderer from '../../nav-renderer/customNavRenderer';
import { ConfigService } from '../../services/config.service';

@Component({
  selector: 'app-nav-renderer',
  standalone: true,
  imports: [CommonModule, RouterModule, SafeHtmlPipe],
  templateUrl: './nav-renderer.html',
})
export class NavRenderer implements OnInit {
  private markdownService = inject(MarkdownService);
  private configService = inject(ConfigService);
  private urlService = inject(UrlService);

  private urlSub?: Subscription;
  urls: string[] = [];

  constructor(private router: Router) {}

  rawMarkdown: string = '';
  htmlContent: string = ''; // This will hold the processed HTML content

  ngOnInit() {
    this.urlSub = this.urlService.getUrls().subscribe((urls: string[]) => {
      this.onUrlChanged(urls);
    });

    this.configService.fetchConfig().subscribe((config) => {
      this.markdownService
        .getMarkdownContent(`/nav.md`)
        .subscribe(async (content) => {
          let DocTitle = config.docname;
          this.rawMarkdown = `# ${DocTitle} {badge(${config.appversion})} \n${content}`;
          this.htmlContent = await this.convertMarkdownToHtml(this.rawMarkdown);
        });
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

  onUrlChanged(urls: string[]) {
    if (!urls || urls.length === 0) {
      console.warn('No valid URLs found!');
      return;
    }

    const segments = [...urls];
    segments[segments.length - 1] = segments[segments.length - 1];
  }

  async convertMarkdownToHtml(markdown: string): Promise<string> {
    const markedrender = new Marked();
    markedrender.use(customNavRenderer);
    let html = markedrender.parse(markdown).toString();

    // Replace href with routerLink for internal links
    return html;
  }
}
