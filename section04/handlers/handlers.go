package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// Goのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world\n")
}

// PostArticleHandler関数
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// レスポンス返却
	io.WriteString(w, "Posting Article...\n")
}

// ArticleListHandler関数(Get /article/listのハンドラ)
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// レスポンス返却
	io.WriteString(w, "Article List\n")
}

// ArticleDataiLHandler関数(Get /article/1のハンドラ)
func ArticleDataiLHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

// PostNiceHandler関数(POST /article/niceのハンドラ)
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

// PostCommentHandler関数
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}