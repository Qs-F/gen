package main

type Variables map[string]interface{}

type Doc struct {
	Variables Variables

	BasePath   string
	LayoutPath string
}
