package mdrepo

import (
	"encoding/json"
	"fmt"
	"mdlook/source/internal/types"
	"os"
	"path/filepath"
	"strings"
)

func (mdlook *MDLookManager) GetAllPaths() ([]string, error) {
	var allFilesAndFolders []string
	docsFolderPath := mdlook.GetDocsFolderPath()

	// Walk through the docs folder and all subfolders
	err := filepath.Walk(docsFolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // if there's an error accessing the file/folder, return it
		}

		targetPath := filepath.ToSlash(path)
		// Append the path to the list (both files and directories)
		allFilesAndFolders = append(allFilesAndFolders, targetPath)
		return nil
	})

	if err != nil {
		return nil, err // if an error occurred, return it
	}

	return allFilesAndFolders, nil
}

// ScanNavTree converts all paths into a hierarchical NavNode structure.
func (mdlook *MDLookManager) ScanNavTree() types.NavNode {
	paths, _ := mdlook.GetAllPaths()

	if len(paths) == 0 {
		return types.NavNode{}
	}

	rootNode := types.NewNavNode(paths[0], true)
	dirMap := make(map[string]*types.NavNode)
	dirMap[rootNode.Path] = rootNode

	for _, path := range paths {
		parts := strings.Split(path, "/")
		var parentNode *types.NavNode = rootNode

		for _, part := range parts[1:] {
			currentPath := filepath.Join(parentNode.Path, part)
			currentPath = filepath.ToSlash(currentPath)
			isDir := isDirectory(currentPath)

			if existingNode, exists := dirMap[currentPath]; exists {
				parentNode = existingNode
			} else {
				childNode := types.NewNavNode(currentPath, isDir)
				parentNode.AddChild(childNode)

				if isDir {
					dirMap[currentPath] = childNode
				}

				parentNode = childNode
			}
		}
	}

	return *rootNode
}

// Helper function to determine if a path is a directory
func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false // If there's an error, assume it's not a directory
	}
	return info.IsDir()
}

// PrintNavNodeAsJSON pretty prints the NavNode structure as JSON.
func (mdlook *MDLookManager) RenderNavTreeJson(node types.NavNode) {
	// Marshal the NavNode into JSON
	jsonData, err := json.MarshalIndent(node, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))
}

// PrintNavNodeWithDashes prints the NavNode structure with dashes, skipping the root node and starting from its children.
func (mdlook *MDLookManager) RenderNavNode(node types.NavNode, depth int) string {
	var result strings.Builder

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
		result.WriteString(mdlook.RenderNavNode(*child, depth+1))
	}

	return result.String()
}

func (mdlook *MDLookManager) RenderNavTree(node types.NavNode) string {
	return mdlook.RenderNavNode(node, 0)
}
