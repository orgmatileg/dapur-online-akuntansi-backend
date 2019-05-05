package http

import (
	"encoding/json"
	"github.com/orgmatileg/SOLO-YOLO-BACKEND/helper"
	"github.com/orgmatileg/SOLO-YOLO-BACKEND/module/auth"
	"github.com/orgmatileg/SOLO-YOLO-BACKEND/module/users/model"
	"golang.org/x/oauth2"

	"net/http"

	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
)

// AuthHandler struct
type AuthHandler struct {
	AUsecase auth.Usecase
}

func NewAuthHttpHandler(r *mux.Router, au auth.Usecase) {

	handler := AuthHandler{
		AUsecase: au,
	}

	r.HandleFunc("/login", handler.AuthLoginJWTHTTPHandler).Methods("POST")
	r.HandleFunc("/oauth/facebook/login", handler.Oauth2FacebookLoginHTTPHandler).Methods("GET")
	r.HandleFunc("/oauth/facebook/callback", handler.Oauth2FacebookCallbackHTTPHandler).Methods("GET")

}

// AuthLoginJWTHTTPHandler handler
func (a *AuthHandler) AuthLoginJWTHTTPHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var mu model.User

	err := decoder.Decode(&mu)

	res := helper.Response{}

	defer res.ServeJSON(w, r)

	if err != nil {
		res.Err = err
		return
	}

	payload, err := a.AUsecase.LoginJWT(&mu)

	if err != nil {
		res.Err = err
		return
	}

	res.Body.Payload = payload

}

func (a *AuthHandler) Oauth2FacebookLoginHTTPHandler(w http.ResponseWriter, r *http.Request) {

	res := helper.Response{}

	defer res.ServeJSON(w, r)

	oauth2ConfigFb, _ := a.AUsecase.Oauth2FacebookLogin()
	url := oauth2ConfigFb.AuthCodeURL("1234")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

func (a *AuthHandler) Oauth2FacebookCallbackHTTPHandler(w http.ResponseWriter, r *http.Request) {

	oauth2ConfigFb, _ := a.AUsecase.Oauth2FacebookLogin()

	state := r.FormValue("state")
	if state != "1234" {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", "1234", state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := oauth2ConfigFb.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	response, err := http.Get("https://graph.facebook.com/v3.2/me?fields=id,name,email&access_token=" + token.AccessToken)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "Content: %s\n", contents)

}
