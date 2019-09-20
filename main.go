package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	//"reflect"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func isZeroString(formstr string) bool {
	//fmt.Println("len: ", len(formstr))
	if len(formstr) == 0 {
		return false
	}
	return true
}

func outErrorPage(w http.ResponseWriter) {
	t, _ := template.ParseFiles("./views/error.gtpl")
	t.Execute(w, nil)
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
		fmt.Println(Name)
	}
	//fmt.Println(Name)
	//fmt.Fprintf(w, Name)
	http.Redirect(w, r, "http://localhost:9090/login", 301)
}

func passwdCheck(username string, passwd string) {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Query("select id from user where passwd=" + passwd)
	fmt.Println(res)
	defer db.Close()
}

func searchId(mail string) int {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "select id from user where mail=?"
	res, err := db.Query(sql, mail)
	if err != nil {
		log.Fatal(err)
	}

	var id int

	for res.Next() {
		err := res.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("ID :", id)
	}

	fmt.Println(id)
	return id
}

func checkPasswd(id int, passwd string) string {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	sql := "select name from user where id=? and passwd=?"
	res, err := db.Query(sql, id, passwd)
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		err := res.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name)
	}

	return name
}

type Person struct {
	UserName string
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method ", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {

		r.ParseForm()
		if isZeroString(r.FormValue("mail")) && isZeroString(r.FormValue("passwd")) {
			fmt.Println("passwd", r.Form["passwd"])
			fmt.Println("mail", r.Form["mail"])

			mail := r.FormValue("mail")
			id := searchId(mail)

			var viewsFile string
			if id != 0 {
				passwd := r.FormValue("passwd")
				name := checkPasswd(id, passwd)

				if name != "" {
					fmt.Println(name)
					viewsFile = "./views/logined.gtpl"
					t, _ := template.ParseFiles(viewsFile)
					p := Person{UserName: name}
					t.Execute(w, p)
				} else {
					fmt.Println(name)
					viewsFile = "./views/error.gtpl"
					t, _ := template.ParseFiles(viewsFile)
					t.Execute(w, nil)
				}
			} else {
				viewsFile = "./views/error.gtpl"
				t, _ := template.ParseFiles(viewsFile)
				t.Execute(w, nil)
			}

		} else {
			fmt.Println("username or passwd are empty")
			outErrorPage(w)
		}
	} else {
		http.NotFound(w, r)
	}
}

func newUserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/new.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.FormValue("mail"))
		fmt.Println(r.FormValue("name"))
		fmt.Println(r.FormValue("age"))
		fmt.Println(r.FormValue("passwd"))

		db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
		if err != nil {
			log.Fatal(err)
		}

		var count int
		if err = db.QueryRow("select count(mail) from user where mail=?", r.FormValue("mail")).Scan(&count); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(count)
		}

		if count != 0 {
			t, _ := template.ParseFiles("./views/register_error.gtpl")
			t.Execute(w, nil)
		} else {

			age, err := strconv.Atoi(r.FormValue("age"))
			if err != nil {
				log.Fatal(err)
			}

			_, err = db.Exec("insert into user (name,mail,age,passwd) value(?,?,?,?)", r.FormValue("name"), r.FormValue("mail"), age, r.FormValue("passwd"))
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("register successful!!")
			}
			defer db.Close()
		}

	} else {
		http.NotFound(w, r)
	}

}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", sayYourName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/new", newUserRegister)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
