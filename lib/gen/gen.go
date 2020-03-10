// Package gen is the packeg of internal gen
package gen

// Variables is the type expressing front matters
type Variables map[string]interface{}

// Gen is the struct for gen cmd
type Gen struct {
	BasePath   string
	LayoutPath string
}

// New returns new *Gen
func New(basePath string, layoutPath string) *Gen {
	return &Gen{
		BasePath:   basePath,
		LayoutPath: layoutPath,
	}
}
