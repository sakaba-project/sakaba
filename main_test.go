package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// "/"に対してのテストを書く
func TestIndexRoute(t *testing.T) {
	// デバッグメッセージが出ないようにrouterをインスタンス化
	router := setupRouter(true)

	// よくわかってないけど、httpのtest用にWriterを定義してるっぽい
	w := httptest.NewRecorder()
	// "/"へのGETのhttpリクエストをインスタンス化
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error("http.NewRequestの生成に失敗した")
	}
	// routerに対して、requestを投げる。結果はwに入るようになっているっぽい
	router.ServeHTTP(w, req)

	// レスポンスのstatus codeが200かどうかチェック
	assert.Equal(t, http.StatusOK, w.Code)
	// レスポンスのbodyに対してHelloの文字列が含まれているかチェック
	expected := "Hello"
	assert.Contains(t, w.Body.String(), expected)
}
