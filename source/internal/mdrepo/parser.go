package mdrepo

import (
	"mdlook/source/internal/types"
	"regexp"
	"strings"
)

func (mdlook *MDLookManager) NavParser(navContent string) types.NavRenderStruct {
	var navItems []types.NavItem
	var headerTitle string
	lines := strings.Split(navContent, "\n")

	// Regex pattern to match markdown links
	linkPattern := regexp.MustCompile(`- \[([^\]]+)\]\(([^)]+)\)`)

	// Iterate through each line and process header and nav items
	for _, line := range lines {
		// Parse header (line starts with #)
		if strings.HasPrefix(line, "# ") && headerTitle == "" {
			headerTitle = strings.TrimSpace(strings.TrimPrefix(line, "# "))
			continue
		}

		// Parse navigation items using the regex
		matches := linkPattern.FindStringSubmatch(line)
		if len(matches) == 3 {
			// matches[1] is the title, matches[2] is the path
			title := matches[1]
			path := matches[2]

			navItems = append(navItems, types.NavItem{Title: title, Path: path})
		}
	}

	// If no header title is found, use a default title
	if headerTitle == "" {
		headerTitle = "Documentation"
	}

	// Return the parsed navigation structure
	return types.NavRenderStruct{
		HeaderTitle: headerTitle,
		NavItems:    navItems,
	}
}
