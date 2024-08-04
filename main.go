package main

import (
	"godoc"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("godoc", func(w *unison.Window) {
		godoc.New("D:\\workspace\\workspace\\app").Layout()
		w.Content().AddChild(godoc.New("D:\\workspace\\workspace\\branch\\golibrary").Layout())
	})
}
