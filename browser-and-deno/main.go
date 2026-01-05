package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-jsonnet"
	"github.com/google/go-jsonnet/ast"
)

func httpgetFunc() *jsonnet.NativeFunction {
	return &jsonnet.NativeFunction{
		Name:   "httpget",
		Params: ast.Identifiers{"url"},
		Func: func(params []any) (res any, err error) {
			if len(params) != 1 {
				return nil, errors.New("httpget accepts exactly 1 argument")
			}
			url, ok := params[0].(string)
			if !ok {
				return nil, errors.New("wrong parameters for 'httpget': expecting url string")
			}
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			return string(body), err
		},
	}
}

func main() {
	vm := jsonnet.MakeVM()
	vm.NativeFunction(httpgetFunc())
	const code = `
local httpget = std.native("httpget");
local res = httpget("https://httpbin.org/anything");
{resp: std.parseJson(res)}
`
	output, err := vm.EvaluateAnonymousSnippet("dummy.jsonnet", code)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
