package body_example

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/ArseniySavin/vegetable/pkg/attacker"
	gofakeit "github.com/brianvoe/gofakeit/v6"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func Payment() {
	token := "Bearer {token}"
	url := "http://localhost:8080/v1/payment/"

	cfg := attacker.LoadCfg{
		DataPath:       "./targets/payment/data",
		Name:           "payment",
		Duration:       attacker.SetDuration("1s"),
		Per:            attacker.SetPerSeconds("1s"),
		Timeout:        attacker.SetTimeout("30s"),
		Freq:           2,
		MaxConnections: 65000,
		ResponseSave:   true,
	}

	targetDatas := attacker.ReadDataTargets(cfg.DataPath)
	temp := attacker.ReadTemplate(cfg.DataPath)

	headers := http.Header{}
	headers.Add("Authorization", token)
	headers.Add("Content-Type", "application/json")

	var vegetaTargets []vegeta.Target
	for _, v := range targetDatas {
		vegetaTargets = append(vegetaTargets, vegeta.Target{
			Method: "GET",
			URL:    url,
			Header: headers,
			Body:   RndReq(v, temp),
		})
	}

	responseBody := make(chan []byte)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		responseFile, _ := os.Create(fmt.Sprintf("%s/%s_%s.json", cfg.DataPath, attacker.ResponseFileName, time.Now().Format(time.RFC3339)))
		defer responseFile.Close()

		for v := range responseBody {
			v = append(v, byte('\n'))
			responseFile.Write(v)
		}
	}()

	go func() {
		wg.Wait()
		close(responseBody)
	}()

	report := attacker.Attacker(cfg, vegetaTargets, responseBody)

	attacker.TextReport(cfg.Name, cfg.DataPath, report)

}

type PayReq struct {
	Acc    string    `json:"pan"`
	Amount float64   `json:"amount"`
	Stamp  time.Time `json:"stamp"`
}

func RndReq(param map[string]string, body []byte) []byte {

	var req PayReq
	bufReq := bytes.NewBuffer(body)
	defer bufReq.Reset()
	err := json.NewDecoder(bufReq).Decode(&req)
	if err != nil {
		log.Fatal(err)
	}

	req.Acc = param["acc"]
	req.Amount = gofakeit.Price(100, 5000)
	req.Stamp = time.Now()

	bufOut := bytes.NewBuffer(nil)
	defer bufOut.Reset()
	json.NewEncoder(bufOut).Encode(&req)

	return bufOut.Bytes()

}
