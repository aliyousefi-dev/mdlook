package types

// NavItem represents an individual navigation item
type NavItem struct {
	Title  string     `json:"title"`  // Title of the navigation item
	Path   string     `json:"path"`   // Path to the corresponding markdown file
	Childs []*NavItem `json:"childs"` // Child navigation item
	IsDir  bool       `json:"is_dir"` // Indicates if the item is a directory
}

// NavRenderStruct represents the structure for rendering the navigation
type NavRenderStruct struct {
	HeaderTitle string    `json:"header_title"` // Title for the navigation header
	NavItems    []NavItem `json:"nav_items"`    // List of navigation items
}

func NewNavItem(title, path string) NavItem {
	return NavItem{
		Title: title,
		Path:  path,
	}
}

// NewNavRender creates a new instance of NavRender and initializes data storage.
func NewNavRender(headerTitle string, navItems []NavItem) NavRenderStruct {
	navStruct := NavRenderStruct{
		HeaderTitle: headerTitle,
		NavItems:    navItems,
	}
	return navStruct
}
