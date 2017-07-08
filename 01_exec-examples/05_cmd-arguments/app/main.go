package main

import (
	"fmt"
	"bytes"
	"os/exec"
	"MikeAustin71/exec_go/01_exec-examples/05_cmd-arguments/common"
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

	args := []string {"/c", "DEL", "/Q", target}

	cmd := exec.Command(cmdStr, args ...)
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	fmt.Println("-------------- Result --------------")
	fmt.Println("Result DeleteTargetFiles(): " )
	fmt.Println(out.String())
	fmt.Println("stderr: ", stderr.String())
	fmt.Println("------------------------------------")
	fmt.Println("")
	if err != nil {
		fmt.Println("!!!!!!!!!!!!!! ERROR !!!!!!!!!!!!!!!")
		fmt.Println("*** ERROR DeleteTargetFiles() ***")
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		return err
	}


	return nil
}

func CopyFiles() error {
	fh := common.FileHelper{}
	srcDir, _ := fh.MakeAbsolutePath("../T03")

	src := fh.AdjustPathSlash(srcDir + "/*.*")

	d, _  := fh.MakeAbsolutePath("../T04")

	destDir := fh.AdjustPathSlash(d)

	cmdStr := "cmd.exe"
	args := []string {"/c", "COPY ", src, destDir}
	cmd := exec.Command(cmdStr, args ... )

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	fmt.Println("-------------- Result --------------")
	fmt.Println("Result CopyFiles() : " )
	fmt.Println(out.String())
	fmt.Println("------------------------------------")
	fmt.Println()

	if err != nil {
		fmt.Println("!!!!!!!!!!!!!! ERROR !!!!!!!!!!!!!!!")
		fmt.Println("*** ERROR CopyFiles() ***")
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		return err
	}



	return nil

}