package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func isAudio(name string) bool {
	extlist := []string{"mp3", "mp4", "wav", "flac"}
	items := strings.Split(name, ".")
	ext := strings.ToLower(items[len(items)-1])
	for _, audioExt := range extlist {
		if ext == audioExt {
			return true
		}
	}
	return false
}

func getTargetName(path string, namelist []string) string {
	items := strings.Split(path, "/")
	itemslen := len(items)
	var tgtName string
	start := itemslen - 3
	if itemslen < 5 {
		start = 2
	}
	tgtName = strings.Join(items[start:itemslen], "_")

	return tgtName
}

func copy(src, dest string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		fmt.Println("write err: ", src, dest, err)
		return err
	}
	return nil
}

func main() {
	var target string
	basepath := "/Volumes/ElementsSE/"
	const MAX_COPY_NUM = 200

	for ; len(target) <= 0; target = strings.TrimSpace(target) {
		fmt.Print("输入搜索关键词：")
		fmt.Scanln(&target)
	}
	resultpath := "/Users/chloelee/Desktop/搜索结果/" + target
	cmd := exec.Command("locate", "-d", basepath+"locate.database", target)
	// cmd := exec.Command("ls", "-a -l")
	var outline, errline bytes.Buffer
	cmd.Stdout = &outline
	cmd.Stderr = &errline
	cmd.Env = os.Environ()
	err := cmd.Run()
	if err != nil {
		fmt.Println(errline)
		log.Fatal("locate err: ", err)
	}

	os.MkdirAll(resultpath, os.ModePerm)
	err = exec.Command("open", resultpath).Run()
	if err != nil {
		log.Fatal("open err: ", err)
	}
	// fmt.Println(outline.String())
	lines := strings.Split(outline.String(), "\n")
	count := 0
	namelist := []string{}
	for _, line := range lines {
		items := strings.Split(line, "/")
		filename := items[len(items)-1]
		if strings.HasPrefix(filename, ".") {
			continue
		}
		if !isAudio(filename) {
			continue
		}
		tgtName := getTargetName(line, namelist)
		namelist = append(namelist, tgtName)
		fmt.Println(tgtName, "-->", line)
		copy(line, resultpath+"/"+tgtName)
		count += 1
		if count > MAX_COPY_NUM {
			break
		}
	}

}
