// customHeadingRenderer.ts
import { Tokens } from 'marked';

function convertNestedUlToDetails(html: string): string {
  // This regex specifically targets:
  // 1. An <li> tag.
  // 2. Which contains an <h3> tag (the title).
  // 3. Immediately followed by a nested <ul> tag (the submenu).
  // The content of the <h3> is captured for the <summary>.
  const h3LiWithUlRegex =
    /<li>\s*(<h3.*?>(.*?)<\/h3>)\s*(<ul.*?<\/ul>)\s*<\/li>/gs;

  let transformed = html;

  transformed = transformed.replace(
    h3LiWithUlRegex,
    (match, fullH3Tag, titleContent, nestedUl) => {
      // titleContent is the text inside the <h3> (e.g., "commands")
      const summaryContent = titleContent.trim();

      // We use the text content for the summary, not the full <h3> tag
      return `
<li><details>
    <summary>${summaryContent}</summary>
    ${nestedUl}
</details></li>`;
    }
  );

  return transformed;
}

function addMenuClassToUl(html: string): string {
  return html.replace(
    /<(ul|ol)(.*?)>/,
    '<$1$2 class="menu w-full bg-base-100">'
  );
}

function addBadges(html: string): string {
  return html.replace(/{badge\((.*?)\)}/g, (match, badgeText) => {
    return `<span class="badge badge-xs badge-dash">${badgeText}</span>`;
  });
}

export function customListRenderer(
  this: any,
  token: Tokens.List
): string | false {
  const isOrdered = token.ordered;
  const listTag = isOrdered ? 'ol' : 'ul';

  let rawListHtml = `<${listTag}>`;

  token.items.forEach((item: Tokens.ListItem) => {
    let listItemHtml = '<li>';

    if (this.parser && typeof this.parser.parse === 'function') {
      listItemHtml += this.parser.parse(item.tokens);
    } else {
      listItemHtml += item.text;
    }

    listItemHtml += '</li>';
    rawListHtml += listItemHtml;
  });

  rawListHtml += `</${listTag}>`;

  let transformedHtml = rawListHtml;
  // Apply both transformations
  transformedHtml = convertNestedUlToDetails(transformedHtml);

  transformedHtml = addMenuClassToUl(transformedHtml);

  transformedHtml = addBadges(transformedHtml);

  return transformedHtml;
}
