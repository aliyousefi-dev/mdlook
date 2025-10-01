package mdrepo

import (
	"fmt"
	"log"
	"mdlook/source/internal/types"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type MarkdownBlock struct {
	SortKey string
	Content string
}

// NavRender generates the markdown content from the NavRenderStruct and returns it as a string
func (mdlook *MDLookManager) NavRender(navRenderStruct types.NavRenderStruct) string {
	// Initialize a strings.Builder to efficiently build the markdown content
	var markdownContent strings.Builder

	// Write the header (Markdown header)
	_, err := markdownContent.WriteString(fmt.Sprintf("# %s\n\n", navRenderStruct.HeaderTitle))
	if err != nil {
		log.Fatalf("error writing header: %v", err)
	}

	pages := mdlook.SortPageSection(mdlook.RenderPageSection(navRenderStruct.NavItems))

	markdownContent.WriteString(pages)

	// Return the generated markdown content as a string
	return markdownContent.String()
}

func (mdlook *MDLookManager) RenderPageSection(navItem []types.NavItem) string {
	var markdownContent strings.Builder

	categories := make(map[string][]types.NavItem)

	for _, item := range navItem {
		relativePath := strings.TrimPrefix(item.Path, "docs/")
		parts := strings.Split(relativePath, "/")

		if len(parts) == 1 {
			categoryName := parts[0]
			dirCandidate := strings.TrimSuffix(categoryName, filepath.Ext(categoryName))

			// If the filename (without extension) matches the name of a possible directory,
			// group it under the directory name instead of its full filename.
			// This prevents the file from being grouped as a standalone root item.
			isDirIndexCandidate := false
			for _, otherItem := range navItem {
				if otherItem.Path != item.Path && strings.HasPrefix(otherItem.Path, "docs/"+dirCandidate+"/") {
					isDirIndexCandidate = true
					break
				}
			}

			if isDirIndexCandidate {
				categories[dirCandidate] = append(categories[dirCandidate], item)
			} else {
				categories[categoryName] = append(categories[categoryName], item)
			}
		} else {
			// Sub-item always groups by directory name
			categories[parts[0]] = append(categories[parts[0]], item)
		}
	}

	for categoryName, items := range categories {
		if strings.HasSuffix(categoryName, ".md") {
			// Case 1: Simple root file (e.g., "quick-start.md")
			if len(items) == 1 {
				item := items[0]
				_, err := markdownContent.WriteString(fmt.Sprintf("- [%s](%s)\n", item.Title, item.Path))
				if err != nil {
					log.Fatalf("error writing direct file link: %v", err)
				}
			}
		} else {
			// Case 2: Directory (e.g., "web-components")
			var categoryLinkItem *types.NavItem
			var subItems []types.NavItem

			for i := range items {
				item := items[i]

				relativePath := strings.TrimPrefix(item.Path, "docs/")
				parts := strings.Split(relativePath, "/")

				if len(parts) == 1 {
					// This is the directory index file (e.g., "docs/web-components.md")
					categoryLinkItem = &item
				} else {
					// This is a sub-item (e.g., "docs/web-components/nav-renderer.md")
					subItems = append(subItems, item)
				}
			}

			// Render the category header link (will be the index file if found)
			if categoryLinkItem != nil {
				_, err := markdownContent.WriteString(fmt.Sprintf("- [%s](%s)\n", categoryLinkItem.Title, categoryLinkItem.Path))
				if err != nil {
					log.Fatalf("error writing directory link: %v", err)
				}
			} else {
				// Fallback if no index file is found, but a directory group exists
				_, err := markdownContent.WriteString(fmt.Sprintf("- <span>%s</span>\n", categoryName))
				if err != nil {
					log.Fatalf("error writing directory name: %v", err)
				}
			}

			// Render the sub-items
			for _, item := range subItems {
				_, err := markdownContent.WriteString(fmt.Sprintf("\t- [%s](%s)\n", item.Title, item.Path))
				if err != nil {
					log.Fatalf("error writing sub-item: %v", err)
				}
			}
		}
	}

	return markdownContent.String()
}

// SortPageSection sorts the top-level list items in the generated markdown content alphabetically.
func (mdlook *MDLookManager) SortPageSection(content string) string {
	// Regex to extract the sort key from the first line of a block:
	// Captures 1: [Link Title](path) or 2: <span>Category</span> or 3: Plain Text
	re := regexp.MustCompile(`^- (?:\[([^\]]+)\]\(.+?\)|<span>(.+?)<\/span>|(.+))`)

	lines := strings.Split(strings.TrimSpace(content), "\n")
	if len(lines) == 0 {
		return ""
	}

	var blocks []MarkdownBlock
	var currentBlock strings.Builder
	var currentKey string

	// Parse content into blocks
	for i, line := range lines {
		if strings.HasPrefix(line, "- ") {
			// Found the start of a new top-level block

			if currentBlock.Len() > 0 {
				// Save the previous block
				blocks = append(blocks, MarkdownBlock{SortKey: currentKey, Content: strings.TrimSpace(currentBlock.String())})
				currentBlock.Reset()
				currentKey = ""
			}

			// Extract the sort key from the new line
			key := ""
			matches := re.FindStringSubmatch(line)
			if len(matches) > 0 {
				// Use the first non-empty match group:
				if matches[1] != "" { // Link title
					key = matches[1]
				} else if matches[2] != "" { // Span text
					key = matches[2]
				} else if matches[3] != "" { // Plain text
					key = strings.TrimSpace(matches[3])
				}
			}
			currentKey = key
			currentBlock.WriteString(line + "\n")
		} else if strings.HasPrefix(line, "\t- ") {
			// Sub-item belongs to the current block
			currentBlock.WriteString(line + "\n")
		} else if i == len(lines)-1 && currentBlock.Len() > 0 {
			// If the last line wasn't a list item and we have a block, ensure it's saved.
		}
	}

	// Save the final block
	if currentBlock.Len() > 0 {
		blocks = append(blocks, MarkdownBlock{SortKey: currentKey, Content: strings.TrimSpace(currentBlock.String())})
	}

	// Sort the blocks alphabetically by the key
	sort.Slice(blocks, func(i, j int) bool {
		return strings.ToLower(blocks[i].SortKey) < strings.ToLower(blocks[j].SortKey)
	})

	// Reassemble the sorted content
	var sortedContent strings.Builder
	for _, block := range blocks {
		sortedContent.WriteString(block.Content + "\n")
	}

	return strings.TrimSpace(sortedContent.String()) + "\n"
}
