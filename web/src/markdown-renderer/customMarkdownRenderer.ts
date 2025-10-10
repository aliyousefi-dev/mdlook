// customNavRenderer.ts
import { MarkedExtension } from 'marked';
import { customCodeRenderer } from './code';
import { customParagraphRenderer } from './paragraph';
import { customHeadingRenderer } from './heading';
import mermaid from 'mermaid';

const customMarkdownRenderer: MarkedExtension = {
  renderer: {
    code: customCodeRenderer,
    heading: customHeadingRenderer,
  },
  hooks: {
    postprocess: postprocess,
  },
};

function postprocess(html: string): string {
  return html;
}

export default customMarkdownRenderer;
