package main

import (
	"context"
	"encoding/json"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	errLenMsg = "the request has more than 20 urls inside"
	netClient = &http.Client{
		Timeout: 5 * time.Second,
	}
)

type Model struct {
	Urls []string `json:"urls"`
}

type Response struct {
	RespMap map[string]string `json:"Response,omitempty"`
}

func NewResponse() *Response {
	var r Response
	r.RespMap = make(map[string]string)
	return &r
}

func sendReq(ctx context.Context, respStruct *Response, url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	res, err := netClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	respStruct.RespMap[url] = res.Status
	return nil
}

func reqHandler(w http.ResponseWriter, req *http.Request) {
	var m Model
	data := NewResponse()
	chMsg := make(chan string, 20)

	err := json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if len(m.Urls) > 20 {
		log.Println(errLenMsg)
		err = errors.New(errLenMsg)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go func() {
		for _, v := range m.Urls {
			chMsg <- v
		}
		close(chMsg)
	}()

	eg, ctx := errgroup.WithContext(req.Context())

	for i := 0; i < 4; i++ {
		eg.Go(func() error {
			for url := range chMsg {
				err := sendReq(ctx, data, url)
				if err != nil {
					log.Println(err)
					return err
				} else {
					select {
					case <-ctx.Done():
						return ctx.Err()
					default:
						continue
					}
				}
			}
			return nil
		})
	}

	err = eg.Wait()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	js, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func maxClients(f http.HandlerFunc, max int) http.HandlerFunc {
	sema := make(chan struct{}, max)

	return func(w http.ResponseWriter, req *http.Request) {
		sema <- struct{}{}
		defer func() { <-sema }()
		f(w, req)
	}
}

func serve(ctx context.Context) (err error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", maxClients(reqHandler, 100))
	srv := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	log.Printf("Server is live")
	<-ctx.Done()
	log.Printf("Server's going down")
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server Shutdown failed: %+v", err)
	}
	log.Print("Server exited normally")
	if err == http.ErrServerClosed {
		err = nil
	}
	return
}

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-sigCh
		log.Printf("System call: %+v", oscall)
		cancel()
	}()

	if err := serve(ctx); err != nil {
		log.Printf("Failed to start: %+v\n", err)
	}
}
