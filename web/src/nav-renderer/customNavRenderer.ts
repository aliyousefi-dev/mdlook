// customNavRenderer.ts
import { MarkedExtension } from 'marked';
import { customHeadingRenderer } from './heading'; // Adjust the path as needed
import { customLinksRenderer } from './links';
import { customListRenderer } from './lists';
import { customParagraphRenderer } from './paragrapgh';

const customNavRenderer: MarkedExtension = {
  renderer: {
    heading: customHeadingRenderer,
    link: customLinksRenderer,
    list: customListRenderer,
    paragraph: customParagraphRenderer,
  },
  hooks: {
    postprocess: postprocess,
  },
};

function postprocess(html: string): string {
  return html;
}

export default customNavRenderer;
