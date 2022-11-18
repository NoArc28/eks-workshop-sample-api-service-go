package main

import (
   "fmt"
   "net/http"
)

func  main(){

	http.HandleFunc("/html", uploadsHandler)
/* 업로드 경로 에 대한 핸들러를 만들어줌 */ 


	http.Handle("/",http.FileServer(http.Dir("public")))
/* 인덱스 경로 - 파일서버 (퍼블릭 경로) */
 
	http.LisetAndServe(":3000",nil)

/* 가장 고전적인 파일 웹서버 - public 폴더 안에 있는 파일들을 엑세스 할수 있도록 열어주는것 */
}
