package main

import(
	"io"
	"net/http"
)

func Print1to20() int{
	res := 0
	for i := 1; i<= 20; i++{
		res += i
	}
	return res
}

func firstPage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<h1>Hello, this is my first page!</h1>")
}

func main(){
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil)
}