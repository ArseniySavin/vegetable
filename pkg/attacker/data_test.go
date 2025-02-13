package attacker

import (
	"fmt"
	"testing"
)

func Test_read(t *testing.T) {
	targets := targetDataFile([]byte(example))

	if len(targets) == 0 {
		t.Fail()
	}
	fmt.Println(targets)
}

func Test_Read(t *testing.T) {
	targets := ReadDataTargets("./data")

	if len(targets) == 0 {
		t.Fail()
	}

	fmt.Println(targets)
}
