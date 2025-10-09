package types

import (
	"encoding/json"
	"fmt"
)

// MarkdownMetaData represents the metadata structure for navigation
type MarkdownMetaData struct {
	NavOrder int      `json:"nav_order"`           // Order of the navigation item (default to 0 if not provided)
	NavTitle string   `json:"nav_title,omitempty"` // Title of the navigation item (default to "" if not provided)
	NavTags  []string `json:"nav_tags,omitempty"`  // Tags associated with the navigation item (default to empty slice if not provided)
}

func NewMarkdownMetaData() MarkdownMetaData {
	return MarkdownMetaData{
		NavOrder: 0,
		NavTitle: "",
		NavTags:  []string{},
	}
}

func UnmarshalMarkdownMetaDataJSON(data string) (*MarkdownMetaData, error) {

	var metadata MarkdownMetaData = NewMarkdownMetaData()
	if err := json.Unmarshal([]byte(data), &metadata); err != nil {
		fmt.Println("Error unmarshalling metadata JSON:", err)
		return nil, err
	}

	// Ensure that NavOrder defaults to 0 if not provided
	if metadata.NavOrder == 0 {
		// You can set a default value here explicitly if desired, but it's redundant because 0 is already the default for int
		metadata.NavOrder = 0
	}

	return &metadata, nil
}
