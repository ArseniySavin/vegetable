package attacker

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func ReadDataTargets(path string) []map[string]string {
	fullPath := filepath.Join(path, dataFilename)

	_, err := os.Stat(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	return targetDataFile(data)
}

func targetDataFile(data []byte) []map[string]string {
	var targets []map[string]string

	err := json.Unmarshal(data, &targets)
	if err != nil {
		log.Fatal(err, ".	Use example ", example)
	}

	return targets
}
