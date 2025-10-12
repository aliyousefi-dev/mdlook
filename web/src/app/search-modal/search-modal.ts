import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'doc-search-modal',
  templateUrl: './search-modal.html',
  imports: [CommonModule],
})
export class SearchModalComponent {
  loading = false;
  private loadingTimeout: any;

  onInput() {
    this.loading = true;
    clearTimeout(this.loadingTimeout);
    this.loadingTimeout = setTimeout(() => {
      this.loading = false;
    }, 1000); // 1 second
  }
}
