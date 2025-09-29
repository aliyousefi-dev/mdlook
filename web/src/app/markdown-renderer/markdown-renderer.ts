import {
  Component,
  inject,
  OnInit,
  OnDestroy,
  ModuleWithComponentFactories,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { MarkdownComponent } from 'ngx-markdown';
import { MarkdownService } from '../services/markdown.service';
import { UrlService } from '../services/url-service';
import { Subscription, skip } from 'rxjs';
import { marked } from 'marked'; // Import marked

@Component({
  selector: 'app-markdown-renderer',
  standalone: true,
  imports: [CommonModule, MarkdownComponent],
  templateUrl: './markdown-renderer.html',
})
export class MarkdownRenderer implements OnInit, OnDestroy {
  private markdownService = inject(MarkdownService);
  private urlService = inject(UrlService);

  private urlSub?: Subscription;
  urls: string[] = [];

  htmlContent: string = '';

  ngOnInit() {
    console.log('Markdown Renderer Initialized');
    this.urlSub = this.urlService.getUrls().subscribe((urls: string[]) => {
      this.onUrlChanged(urls);
    });
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
    });
  }

  convertMarkdownToHtml(markdown: string): string {
    // Convert markdown to HTML
    const html = marked(markdown).toString();
    // Add 'text-xl' class to the first <h1>

    // Replace href with routerLink for internal links
    return html;
  }

  ngOnDestroy() {
    this.urlSub?.unsubscribe();
  }
}
