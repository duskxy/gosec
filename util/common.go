package util

import (
	"math/rand"
	"time"
	"os"
	"strings"
	"path/filepath"
)

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// GetCurrentPath 返回项目路径
func GetCurrentPath() string {
    s, err := os.Getwd()
    if err != nil {
		panic(err)
	}
	path := strings.Replace(s, string(filepath.Separator), "/", -1)
    return path
}
