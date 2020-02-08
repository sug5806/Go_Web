package main

import (
	"log"
	"net/http"
)

/* response 바디에 Hello from Snippetbox를 포함한
바이트 슬라이스를 작성하는 홈 핸들러 함수
*/
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	/* NewServeMux 함수를 초기화하고 home 함수와 "/" URL 패턴에 등록함 */
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	/* ListenAndServe 함수를 사용하여 웹 서버를 시작,
	에러가 발생하면 오류를 출력하고 종료한다
	*/
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
