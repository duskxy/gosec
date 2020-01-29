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

func CmdExe(cmd string,dstcmd []string,mt interface{},ws interface{}) {
	c := exec.Command(cmd, dstcmd...)
	// out, err := c.CombinedOutput()
	stdout, err := c.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	c.Start()
	scanner := bufio.NewScanner(stdout)
    for scanner.Scan() {
        ws.WriteMessage(mt,[]byte(scanner.Text()))
    }
	// return string(out)
}

