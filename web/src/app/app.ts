import { Component, OnInit, OnDestroy, inject } from '@angular/core';
import { MarkdownService } from './services/markdown.service';
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
    PrintButtonComponent,
    GithubLinkButtonComponent,
  ],
  templateUrl: './app.html',
  styles: [
    `
      .sidebar-container {
        background: var(--b1);
      }

      .sidebar-sidenav {
        background: var(--b1);
        border-radius: 0px;
      }
    `,
  ],
})
export class App implements OnInit, OnDestroy {
  private markdownService = inject(MarkdownService);

  indexContent: string = '';

  drawerMode: 'over' | 'side' = 'side';
  drawerOpened = true;

  constructor() {
    this.updateDrawerMode();
  }

  ngOnInit() {
    this.updateDrawerMode();
    window.addEventListener('resize', this.updateDrawerMode.bind(this));
  }

  ngOnDestroy() {
    window.removeEventListener('resize', this.updateDrawerMode.bind(this));
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
