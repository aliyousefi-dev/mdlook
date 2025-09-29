import { Injectable } from '@angular/core';
import { Router, ActivatedRoute, NavigationEnd } from '@angular/router';
import { BehaviorSubject, filter, skip } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class UrlService {
  private urls$ = new BehaviorSubject<string[]>([]);

  constructor(private router: Router, private activatedRoute: ActivatedRoute) {
    this.router.events
      .pipe(filter((e) => e instanceof NavigationEnd))
      .subscribe((event) => {
        const urls = this.buildUrls(this.activatedRoute.root);
        this.urls$.next(urls);
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
