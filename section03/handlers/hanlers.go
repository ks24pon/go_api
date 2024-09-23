package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// Goのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// もしreqの中のMethodがGetだった場合
	if req.Method == http.MethodGet {
		// レスポンス返却
		io.WriteString(w, "Hwllo, World\n")
	} else {
		// レスポンスを405番のステータスコードと共に返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// PostArticleHandler関数
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// もしreqの中のMethodがPOSTだった場合
	if req.Method == http.MethodPost {
		// レスポンス返却
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
// ArticleListHandler関数(Get /article/listのハンドラ)
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// もしGetだった場合
	if req.Method == http.MethodGet {
		// レスポンス返却
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
// ArticleDataiLHandler関数(Get /article/1のハンドラ)
func ArticleDataiLHandler(w http.ResponseWriter, req *http.Request) {
	// もしGetだった場合
	if req.Method == http.MethodGet {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
// PostNiceHandler関数(POST /article/niceのハンドラ)
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// もしPOSTだったら
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
// PostCommentHandler関数
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// もしPOSTだったら
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}