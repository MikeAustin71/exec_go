package main

import (
	"context"
	"fmt"
	"time"
	"os/exec"
	"bytes"
)

func main()  {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, "sleep", "5")

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	fmt.Println("=========== Result ===========")
	fmt.Println("Stdout: ", out.String())
	fmt.Println("Stderr: ", stderr.String())
	fmt.Println("=========== Result ===========")
	fmt.Println()

	if err != nil {
		// This will fail after 100 milliseconds. The 5 second sleep
		// will be interrupted.
		fmt.Println("!!!!!!!!!!!! Error !!!!!!!!!!!!")
		fmt.Println("err: ", err.Error())
		fmt.Println("Stderr: ", stderr.String())
		fmt.Println("Stdout: ", out.String())
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	}


	fmt.Println()


}

