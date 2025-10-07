// customHeadingRenderer.ts
import { Tokens } from 'marked';

export function customHeadingRenderer(
  this: any,
  token: Tokens.Heading
): string | false {
  const { tokens, depth } = token;

  const text = this.parser.parseInline(tokens);

  if (depth === 1) {
    return `
      <div class="navbar h-20 sticky top-0 z-20 items-center gap-2 px-5 py-0 bg-base-100/90 backdrop-blur">
        <h1 class="text-xl">
          ${text}
        </h1>
      </div>
    `;
  }

    if (depth === 2) {
      return `
        <h1 class="menu-title">
          ${text}
        </h1>
    `;
    }

  // Standard rendering for other headings (and subsequent H1s)
  return `
    <h${depth}>
      ${text}
    </h${depth}>
  `;
}
