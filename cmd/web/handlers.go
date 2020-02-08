package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// template.ParseFiles 함수를 사용하여 템플릿 파일을 읽고
	// 오류가 있는 경우 오류 메시지를 기록한다
	// http.Error 함수를 사용하여 일반 500 내부 서버 오류 메시지를 사용자에게 보냄
	//ts, err := template.ParseFiles("C:/Users/sug58/go/src/code/snippetbox/cmd/ui/html/home.page.tmpl")
	//if err != nil {
	//	log.Println(err.Error())
	//	http.Error(w, "Internal Server Error", 500)
	//	return
	//}

	//Upgrade
	files := []string{
		"C:/Users/sug58/go/src/code/snippetbox/cmd/ui/html/home.page.tmpl",
		"C:/Users/sug58/go/src/code/snippetbox/cmd/ui/html/base.layout.tmpl",
		"C:/Users/sug58/go/src/code/snippetbox/cmd/ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 템플릿 세트에서 Execute() 메소드를 사용하여 리스폰스 바디로 템플릿 컨텐츠를 보낸다
	// Execute 의 마지막 매개 변수는 전달하려는 동적 데이터를 나타내며 현재는 nil 로 남겨둔다
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
