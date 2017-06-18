package main

import (
	"fmt"
	"MikeAustin71/exec_go/01_exec-examples/03_cmd-format/common"
	"bytes"
	"os/exec"
)

func main()  {

	if DeleteTargetFiles() != nil {
		return
	}

	CopyFiles()

}

func DeleteTargetFiles()  error {

	fh := common.FileHelper{}
	dir, _ := fh.MakeAbsolutePath("../T04")
	s := dir + "/*.*"
	target := fh.AdjustPathSlash(s)
	cmdStr := "cmd.exe"
	fmt.Println("Target: ", target)
	cmd := exec.Command(cmdStr, "/c", "DEL", "/Q", target)
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("*** ERROR DeleteTargetFiles() ***")
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println(out.String() )
		return err
	}

	fmt.Println("Result DeleteTargetFiles(): " )
	fmt.Println(out.String())

	return nil
}

func CopyFiles() error {
	fh := common.FileHelper{}
	srcDir, _ := fh.MakeAbsolutePath("../T03")

	src := fh.AdjustPathSlash(srcDir + "/*.*")

	d, _  := fh.MakeAbsolutePath("../T04")

	destDir := fh.AdjustPathSlash(d)


	cmd := exec.Command("cmd.exe", "/c", "COPY ", src, destDir)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("*** ERROR CopyFiles() ***")
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println(out.String() )
		return err
	}

	fmt.Println("Result CopyFiles() : " )
	fmt.Println(out.String())


	return nil

}