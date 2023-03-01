package voiceAlert

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkVoiceAlert(b *testing.B) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < b.N; i++ {
		b.Log(b.N)
	}
}
func TestCustomizedOnMacWithChinese(t *testing.T) {
	words := "我有一对大山雀和一只小猫咪"
	Customize(words, Lilian)
}

func TestCustomizedOnMacWithEnglish(t *testing.T) {
	words := "i have a pair of big tits and a small pussy"
	customizedOnMac(Victoria, words)
}
