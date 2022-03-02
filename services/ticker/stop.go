package ticker

import "github.com/reasje/go_discord_thread_saver/internal/storage"

func StopTicker() {
	storage.Ticker.Stop()
}
