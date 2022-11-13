package config

import "time"

type Option func(*JS)

func withConnectTimeout(dur time.Duration) Option {
	return func(cfg *JS) {
		cfg.ConnectTimeout = dur
	}
}
func withMaxReconnect(count int) Option {
	return func(cfg *JS) {
		cfg.MaxReconnects = count
	}
}

func newJSConf(ops ...Option) JS {
	var js = JS{}
	for _, op := range ops {
		op(&js)
	}
	return js
}
