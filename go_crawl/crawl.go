package main

import (
	"fmt"
	"net"
)

func main() {
	var url string
	fmt.Scan(&url)
	ip, _ := net.LookupIP(url)
	fmt.Println(ip)
}
