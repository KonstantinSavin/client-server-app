package main

import (
	"encoding/json"
	"fmt"
	"mtg_client/internal/model"
	"mtg_client/internal/service"
	"os"
)

func main() {
	var start string

	err := os.MkdirAll("data", 0644)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.OpenFile("data/all.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	for {
		fmt.Print("Запросить данные? [y/n] ")
		fmt.Fscan(os.Stdin, &start)
		if start == "y" {
			dataCh := make(chan model.Data)
			go service.GetData(dataCh)

			go func(ch chan model.Data, file *os.File) {
				for v := range ch {
					data, err := json.MarshalIndent(v, "", " ")
					if err != nil {
						panic(err)
					}
					file.Write(data)
				}
			}(dataCh, outputFile)
		} else {
			break
		}
	}

}
