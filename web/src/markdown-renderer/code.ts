// customHeadingRenderer.ts
import { Tokens } from 'marked';
import js from '@shikijs/langs/javascript';
import nord from '@shikijs/themes/nord';
import { createHighlighterCoreSync } from 'shiki/core';
import { createJavaScriptRegexEngine } from 'shiki/engine/javascript';

export function customCodeRenderer(
  this: any,
  token: Tokens.Code
): string | false {
  const { text, type } = token;

  const html = `
  <div class="relative">
        <div class="absolute top-2 right-2 flex items-center justify-end h-full">
        <div class="tooltip tooltip-left self-start" data-tip="Copy to clipboard">
            <button class="btn btn-circle btn-sm" aria-label="Copy to clipboard">
                <i class="bx bx-clipboard bx-xs"></i>
            </button>
        </div>
    </div>
  ${text}
  </div>`;

  // Return a placeholder immediately (since rendering is async)
  return html;
}
