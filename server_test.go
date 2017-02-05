package main

import (
    "testing"
    "bytes"
    "net/http"
    "net/http/httptest"
    "os"
    "strings"
)

func TestEncode(t *testing.T) {
    expected := "5yc1t"
    actual := convertToBase36(10000001)
    if (actual != expected){
        t.Error("TestEncode Failed")
    }
}

func TestDecode(t *testing.T){
    var expected int64 = 45000012
    actual := convertFromBase36("qsi8c")
    if (actual != expected){
        t.Error("TestDecode Failed")
    }
}

func TestEncodeDecode(t *testing.T){
    var expected int64 = 45000012
    actual := convertFromBase36(convertToBase36(expected))
    if ( actual != expected ){
        t.Error("TestEncodeDecode Failed")
    }
}

func TestShortenURL(t *testing.T){
    os.Remove("./urldb.sqlite");
    InitDB()
    var jsonStr = []byte(`{"url":"http://averylongurl"}`)
    req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonStr ))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(shortenUrl)
    handler.ServeHTTP(rr, req)
    if (rr.Code != 200) {
        t.Error("TestShortenUrl Failed")
    }
    expected := `{"short":"http://5yc1t"}`
    actual := strings.Trim(rr.Body.String(),"\n")
    if (actual != expected) {
        t.Errorf("TestShortenUrl Failed, Expected:%s: Got:%s:", expected, actual);
    }
}

func TestOriginalURL(t *testing.T){
    var jsonStr = []byte(`{"url":"http://5yc1t"}`)
    req, _ := http.NewRequest("GET", "/original", bytes.NewBuffer(jsonStr ))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(originalUrl)
    handler.ServeHTTP(rr, req)
    if (rr.Code != 200) {
        t.Error("TestShortenUrl Failed")
    }
    expected := `{"short":"http://averylongurl"}`
    actual := strings.Trim(rr.Body.String(),"\n")
    if (actual != expected) {
         t.Errorf("TestShortenUrl Failed, Expected:%s: Got:%s:", expected, actual);
    }
}

func LastFunction(t *testing.T){
    os.Remove("./urldb.sqlite");

}


            



