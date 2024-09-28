package exec

import (
	"fmt"
	"os/exec"
)

func Run_multiple_args() {
	prg := "echo"

	args1 := "there"
	args2 := "are slurp"
	args3 := "in Slurpland"

	
	cmd := exec.Command(prg, args1, args2, args3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}