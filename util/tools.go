package util

import (
	"bufio"
	"fmt"
	"os/exec"
)

// CmdExe 命令执行
func CmdExe(cmd string, dstcmd []string) (*exec.Cmd, *bufio.Scanner, error) {

	c := exec.Command(cmd, dstcmd...)
	stdout, err := c.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	_, oerr := c.StderrPipe()
	if oerr != nil {
		fmt.Println(oerr)
	}
	c.Start()
	scanner := bufio.NewScanner(stdout)
	return c, scanner, nil
}
