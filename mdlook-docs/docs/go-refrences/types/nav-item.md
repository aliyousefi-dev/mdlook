# Nav Item

This is part of frontend and client.

This convert the markdown to a custom html for nav support the daisy ui.

## Data Structure

```
// NavItem represents an individual navigation item
type NavItem struct {
	Title string `json:"title"`
	Path  string `json:"path"`
}

// NavRenderStruct represents the structure for rendering the navigation
type NavRenderStruct struct {
	HeaderTitle string    `json:"header_title"`
	NavItems    []NavItem `json:"nav_items"`
}
```
