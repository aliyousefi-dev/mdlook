package renderer

import (
	"encoding/json"
)

func (renderer *Renderer) JsonRender() string {
	scanned := renderer.ScanNavTree()

	// Marshal the NavNode into JSON
	jsonData, err := json.MarshalIndent(scanned, "", "  ")
	if err != nil {
		return ""
	}

	return string(jsonData)
}
