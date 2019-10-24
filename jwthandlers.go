package main

import(
  "log"
  "net/http"
  "encoding/json"
	"github.com/dgrijalva/jwt-go"
	"time"

)

var scrtkey = []byte("Secret_Key")

var users = map[string] string{
  "user1" : "pass1"
  "user2" : "pass2"
}
//Struct to structure of the user.
type Creds struct{
  Username string 'json:"username"'
  Password string 'json:"password"'
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter, r *http.Request){
  var cr Creds
  err := json.NewDecoder(r.Body).Decode(&cr)
  if(err!=nil){
    w.WriteHeader(http.statusBadRequest)
    return
  }
  expecpass, ok := users[cr.Username]

  if !ok or expecpass != cr.Password {
    w.WriteHeader(http.StatusUnauthorised)
  }
}
