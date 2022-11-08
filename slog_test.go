package voiceAlert

import (
	"golang.org/x/exp/slog"
	"net"
	"os"
	"testing"
)

func TestDemo(t *testing.T) {
	demo()
}
func demo() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout)))

	slog.Info("info", "name", "Al")
	slog.Debug("debug", "name", "zen")
	slog.Warn("warning", "name", "none")
	slog.Error("oops", net.ErrClosed, "status", 500)

	slog.LogAttrs(slog.ErrorLevel, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))
}
