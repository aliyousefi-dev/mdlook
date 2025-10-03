import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'doc-print-button',
  templateUrl: './print-button.html',
  imports: [CommonModule],
})
export class PrintButtonComponent {
  printPage() {
    window.print(); // This will trigger the browser's print dialog
  }
}
