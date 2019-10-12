package post

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"
	"text/template"

	"../cookie"
)

type User struct {
	UserId   int
	UserName string
	Mail     string
	Age      int
	Image    string
	Address  string
	Animal   string
	Word     string
}

type Posts struct {
    UserPosts []string
    Created_at []string
}

func StorePost(uid int, postText string) {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into vulnapp.posts(uid,post) values (?,?)", uid, postText)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
}

func GetPost(uid int) ([]string, []string, error) {
    db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		fmt.Printf("%+v\n", err)
		return []string{},[]string{},err
	}
	defer db.Close()

    res, err := db.Query("select post,created_at from posts where uid=? order by created_at desc",uid)
    if err != nil {
        fmt.Printf("%+v\n",err)
        return []string{},[]string{},err
    }
    defer res.Close()

    var posts []string
    var create_ats []string

    var post string
    var create_at string
    for res.Next() {
        if err := res.Scan(&post,&create_at); err != nil {
            fmt.Println(err)
        }
        posts = append(posts,post)
        create_ats = append(create_ats,create_at)
    }


    return posts, create_ats, nil
}


func ShowAddPostPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if cookie.CheckSessionID(r) {
          _, _, uid, err := cookie.GetCookieValue(r)
			if err != nil {
				fmt.Printf("%+v\n", err)
				http.NotFound(w, nil)
				return
			}

            posts, creates, err := GetPost(uid)
            if err != nil {
                fmt.Printf("%+v\n",err)
                return
            }

            p := Posts{UserPosts: posts, Created_at: creates}

            fmt.Println(p)

			t, _ := template.ParseFiles("./views/postadd.gtpl")
			t.Execute(w, p)

		}

	} else if r.Method == "POST" {
		if cookie.CheckSessionID(r) {
			_, _, uid, err := cookie.GetCookieValue(r)
			if err != nil {
				fmt.Printf("%+v\n", err)
				http.NotFound(w, nil)
				return
			}

			fmt.Println(uid, r.FormValue("post"))

			postText := r.FormValue("post")

			fmt.Println(reflect.TypeOf(postText))

			StorePost(uid, postText)

			http.Redirect(w, r, "/post", 301)

		}
	} else {
		http.NotFound(w, nil)
	}
}