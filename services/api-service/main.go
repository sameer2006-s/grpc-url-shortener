package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/sameer2006-s/grpc-url-shortner/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewLinkServiceClient(conn)

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")

		res, err := client.CreateLink(context.Background(), &pb.CreateLinkRequest{Url: url})
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte(res.ShortUrl))
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		res, err := client.GetLink(context.Background(), &pb.GetLinkRequest{ShortUrl: code})
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte(res.Url))
	})

	log.Println("api :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
