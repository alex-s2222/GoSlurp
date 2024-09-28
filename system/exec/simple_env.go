package exec

import (
	"fmt"
	"os"
)

func Simple_envs() {

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	for _, e := range os.Environ(){
		fmt.Println(e)
	}
}