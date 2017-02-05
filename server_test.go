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
    actual := ConvertToBase36(10000001)
    if (actual != expected){
        t.Error("TestEncode Failed")
    }
}

func TestDecode(t *testing.T){
    var expected int64 = 45000012
    actual := ConvertFromBase36("qsi8c")
    if (actual != expected){
        t.Error("TestDecode Failed")
    }
}

func TestEncodeDecode(t *testing.T){
    var expected int64 = 45000012
    actual := ConvertFromBase36(ConvertToBase36(expected))
    if ( actual != expected ){
        t.Error("TestEncodeDecode Failed")
    }
}

func TestShortenURLHandler(t *testing.T){
    os.Remove("./urldb.sqlite");
    InitDB()
    var jsonStr = []byte(`{"url":"http://averylongurl"}`)
    req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonStr ))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(ShortenUrlHandler)
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

func TestOriginalURLHandler(t *testing.T){
    var jsonStr = []byte(`{"short":"http://5yc1t"}`)
    req, _ := http.NewRequest("POST", "/original", bytes.NewBuffer(jsonStr ))
    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(OriginalUrlHandler)
    handler.ServeHTTP(rr, req)
    if (rr.Code != 200) {
        t.Error("TestShortenUrl Failed")
    }
    expected := `{"original":"http://averylongurl"}`
    actual := strings.Trim(rr.Body.String(),"\n")
    if (actual != expected) {
         t.Errorf("TestShortenUrl Failed, Expected:%s: Got:%s:", expected, actual);
    }
}

func LastFunction(t *testing.T){
    os.Remove("./urldb.sqlite");
}
