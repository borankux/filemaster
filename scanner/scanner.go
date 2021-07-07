package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type FileInfo struct {
	Name     string
	FullName string
	IsDir    bool
	Size     int64
	Mode     fs.FileMode
	Time     time.Time
	Extension string
}

func full(root string, path string) string {
	last := root[len(root)-1:]
	middle := ""

	//if ends with /, then let it be the middle
	if last != "/" {
		middle = "/"
	}

	return root + middle + path
}

func (file *FileInfo) fromFs(fi fs.FileInfo, fullName string) {
	file.Name = fi.Name()
	file.Mode = fi.Mode()
	file.IsDir = fi.IsDir()
	file.Time = fi.ModTime()
	file.FullName = fullName
	file.Extension = filepath.Ext(fullName)
}

func walk(root string, depth int) []FileInfo {
	//if it is 0, then it means end
	if depth == 0 {
		return nil
	}
	depth--

	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var set []FileInfo
	for _, f := range files {
		fullName := full(root, f.Name())

		//this is where it prints stuff
		newFi := FileInfo{}
		newFi.fromFs(f, fullName)
		//fmt.Println(strconv.Itoa(int(f.Size())) + "," + newFi.FullName + ","+ newFi.Extension)

		set = append(set, newFi)
		if f.IsDir() {
			walked := walk(fullName, depth)
			set = append(set, walked...)
		}
	}
	return set
}

func scan(root string, wg *sync.WaitGroup, cs chan string, depth int) {
	defer wg.Done()
	walk(root, depth)
	cs <- "... Done ..."
}

func monitorScans(wg *sync.WaitGroup, cs chan string) {
	wg.Wait()
	close(cs)
}

func printWorkers(cs<-chan string, done chan<-bool) {
	for i:= range cs {
		fmt.Println(i)
	}
	done<-true
}

func scanForCMD() {
	root := os.Args[1]
	depth, _ := strconv.Atoi(os.Args[2])
	wg := &sync.WaitGroup{}
	cs := make(chan string)
	wg.Add(1)

	go scan(root, wg, cs, depth)
	go monitorScans(wg, cs)

	done := make(chan bool, 1)
	go printWorkers(cs, done)
	<-done
}

func main() {
	scanForCMD()
}

