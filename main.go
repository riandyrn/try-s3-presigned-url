package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/riandyrn/go-env"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	serverAddr       = ":8765"
	queryKeyFilename = "filename"
)

const (
	envKeyBucketName        = "BUCKET_NAME"
	envKeyURLExpireDuration = "URL_EXPIRE_DURATION"
)

func main() {
	// define s3 service
	svc := s3.New(session.Must(session.NewSession()))
	// define handler for generating presigned upload url
	http.HandleFunc("/upload-url", func(w http.ResponseWriter, r *http.Request) {
		// check if http method is GET, otherwise reject it
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// get filename from query params
		vals := r.URL.Query()
		filename := vals.Get(queryKeyFilename)
		// create request for put object on S3
		req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
			Bucket: aws.String(env.GetString(envKeyBucketName)),
			Key:    aws.String(filename),
			ACL:    aws.String("public-read"),
		})
		// generate presigned url
		uploadURL, err := req.Presign(time.Duration(env.GetInt(envKeyURLExpireDuration)) * time.Second)
		if err != nil {
			msg := fmt.Sprintf("unable to create presigned url due: %v", err)
			log.Printf("[INTERNAL ERROR] %v", msg)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(msg))
			return
		}
		// return success response
		w.Write([]byte(uploadURL))
	})
	// define handler for serving web front-end
	http.Handle("/web", http.RedirectHandler("/web/", 301))
	http.HandleFunc("/web/", func(w http.ResponseWriter, r *http.Request) {
		fs := http.StripPrefix("/web", http.FileServer(http.Dir("./web")))
		fs.ServeHTTP(w, r)
	})
	// start server
	log.Printf("server is listening on %v", serverAddr)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatalf("unable to start server due: %v", err)
	}
}
