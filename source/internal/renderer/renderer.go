package renderer

import "mdlook/source/internal/workstation"

type Renderer struct {
	workstation *workstation.Workstation
}

// NewRenderer creates a new instance of Renderer and initializes data storage.
func NewRenderer(workstation *workstation.Workstation) *Renderer {
	renderer := &Renderer{
		workstation: workstation,
	}
	return renderer
}
