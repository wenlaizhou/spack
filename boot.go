package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

func main() {

	flag.Parse()

	processDir := flag.Arg(0)

	if len(processDir) <= 0 {
		processDir = "."
	}

	err := filepath.Walk(processDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fullName := ""
		if processDir == "." {
			fullName = fmt.Sprintf("%v/%v", path, info.Name())

		} else {
			fullName = fmt.Sprintf("%v/%v/%v", processDir, path, info.Name())
		}
		if filepath.Ext(path) == "js" {
			// process js
			println(fmt.Sprintf("开始处理: %v", fullName))
			appendLine("dist.js", fmt.Sprintf("// processed by wenlai file : %s", info.Name()))
			appendLine("dist.js", readStr(fullName))
			return nil
		}
		if filepath.Ext(path) == "css" {
			// process css
			println(fmt.Sprintf("开始处理: %v", fullName))
			appendLine("dist.css", fmt.Sprintf("/*! processed by wenlai file : %s */", info.Name()))
			appendLine("dist.css", readStr(fullName))
			return nil
		}
		return nil
	})
	if err != nil {
		println(err.Error())
	}

}

func readStr(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return ""
	}
	return string(content)
}

func appendLine(filePath string, s string) (int, error) {
	os.MkdirAll(path.Dir(filePath), os.ModePerm)
	var fw *os.File
	if !exists(filePath) {
		fw, _ = os.OpenFile(filePath, os.O_CREATE, os.ModePerm)
	} else {
		fw, _ = os.OpenFile(filePath, os.O_APPEND, os.ModePerm)
	}
	defer fw.Close()
	return fw.WriteString(fmt.Sprintf("%s\n", s))
}

func exists(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}
