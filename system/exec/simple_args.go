package exec

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"fmt"
)

func Run_app_with_args() {
	cmd := exec.Command("tr", "a-z", "A-Z")

	cmd.Stdin = strings.NewReader("little slurp gois big")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("translated phrase: %q\n", out.String())
}