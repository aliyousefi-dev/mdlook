import { Component, inject, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MarkdownService } from '../services/markdown.service';
import { marked } from 'marked'; // Import marked
import { SafeHtmlPipe } from '../services/safeHtmlPipe';

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
        return `<a routerLink="${cleanHref}" routerLinkActive="menu-active" href="${cleanHref}"${attrs}>${text}</a>`;
      }
    );
  }

  private convertNestedUlToDetails(html: string): string {
    // Regex targets <li> followed by <span> (the folder title) and then <ul> (the submenu)
    // The 's' flag allows the dot '.' to match newlines.
    const nestedLiRegex =
      /<li>\s*(<span>(.*?)<\/span>)(<ul.*?<\/ul>)\s*<\/li>/gs;

    let transformed = html;
    let previous;

    // Loop until no more replacements are made, ensuring all levels of nesting are correctly transformed.
    do {
      previous = transformed;
      transformed = transformed.replace(
        nestedLiRegex,
        (match, spanTag, titleText, nestedUl) => {
          // Replace the entire <li>...</li> block with <details>...
          // 'open' attribute keeps the submenu expanded by default.
          return `
<li><details >
    <summary>${titleText}</summary>
    ${nestedUl}
</details></li>`;
        }
      );
    } while (transformed !== previous);

    return transformed;
  }

  addIconToSpan(html: string): string {
    return html.replace(
      /<span(.*?)>(.*?)<\/span>/g, // Match all <span> elements
      '<span$1>$2</span>' // Prepend the icon inside <span> and preserve the text
    );
  }
  addMenuClassToUl(html: string): string {
    return html.replace(
      /<ul(.*?)>/, // This matches the first <ul> tag
      '<ul$1 class="menu px-5 w-full bg-base-100">'
    );
  }

  async convertMarkdownToHtml(markdown: string): Promise<string> {
    // Convert markdown to HTML
    const html = marked(markdown).toString();
    // Add 'text-xl' class to the first <h1>
    // Add 'text-xl' class to the first <h1> and wrap it in a <div>
    // Add 'text-xl' class to the first <h1> and wrap it in a <div>
    let htmlWithClass = html.replace(
      /<h1(.*?)>/,
      '<div class="navbar h-20 sticky top-0 z-20 items-center gap-2 px-5 py-0 bg-base-100/90 backdrop-blur"><h1$1 class="text-xl">'
    );

    // Close the first <div> immediately after the first <h1> tag
    htmlWithClass = htmlWithClass.replace(
      /(<h1.*?>.*?<\/h1>)/, // Find the first <h1> block
      '$1</div>' // Add the closing </div> right after the first <h1>
    );

    htmlWithClass = this.addMenuClassToUl(htmlWithClass);
    htmlWithClass = this.addIconToSpan(htmlWithClass);
    htmlWithClass = this.convertNestedUlToDetails(htmlWithClass);

    // Replace href with routerLink for internal links
    return this.convertHrefToRouterLink(htmlWithClass);
  }
}
