// app-table-of-contents.component.ts

import { Component, Input, OnChanges } from '@angular/core'; // Add OnChanges
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
// ... other imports

interface Heading {
  level: number;
  text: string;
  id: string;
  children?: Heading[]; // Add optional children array for nesting
}

@Component({
  selector: 'app-table-of-contents',
  standalone: true,
  imports: [CommonModule, RouterModule],
  templateUrl: './table-of-contents.html',
})
// Implement OnChanges lifecycle hook
export class TableOfContents implements OnChanges {
  @Input() mdContent: string = '';

  // Change headings to be a nested structure: an array of root-level headings
  headings: Heading[] = [];

  ngOnChanges() {
    this.generateHeadings();
  }

  // Helper function to convert flat array to nested structure
  private nestHeadings(flatHeadings: Heading[]): Heading[] {
    const root: Heading = { level: 0, text: '', id: '', children: [] };
    const stack: Heading[] = [root];

    for (const heading of flatHeadings) {
      // Pop off headings from the stack that are at the same level or lower (greater number)
      while (
        stack.length > 1 &&
        stack[stack.length - 1].level >= heading.level
      ) {
        stack.pop();
      }

      const parent = stack[stack.length - 1];
      if (!parent.children) {
        parent.children = [];
      }
      parent.children.push(heading);
      stack.push(heading);
    }

    // Return the children of the virtual root (the actual level 1 headings)
    return root.children || [];
  }

  private generateHeadings() {
    let flatHeadings: Heading[] = []; // Use a temporary flat array first
    if (!this.mdContent) {
      this.headings = [];
      return;
    }

    // ... (Your existing code to extract flat headings)
    // ------------------------------------------------------------------
    // Ignore headings inside code blocks
    const codeBlockRegex = /```[\s\S]*?```/g;
    let nonCodeParts: string[] = [];
    let lastIndex = 0;
    let match;
    while ((match = codeBlockRegex.exec(this.mdContent)) !== null) {
      nonCodeParts.push(this.mdContent.slice(lastIndex, match.index));
      lastIndex = codeBlockRegex.lastIndex;
    }
    nonCodeParts.push(this.mdContent.slice(lastIndex));

    const headingRegex = /^(\#{1,6})\s+(.*)$/gm;
    for (const part of nonCodeParts) {
      let hMatch;
      while ((hMatch = headingRegex.exec(part)) !== null) {
        const level = hMatch[1].length;
        const text = hMatch[2].trim();
        // Ensure IDs are unique and URL-friendly
        const id = text.toLowerCase().replace(/[^\w]+/g, '-');
        flatHeadings.push({ level, text, id });
      }
    }
    // ------------------------------------------------------------------

    // Convert the flat list to the final nested structure
    this.headings = this.nestHeadings(flatHeadings);
  }

  // onAnchorClick remains the same
  onAnchorClick(event: Event, id: string) {
    event.preventDefault();
    const el = document.getElementById(id);
    if (el) {
      el.scrollIntoView({
        behavior: 'smooth',
        block: 'center',
        inline: 'center',
      });
    }
  }
}
