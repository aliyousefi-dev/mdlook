import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { UrlService } from '../../services/url-service';
import { Subscription } from 'rxjs';
import { RouterLink } from '@angular/router';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'doc-changelog-button',
  templateUrl: './changelog-button.html',
  imports: [CommonModule, RouterLink, RouterModule],
})
export class ChangelogButtonComponent {}
