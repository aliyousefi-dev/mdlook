package mdrepo

import (
	"fmt"
	"log"
	"mdlook/source/internal/types"
	"strings"
)

// NavRender generates the markdown content from the NavRenderStruct and returns it as a string
func (mdlook *MDLookManager) NavRender(navRenderStruct types.NavRenderStruct) string {
	// Initialize a strings.Builder to efficiently build the markdown content
	var markdownContent strings.Builder

	// Write the header (Markdown header)
	_, err := markdownContent.WriteString(fmt.Sprintf("# %s\n\n", navRenderStruct.HeaderTitle))
	if err != nil {
		log.Fatalf("error writing header: %v", err)
	}

	// Write the navigation items (Markdown list)
	for _, item := range navRenderStruct.NavItems {
		_, err := markdownContent.WriteString(fmt.Sprintf("- [%s](%s)\n", item.Title, item.Path))
		if err != nil {
			log.Fatalf("error writing nav item: %v", err)
		}
	}

	// Return the generated markdown content as a string
	return markdownContent.String()
}
