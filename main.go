package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func isZeroString(formstrs []string)bool{
		formstr := formstrs[0]
		fmt.Println("len: ",len(formstrs))
		if len(formstr)==0{
				return false
		}
		return true
}

func outErrorPage(w http.ResponseWriter){
		t, _ := template.ParseFiles("./views/error.gtpl")
		t.Execute(w,nil)
}

func sayYourName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("r.Form", r.Form)
	fmt.Println("r.Form[name]", r.Form["name"])
	var Name string
	for k, v := range r.Form {
		fmt.Println("key:", k)
		Name = strings.Join(v, ",")
	}
	fmt.Println(Name)
	fmt.Fprintf(w, Name)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method ", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		if ( isZeroString(r.Form["username"]) && isZeroString(r.Form["passwd"]) ){
				fmt.Println("passwd", r.Form["passwd"])
				fmt.Println("username", r.Form["username"])
				t, _ := template.ParseFiles("./views/logined.gtpl")
				t.Execute(w,nil)
		}else{
				fmt.Println("username or passwd are empty")
				outErrorPage(w)
		}
	}
}

func main() {
	http.HandleFunc("/", sayYourName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}