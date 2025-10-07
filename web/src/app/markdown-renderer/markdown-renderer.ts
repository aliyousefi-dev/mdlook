import {
  Component,
  inject,
  OnInit,
  OnDestroy,
  AfterViewInit,
  ElementRef,
  Renderer2,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { MarkdownService } from '../services/markdown.service';
import { UrlService } from '../services/url-service';
import { Subscription } from 'rxjs';
import { marked, Marked } from 'marked'; // Import marked
import { SafeHtmlPipe } from '../services/safeHtmlPipe';
import customMarkdownRenderer from '../../markdown-renderer/customMarkdownRenderer';
import { markedHighlight } from 'marked-highlight';
import { codeToHtml } from 'shiki';
import mermaid from 'mermaid';

@Component({
  selector: 'app-markdown-renderer',
  standalone: true,
  imports: [CommonModule, SafeHtmlPipe],
  templateUrl: './markdown-renderer.html',
})
export class MarkdownRenderer implements OnInit, OnDestroy, AfterViewInit {
  private markdownService = inject(MarkdownService);
  private urlService = inject(UrlService);

  private urlSub?: Subscription;
  urls: string[] = [];

  htmlContent: string = '';

  constructor(private el: ElementRef, private renderer: Renderer2) {}

  ngOnInit() {
    mermaid.initialize({ startOnLoad: true });
    console.log('Markdown Renderer Initialized');
    this.urlSub = this.urlService.getUrls().subscribe((urls: string[]) => {
      this.onUrlChanged(urls);
    });
  }

  ngAfterViewInit() {
    // This method ensures content is processed after the view is initialized
    this.renderLinksWithRouterLink();
  }

  // Assuming this is part of a component/class method
  onUrlChanged(urls: string[]) {
    if (!urls || urls.length === 0) {
      console.warn('No valid URLs found!');
      return;
    }

    this.urls = urls;

    // Append .md to the last segment of the segmented URL array
    const segments = [...urls];
    segments[segments.length - 1] = segments[segments.length - 1] + '.md';
    const mdUrl = '/' + segments.join('/');

    this.markdownService
      .getMarkdownContent(mdUrl)
      .subscribe(async (content) => {
        this.htmlContent = await this.convertMarkdownToHtml(content);

        // Wait for the content to be fully injected into the DOM
        setTimeout(() => {
          mermaid.run({
            querySelector: '.mermaid', // Ensure it targets only the mermaid elements
          });
        }, 200); // A small timeout to ensure rendering completion

        // After the content is fetched and set, call this function to process links
        this.renderLinksWithRouterLink();
      });
  }

  async convertMarkdownToHtml(markdown: string): Promise<string> {
    const marked = new Marked(
      markedHighlight({
        // We set async to true to indicate the highlight function is asynchronous.
        async: true,
        async highlight(code, lang, info) {
          try {
            if (lang === 'mermaid') {
              // For Mermaid syntax, wrap the code in <pre class="mermaid"> tags
              return `<pre class="mermaid flex items-center justify-center bg-base-100">${code}</pre>`;
            }

            const html = await codeToHtml(code, {
              lang: lang,
              theme: 'vitesse-dark',
            });
            return html;
          } catch (error) {
            console.error('Highlighting failed:', error);
            return code;
          }
        },
      })
    );

    marked.use(customMarkdownRenderer);

    const parsedHtml = await marked.parse(markdown);
    let html = parsedHtml.toString();

    return html;
  }

  renderLinksWithRouterLink() {
    // If content is not loaded yet, return
    if (!this.htmlContent) return;

    const container = this.el.nativeElement.querySelector('#markdown-content');
    const parser = new DOMParser();
    const doc = parser.parseFromString(this.htmlContent, 'text/html');

    const links = doc.querySelectorAll('a');
    links.forEach((link) => {
      const href = link.getAttribute('href');
      if (href && href.startsWith('/')) {
        this.renderer.setAttribute(link, 'routerLink', href); // Add routerLink for internal links
      }
    });

    container.innerHTML = doc.body.innerHTML; // Set the modified HTML with routerLink attributes
  }

  ngOnDestroy() {
    this.urlSub?.unsubscribe();
  }
}
