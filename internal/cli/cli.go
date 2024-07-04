// File: internal/cli/cli.go
package cli

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/ufakz/smolpdf/internal/handler"
	"github.com/ufakz/smolpdf/pkg/errors"
)

type CLI struct {
	handler handler.Handler
}

func NewCLI(h handler.Handler) *CLI {
	return &CLI{handler: h}
}

func (c *CLI) Run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("%w: expected 'merge' or 'split' subcommands", errors.ErrInvalidCommand)
	}

	switch os.Args[1] {
	case "merge":
		return c.runMerge(os.Args[2:])
	case "split":
		return c.runSplit(os.Args[2:])
	default:
		return fmt.Errorf("%w: expected 'merge' or 'split' subcommands", errors.ErrInvalidCommand)
	}
}

func (c *CLI) runMerge(args []string) error {
	mergeCmd := flag.NewFlagSet("merge", flag.ExitOnError)
	outputFile := mergeCmd.String("o", "", "Output file name")
	mergeCmd.Parse(args)

	if *outputFile == "" || mergeCmd.NArg() < 2 {
		return errors.ErrInsufficientInputFiles
	}

	ctx := context.Background()
	err := c.handler.Merge(ctx, mergeCmd.Args(), *outputFile)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully merged PDFs into %s\n", *outputFile)
	return nil
}

func (c *CLI) runSplit(args []string) error {
	splitCmd := flag.NewFlagSet("split", flag.ExitOnError)
	inputFile := splitCmd.String("i", "", "Input file name")
	outputDir := splitCmd.String("o", "", "Output directory")
	splitCmd.Parse(args)

	if *inputFile == "" || *outputDir == "" {
		return fmt.Errorf("%w: input file and output directory are required. -i <input_file> -o <output_file>", errors.ErrInvalidCommand)
	}

	ctx := context.Background()
	err := c.handler.Split(ctx, *inputFile, *outputDir)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully split PDF %s into individual pages in %s\n", *inputFile, *outputDir)
	return nil
}
