import { Component, ViewChild, ElementRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HostListener } from '@angular/core';
import { SearchModalComponent } from '../search-modal/search-modal';

@Component({
  selector: 'doc-search-button',
  templateUrl: './search-button.html',
  imports: [CommonModule, SearchModalComponent],
  styleUrl: './search-button.css',
})
export class SearchButtonComponent {
  @ViewChild('searchInput', { static: false }) searchInput:
    | ElementRef
    | undefined;

  // Listen for keyboard events
  @HostListener('document:keydown', ['$event'])
  onKeydown(event: KeyboardEvent): void {
    // Check if Ctrl + K is pressed
    if (event.ctrlKey && event.key === 'k') {
      event.preventDefault(); // Prevent the default browser behavior
      this.focusSearchInput(); // Focus the search input
    }
  }

  // Function to focus the search input
  focusSearchInput(): void {
    if (this.searchInput) {
      this.searchInput.nativeElement.focus();
    }
  }
}
