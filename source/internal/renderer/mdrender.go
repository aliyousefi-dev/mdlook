package renderer

import (
	"mdlook/source/internal/types"
	"sort"
	"strings"
)

func SortNavNodes(nodes []*types.NavNode) {
	sort.Slice(nodes, func(i, j int) bool {
		// Default order is 0 if there is no nav_order
		orderI := 0
		orderJ := 0

		// Check if nav_order exists and assign the values to orderI and orderJ
		if nodes[i].Metadata != nil {
			orderI = nodes[i].Metadata.NavOrder
		}
		if nodes[j].Metadata != nil {
			orderJ = nodes[j].Metadata.NavOrder
		}

		// Directories should always be at the end
		if nodes[i].IsDir && !nodes[j].IsDir {
			return false // i is dir, j is file: file first
		}
		if !nodes[i].IsDir && nodes[j].IsDir {
			return true // i is file, j is dir: file first
		}

		// If both are 0 or undefined, they are treated as equal
		if orderI == 0 && orderJ == 0 {
			return false
		}

		// Ensure nodes with order 0 or undefined are always at the bottom
		if orderI == 0 || orderI < 1 {
			return false
		}
		if orderJ == 0 || orderJ < 1 {
			return true
		}

		// For non-zero values of nav_order, sort by the nav_order value
		return orderI < orderJ
	})
}

// RenderNavNode prints the NavNode structure with dashes, sorting by nav_order.
func (renderer *Renderer) RenderNavNode(node types.NavNode, depth int) string {
	var result strings.Builder

	// Sort the children based on their nav_order metadata
	SortNavNodes(node.Childs)

	// Recursively build the string with all child nodes, increased depth
	for _, child := range node.Childs {
		// Create indentation based on depth
		indent := strings.Repeat("  ", depth)
		if child.IsDir {
			result.WriteString(indent + "- ### " + child.GetNodeTitle() + "\n")
		} else {
			result.WriteString(indent + "- [" + child.GetNodeTitle() + "](" + child.Path + ")\n")
		}

		// Recurse to build the string for the child's children
		result.WriteString(renderer.RenderNavNode(*child, depth+1))
	}

	return result.String()
}

func (renderer *Renderer) MdRender() string {
	return renderer.RenderNavNode(renderer.ScanNavTree(), 0)
}
