package option

import (
	"errors"
	"net/http"
)

const (
	defaultHTTPPort = 8080
)

type Config struct {
	Port int
}

type ConfigBuilder struct {
	port *int
}

func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

// ↓ Config構造体を作成するためのBuildメソッド
func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	if b.port == nil {
		// ←ポート管理の主要ロジック
		cfg.Port = defaultHTTPPort
	} else {
		if *b.port == 0 {
			cfg.Port = randomPort()
		} else if *b.port < 0 {
			return Config{}, errors.New("port should be positive")
		} else {
			cfg.Port = *b.port
		}
	}
	return cfg, nil
}

func NewServer(addr string, config Config) (*http.Server, error) {
	// ...

	return nil, nil
}

func randomPort() int {
	return 0
}
