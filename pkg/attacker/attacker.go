package attacker

import (
	"fmt"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func Attacker(cfg LoadCfg, targets []vegeta.Target, responsefunc func([]byte)) vegeta.Metrics {
	rate := vegeta.Rate{Freq: cfg.Freq, Per: cfg.Per}

	atakerConnection := vegeta.MaxConnections(cfg.MaxConnections)
	atakerTimeout := vegeta.Timeout(cfg.Timeout)
	attacker := vegeta.NewAttacker(atakerTimeout, atakerConnection)

	var metrics vegeta.Metrics

	if cfg.ResponseSave && responsefunc != nil {
		responseFile, _ = os.Create(fmt.Sprintf("%s/%s_%s.json", cfg.DataPath, responseFileName, time.Now().Format(time.RFC3339)))
		defer responseFile.Close()
	}

	for res := range attacker.Attack(vegeta.NewStaticTargeter(targets...), rate, cfg.Duration, cfg.Name) {
		if cfg.ResponseSave && responsefunc != nil {
			responsefunc(res.Body)
		}
		metrics.Add(res)
	}

	metrics.Close()

	return metrics
}

func ResponseSaveBody(body []byte) {
	body = append(body, byte('\n'))
	responseFile.Write(body)
}
