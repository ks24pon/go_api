package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ks24pon/go_api/handlers"
)

func main() {
	// ルータを明示的に作成
	r := mux.NewRouter()
	// パスに対してリクエストがあった際に関数を実行
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/1", handlers.ArticleDataiLHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	// gorilla/muxのルータを明示的に第二引数に渡す
	log.Fatal(http.ListenAndServe(":8080", r))
}
