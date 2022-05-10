package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	url := "Persian.glb"
	file, err := os.Open(url)
	if err != nil {
		return
	}
	data,err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	//copy the relevant headers. If you want to preserve the downloaded file name, extract it with go's url parser.
	w.Header().Set("Content-Disposition", "attachment; filename=Persian.glb")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	w.Write(data)
	//stream the body to the client without fully loading it into memory

}

func main() {
	http.HandleFunc("/", Index)
	fmt.Print("listen on 8080")
	err := http.ListenAndServe("localhost:8000", nil)
	fmt.Print("listen on 8080")
	if err != nil {
		fmt.Println(err)
	}
}