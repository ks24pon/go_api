package main

import (
	"io"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// helloHandlerという名前でハンドラを定義
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	// postArticleHandlerを定義
	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article...\n")
	}
	// 一覧
	articlelistHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}
	// articleDetailHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	articleID := 1
	// 	resString := fmt.Sprintf("Article No.%d\n", articleID)
	// 	io.WriteString(w, resString)
	// }
	articleDetailHandler := func(w http.ResponseWriter, req *http.Request) {
		articleID := 1
		resString := fmt.Sprintf("Article No%d\n", articleID)
		io.WriteString(w, resString)
	}
	// 記事にいいねつけるエンドポイント
	articlelikeHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Artice Nice・・・\n")
	}
	// コメントエンドポイント
	articlecommentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment・・・\n")
	}
	// 定義したhelloHandlerを使うように登録
	http.HandleFunc("/hello", helloHandler)
	// Article..文字列を返す
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", articlelistHandler)
	http.HandleFunc("/article/1", articleDetailHandler)
	http.HandleFunc("/article/nice", articlelikeHandler)
	http.HandleFunc("/comment", articlecommentHandler)
	// サーバー起動時のログ出力
	log.Println("server start at port 8080")
	// ListenAndServeでサーバーを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
