package mdrepo

import "mdlook/source/internal/types"

func (mdlook *MDLookManager) SyncNav() {
	scanResult := mdlook.ScanDirectory()
	navContent := mdlook.LoadNav()
	parsedNavContent := mdlook.NavParser(navContent)
	mergedChanges := mdlook.MergeChanges(scanResult, parsedNavContent)
	renderMarkdown := mdlook.NavRender(mergedChanges)
	mdlook.CleanNav()
	mdlook.WriteNav(renderMarkdown)
}

func (mdlook *MDLookManager) MergeChanges(scannedNav types.NavRenderStruct, parsedNavContent types.NavRenderStruct) types.NavRenderStruct {
	var mergedNavItems []types.NavItem

	// Step 1: Remove items from parsedNavContent that are not in scannedNav
	for _, parsedItem := range parsedNavContent.NavItems {
		exists := false
		for _, scannedItem := range scannedNav.NavItems {
			if parsedItem.Path == scannedItem.Path {
				exists = true
				break
			}
		}
		if exists {
			mergedNavItems = append(mergedNavItems, parsedItem)
		}
	}

	// Step 2: Add new items from scannedNav to the end of mergedNavItems
	for _, scannedItem := range scannedNav.NavItems {
		exists := false
		for _, parsedItem := range parsedNavContent.NavItems {
			if scannedItem.Path == parsedItem.Path {
				exists = true
				break
			}
		}
		if !exists {
			mergedNavItems = append(mergedNavItems, scannedItem)
		}
	}

	return types.NavRenderStruct{
		HeaderTitle: parsedNavContent.HeaderTitle,
		NavItems:    mergedNavItems,
	}
}
