package test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_getNodeIndex(t *testing.T) {
	index, line := getNodeIndex("peer2.org1.example.com")
	t.Logf("index: %d; line: %d", index, line)
}

func getNodeIndex(nodeUUID string) (int, int) {
	args := strings.Split(nodeUUID, ".")
	if len(args) == 0 {
		return 0, 16
	}
	index := 0
	for _, s := range args[0] {
		fmt.Printf("s :%d", s)
		if s >= '0' && s <= '9' {
			index = int(s-'0') + 10*index
		}
	}
	//rune 2
	return index, 28
}
