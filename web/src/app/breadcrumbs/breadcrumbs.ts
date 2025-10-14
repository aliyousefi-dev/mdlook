import { Component, inject, OnInit, OnDestroy } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Subscription } from 'rxjs';
import { UrlService } from '../../services/url-service';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'doc-breadcrumbs',
  templateUrl: './breadcrumbs.html',
  imports: [CommonModule, RouterModule],
})
export class BreadcrumbsComponent implements OnInit, OnDestroy {
  private urlService = inject(UrlService);
  private urlSub?: Subscription;
  urls_segments: string[] = [];

  ngOnInit() {
    this.urlSub = this.urlService.getUrls().subscribe((urls_segments) => {
      this.urls_segments = urls_segments;
    });
  }

  ngOnDestroy() {
    this.urlSub?.unsubscribe();
  }
}
