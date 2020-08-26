package http

import (
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

type cacheHandler struct {
    *Server
}

//cacheHandler实现go中的Handler接口，它有一个方法：ServeHttp
func (ch *cacheHandler)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    key := strings.Split(r.URL.EscapedPath(), "/")[2]
    if len(key) == 0 {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    requestMethod := r.Method
    if requestMethod == http.MethodGet {
        value, getValeuErr := ch.Get(key)
        if getValeuErr != nil {
            log.Println(getValeuErr)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        if len(value) == 0 {
            w.WriteHeader(http.StatusNotFound)
            return
        }
        w.Write(value)
        return
    }
    if requestMethod == http.MethodPut {
        value, readValueErr := ioutil.ReadAll(r.Body)
        if readValueErr != nil {
            log.Println(readValueErr)
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        if len(value) != 0 {
            setValueErr := ch.Set(key, value)
            if setValueErr != nil {
                log.Println(setValueErr)
                w.WriteHeader(http.StatusInternalServerError)
            }
        } else {
            w.WriteHeader(http.StatusBadRequest)
        }
        return
    }
    if requestMethod == http.MethodDelete {
        delErr := ch.Del(key)
        if delErr != nil {
            log.Println(delErr)
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }
    w.WriteHeader(http.StatusMethodNotAllowed)
}