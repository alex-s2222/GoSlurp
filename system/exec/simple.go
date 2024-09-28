package exec

import (
	"log"
	"os/exec"
)

func Run_simple_app() {
	cmd := exec.Command("firefox")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}