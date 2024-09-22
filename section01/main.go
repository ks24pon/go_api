package main
import ( "io"
"log"
"net/http"
)
func main() {
// helloHandlerという名前でハンドラを定義
helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// ハンドラの第一引数として渡されていたhttp.ResponseWriter型の変数wに
		// 何がきても、"Hello,world"を返す処理
        io.WriteString(w, "Hello, world!\n")
    }
	// 定義したhelloHandlerを使うように登録
    http.HandleFunc("/hello", helloHandler)
	// サーバー起動時のログ出力
    log.Println("server start at port 8080")
	// ListenAndServeでサーバーを起動
    log.Fatal(http.ListenAndServe(":8080", nil))
}