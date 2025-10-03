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
import { marked } from 'marked'; // Import marked
import { SafeHtmlPipe } from '../services/safeHtmlPipe';

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
    console.log('Markdown Renderer Initialized');
    this.urlSub = this.urlService.getUrls().subscribe((urls: string[]) => {
      this.onUrlChanged(urls);
    });
  }

  ngAfterViewInit() {
    // This method ensures content is processed after the view is initialized
    this.renderLinksWithRouterLink();
  }

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

    this.markdownService.getMarkdownContent(mdUrl).subscribe((content) => {
      this.htmlContent = this.convertMarkdownToHtml(content);
      // After the content is fetched and set, call this function to process links
      this.renderLinksWithRouterLink();
    });
  }

  convertMarkdownToHtml(markdown: string): string {
    // Convert markdown to HTML
    const html = marked(markdown).toString();

    // Add 'text-xl' class to the first <h1>
    return html;
  }

  renderLinksWithRouterLink() {
    // If content is not loaded yet, return
    if (!this.htmlContent) return;

    const container = this.el.nativeElement.querySelector('.prose');
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
