package main

import (
	"godoc"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("godoc", func(w *unison.Window) {
		// godoc.New("D:\\workspace\\workspace\\app").Layout(w.Content())
		godoc.New("D:\\workspace\\workspace\\branch\\golibrary").Layout(w.Content())
	})
}
