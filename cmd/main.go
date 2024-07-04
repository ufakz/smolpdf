package main

import (
	"fmt"
	"os"

	"github.com/ufakz/smolpdf/internal/cli"
	"github.com/ufakz/smolpdf/internal/handler"
)

func main() {
	pdfHandler := handler.NewPDFHandler()
	cli := cli.NewCLI(pdfHandler)

	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
