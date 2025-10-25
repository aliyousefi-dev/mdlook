import { Component, inject, OnInit, OnDestroy } from '@angular/core';
import { Router, NavigationEnd } from '@angular/router';
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
export class NavRenderer implements OnInit, OnDestroy {
  private markdownService = inject(MarkdownService);
  private configService = inject(ConfigService);
  private urlService = inject(UrlService);

  private urlSub?: Subscription;
  private routerSub?: Subscription;
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

    // Subscribe to router events to detect navigation changes
    this.routerSub = this.router.events.subscribe((event) => {
      if (event instanceof NavigationEnd) {
        this.setActiveLink();
      }
    });
  }

  ngOnDestroy() {
    if (this.urlSub) {
      this.urlSub.unsubscribe();
    }
    if (this.routerSub) {
      this.routerSub.unsubscribe();
    }
  }

  onUrlChanged(urls: string[]) {
    if (!urls || urls.length === 0) {
      console.warn('No valid URLs found!');
      return;
    }

    const fullurl = urls.join('/');
    console.log('URLs changed:', fullurl);
    this.setActiveLink();
  }

  async convertMarkdownToHtml(markdown: string): Promise<string> {
    const filteredMarkdown = this.removeUnwantedLinks(markdown);

    const markedrender = new Marked();
    markedrender.use(customNavRenderer);
    let html = markedrender.parse(filteredMarkdown).toString();

    // Replace href with routerLink for internal links
    return html;
  }

  private removeUnwantedLinks(markdown: string): string {
    // Split markdown into lines
    const lines = markdown.split('\n');

    // Filter out lines that contain 'changelog.md' or 'next-features.md'
    const filteredLines = lines.filter(
      (line) =>
        !line.includes('changelog.md') && !line.includes('next-features.md')
    );

    // Join the filtered lines back into a single string
    return filteredLines.join('\n');
  }

  // Function to check if current URL matches the link and set active class
  private setActiveLink() {
    const currentUrl = this.router.url;
    const links = document.querySelectorAll('#nav-renderer a[routerLink]');

    links.forEach((link: any) => {
      const routerLink = link.getAttribute('routerLink');
      if (routerLink && currentUrl.includes(routerLink)) {
        link.classList.add('menu-active');
      } else {
        link.classList.remove('menu-active');
      }
    });
  }
}
