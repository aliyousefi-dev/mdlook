import { Component, ViewChild, ElementRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HostListener } from '@angular/core';

@Component({
  selector: 'doc-search-button',
  templateUrl: './search-button.html',
  imports: [CommonModule],
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
      // Open the modal
      this.onSearchButtonClick();
    }
  }

  // Function to focus the search input
  focusSearchInput(): void {
    if (this.searchInput) {
      this.searchInput.nativeElement.focus();
    }
    // Also focus the input with id="search"
    const searchInputEl = document.getElementById(
      'search'
    ) as HTMLInputElement | null;
    if (searchInputEl) {
      searchInputEl.focus();
    }
  }

  // Call this method from the button click event
  onSearchButtonClick(): void {
    const modal: any = document.getElementById('my_modal_2');
    if (modal && typeof modal.showModal === 'function') {
      modal.showModal();
      // Focus the input after modal is open and rendered
      setTimeout(() => {
        this.focusSearchInput();
      }, 100);
    }
  }
}
