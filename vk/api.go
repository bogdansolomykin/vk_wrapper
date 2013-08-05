package vk

import (
    "net/url"
    "net/http"
    "io/ioutil"
    "strings"
)

const API_METHOD_URL = "https://api.vk.com/method/"

type Api struct {
    AccessToken string
    UserId      string
    ExpiresIn   string
}

func ParseResponseUrl(responseUrl string) (string, string, string) {
    u, err := url.Parse(strings.Replace(responseUrl, "#", "?", 1))
    if err != nil {
        panic(err)
    }
    
    q := u.Query()    
    return q.Get("access_token"), q.Get("user_id"), q.Get("expires_in")        
}

func (vk Api) Request(methodName string, params map[string]string) string {
    u, err := url.Parse(API_METHOD_URL+methodName)
    if err != nil {
        panic(err)
    }

    q := u.Query()
    for k, v := range params {
        q.Set(k, v)
    }
    q.Set("access_token", vk.AccessToken)
    u.RawQuery = q.Encode()
    
    resp, err := http.Get(u.String())
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    return string(content)
}