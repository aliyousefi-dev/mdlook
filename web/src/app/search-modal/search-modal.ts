import { Component } from '@angular/core';
import { SearchService } from '../../services/search.service';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'doc-search-modal',
  templateUrl: './search-modal.html',
  imports: [CommonModule, FormsModule, RouterModule],
})
export class SearchModalComponent {
  loading = false;
  searchResults: any[] = [];
  query: string = '';
  private loadingTimeout: any;

  constructor(private searchService: SearchService) {}

  // This method will be called when the input changes
  onInput() {
    this.loading = true;
    clearTimeout(this.loadingTimeout);

    // Set timeout to debounce the search and stop loading after 1 second
    this.loadingTimeout = setTimeout(() => {
      this.search(); // Call the search function after the debounce time
      this.loading = false; // Stop loading after search completes
    }, 300);
  }

  // Perform the search using SearchService
  search() {
    if (this.query.trim().length > 0) {
      this.searchService.search(this.query).subscribe((results) => {
        this.searchResults = results; // Store the results from the search
      });
    } else {
      this.searchResults = []; // Clear results if the query is empty
    }
  }

  // Method to extract a snippet from the content and highlight the search term
  getSnippet(content: string): string {
    const snippetLength = 100; // Define the length of the snippet (characters)
    let snippet = content.substring(0, snippetLength);

    // Highlight the search query term in the snippet
    if (this.query.trim().length > 0) {
      const regex = new RegExp(`(${this.query.trim()})`, 'gi'); // Create a case-insensitive regex to find the query term
      snippet = snippet.replace(regex, '<mark>$1</mark>'); // Wrap matched search term in <mark> tag
    }

    // Add "..." if the snippet is truncated
    return snippet.length === snippetLength ? snippet + '...' : snippet;
  }

  // Call this method from the button click event
  CloseModal(): void {
    const modal: any = document.getElementById('my_modal_2');
    if (modal && typeof modal.showModal === 'function') {
      modal.close(); // Close the modal
    }

    setTimeout(() => {
      // Clear the input field by resetting the query
      this.query = ''; // This will reset the input field
      this.searchResults = []; // Optionally, you can also clear the search results
    }, 1000);
  }
}
