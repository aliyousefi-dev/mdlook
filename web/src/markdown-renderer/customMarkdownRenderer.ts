// customNavRenderer.ts
import { MarkedExtension } from 'marked';
import { customCodeRenderer } from './code';
import mermaid from 'mermaid';

const customMarkdownRenderer: MarkedExtension = {
  renderer: {
    code: customCodeRenderer,
  },
  hooks: {
    postprocess: postprocess,
  },
};

function postprocess(html: string): string {
  return html;
}

export default customMarkdownRenderer;
