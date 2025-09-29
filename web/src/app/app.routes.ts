import { Routes } from '@angular/router';
import { App } from './app';
import { MarkdownRenderer } from './markdown-renderer/markdown-renderer';

export const routes: Routes = [
  {
    path: '**',
    component: MarkdownRenderer,
  },
];
