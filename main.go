package main

import (
	"godoc"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("godoc", func(w *unison.Window) {
		godoc.New().Layout(w.Content())
	})
}
