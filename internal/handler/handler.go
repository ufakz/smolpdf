package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/ufakz/smolpdf/pkg/errors"
)

type Handler interface {
	Merge(ctx context.Context, inputFiles []string, outputFile string) error
	Split(ctx context.Context, inputFile string, outputDir string) error
}

type PDFcpuHandler struct {
}

func NewPDFHandler() *PDFcpuHandler {
	return &PDFcpuHandler{}
}

func (h *PDFcpuHandler) Merge(ctx context.Context, inputFiles []string, outputFile string) error {
	if len(inputFiles) < 2 {
		return errors.ErrInsufficientInputFiles
	}

	conf := model.NewDefaultConfiguration()
	err := api.MergeCreateFile(inputFiles, outputFile, false, conf)

	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrMergeFailed, err)
	}

	return nil
}

func (h *PDFcpuHandler) Split(ctx context.Context, inputFile string, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("%w: %v", errors.ErrCreateOutputDir, err)
	}

	conf := model.NewDefaultConfiguration()
	err := api.SplitFile(inputFile, outputDir, 1, conf)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrSplitFailed, err)
	}

	return nil
}
