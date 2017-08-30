package pubpub

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

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

func writeResource(resourcePath string, resourceOutput string) {
	data, err := Asset(resourcePath)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(resourceOutput, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func runPrintScript(cmd *cobra.Command, args []string) {
	// db.GetAllRepresentatives()
	// log.Printf("Running print command on %s\n", notebookFile)
	tdir, err := ioutil.TempDir("/tmp", "")
	if err != nil {
		log.Fatal(err)
	}

	pyPrinter := tdir + "/jupyter_to_print.py"
	writeResource("resources/jupyter_to_print.py", pyPrinter)
	writeResource("resources/embed_html.py", tdir+"/embed_html.py")
	writeResource("resources/no_code.tpl", tdir+"/no_code.tpl")
	writeResource("resources/custom.css", tdir+"/custom.css")
	writeResource("resources/environment.yml", tdir+"/environment.yml")

	// c := exec.Command("/bin/bash", "source activate pubpub", "&&", pyPrinter, "-pm", notebookFile)
	c := fmt.Sprintf("python \"%s\" -pm \"%s\"", pyPrinter, notebookFile)
	out, err := exec.Command("bash", "-c", c).Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("%s\n", out))
}
