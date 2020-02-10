package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gosec/util"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var pytools = map[string]string{
	"dirsearch": "/extra/dirsearch/dirsearch.py",
	"Sublist3r": "/extra/Sublist3r/sublist3r.py",
}

var (
	ebuf   *bufio.Scanner
	dircmd []string
)

type data struct {
	Keyword string `json:"keyword"`
	Select1 string `json:"select1"`
	Switch1 bool   `json:"switch1"`
}

func MessCmd(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		msg := &data{}
		err = json.Unmarshal(message, msg)
		if err != nil {
			fmt.Println("接收消息失败 ", err)
			continue
		}
		ppath := util.GetCurrentPath()
		fmt.Println(msg.Keyword)
		if msg.Select1 == "dirsearch" {
			dircmd = []string{ppath + pytools[msg.Select1], "-u", msg.Keyword, "-e *"}
		} else if msg.Select1 == "Sublist3r" {
			dircmd = []string{ppath + pytools[msg.Select1], "-d", msg.Keyword,"-p 80,443"}
		}
		c, ebuf, err := util.CmdExe(os.Getenv("PYEXE"), dircmd)
		if err != nil {
			fmt.Println("命令执行失败 ", err)
			continue
		}

		for ebuf.Scan() {
			if !strings.Contains(ebuf.Text(), "Last request to") {
				if !strings.Contains(ebuf.Text(), "Error Log") {
					fmt.Println(ebuf.Text())
					err = ws.WriteMessage(mt, []byte(ebuf.Text()+"\r\n"))
					if err != nil {
						break
					}
				}
			}
		}
		c.Wait()

	}
}
