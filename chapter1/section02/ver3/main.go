package main

import (
	"log"
	"net/http"

	"github.com/ks24pon/go_api/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)

	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDataiLHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)

	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
