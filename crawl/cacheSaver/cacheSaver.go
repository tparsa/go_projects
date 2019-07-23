package cacheSaver

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net"
	"os"
)

func Save(cache map[string][]net.IP){
	file, err := os.OpenFile("cache.cache",  os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for k, v := range cache {
		file.WriteString(k)
		for _, ip := range v {
			_, _ = file.WriteString(",")
			_, _ = file.WriteString(net.IP.String(ip))
		}
		_, err := file.WriteString("\n")
		//bufferWriter.Write()
		if err != nil {
			panic(err)
		}

	}
}

func Load(cache *map[string][]net.IP){
	file, err := os.Open("cache.cache")
	if err != nil {
		fmt.Println("No cache available")
		return
	}
	defer file.Close()
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
				if i > 0 {
					(*cache)[record[0]] = append((*cache)[record[0]], net.ParseIP(record[i]))
				}
			}
		}
	}
}