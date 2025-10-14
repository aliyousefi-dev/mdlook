import { Injectable } from '@angular/core';
import { Router, ActivatedRoute, NavigationEnd } from '@angular/router';
import { ReplaySubject, filter } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class UrlService {
  // Use ReplaySubject, it won't emit a value until .next() is called at least once
  private urls$ = new ReplaySubject<string[]>(1);

  constructor(private router: Router, private activatedRoute: ActivatedRoute) {
    this.router.events
      .pipe(filter((e) => e instanceof NavigationEnd))
      .subscribe(() => {
        // ... (rest of subscription logic)
        const urls = this.buildUrls(this.activatedRoute.root);
        this.urls$.next(urls); // <-- First value is set here, so it only emits once
      });
  }

  getUrls() {
    return this.urls$.asObservable();
  }

  private buildUrls(
    route: ActivatedRoute,
    url: string = '',
    urls: string[] = []
  ): string[] {
    const children: ActivatedRoute[] = route.children;
    if (children.length === 0) {
      return urls;
    }
    for (const child of children) {
      const segments = child.snapshot.url
        .map((segment) => segment.path)
        .filter(Boolean);
      if (segments.length > 0) {
        for (const seg of segments) {
          url += `/${seg}`;
          urls.push(seg);
        }
      }
      return this.buildUrls(child, url, urls);
    }
    return urls;
  }
}
