import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { UrlService } from '../services/url-service';
import { Subscription } from 'rxjs';
import { OnInit, OnDestroy } from '@angular/core';

@Component({
  selector: 'doc-markdown-button',
  templateUrl: './markdown-button.html',
  imports: [CommonModule],
})
export class MarkdownButtonComponent {
  private urlService = inject(UrlService);
  private urlSub?: Subscription;
  mdUrl: string = '';

  ngOnInit() {
    console.log('Markdown Renderer Initialized');
    this.urlSub = this.urlService.getUrls().subscribe((urls: string[]) => {
      this.onUrlChanged(urls);
    });
  }

  ngOnDestroy() {
    this.urlSub?.unsubscribe();
  }

  onUrlChanged(urls: string[]) {
    if (!urls || urls.length === 0) {
      console.warn('No valid URLs found!');
      return;
    }

    // Append .md to the last segment of the segmented URL array
    const segments = [...urls];
    segments[segments.length - 1] = segments[segments.length - 1] + '.md';
    this.mdUrl = this.mdUrl = '/' + segments.join('/');
  }
}
