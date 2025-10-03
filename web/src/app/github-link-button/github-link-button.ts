import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Input } from '@angular/core';

@Component({
  selector: 'doc-github-link-button',
  templateUrl: './github-link-button.html',
  imports: [CommonModule],
})
export class GithubLinkButtonComponent {
  @Input() repoUrl: string = 'https://github.com/your-repo';
}
