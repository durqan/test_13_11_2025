package services

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

func GeneratePDFReport(linksList []int) []byte {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Links Status Report")
	pdf.Ln(15)

	pdf.SetFont("Arial", "", 12)

	for _, id := range linksList {
		linkSet, exists := GetLinksSet(id)
		if !exists {
			continue
		}

		pdf.Cell(40, 10, "Set #"+strconv.Itoa(id))
		pdf.Ln(8)

		for url, status := range linkSet.Links {
			statusText := "FAIL"
			if status {
				statusText = "OK"
			}
			urlString := fmt.Sprintf("%s - %s", url, statusText)

			pdf.Cell(40, 8, urlString)
			pdf.Ln(6)
		}
		pdf.Ln(10)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return []byte{}
	}

	return buf.Bytes()
}
