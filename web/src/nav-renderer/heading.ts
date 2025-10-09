// customHeadingRenderer.ts
import { Tokens } from 'marked';

export function customHeadingRenderer(
  this: any,
  token: Tokens.Heading
): string | false {
  const { tokens, depth, text } = token;

  if (depth === 1) {
    let heading1 = `      
    <div class="navbar h-20 sticky top-0 z-20 items-center gap-2 px-5 py-0 bg-base-100/90 backdrop-blur">
        <h1 class="text-xl">
          ${text}
        </h1>
      </div>`;

    heading1 = addBadges(heading1);

    return heading1;
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

function addBadges(html: string): string {
  return html.replace(/{badge\((.*?)\)}/g, (match, badgeText) => {
    return `<span class="badge badge-xs badge-ghost ml-2">${badgeText}</span>`;
  });
}
