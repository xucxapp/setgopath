package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// 获取当前运行目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前目录失败: %v\n", err)
		return
	}

	// 执行 go env -w 命令设置 GOPATH
	cmd := exec.Command("go", "env", "-w", "GOPATH="+currentDir)
	if err := cmd.Run(); err != nil {
		fmt.Printf("设置 GOPATH 失败: %v\n", err)
		return
	}

	fmt.Printf("成功设置 GOPATH 为: %s\n", currentDir)
}
