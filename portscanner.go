package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

var resultsList []string
var ip string = "172.25.85.55" // CHANGE ME !!!

func portScanner(wg *sync.WaitGroup, ip string, port string) bool {
	fullHost := ip + ":" + string(port)
	conn, err := net.Dial("tcp", fullHost)
	defer wg.Done()
	if err != nil {
		fmt.Printf("Port %s is closed\n", port)
		return false
	} else {

		fmt.Printf("Port %s open \n", port)
		resultsList = append(resultsList, port)
		defer conn.Close()
		return true
	}
}
func main() {
	fmt.Println("Port Scanner")
	maxPort := 65535
	wg := new(sync.WaitGroup)
	fmt.Println("Scanning the port")
	for i := 1; i <= maxPort; i++ {
		wg.Add(1)
		go portScanner(wg, ip, strconv.Itoa(i))

	}
	wg.Wait()
	println("The open ports are:")
	for _, port := range resultsList {
		println(port)
	}
}
