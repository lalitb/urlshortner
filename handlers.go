package main

import (

   "encoding/json"
   "fmt"
   "net/http"
   "github.com/gorilla/mux"
)

func shortenURL( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars(r)
    longURL := vars["original"]
    shortenURL := 

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
    return
}

func originalURL( w http.ResponseWriter, r *http.Request) {

}



                    
