package report

import (
	"bytes"
	"fmt"
	"proj/internal/models"

	"github.com/jung-kurt/gofpdf"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(links []models.LinksResponse) (bytes.Buffer, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 24)
	pdf.Cell(0, 20, "List links")
	pdf.Ln(25)

	pdf.SetFont("Arial", "", 18)
	for _, link := range links {
		pdf.Cell(0, 15, fmt.Sprintf("[%d]", link.LinkNums))
		pdf.Ln(10)
		for ln, status := range link.Links {
			pdf.Cell(0, 15, ln+"   --->   "+status)
			pdf.Ln(10)
		}
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	return buf, err
}
