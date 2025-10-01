import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ConfigData } from '../types/configdata';

@Injectable({
  providedIn: 'root',
})
export class ConfigService {
  constructor(private http: HttpClient) {}

  // Fetch the markdown content by file path
  fetchConfig(): Observable<ConfigData> {
    return this.http.get<ConfigData>('config.json', { responseType: 'json' });
  }
}
