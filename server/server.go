package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	helloworld "gitlab.com/mc_go/firebase_example"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(helloworld.NewExecutableSchema(helloworld.Config{Resolvers: &helloworld.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//Goのコード生成コマンド「go generate ./...」から以下コメントのコマンドを実行する
//「./...」でどういう挙動になるのかわからない・・・

//go:generate go run github.com/99designs/gqlgen ../
//go:generate sqlboiler --wipe mysql --output ../domain/entity -d -c ../sqlboiler.yaml --no-tests
