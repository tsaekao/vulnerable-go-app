package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"../cookie"
)

type User struct {
	UserName string
	Mail     string
	Age      int
}

func ShowUserProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if cookie.CheckSessionID(r) {
			_, userName, uid, err := cookie.GetCookieValue(r)
			if err != nil {
				fmt.Println(err)
			}
			mail, age, err := GetUserInfos(uid)
			if err != nil {
				fmt.Println(err)
			}
			u := User{UserName: userName, Mail: mail, Age: age}
			t, _ := template.ParseFiles("./views/users.gtpl")
			t.Execute(w, u)
		} else {
			http.NotFound(w, nil)
		}

	} else {
		t, _ := template.ParseFiles("./views/error.gtpl")
		t.Execute(w, nil)
	}
}

func GetUserInfos(uid int) (userMail string, userAge int, err error) {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		fmt.Println(err)
		return "", 0, err
	}
	defer db.Close()

	var mail string
	var age int

	res, err := db.Query("select mail,age from vulnapp.user where id=?", uid)
	if err != nil {
		fmt.Println(err)
		return "", 0, err
	}

	for res.Next() {
		if err := res.Scan(&mail, &age); err != nil {
			fmt.Println(err)
			return "", 0, err
		}
	}

	return mail, age, nil
}
