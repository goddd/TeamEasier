package request

import (
	"fmt"
	"github.com/goddd/TeamEasier/config"
	"gopkg.in/resty.v0"
)

type Req struct {
	Url       string
	Call_type string
	Payload   string
	Key       string
	Base_url  string
	full_url  string
}

func Get(url string) (*resty.Response, error) {

	req := Req{
		Call_type: "GET", Url: url}

	return MakeRequest(req)
}

func Put(url string, payload string) (*resty.Response, error) {

	req := Req{
		Call_type: "PUT", Url: url}

	if payload != "" {

		req.Payload = payload
	}

	return MakeRequest(req)
}

func Post(url string, payload string) (*resty.Response, error) {

	req := Req{
		Call_type: "POST", Url: url, Payload: payload}

	return MakeRequest(req)
}

func Delete(url string) (*resty.Response, error) {

	req := Req{
		Call_type: "DELETE", Url: url}

	return MakeRequest(req)
}

func MakeRequest(req Req) (*resty.Response, error) {

	r := resty.R()

	if req.Base_url == "" {

		req.full_url = config.BaseUrl + req.Url
	} else {

		req.full_url = req.Base_url + req.Url
	}

	fmt.Println("Request obj", req)

	if req.Key == "" {

		req.Key = config.ApiKey
	}

	r.SetBasicAuth(req.Key, "xxx")

	var resp *resty.Response
	var err error

	if req.Call_type == "GET" {

		resp, err = r.Get(req.full_url)

	} else if req.Call_type == "POST" {

		resp, err = r.SetBody(req.Payload).Post(req.full_url)
	} else if req.Call_type == "PUT" {

		resp, err = r.SetBody(req.Payload).Put(req.full_url)
	} else if req.Call_type == "DELETE" {

		resp, err = r.Delete(req.full_url)
	}

	if err != nil {

		fmt.Println("Got error", err)
	}

	fmt.Println("Body", resp.Body())
	fmt.Println("Status Code", resp.StatusCode())

	return resp, err
}

func makeUrl(req Req) string {

	return config.BaseUrl + req.Url

}
