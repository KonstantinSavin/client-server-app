package service

import (
	"encoding/json"
	"fmt"
	"io"
	"mtg_client/internal/model"
	"net/http"
)

func GetData(ch chan<- model.Data) {
	req, err := http.NewRequest(
		"POST", "http://localhost:8000/data", nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	dataArr := []model.Data{}
	err = json.Unmarshal(b, &dataArr)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range dataArr {
		ch <- v
	}
}
