// customHeadingRenderer.ts
import { Tokens } from 'marked';

export function customHeadingRenderer(
  this: any,
  token: Tokens.Heading
): string | false {
  const { tokens, depth, text } = token;
  const parsed = this.parser.parseInline(tokens);
  const slug = text
    .toLowerCase()
    .replace(/[^\w\s-]/g, '') // Remove non-word except space/hyphen
    .replace(/\s+/g, '-') // Replace spaces with hyphens
    .replace(/-+/g, '-') // Collapse multiple hyphens
    .replace(/^-|-$/g, ''); // Trim leading/trailing hyphens

  return `  <h${depth} id="${slug}">
              ${parsed}
            </h${depth}>`;
}
