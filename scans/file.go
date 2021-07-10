package fs

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
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


func full(root string, path string) string {
	last := root[len(root)-1:]
	middle := ""

	//if ends with /, then let it be the middle
	if last != "/" {
		middle = "/"
	}

	return root + middle + path
}

