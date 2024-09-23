package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
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

// ArticleDataiLHandler関数(Get /article/{id}のハンドラ)
func ArticleDataiLHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
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