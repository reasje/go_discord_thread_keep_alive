package ticker

import (
	"time"

	"github.com/reasje/go_discord_thread_saver/internal/storage"
)

// This will start the operation and assign the ticker
// to the active ticker
func StartTicker() {
	KeepAlive()
	storage.Ticker = time.NewTicker(24 * time.Hour)
	for _ = range storage.Ticker.C {
		KeepAlive()
	}
}
