package renderer

import (
	"mdlook/source/internal/types"
	"path/filepath"
	"strings"
)

func (renderer *Renderer) ScanNavTree() types.NavNode {
	paths, _ := renderer.GetAllPaths()

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
				childNode.Metadata, _ = renderer.LoadMetadataFromMarkdown(currentPath)

				parentNode.AddChild(childNode)

				if isDir {
					dirMap[currentPath] = childNode
				} else {
					childNode.Metadata, _ = renderer.LoadMetadataFromMarkdown(currentPath)
				}

				parentNode = childNode
			}

		}
	}

	return *rootNode
}
