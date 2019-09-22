package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
    login "github.com/Snow-HardWolf/Vulnerability-goapp/pkg/login"

	"./pkg/register"
)

type Person struct {
	UserName string
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

func showUserTopPage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println("Cookie :", err)
		return
	}

	if r.Method == "GET" {
		fmt.Println(cookie)
		userName, err := r.Cookie("UserName")
		if err != nil {
			fmt.Println(err)
		}

		userMail, err := r.Cookie("SessionID")
		if err != nil {
			fmt.Println(err)
		}
		decodeMail, err := base64.StdEncoding.DecodeString(userMail.Value)
		if err != nil {
			fmt.Println(err)
		}
		mail := string(decodeMail)
		fmt.Println(mail)

		db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		var userID int
		if err := db.QueryRow("select id from user where mail=?", mail).Scan(&userID); err != nil {
			log.Fatal(err)
		}

		log.Println(userID)
		uid := strconv.Itoa(userID)
		cookieUserID := &http.Cookie{
			Name:  "UserID",
			Value: uid,
		}

		http.SetCookie(w, cookieUserID)
		p := Person{UserName: userName.Value}
		t, _ := template.ParseFiles("./views/top.gtpl")
		t.Execute(w, p)

	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", sayYourName)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/new", register.NewUserRegister)
	http.HandleFunc("/top", showUserTopPage)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
