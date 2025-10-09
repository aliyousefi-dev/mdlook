package types

// NavItem represents an individual navigation item
type NavItem struct {
	Title    string           `json:"title"`    // Title of the navigation item
	Path     string           `json:"path"`     // Path to the corresponding markdown file
	Metadata MarkdownMetaData `json:"metadata"` // Metadata associated with the navigation item
	Childs   []*NavItem       `json:"childs"`   // Child navigation item
	IsDir    bool             `json:"is_dir"`   // Indicates if the item is a directory
}

func NewNavItem(title, path string) NavItem {
	return NavItem{
		Title: title,
		Path:  path,
	}
}
