package cacheSaver

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net"
	"os"
)

func Save(cache map[string][]net.IP){

	file, err := os.Create("cache.cache")
	if err != nil {
		panic(err)
	}
	bufferWriter := bufio.NewWriter(file)
	for k, v := range cache {
		buffer := new(bytes.Buffer)
		buffer.WriteString(k)
		for _, ip := range v {
			_, _ = buffer.WriteString(",")
			_, _ = buffer.WriteString(net.IP.String(ip))
		}
		line := buffer.String()
		_, _ = bufferWriter.WriteString(line)

	}
	defer file.Close()
}

func Load(cache *map[string][]net.IP){
	file, err := os.Open("cache.cache")
	defer file.Close()
	if err != nil {
		fmt.Println("No cache available")
		return
	}
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			fmt.Println("Cache loaded successfully")
			break
		} else if err != nil {
			panic(err)
		} else {
			for i := range record {
				(*cache)[record[0]] = append((*cache)[record[0]], net.ParseIP(record[i]))
			}
		}
	}
}