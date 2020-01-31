package helloworld

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//テスト対象のハンドラ
var sampleHandler = http.HandlerFunc(HelloHTTP)

func TestCanNotDecodeE2E(t *testing.T) {
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


func TestCanNotDecode(t *testing.T) {

	req := httptest.NewRequest("POST", "/", nil)
	rec := httptest.NewRecorder()

	HelloHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "jsonがデコードできなかったよHello, World!", rec.Body.String())
}


func TestNoName(t *testing.T) {

	body := strings.NewReader(`{"name":""}`)

	req := httptest.NewRequest("POST", "/", body)
	rec := httptest.NewRecorder()

	HelloHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "名前が空文字だよHello, World!", rec.Body.String())
}


func TestExistName(t *testing.T) {

	body := strings.NewReader(`{"name":"えびえび"}`)

	req := httptest.NewRequest("POST", "/", body)
	rec := httptest.NewRecorder()

	HelloHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, えびえび!", rec.Body.String())
}


////json用
//type Person struct {
//	Name string `json:"name"`
//}
//
//func TestNotName(t *testing.T) {
//	ts := httptest.NewServer(sampleHandler)
//	defer ts.Close()
//
//	//リクエスト作成
//	request, err := http.NewRequest("POST", ts.URL, nil)
//	if err != nil{
//		log.Fatal(err)
//	}
//
//	dummy := []byte(`{"name":""}`)
//
//	// JSONデコード
//	var person []Person
//	if err := json.Unmarshal(dummy, &person); err != nil {
//		log.Fatal(err)
//	}
//
//	//クエリパラメータ
//	params := request.URL.Query()
//	request.URL.RawQuery = params.Encode()
//
//	request.
//
//	//作ったクエリを確認
//	fmt.Println(request.URL.String())
//
//	timeout := time.Duration(5 * time.Second)
//	client := &http.Client{
//		Timeout: timeout,
//	}
//
//	//レスポンス取得
//	response, err := client.Do(request)
//	if err != nil{
//		log.Fatal(err)
//	}
//	defer response.Body.Close()
//
//	//レスポンス読み込み
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//
//	//assert
//	if "名前が空文字だよHello, World!" != string(string(body)) {
//		t.Fatalf("Data Error. %v", string(body))
//	}
//
//}