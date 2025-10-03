# Architecture

it serve all markdowns. and then render them on client frontend locally. you can serve the output with anything.

---

## Static Serve

it serve all markdown as public statics

```
nav.md
- docs
    - content01.md
    - content02.md
    - content03.md
- assets
    - img.png
```

## Internal Renderer

it uses the `MarkdownRendererComponent` to render the markdown.
it render the markdown internally. in client frontend.
also wrap that and create a custom language for that also.

## Sync Nav

The Sync function sync the `nav.md` file.

### Nav Structure

here you can see the nav structre file `nav.md`

```
# MDLook <!-- Doc Header Title -->

<!-- Pages -->
- [Overview](docs/overview.md)
- [Quick Start](docs/quick-start.md)
- [Architecture](docs/architecture.md)
- [Web Components](docs/web-components.md)
- [nav](docs/nav.md)
- [assets](docs/assets.md)
- [nav-renderer](docs/nav-renderer.md)
- [init-doc](docs/init-doc.md)
- [export-doc](docs/export-doc.md)

```

> by default Sync is enabled for auto nav. you can disable that on the `config.json` file.

### Nested َAlgorythm

here show how nested file is rendererd on the `nav.md`

if there is a same name on outside the sub pages going under that with intend.
if there is not any same name as folder. it create a title and they going under that title.
