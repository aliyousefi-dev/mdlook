import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Input } from '@angular/core';

@Component({
  selector: 'app-table-of-contents',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './table-of-contents.html',
})
export class TableOfContents {
  @Input() mdContent: string = '';

  headings: { level: number; text: string; id: string }[] = [];

  ngOnChanges() {
    this.generateHeadings();
  }

  private generateHeadings() {
    this.headings = [];
    if (!this.mdContent) return;

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
        const id = text.toLowerCase().replace(/[^\w]+/g, '-');
        this.headings.push({ level, text, id });
      }
    }
  }
}
