package pubpub

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	libsass "github.com/wellington/go-libsass"
)

var cmdRenderSass = &cobra.Command{
	Use:     "sass",
	Aliases: []string{"s"},
	Short:   "Render sass",
	Run:     runRenderSass,
}

var sassFile string
var outputFile string

func init() {
	cmdRenderSass.Flags().StringVarP(&sassFile, "sassfile", "s", "", "File to convert into sass")
	cmdRenderSass.Flags().StringVarP(&outputFile, "output", "o", "", "Output file")
}

func runRenderSass(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", "Running print script")

	fi, err := os.Open(sassFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	fo, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	sassDir := filepath.Dir(sassFile)

	p := libsass.IncludePaths([]string{sassDir})
	s := libsass.OutputStyle(libsass.COMPRESSED_STYLE)

	// Compile
	compiler, err := libsass.New(fo, fi, p, s)
	if err != nil {
		log.Fatal(err)
	}

	if err := compiler.Run(); err != nil {
		log.Fatal(err)
	}
}
