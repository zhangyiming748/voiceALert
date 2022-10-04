package util

import "testing"

func TestHaveSpoker(t *testing.T) {
	ret := hasSpoker("Kate")
	t.Log(ret)
}
