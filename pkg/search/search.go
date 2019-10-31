package search

import (
	"fmt"
	"net/http"
	"os/exec"
)

type SearchResults struct {
	PostResults []string
	Created_at  []string
	UserImages  []string
}

func SearchPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		searchWord := r.FormValue("post")
		fmt.Println("value : ", searchWord)

		/*unionStr := "mysql -h mysql -u root -prootwolf -e 'select post,\",\",created_at,\",\",userimage,\",\" from vulnapp.posts inner join vulnapp.userdetails on vulnapp.posts.uid = vulnapp.userdetails.uid where vulnapp.posts.post like \"%" + searchWord + "%\";'"
				fmt.Println(unionStr)
				postRes, err := exec.Command("sh", "-c", unionStr).Output()
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(string(postRes))

				var posts, createdTimes, userImages []string

				splitedRes := strings.Split(string(postRes), ",")

				for i, v := range splitedRes {
					if i > 2 {
						judge := i % 3
						switch judge {
						case 0:
							trimed := strings.TrimSpace(v)
							posts = append(posts, trimed)
						case 1:
							trimed := strings.TrimSpace(v)
							createdTimes = append(createdTimes, trimed)
						case 2:
							trimed := strings.TrimSpace(v)
							userImages = append(userImages, trimed)
						default:
							fmt.Println("error in split")
						}
						trimed := strings.TrimSpace(v)
						fmt.Println(i, trimed)
					}
				}

				fmt.Println("posts : ", posts)
				fmt.Println("createdTimes : ", createdTimes)
				fmt.Println("userImages : ", userImages)

		        p := SearchResults{PostResults: posts, Created_at: createdTimes, UserImages: userImages}

		        fmt.Println(p)*/

		testStr := "mysql -h mysql -u root -prootwolf -e 'select post,created_at from vulnapp.posts where post like \"%" + searchWord + "%\"'"

        fmt.Println(testStr)

		testres, err := exec.Command("sh", "-c", testStr).Output()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(testres))

		//t, _ := template.ParseFiles("./views/search/searchpost.gtpl")
		//t.Execute(w, p)

	} else {
		http.NotFound(w, nil)
	}
}