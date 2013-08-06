vk_wrapper
==========

go (golang) api wrapper for most popular Russian social network vk.com

#How to use

> go get github.com/bogdansolomykin/vk_wrapper

Import package

```Go

import "github.com/bogdansolomykin/vk_wrapper/vk"
```

next you need to initialize a struct with your application params

```Go
auth := vk.Auth{
    AppId: "11111", 
    Scope: "friends, offline", //more access permissions https://vk.com/dev/permissions
    RedirectUri: "yoursite.com/get_access_token",
    ResponseType: "token",
}

//and redirect user to the formed url
authUrl := auth.GetAuthUrl()
YourRedirectFunc(authUrl)
```

if user will confirm requested application access permissions, he will be redirected to the specified url with access_token in params

```Go
currentUrl := getCurrentUrl() // for example "yoursite.com/get_access_token#access_token=3304fdb7c3b69ace6b055c6cba34e5e2f0229f7ac2ee4ef46dc9f0b241143bac993e6ced9a3fbc111111&expires_in=0&user_id=1"
accessToken, userId, expiresIn := vk.ParseResponseUrl(currentUrl)

api := vk.Api{
    AccessToken: accessToken,
    UserId: userId,
    ExpiresIn: expiresIn,
}
```

you can call "Request" method with api method name as first argument
and map with string keys and values as second argument if api method requires a params

```Go
m := make(map[string]string)
m["uid"] = userId

//you will get string in json format that can be parsed with any json lib
stringResponse := api.Request("getProfiles", m) //{"response":[{"uid":1,"first_name":"Pavel","last_name":"Durov"}]}
```

you can find all api methods on https://vk.com/dev/methods
