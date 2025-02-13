package attacker

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type Inventory struct {
	Material string
	Count    uint
}

func readTemplate(path string) []byte {
	fullPath := filepath.Join(path, targetFilename)

	_, err := os.Stat(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func ReadTemplate(path string) []byte {
	return readTemplate(path)
}

func FillTargetBody(path string, target map[string]string) []byte {
	templateBody = readTemplate(path)

	return fillTemplates(target)
}

func fillTemplates(target map[string]string) []byte {
	tmpl, err := template.New(targetFilename).Parse(string(templateBody))
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(nil)
	err = tmpl.Execute(buf, target)
	if err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}
