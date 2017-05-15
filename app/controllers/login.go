package controllers

import (
	"fmt"
	"github.com/antonholmquist/jason"
	"github.com/omar-ozgur/flock-api/app/models"
	"github.com/omar-ozgur/flock-api/utilities"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"encoding/json"

)

func LoginWithFacebook(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	tok, err := utilities.FbConfig.Exchange(oauth2.NoContext, code)
	fmt.Println(tok)

	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Get("https://graph.facebook.com/me?access_token=" + tok.AccessToken + "&fields=email,first_name,last_name,id,friends")

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)


	user, _ := jason.NewObjectFromBytes([]byte(body))

	first_name, _ := user.GetString("first_name")
	last_name, _ := user.GetString("last_name")
	email, _ := user.GetString("email")
	fb_id_string, _ := user.GetString("id")

	status, message, app_token := models.ProcessFBLogin(first_name, last_name, email, fb_id_string)

	if status != "success" {
		fmt.Println(message)
	}
		
	JSON, _ := json.Marshal(map[string]interface{}{
		"status":  status,
		"message": message,
		"fb_token": tok.AccessToken,
		"app_token": app_token,
	})

	w.Write(JSON)

}

type LoggedInPageAttr struct {
	Name string
	URL  string
}
