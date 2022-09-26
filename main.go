package main

import (
	"path/filepath"
	"os"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/sirupsen/logrus"
)

/*
pdfcpu stamp add -pages 1 -mode pdf -- "CollationOMRFirstPage.pdf:1" "rot:0, scale:1 rel" duplicate.pdf out-cli.pdf
*/

func main() {
	currentDir, _ := os.Getwd()
	omr := filepath.Join(currentDir, "omr", "CollationOMROnePage.pdf")
	inFile := filepath.Join(currentDir, "duplicate.pdf")

	f, err := os.Open(inFile)
	if err != nil {
		logrus.Errorf("%v. err opening file", err)
		panic(err)
	}

	out, err := os.Create("out-api.pdf")
	if err != nil {
		logrus.Errorf("%v. err creating output file", err)
		panic(err)
	}

	test, err := api.PDFWatermark(omr, "rot:0, scale:1 rel", true, true, pdfcpu.POINTS)
	if err != nil {
		logrus.Errorf("error, %v", err)
		panic(err)
	}

	logrus.Info("%v", test)

	err = api.AddWatermarks(f, out, []string{"1"}, test, nil)
	if err != nil {
		logrus.Errorf("error adding watermark, %v", err)
		panic(err)
	}
}
