import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, of } from 'rxjs';
import { MarkdownService } from './markdown.service';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class NavService {
  constructor(
    private http: HttpClient,
    private markdownService: MarkdownService
  ) {}

  // Fetch the first URL from a file
  getFirstUrl(filePath: string = '/nav.md'): Observable<string | null> {
    return this.markdownService.getMarkdownContent(filePath).pipe(
      map((content: string) => {
        // Match the first markdown link: [text](url)
        const match = content.match(/\[.*?\]\((.*?)\)/);
        if (match && match[1]) {
          // Remove .md extension if present
          return match[1].replace(/\.md$/, '');
        }
        return null;
      })
    );
  }
}
