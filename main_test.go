package voiceAlert

import (
	"math/rand"
	"testing"
	"time"
)

func TestVoice(t *testing.T) {
	VoiceAlert(1)
}
func BenchmarkVoiceAlert(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		b.Log(b.N)
	}
}
