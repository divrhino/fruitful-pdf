package main

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
}
