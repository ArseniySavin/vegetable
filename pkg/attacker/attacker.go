package attacker

import (
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func Attacker(cfg LoadCfg, targets []vegeta.Target, body chan<- []byte) vegeta.Metrics {
	rate := vegeta.Rate{Freq: cfg.Freq, Per: cfg.Per}

	atakerConnection := vegeta.MaxConnections(cfg.MaxConnections)
	atakerTimeout := vegeta.Timeout(cfg.Timeout)
	attacker := vegeta.NewAttacker(atakerTimeout, atakerConnection)

	var metrics vegeta.Metrics

	for res := range attacker.Attack(vegeta.NewStaticTargeter(targets...), rate, cfg.Duration, cfg.Name) {
		if cfg.ResponseSave {
			body <- res.Body
		}
		metrics.Add(res)
	}

	metrics.Close()

	return metrics
}
