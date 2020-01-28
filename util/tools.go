package util

import (
	"os/exec"
	"fmt"
)

func CmdExe(cmd string,domain string) {
	dstcmd := []string{"E://devops//Go//gosec//extra//dirsearch//dirsearch.py","-u","www.xiaoyi.com","-e *"}
	c := exec.Command(cmd, dstcmd...)
	out, err := c.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	
}

func main() {
	CmdExe("d://Python36//python.exe","www.xiaoyi.com")
}
