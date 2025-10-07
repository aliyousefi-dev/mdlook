// customHeadingRenderer.ts
import { Tokens } from 'marked';

export function customParagraphRenderer(
  this: any,
  token: Tokens.Paragraph
): string | false {
  const { tokens, text } = token;

  return `${text}`;
}
