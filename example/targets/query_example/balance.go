package query_example

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/ArseniySavin/vegetable/pkg/attacker"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func Balance() {
	token := "Bearer {TOKEN}"
	url := "http://localhost:8080/v1/balance/"

	cfg := attacker.LoadCfg{
		DataPath:       "./targets/example/data",
		Name:           "balance",
		Duration:       attacker.SetDuration("1s"),
		Per:            attacker.SetPerSeconds("1s"),
		Timeout:        attacker.SetTimeout("30s"),
		Freq:           2,
		MaxConnections: 65000,
		ResponseSave:   true,
	}

	targetDatas := attacker.ReadDataTargets(cfg.DataPath)

	headers := http.Header{}
	headers.Add("Authorization", token)
	headers.Add("Content-Type", "application/json")

	var vegetaTargets []vegeta.Target
	for _, v := range targetDatas {
		vegetaTargets = append(vegetaTargets, vegeta.Target{
			Method: "GET",
			URL:    url + v["pan"],
			Header: headers,
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
