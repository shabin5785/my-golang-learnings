package gobyex

import (
	"fmt"
	"io"
	"os/exec"
)

func SpwanProcess() {
	dateCmd := exec.Command("date")
	date, err := dateCmd.Output()
	check(err)
	fmt.Println(string(date))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsout, err := lsCmd.Output()
	check(err)
	fmt.Println(string(lsout))

	// grepCmd1 := exec.Command("grep", "code")

	// grepCmd1.Start()
	// grepIn1, _ := grepCmd1.StdinPipe()
	// grepOut1, _ := grepCmd1.StdoutPipe()
	// grepIn1.Write(lsout)
	// grepIn1.Close()
	// grepBytes, _ = io.ReadAll(grepOut1)
	// grepCmd.Wait()
	// fmt.Println(string(grepBytes))

}
