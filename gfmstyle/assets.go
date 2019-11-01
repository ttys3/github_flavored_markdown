// Package gfmstyle contains CSS styles for rendering GitHub Flavored Markdown.
package gfmstyle

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

func AssetsHandler() http.Handler {
	box := rice.MustFindBox("./_data")
	handler := http.FileServer(box.HTTPBox())
	return handler
}
