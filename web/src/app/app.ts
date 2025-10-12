import { Component, OnInit, OnDestroy, inject } from '@angular/core';
import { MatSidenavModule } from '@angular/material/sidenav';
import { RouterOutlet } from '@angular/router';
import { RouterModule } from '@angular/router';
import { NavRenderer } from './nav-renderer/nav-renderer';
import { BreadcrumbsComponent } from './breadcrumbs/breadcrumbs';
import { ThemeDropdownComponent } from './theme-dropdown/theme-dropdown';
import { MarkdownButtonComponent } from './markdown-button/markdown-button';
import { SearchButtonComponent } from './search-button/search-button';
import { PrintButtonComponent } from './print-button/print-button';
import { GithubLinkButtonComponent } from './github-link-button/github-link-button';
import { SearchModalComponent } from './search-modal/search-modal';
import { UrlService } from './services/url-service';
import { Subscription } from 'rxjs';
import { Router } from '@angular/router';
import { NavService } from './services/nav.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [
    RouterOutlet,
    MatSidenavModule,
    RouterModule,
    NavRenderer,
    BreadcrumbsComponent,
    ThemeDropdownComponent,
    MarkdownButtonComponent,
    SearchButtonComponent,
    SearchModalComponent,
    PrintButtonComponent,
    GithubLinkButtonComponent,
  ],
  templateUrl: './app.html',
  styleUrl: './app.css',
})
export class App implements OnInit, OnDestroy {
  indexContent: string = '';
  private urlService = inject(UrlService);
  private router = inject(Router);
  private navService = inject(NavService);

  private urlSub?: Subscription;
  urls: string[] = [];

  drawerMode: 'over' | 'side' = 'side';
  drawerOpened = true;

  constructor() {
    this.updateDrawerMode();
  }

  ngOnInit() {
    this.updateDrawerMode();
    window.addEventListener('resize', this.updateDrawerMode.bind(this));

    this.urlSub = this.urlService.getUrls().subscribe((urls: string[]) => {
      this.onUrlChanged(urls);
    });
  }

  onUrlChanged(urls: string[]) {
    this.urls = urls;
    console.log('URLs updated:', this.urls);
    if (urls.length === 0) {
      this.navService.getFirstUrl().subscribe((url: string | null) => {
        this.router.navigate([url]);
      });
    }
  }

  ngOnDestroy() {
    window.removeEventListener('resize', this.updateDrawerMode.bind(this));
    this.urlSub?.unsubscribe();
  }

  private updateDrawerMode() {
    if (window.innerWidth < 1024) {
      this.drawerMode = 'over';
      this.drawerOpened = false;
    } else {
      this.drawerMode = 'side';
      this.drawerOpened = true;
    }
  }
}
