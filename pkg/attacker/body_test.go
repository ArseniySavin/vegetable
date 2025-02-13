package attacker

import (
	"fmt"
	"testing"
)

func Test_Body(t *testing.T) {
	path := "./data"

	targets := ReadDataTargets(path)

	if len(targets) == 0 {
		t.Fail()
	}

	body := FillTargetBody(path, targets[0])

	fmt.Print(string(body))

}
