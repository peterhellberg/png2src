package main

import (
	"flag"
	"fmt"
	"html/template"
	"image"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	template string
	filename string
}

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func parseConfig(args []string) (*Config, error) {
	var cfg Config

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	flags.StringVar(&cfg.template, "template", "image.tpl", "The title of the ART project")

	if err := flags.Parse(args[1:]); err != nil {
		return nil, err
	}

	rest := flags.Args()

	if len(rest) == 0 {
		return nil, fmt.Errorf("no image filename name given as the first argument")
	}

	if !strings.HasSuffix(rest[0], ".png") {
		return nil, fmt.Errorf("not a .png file")
	}

	cfg.filename = rest[0]

	if cfg.filename == "" {
		return nil, fmt.Errorf("no filename")
	}

	return &cfg, nil
}

func run(args []string, stdout io.Writer) error {
	cfg, err := parseConfig(args)
	if err != nil {
		return err
	}

	tpl, err := template.ParseFiles(cfg.template)
	if err != nil {
		return err
	}

	r, err := os.Open(cfg.filename)
	if err != nil {
		return err
	}
	defer r.Close()

	m, err := png.Decode(r)
	if err != nil {
		return err
	}

	n, ok := m.(*image.NRGBA)
	if !ok {
		return fmt.Errorf("expected *image.NRGBA")
	}

	name := strings.TrimSuffix(filepath.Base(cfg.filename), ".png")

	size := m.Bounds().Max

	return tpl.Execute(stdout, Value{
		Name:        name,
		Length:      len(n.Pix),
		Bytes:       n.Pix,
		BytesString: bytesString(n.Pix),
		Width:       size.X,
		Height:      size.Y,
	})
}

func bytesString(bs []byte) string {
	elems := []string{}

	for _, b := range bs {
		elems = append(elems, fmt.Sprintf("%d", b))
	}

	return strings.Join(elems, ", ")
}

type Value struct {
	Name        string
	Length      int
	Bytes       []byte
	BytesString string
	Width       int
	Height      int
}
