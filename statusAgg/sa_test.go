package statusAggregator

import (
	"testing"
	"time"
)

func TestStatusResponseToString(t *testing.T) {
	now := time.Now()

	s := StatusResponse{"Hello", 5, now}.String()

	matchString := now.Format("2006-01-02T15:04:05.999 MST") + " 5 Hello"

	if s != matchString {
		t.Error(
			"expected {",
			matchString,
			"} got {",
			s,
			"}",
		)
	}
}

func TestStatusMapStartsEmpty(t *testing.T) {
	smLen := len(StatusMap)
	if smLen != 0 {
		t.Error(
			"expected {",
			0,
			"} got {",
			smLen,
			"}",
		)
	}
}
