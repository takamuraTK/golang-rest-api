package main

import (
	"log"
	"net/http"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func() {
		api := rest.NewApi()
		api.Use(rest.DefaultDevStack...)
		api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(map[string]string{
				"fruit": "Apple",
				"size": "Large",
				"color": "Red",
			})
		}))
		log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
	})
}


