package pubpub

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cmdPrint = &cobra.Command{
	Use:     "print",
	Aliases: []string{"p"},
	Short:   "print",
	Run:     runPrintScript,
}

var (
	notebookFile string
)

func init() {
	cmdPrint.Flags().StringVarP(&notebookFile, "notebook", "n", "", "Notebook file")
}

func runPrintScript(cmd *cobra.Command, args []string) {
	// db.GetAllRepresentatives()
	log.Printf("Running print command on %s\n", notebookFile)

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exPath := filepath.Dir(ex)
	pyPrinter := exPath + "/resources/jupyter_to_print.py"
	fmt.Println(pyPrinter)

	// c := exec.Command("/bin/bash", "source activate pubpub", "&&", pyPrinter, "-pm", notebookFile)
	c := fmt.Sprintf("python \"%s\" -pm \"%s\"", pyPrinter, notebookFile)
	out, err := exec.Command("bash", "-c", c).Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("%s\n", out))
}
