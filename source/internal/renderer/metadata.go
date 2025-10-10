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
	re := regexp.MustCompile(`(?s)<!--\s*(\{.*\})\s*-->`)
	frontmatterMatches := re.FindAllSubmatchIndex(data, -1)

	// Find code fence regions
	codeFenceRe := regexp.MustCompile("(?m)^```")
	codeFenceMatches := codeFenceRe.FindAllIndex(data, -1)
	var codeFenceRegions [][2]int
	for i := 0; i+1 < len(codeFenceMatches); i += 2 {
		codeFenceRegions = append(codeFenceRegions, [2]int{codeFenceMatches[i][0], codeFenceMatches[i+1][1]})
	}
	isInsideCodeFence := func(pos int) bool {
		for _, region := range codeFenceRegions {
			if pos >= region[0] && pos < region[1] {
				return true
			}
		}
		return false
	}

	var metadataJSON string
	for _, m := range frontmatterMatches {
		start := m[0]
		if !isInsideCodeFence(start) {
			metadataJSON = string(data[m[2]:m[3]])
			break
		}
	}

	if metadataJSON == "" {
		return &types.MarkdownMetaData{
			NavOrder: 0, // Default NavOrder if no metadata found
		}, nil
	}

	metadata, _ := types.UnmarshalMarkdownMetaDataJSON(metadataJSON)
	return metadata, nil
}
