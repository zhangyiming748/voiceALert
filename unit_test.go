package voiceAlert

import (
	"testing"
)

func TestCustomizedOnMacWithChinese(t *testing.T) {
	words := "我有一对大山雀和一只小猫咪"
	Customize(words, Lilian)
}

func TestCustomizedOnMacWithEnglish(t *testing.T) {
	words := "i have a pair of big tits and a small pussy"
	customizedOnMac(Ava, words)
}
