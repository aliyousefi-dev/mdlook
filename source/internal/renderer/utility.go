package renderer

import "os"

// Helper function to determine if a path is a directory
func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false // If there's an error, assume it's not a directory
	}
	return info.IsDir()
}
