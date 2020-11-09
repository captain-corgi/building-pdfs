// Link: https://courses.calhoun.io/lessons/les_goph_127
package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

var (
	output = "output/p1.pdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "Arial")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 18)
	pdf.Cell(40, 10, "Hello World")

	if err := pdf.OutputFileAndClose(output); err != nil {
		log.Fatalf("something went wrong: %v", err)
	}
}
