package voiceAlert

import (
	"math/rand"
	"testing"
	"time"
)

func TestVoice(t *testing.T) {
	Voice(1)
}
func BenchmarkVoiceAlert(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		b.Log(b.N)
	}
}
func TestCustomizedOnMacWithChinese(t *testing.T) {
	words := "我有一对大山雀和一只小猫咪"
	CustomizedOnMac(Shasha, words)
}

func TestCustomizedOnMacWithEnglish(t *testing.T) {
	words := "i have a pair of big tits and a small pussy"
	CustomizedOnMac(Victoria, words)
}
