package integration

//import (
//	"bytes"
//	"io"
//	"net/http"
//	"testing"
//)
//
//func Init() {
//
//}
//
//func TestRouter(t *testing.T) {
//	client := http.Client{}
//
//	type test struct {
//		method   string
//		body     []byte
//		endpoint string
//		response string
//	}
//
//	tests := []test{
//		{
//			method:   "GET",
//			endpoint: "http://localhost:8080/",
//			body:     nil,
//			response: `{"message":"hello world"}`,
//		},
//	}
//
//	for i := 0; i < len(tests); i++ {
//		req, err := http.NewRequest(tests[i].method, tests[i].endpoint, bytes.NewBuffer(tests[i].body))
//		if err != nil {
//			t.Fatal(err)
//		}
//		res, err := client.Do(req)
//		if err != nil {
//			t.Fatal(err)
//		}
//		defer res.Body.Close()
//		body, err := io.ReadAll(res.Body)
//		if err != nil {
//			t.Fatal(err)
//		}
//		if string(body) != tests[i].response {
//			t.Fatalf("expect '%s', got '%s'", tests[i].response, string(body))
//		}
//
//	}
//
//}
