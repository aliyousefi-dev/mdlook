// customHeadingRenderer.ts
import { Tokens } from 'marked';

export function customParagraphRenderer(
  this: any,
  token: Tokens.Paragraph
): string | false {
  const { text, pre } = token;

  // Return a placeholder immediately (since rendering is async)
  return `<p>${text}</p>`;
}


function addBadges(html: string): string {
  return html.replace(/{badge\((.*?)\)}/g, (match, badgeText) => {
    return `<span class="badge badge-xs badge-ghost ml-2">${badgeText}</span>`;
  });
}