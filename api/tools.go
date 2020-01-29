package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
	"bufio"
	"gosec/util"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var pytools = map[string]string{
	"dirsearch": "/extra/dirsearch/dirsearch.py",
}

var ebuf *bufio.Scanner

const (
	pyexe string = "d:/Python36/python.exe"
)

type data struct {
	Keyword string `json:"keyword"`
	Select1 string `json:"select1"`
	Switch1 bool `json:"switch1"`
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
		ppath := util.GetCurrentPath()
		fmt.Println(msg.Keyword)
		dircmd := []string{ppath + pytools["dirsearch"],"-u",msg.Keyword,"-e *"}
		ebuf = util.CmdExe(pyexe,dircmd)
		if err != nil {
			fmt.Println("接收消息失败 ", err)
			continue
		}

		for ebuf.Scan() {
			// fmt.Println(ebuf.Text())
			err = ws.WriteMessage(mt,[]byte(ebuf.Text() + "\r\n"))
		}

		// err = ws.WriteMessage(mt,[]byte("hello"))
		
		if err != nil {
			break
		}
	}
}
