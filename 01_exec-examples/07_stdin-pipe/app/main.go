package main

import (
	"fmt"
	"os/exec"
	"log"
	"io"
)

func main()  {
	//args := []string {"/c", "dir"}
	//cmd := exec.Command("cmd.exe", "/c")
	cmd := exec.Command("./rat.exe", "")
	stdin, err := cmd.StdinPipe()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "My rabbit has fleas!")
	}()

	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output: %s\n", out)



}
