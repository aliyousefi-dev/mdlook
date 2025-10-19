import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'doc-next-features-button',
  templateUrl: './next-features-button.html',
  imports: [CommonModule, RouterLink, RouterModule],
})
export class NextFeaturesButtonComponent {}
