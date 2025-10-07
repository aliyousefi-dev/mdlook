// customHeadingRenderer.ts
import { Tokens } from 'marked';

export function customLinksRenderer(
  this: any,
  token: Tokens.Link
): string | false {
  const { href,text } = token;

  const cleanHref = href.replace(/\.md$/, '');

  // Standard rendering for other headings (and subsequent H1s)
  return `
<a routerLink="${cleanHref}" routerLinkActive="menu-active" href="${cleanHref}">${text}</a>
  `;

}
