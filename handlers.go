package main

import (
   "encoding/json"
   "net/http"
   "io/ioutil"
   "io"
   "fmt"
)
type LongUrl struct{
    Url string  `json:"url"`
}

type ShortUrl struct{
    Short string `json:"short"`
}

type OriginalUrl struct{
    Original string `json:"original"`
}

func shortenUrl( w http.ResponseWriter, r *http.Request ) {
    var longUrl LongUrl
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &longUrl); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    db := InitDB()
    defer db.Close()
    id :=  insertLongUrl(db, longUrl.Url)
    fmt.Printf("%d",id)
    short := convertToBase36(id)
    short = "http://" + short
    s := ShortUrl{Short:short}

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(s); err != nil {
        panic(err)
    }
    return
}

func originalUrl( w http.ResponseWriter, r *http.Request) {
 //   vars := mux.Vars(r)
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    var shortUrl ShortUrl
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &shortUrl); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    db := InitDB()
    defer db.Close()
    short := shortUrl.Short[7:len(shortUrl.Short)]
    id := convertFromBase36(short)
    fmt.Printf("ID IS %d SHORT IS : %s", id, short)
    long := getLongUrl(db, id)
    orig := OriginalUrl{Original:long}
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(orig); err != nil {
        panic(err)
    }
}
