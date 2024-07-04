package handler

import "context"

type Handler interface {
	Merge(ctx context.Context, inputFiles []string, outputFile string) error
	Split(ctx context.Context, inputFile string, outputDir string) error
}

type PDFcpuHandler struct {
}

func NewPDFHandler() *PDFcpuHandler {
	return &PDFcpuHandler{}
}
