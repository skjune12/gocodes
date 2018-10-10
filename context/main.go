package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Server listen on port :8888")

	http.HandleFunc("/", indexHandler)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		os.Exit(1)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		errCh <- request(ctx)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			fmt.Fprintln(w, "failed:", err)
		}
	}

	fmt.Fprintln(w, "success")
}

func request(ctx context.Context) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		return err
	}

	errCh := make(chan error, 1)
	go func() {
		res, err := client.Do(req)
		if err != nil {
			errCh <- err
		}
		defer res.Body.Close()

		byteArray, err := ioutil.ReadAll(res.Body)
		if err != nil {
			errCh <- err
		}

		fmt.Println(string(byteArray))
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-errCh
		return ctx.Err()
	}

	return nil
}
