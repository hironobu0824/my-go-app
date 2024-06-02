package main

import (
	"fmt" // フォーマットされたI/Oを提供する
	"net/http" // HTTPクライアントおよびサーバーの実装を提供する
)

// w http.ResponseWriter：HTTPレスポンスを作成するためのインターフェース。このインターフェースを使用して、クライアントにデータを送信します。
// r *http.Request：HTTPリクエストの情報を保持する構造体。この構造体を使用して、リクエストの詳細を取得します。
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// http.HandleFunc("/", handler)：http.HandleFunc関数を使用して、特定のパスに対するリクエストを処理するハンドラーを登録します。この場合、ルートパス（"/"）に対するリクエストはhandler関数で処理されます。
// http.ListenAndServe(":8080", nil)：HTTPサーバーを起動し、ポート8080でリクエストを待ち受けます。この関数はブロッキング関数であり、サーバーが停止するまで実行され続けます。
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
