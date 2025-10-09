package types

import (
	"path/filepath"
	"strings"
)

type NavNode struct {
	Path     string            `json:"path"`     // Path to the corresponding markdown file
	Childs   []*NavNode        `json:"childs"`   // Child navigation item
	Metadata *MarkdownMetaData `json:"metadata"` // Metadata associated with the navigation item
	IsDir    bool              `json:"is_dir"`   // Indicates if the item is a directory
}

func NewNavNode(path string, isDir bool) *NavNode {
	return &NavNode{
		Path:   path,
		Childs: []*NavNode{},
		IsDir:  isDir,
	}
}

func (n *NavNode) GetNodeTitle() string {
	// Use filepath.Base to get the last part of the path (filename or directory name)
	if n.IsDir {
		return filepath.Base(n.Path)
	} else {
		return strings.TrimSuffix(filepath.Base(n.Path), filepath.Ext(n.Path))
	}

}

func (n *NavNode) AddChild(child *NavNode) {
	// Append the given child node to the Childs slice
	n.Childs = append(n.Childs, child)
}
