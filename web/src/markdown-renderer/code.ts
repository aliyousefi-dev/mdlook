// customHeadingRenderer.ts
import { Tokens } from 'marked';

export function customCodeRenderer(
  this: any,
  token: Tokens.Code
): string | false {
  const { text, lang } = token;

  if (lang === 'mermaid') {
    // For Mermaid syntax, wrap the code in <div class="mermaid"> tags
    return `${text}`;
  }

  // Replace inline style in <pre> tag with class
  const codeHtml = text.replace(
    /<pre.*?>/,
    '<pre class="shiki bg-neutral" tabindex="0">'
  );

  const html = `
<div class="relative group">
  <!-- Copy Button -->
  <div class="absolute top-2 right-5 flex items-center justify-end h-full opacity-0 group-hover:opacity-100 group-hover:block transition-opacity duration-300 z-10">
    <div class="tooltip tooltip-right self-start" data-tip="Copy to clipboard">
      <button class="btn btn-circle btn-sm" aria-label="Copy to clipboard">
        <i class='bx bx-copy-alt bx-xs' ></i>
      </button>
    </div>
  </div>

  <!-- Language Text -->
  <div class="absolute  top-2 right-5 opacity-100 group-hover:opacity-0 transition-opacity duration-300 select-none">
    <span class="text-xs text-white opacity-50">${lang}</span>
  </div>

  ${codeHtml}
</div>
`;

  // Return a placeholder immediately (since rendering is async)
  return html;
}
