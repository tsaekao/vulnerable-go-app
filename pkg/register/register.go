package register

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
    "reflect"
)

type Person struct {
	UserName string
}

func NewUserRegister(w http.ResponseWriter, r *http.Request) {
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
			fmt.Println("db error : ",err)
		} else {
			fmt.Println(reflect.TypeOf(count))
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
				name := r.FormValue("name")
				mail := r.FormValue("mail")
				encodeMail := base64.StdEncoding.EncodeToString([]byte(mail))
				fmt.Println("register successful!!")
				cookieSID := &http.Cookie{
					Name:  "SessionID",
					Value: encodeMail,
				}
				cookieUserName := &http.Cookie{
					Name:  "UserName",
					Value: name,
				}
				http.SetCookie(w, cookieSID)
				http.SetCookie(w, cookieUserName)
                p := Person{UserName: name}
				t, _ := template.ParseFiles("./views/success_register.gtpl")
				t.Execute(w,p)
			}
			defer db.Close()
		}

	} else {
		http.NotFound(w, r)
	}

}
