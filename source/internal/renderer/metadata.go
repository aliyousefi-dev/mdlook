package renderer

import (
	"fmt"
	"mdlook/source/internal/types"
	"os"
	"regexp"
)

// LoadMetadataFromMarkdown reads a Markdown file and loads the metadata
func (renderer *Renderer) LoadMetadataFromMarkdown(filepath string) (*types.MarkdownMetaData, error) {
	// Read the entire file content using os.ReadFile
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Regex pattern to find the frontmatter block (JSON within <!-- -->)
	// This matches everything between <!-- and -->, including the curly braces and whitespace
	re := regexp.MustCompile(`(?s)<!--\s*(\{.*\})\s*-->`)

	// Find the first match
	matches := re.FindSubmatch(data)
	if len(matches) < 2 {
		return &types.MarkdownMetaData{
			NavOrder: 0, // Default NavOrder if no metadata found
		}, nil
	}

	// The first capture group contains the JSON string
	metadataJSON := string(matches[1])

	// Unmarshal the JSON into the MarkdownMetaData struct
	metadata, _ := types.UnmarshalMarkdownMetaDataJSON(metadataJSON)

	return metadata, nil
}
