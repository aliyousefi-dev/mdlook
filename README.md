## Beta - Under Development

MDLook is currently in beta and under active development. It provides a command-line interface for converting Markdown files to HTML documentation, with additional features for organizing, processing, and managing your documentation workflow.

https://github.com/user-attachments/assets/7855e753-6067-46fd-97e5-901b9d767bb4

## How it Works?

<img width="1467" height="897" alt="image" src="https://github.com/user-attachments/assets/6c2e7c23-9f57-4cfa-a48f-a1fcca1b8c0d" />

it serve all markdowns. and then render them on client frontend locally.
you can serve the output with anything.
it using two renderer . the nav renderer that read from nav.md and the markdown renderer that read from your markdowns.
and convert them to an html page.

## Quick Start

use `mdlook init <path>`
to create workstation

use `mdlook watch .` on worksation to live reload

use `mdlook export` on workstation to export your docs as index.html servable.

you can serve the page with any language you prefer.
also mdlook provide `mdlook serve .` for production serve if you want.
