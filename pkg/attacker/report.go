package attacker

import (
	"fmt"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func TextReport(name, path string, metrics vegeta.Metrics) {
	makeTextReport(name, path, metrics)
}

func makeTextReport(name, path string, metrics vegeta.Metrics) {
	f, _ := os.Create(fmt.Sprintf("%s/%s_%s.report", path, name, time.Now().Format(time.RFC3339)))
	defer f.Close()

	err := vegeta.NewTextReporter(&metrics).Report(f)
	if err != nil {
		fmt.Println(err)
	}
}
