package mdrepo

func (mdlook *MDLookManager) SyncNav() {
	renderMarkdown := mdlook.Renderer.MdRender()
	mdlook.Workstation.CleanNavFile()
	mdlook.Workstation.UpdateNavFile(renderMarkdown)
}
