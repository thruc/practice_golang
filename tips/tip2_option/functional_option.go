package option

import (
	"errors"
	"net/http"
)

const (
	defaultHTTPPort = 8080
)

type options struct {
	// ←設定構造体
	port *int
}

type Option func(options *options) error

// ←設定構造体を更新する関数型を表す
func WithPort(port int) Option {
	// ←ポートを更新する設定関数
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

func NewServer(addr string, opts ...Option) (*http.Server, error) {
	var options options

	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}

	var port int
	if options.port == nil {
		port = defaultHTTPPort
	} else {
		if *options.port == 0 {
			port = randomPort()
		} else {
			port = *options.port
		}
	}

	// ...
	_ = port
	return nil, nil
}

func randomPort() int {
	return 0
}
