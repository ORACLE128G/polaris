package controller

import (
	"context"
	"github.com/olivere/elastic"
	"net/http"
	"polaris/crawler/engine"
	"polaris/crawler/frontend/model"
	"polaris/crawler/frontend/view"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

// Handle localhost:80/search?q=男 已购房 已购车&from=10
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
func (h SearchResultHandler) getSearchResult(s string, i int) (model.SearchResult, error) {
	var result model.SearchResult
	// fill in query string
	result.Query = s
	resp, err := h.client.Search("polaris").
		Query(elastic.NewQueryStringQuery(
		rewriteQueryString(s))).From(i).Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = i
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	// Prev page
	result.PrevFrom = result.Start - len(result.Items)
	// Next page
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}

var client, err = elastic.NewClient(elastic.SetSniff(false))

func CreateSearchResultHandler(template string) SearchResultHandler {
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

var qRegExp = regexp.MustCompile(`([A-Z][a-z]*):`)

func rewriteQueryString(q string) string {
	return qRegExp.ReplaceAllString(q, "PayLoad.$1:")
}
