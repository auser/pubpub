package pubpub

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var cmdSetupSystem = &cobra.Command{
	Use:     "setup",
	Aliases: []string{"b"},
	Short:   "Setup ",
	Run:     runSetupSystem,
}

func runSetupSystem(cmd *cobra.Command, args []string) {
	tmp, err := ioutil.TempDir("pubpub", "")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmp)

	writeResource("resources/environment.yml", tmp+"/environment.yml")

	c := fmt.Sprintf("conda env create --name pubpub --file %s", tmp+"/environment.yml")

	out, err := exec.Command("bash", "-c", c).Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(out)
}
