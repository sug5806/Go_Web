package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	// r.Method를 사용해 POST인지 아닌지 체크한다
	if r.Method != http.MethodPost {
		// Header().Set() 메소드를 통해 허용하는 메소드가 무엇인지 추가해준다
		w.Header().Set("Allow", http.MethodPost)
		/* POST가 아니면 405 상태코드를 보내고
		 w.Write를 사용해 Method Not Allowed 를 작성한다
		\*

		//w.WriteHeader(405)
		//w.Write([]byte("Method Not Allowed"))

		/* Upgrade */
		http.Error(w, "Mehtod Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet...."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
