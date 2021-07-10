package main

import (
	"fmt"
	"github.com/borankux/filemaster/msg"
	"github.com/borankux/filemaster/scans"
	"os"
	"strconv"
	"sync"
)


func scan(root string, wg *sync.WaitGroup, cs chan string, depth int) {
	defer wg.Done()
	scans.Walk(root, depth)
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
	//scanForCMD()
	request := msg.HelloRequest{Name: "Hello World"}
	fmt.Println(request)
}

