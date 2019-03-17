package utils

import (
	"net/url"
	"strconv"

	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
)

var (
	hashKey  = []byte("the-big-and-secret-fash-key-here")
	blockKey = []byte("lot-secret-of-characters-big-too")
	sc       = securecookie.New(hashKey, blockKey)
)

// GetUser is yes
func GetUser(ctx iris.Context) (int, string) {
	if Check(ctx) {
		id, _ := strconv.Atoi(ctx.GetCookie("id", iris.CookieDecode(sc.Decode)))
		nickname, _ := url.QueryUnescape(ctx.GetCookie("username", iris.CookieDecode(sc.Decode)))
		return id, nickname
	}

	return -1, ""
}

// Check is yes
func Check(ctx iris.Context) bool {
	id := ctx.GetCookie("id", iris.CookieDecode(sc.Decode))
	userid := ctx.GetCookie("userid", iris.CookieDecode(sc.Decode))
	username := ctx.GetCookie("username", iris.CookieDecode(sc.Decode))
	token := ctx.GetCookie("token", iris.CookieDecode(sc.Decode))
	if id == "" || userid == "" || username == "" && token == "" {
		return false
	}

	encode := Encryption(id, userid)
	if encode == token {
		return true
	}
	return false
}
