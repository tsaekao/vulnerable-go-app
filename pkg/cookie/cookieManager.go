package cookie

import (
	"errors"
	"fmt"
	"net/http"
)

func CheckCookieOnlyLogin(r *http.Request) (userNameCookie string, sessionIDCookie string, err error) {
	userName, err := r.Cookie("UserName")
	if err != nil {
		fmt.Println(err)
	}

	sessionID, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userName, sessionID)

	if userName.Value == "" && sessionID.Value == "" {
		return "", "", errors.New("Cookie not exsit")
	} else {
		return "test", "test", nil
	}
}
