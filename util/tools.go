package util

import (
	"os/exec"
	"fmt"
	"bufio"
)

// func CmdExe(cmd string,exepath string,domain string) string {
// 	dstcmd := []string{exepath,"-u",domain,"-e *"}
// 	c := exec.Command(cmd, dstcmd...)
// 	out, err := c.CombinedOutput()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// fmt.Println(string(out))
// 	return string(out)
// }

func CmdExe(cmd string,dstcmd []string) *bufio.Scanner {
	c := exec.Command(cmd, dstcmd...)
	// out, err := c.CombinedOutput()
	stdout, err := c.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	c.Start()
	scanner := bufio.NewScanner(stdout)
	return scanner
    // for scanner.Scan() {
    //     fmt.Println(scanner.Text())
    // }
	// return string(out)
}

