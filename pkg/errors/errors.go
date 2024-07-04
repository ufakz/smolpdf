package errors

import "errors"

var (
	ErrInsufficientInputFiles = errors.New("at least two input files are required")
	ErrMergeFailed            = errors.New("failed to merge PDFs")
	ErrSplitFailed            = errors.New("failed to split PDF")
	ErrCreateOutputDir        = errors.New("failed to create output directory")
	ErrInvalidCommand         = errors.New("invalid command")
)
