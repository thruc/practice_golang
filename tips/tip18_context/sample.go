package main

import (
	"context"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	go func() {
		err = publish(r.Context(), response)
	}()

	w.Write([]byte("OK"))
}

func doSomeTask(ctx context.Context, r *http.Request) (string, error) {
	// ...
	return "", nil
}

func publish(ctx context.Context, response string) error {
	// ...
	return nil
}
