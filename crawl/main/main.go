package main

import (
	. "crawl/cacheSaver"
	"crawl/urlValidator"
	"fmt"
	"net"
)

var cache map[string][]net.IP

func findIP(url string) {
	var ipChannel, cacheIPChannel chan []net.IP
	ipChannel = make(chan []net.IP)
	cacheIPChannel = make(chan []net.IP)
	go dnsLoopUp(url, ipChannel)
	go cacheLookUp(url, cacheIPChannel)

	select {
	case ip := <- cacheIPChannel:
		fmt.Println(ip)
	case ip := <- ipChannel:
		fmt.Println(ip)
		cache[url] = ip
	}
}

func dnsLoopUp(url string, ipChannel chan []net.IP){
	ip, _ :=  net.LookupIP(url)
	ipChannel <- ip
}

func cacheLookUp(url string, cacheIPChannel chan []net.IP){
	if ip, ok := cache[url]; ok {
		cacheIPChannel <- ip
	}
}

func main() {
	var url string
	cache = make(map[string][]net.IP)
	Load(&cache)
	defer Save(cache)
	fmt.Println(cache)
	for {
		_, _ = fmt.Scan(&url)
		if urlValidator.Validate(url) {
			findIP(url)
		} else if url != "exit" {
			fmt.Println("Invalid Url")
		} else {
			break
		}
	}
}

