package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/valyala/fasthttp"
	"net/http"
	"net/url"
	"os"
)

func Crawl(q string) ([]MusMeta, error) {
	link := "https://ruv.hotmo.org/search?q=" + url.QueryEscape(q)
	var client http.Client
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("UserAgent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	html, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	var out []MusMeta
	html.Find(`#pjax-container > div > div > ul.tracks__list > li`).Each(func(i int, selection *goquery.Selection) {
		attr, ok := selection.Attr("data-musmeta")
		if !ok {
			err = errors.New("could not find tracks")
			return
		}
		var musmeta MusMeta
		err := json.Unmarshal([]byte(attr), &musmeta)
		if err != nil {
			return
		}
		if musmeta.Img == "" {
			musmeta.Img = "https://data.kamakepar.ru/vinyl.png"
		}
		out = append(out, musmeta)
	})
	return out, err
}


type MusMeta struct {
	Artist string `json:"artist"`
	Title string `json:"title"`
	Url string `json:"url"`
	Img string `json:"img"`
}

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")
		switch string(ctx.Path()) {
		case "/search":
			q := string(ctx.QueryArgs().Peek("q"))
			if q == "" {
				fmt.Fprint(ctx, `enter "q" parameter"`)
				ctx.SetStatusCode(400)
				return
			}
			data, err := Crawl(q)
			if err != nil {
				fmt.Fprint(ctx, err)
				return
			}
			outJson, _ := json.Marshal(data)
			_, err = fmt.Fprintln(ctx, string(outJson))
			if err != nil {
				fmt.Fprint(ctx, err)
				fmt.Println(err)
				fmt.Fprint(os.Stdout, err)
			}
		default:
			fmt.Fprint(ctx, "go to /search?q=SomeMusic")
			return
		}
	}
	fmt.Println("Server started.")
	fmt.Fprintln(os.Stdout, "Server started.")
	if err := fasthttp.ListenAndServe(":1024", requestHandler); err != nil {
		panic(err)
	}

}

