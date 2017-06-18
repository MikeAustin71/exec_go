package main

import (
	"fmt"
	"os/exec"
	"bytes"
	"io"

)

func main()  {

	Ex01()


}

func Ex01()  {

	cmdStr := "cmd.exe"
	cmd := exec.Command(cmdStr, "/c", "DEL", "D:\\T04\\*.*")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err) //replace with logger, or anything you want
	}

	defer stdin.Close() // the doc says subProcess.Wait will close it, but there is some doubt
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	cmd.Start()
	io.WriteString(stdin, "Y\n")

	cmd.Wait()

	fmt.Println("StdOut: ", out.String())
	fmt.Println("StdErr: ", stderr.String())

	/*
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	} else {
		fmt.Println(string(output))
	}
 */

}

func Ex02() {
	cmdStr := "cmd.exe"
	cmd := exec.Command(cmdStr, "/c", "DEL", "D:\\T04\\*.*")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	} else {
		fmt.Println(string(output))
	}

}

func Ex03()  {

	cmdStr := "cmd.exe"
	cmd := exec.Command(cmdStr, "/c", "DEL", "/Q", "D:\\T04\\*.*")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	} else {
		fmt.Println(string(output))
	}

}

func Ex04()  {
	cmdStr := "cmd.exe"

	cmd := exec.Command(cmdStr, "/c", "COPY ", "D:\\T03\\*.*", "D:\\T04")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	} else {
		fmt.Println(string(output))
	}

	fmt.Println()

}

func Ex05()  {

	cmdStr := "cmd.exe"

	cmd := exec.Command(cmdStr, "/c", "COPY ", "D:\\T99\\*.*", "D:\\T04")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	} else {
		fmt.Println(string(output))
	}

	fmt.Println()
}

func Ex06()  {

	cmdStr := "cmd.exe"

	cmd := exec.Command(cmdStr, "/c", "COPY ", "D:\\T99\\*.*", "D:\\T04")

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("*** ERROR ***")
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println(out.String() )
		return
	}

	fmt.Println("Result: " )
	fmt.Println(out.String())

}