import { Component, inject, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Input } from '@angular/core';
import { ConfigService } from '../../services/config.service';

@Component({
  selector: 'doc-github-link-button',
  templateUrl: './github-link-button.html',
  imports: [CommonModule],
})
export class GithubLinkButtonComponent implements OnInit {
  @Input() repoUrl: string = '';

  constructor(private configService: ConfigService) {}

  ngOnInit() {
    this.githubUrl();
  }

  githubUrl() {
    this.configService.fetchConfig().subscribe((config) => {
      this.repoUrl = config.GitUrl;
    });
  }
}
