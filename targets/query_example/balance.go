package query_example

import (
	"net/http"

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

	report := attacker.Attacker(cfg, vegetaTargets, attacker.ResponseSaveBody)

	attacker.TextReport(cfg.Name, cfg.DataPath, report)

}
