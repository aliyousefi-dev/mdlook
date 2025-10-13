import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';

interface SearchIndex {
  title: string;
  content: string;
  path: string;
}

@Injectable({
  providedIn: 'root',
})
export class SearchService {
  private searchIndexUrl = '/search_index.json'; // URL to fetch the search index from
  private searchIndexCache: SearchIndex[] | null = null; // Cache to store the search index
  private searchIndexSubject = new BehaviorSubject<SearchIndex[]>([]); // Observable for the search index

  constructor(private http: HttpClient) {}

  // Fetch the search index from the server (if not cached)
  private fetchSearchIndex(): Observable<SearchIndex[]> {
    return this.http.get<SearchIndex[]>(this.searchIndexUrl).pipe(
      map((data) => {
        this.searchIndexCache = data; // Cache the data
        this.searchIndexSubject.next(data); // Push data to the observable
        return data;
      })
    );
  }

  // Get the cached search index or fetch it if not cached
  private getSearchIndex(): Observable<SearchIndex[]> {
    if (this.searchIndexCache) {
      // Return the cached data if available
      return this.searchIndexSubject.asObservable();
    } else {
      // Fetch and cache the data if not available
      return this.fetchSearchIndex();
    }
  }

  // Perform a search query on the search index
  search(searchQuery: string): Observable<SearchIndex[]> {
    return this.getSearchIndex().pipe(
      map((index) => {
        return index.filter((entry) => {
          // Check if the title or content includes the search query
          return (
            entry.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
            entry.content.toLowerCase().includes(searchQuery.toLowerCase())
          );
        });
      })
    );
  }
}
