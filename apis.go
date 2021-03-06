// Package daumapi implements API functions for Daum search service.
// The package basically wraps REST API from Daum search service.
// For the service details, refer to https://developers.kakao.com/docs/restapi/search
package daumapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type serviceFn func(string, string) string

const KAKAO_REST_API_URL = "https://dapi.kakao.com/v2/search"

var (
	buf bytes.Buffer
	logger = log.New(&buf, "INFO: ", log.Lshortfile)
)

// Perform a GET request with a custom header. The response
// is read into []byte and returned. If any error occurs, the
// the entire program will terminate with os.Exit(1).
func getResult(appkey string, url string) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", appkey)
	logger.Printf("Header: %v", req.Header)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		logger.Fatal(err)
	}	
	logger.Printf("resp.Body: %s", body)
	return body
}

// Parse a byte string into given Response type which are declared in
// github.com/hyunchel/response.go file. Returns the same Response type
// which the function was given as an argument.
// Any error will result in termination of the program with os.Exit(1).
func decodeJSON(responseType interface{}, contents []byte) interface{} {
	err := json.Unmarshal(contents, &responseType)
	if err != nil {
		logger.Fatal(err)
	}
	return responseType
}

// Stringify a generic Response type. Returns a byte string
// Any error will result in termination of the program with os.Exit(1).
func encodeJSON(responseType interface{}) []byte {
	b, err := json.Marshal(responseType)
	if err != nil {
		logger.Fatal(err)
	}
	return b
}

// Create a complete URL with corresponding service name and provided keyword.
func composeURL(service string, keyword string) string {
	url := fmt.Sprintf("%s/%s?query=%s", KAKAO_REST_API_URL, service, keyword)
	logger.Printf("Composed URL: %s", url)
	return url
}

func Web(appkey string, keyword string) string {
	service := "web"
	logger.Printf("Running %v function.", service)
	url := composeURL(service, keyword)
	rawResp := getResult(appkey, url)
	decodedResp := decodeJSON(WebResponse{}, rawResp)
	encodedResp := encodeJSON(decodedResp)
	return fmt.Sprintf("%s", encodedResp)
}

func Vclip(appkey string, keyword string) string {
	service := "vclip"
	logger.Printf("Running %v function.", service)
	url := composeURL(service, keyword)
	rawResp := getResult(appkey, url)
	decodedResp := decodeJSON(VclipResponse{}, rawResp)
	encodedResp := encodeJSON(decodedResp)
	return fmt.Sprintf("%s", encodedResp)
}

func Image(appkey string, keyword string) string {
	service := "image"
	logger.Printf("Running %v function.", service)
	url := composeURL(service, keyword)
	rawResp := getResult(appkey, url)
	decodedResp := decodeJSON(ImageResponse{}, rawResp)
	encodedResp := encodeJSON(decodedResp)
	return fmt.Sprintf("%s", encodedResp)
}

func Blog(appkey string, keyword string) string {
	service := "blog"
	logger.Printf("Running %v function.", service)
	url := composeURL(service, keyword)
	rawResp := getResult(appkey, url)
	decodedResp := decodeJSON(BlogResponse{}, rawResp)
	encodedResp := encodeJSON(decodedResp)
	return fmt.Sprintf("%s", encodedResp)
}

func Tip(appkey string, keyword string) string {
	service := "tip"
	logger.Printf("Running %v function.", service)
	url := composeURL(service, keyword)
	rawResp := getResult(appkey, url)
	decodedResp := decodeJSON(TipResponse{}, rawResp)
	encodedResp := encodeJSON(decodedResp)
	return fmt.Sprintf("%s", encodedResp)
}

func Book(appkey string, keyword string) string {
	service := "book"
	logger.Printf("Running %v function.", service)
	url := composeURL(service, keyword)
	rawResp := getResult(appkey, url)
	decodedResp := decodeJSON(BookResponse{}, rawResp)
	encodedResp := encodeJSON(decodedResp)
	return fmt.Sprintf("%s", encodedResp)
}

func Cafe(appkey string, keyword string) string {
	service := "cafe"
	logger.Printf("Running %v function.", service)
	url := composeURL(service, keyword)
	rawResp := getResult(appkey, url)
	decodedResp := decodeJSON(CafeResponse{}, rawResp)
	encodedResp := encodeJSON(decodedResp)
	return fmt.Sprintf("%s", encodedResp)
}

// Run the given function with all the logs printed.
func PrintLog(fn serviceFn, appkey string, keyword string) string {
	logger.Println("Running PrintLog function")
	result := fn(appkey, keyword)
	fmt.Print(&buf)
	return result
}
