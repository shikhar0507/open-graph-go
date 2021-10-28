package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocolly/colly/v2"
	"github.com/shikhar0507/requestJSON"
)

type Request struct {
	Url string `json:url`
}

type ErrorResponse struct {
	Message    string
	StatusCode int
}

func main() {
	parseRequest := colly.NewCollector(colly.CheckHead(), colly.ParseHTTPErrorResponse())

	http.HandleFunc("/linkPreview", func(w http.ResponseWriter, r *http.Request) {
		websiteRequest(w, r, parseRequest)
	})
	http.ListenAndServe(":5000", nil)

}

func websiteRequest(w http.ResponseWriter, r *http.Request, parseRequest *colly.Collector) {

	allowed := false
	switch r.Method {
	case "options":
		return
	case "POST":
		allowed = true
	default:
		errResp := ErrorResponse{Message: "Method not allowed", StatusCode: http.StatusMethodNotAllowed}
		sendResponse(w, r, errResp)
	}
	if !allowed {
		return
	}
	fmt.Println(r.Method)
	var request Request
	result := requestJSON.Decode(w, r, &request)

	if result.Status != 200 {
		errorResponse := ErrorResponse{Message: result.Message, StatusCode: result.Status}
		sendResponse(w, r, errorResponse)
		return
	}
	parseRequest.OnHTML("meta", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("property"), e.Attr("content"))
	})

	parseRequest.OnRequest(func(collyReq *colly.Request) {

		fmt.Println("visiting", collyReq.URL.String())
	})

	parseRequest.OnResponse(func(collyResp *colly.Response) {
		fmt.Println("parsing response", collyResp.StatusCode)
	})
	parseRequest.OnScraped(func(r *colly.Response) {
		fmt.Println("parsed", r.StatusCode)
	})
	fmt.Println("parsing", r.URL.String())
	parseRequest.Visit(request.Url)
}

func sendResponse(w http.ResponseWriter, r *http.Request, body ErrorResponse) {
	data, err := json.Marshal(&body)
	w.WriteHeader(body.StatusCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w)
		return
	}
	fmt.Fprint(w, string(data))
}
