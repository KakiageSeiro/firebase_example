package helloworld

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//テスト対象のハンドラ
var sampleHandler = http.HandlerFunc(HelloHTTP)

func TestCanNotDecode(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	// リクエストの送信先はテストサーバのURLへ。
	r, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}

	//レスポンス読み込み
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}

	//assert
	if "jsonがデコードできなかったよHello, World!" != string(data) {
		t.Fatalf("Data Error. %v", string(data))
	}
}


func TestNotName(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	//リクエスト作成
	request, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil{
		log.Fatal(err)
	}

	//クエリパラメータ
	params := request.URL.Query()
	params.Add("name","")
	request.URL.RawQuery = params.Encode()

	//作ったクエリを確認
	fmt.Println(request.URL.String())

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil{
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
	log.Fatal(err)
	}


	//assert
	if "名前が空文字だよHello, World!" != string(string(body)) {
		t.Fatalf("Data Error. %v", string(body))
	}

}