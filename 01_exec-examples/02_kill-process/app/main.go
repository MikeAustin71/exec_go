package main

import (
	"fmt"
	"os/exec"
	"time"

)

func main()  {
	cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		fmt.Println("Fatal Error: ",err)
		return
	}

	done := make(chan error, 1)

	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(3 * time.Second):
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println("failed to kill: ", err)
			return
		}
		fmt.Println("process killed as timeout reached")
	case err := <-done:
		if err != nil {
			fmt.Printf("process done with error = %v /n", err)
		} else {
			fmt.Printf("process done gracefully without error /n")
		}
	}

	fmt.Println()
}

