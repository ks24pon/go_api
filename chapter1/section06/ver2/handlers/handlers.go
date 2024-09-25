package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// GET /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// POST /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

// ArticleListHandler関数(Get /article/listのハンドラ)
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// クエリパラメータをreq.URLフィールドとURL型Queryメソッドから取得
	queryMap := req.URL.Query()
	// 「何ページ目の記事一覧が欲しいのか」を変数pageに格納
	var page int
	// パラメータpageが1個以上あるなら
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])

		// 数値に変換できない値だった場合は404番エラーを返す
		if err != nil {
			http.Error(w, "Invalid query paprameter", http.StatusBadRequest)
			return
		}
	// パラメータpageが存在しなかった場合
	} else {
		// page１がついていた時と同じ処理をしたい
		page = 1
	}
	resString := fmt.Sprintf("Article List (page %d)\n", page)
	// レスポンス返却
	io.WriteString(w, resString)

}

// GET /article/{id} のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

// POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

// POST /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}