package main

import (
	"fmt"
	"bytes"
	"os/exec"
	"strings"
	"strconv"
)

func main()  {

	cmd := exec.Command("cmd.exe", "/c", "COPY ", "D:\\T99\\*.*", "D:\\T04")

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("*** ERROR ***")
		errStr := err.Error()
		errAry := strings.Split(errStr," ")

		errCode, _ := strconv.Atoi(errAry[2])

		fmt.Println(errStr + ": " + stderr.String())
		fmt.Println(out.String() )
		fmt.Println("Error Exit Code: ", errCode)
		return
	}

	fmt.Println("Result: " )
	fmt.Println(out.String())


	fmt.Println()

}

