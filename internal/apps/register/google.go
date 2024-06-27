package controllers

import (
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "github.com/Siddheshk02/go-oauth2/pkce"
    "github.com/Siddheshk02/go-oauth2/config"
    "golang.org/x/oauth2"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
    codeVerifier, codeChallenge, err := pkce.GenerateCodeVerifierAndChallenge()
    if err != nil {
        http.Error(w, "Failed to generate PKCE code challenge", http.StatusInternalServerError)
        return
    }

    url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate", oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("code_challenge", codeChallenge), oauth2.SetAuthURLParam("code_challenge_method", "S256"))

    http.SetCookie(w, &http.Cookie{
        Name:  "oauth2_code_verifier",
        Value: codeVerifier,
        Path:  "/",
    })

    http.Redirect(w, r, url, http.StatusSeeOther)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
    state := r.URL.Query().Get("state")
    if state != "randomstate" {
        http.Error(w, "States don't match", http.StatusBadRequest)
        return
    }

    code := r.URL.Query().Get("code")

    cookie, err := r.Cookie("oauth2_code_verifier")
    if err != nil {
        http.Error(w, "Missing code verifier", http.StatusBadRequest)
        return
    }
    codeVerifier := cookie.Value

    googleCon := config.AppConfig.GoogleLoginConfig

    token, err := googleCon.Exchange(context.Background(), code, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
    if err != nil {
        http.Error(w, "Code-Token Exchange Failed", http.StatusInternalServerError)
        return
    }

    resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
    if err != nil {
        http.Error(w, "User Data Fetch Failed", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    userData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, "JSON Parsing Failed", http.StatusInternalServerError)
        return
    }

    var result map[string]interface{}
    if err := json.Unmarshal(userData, &result); err != nil {
        http.Error(w, "JSON Unmarshal Failed", http.StatusInternalServerError)
        return
    }

    email := result["email"].(string)
    fmt.Fprintf(w, "User email: %s", email)
}

