package vk

import "net/url"

const AUTH_HOST = "https://oauth.vk.com/authorize"

type Auth struct {
    AppId        string
    Scope        string
    RedirectUri  string
    ResponseType string 
}

func (a Auth) GetAuthUrl() string {
    u, err := url.Parse(AUTH_HOST)
    if err != nil {
        panic(err)
    }

    q := u.Query()
    q.Set("client_id", a.AppId)
    q.Set("scope", a.Scope)
    q.Set("redirect_uri", a.RedirectUri)
    q.Set("response_type", a.ResponseType)
    u.RawQuery = q.Encode()

    return u.String()
}

