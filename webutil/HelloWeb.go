/*
@Time : 2019-01-15 17:32 
@Author : caozg
@File : HelloWeb
*/
package webutil

import (
	"io"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	http.ListenAndServe(":8080", nil)
}
