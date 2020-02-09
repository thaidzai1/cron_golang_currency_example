package httpreq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	cm "gido.vn/gic/cron/auto-edit-price/models"
)

// API ...
type API struct {
	Host   string
	URI    string
	Query  Query
	Method string
	Body   interface{}
	Token  string
}

// Query ...
type Query struct {
	Filter cm.Filter
	Paging cm.Paging
}

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (a *API) makeAPIURL() (api string) {
	queryString := a.makeQueryString()
	api = a.Host + a.URI
	if queryString == "" {
		return
	}
	api += "?" + queryString
	return
}

func (a *API) makeQueryString() string {
	var pagingQuery, filterQuery string

	if (cm.Filter{}) == a.Query.Filter && (cm.Paging{}) == a.Query.Paging {
		return ""
	}

	if (cm.Filter{}) != a.Query.Filter {
		filterValue := reflect.ValueOf(a.Query.Filter)
		filterFields := reflect.TypeOf(a.Query.Filter)
		lengthStruct := filterFields.NumField()

		for i := 0; i < lengthStruct; i++ {
			if filterValue.Field(i).String() != "" {
				query := "filters." + toSnakeCase(filterFields.Field(i).Name) + "=" + filterValue.Field(i).String()
				filterQuery += query
				if i != lengthStruct-1 {
					filterQuery += "&"
				}
			}
		}
	}

	if (cm.Paging{}) != a.Query.Paging {
		pagingValue := reflect.ValueOf(a.Query.Paging)
		pagingFields := reflect.TypeOf(a.Query.Paging)
		lengthStruct := pagingFields.NumField()

		for i := 0; i < lengthStruct; i++ {
			if pagingValue.Field(i).Int() != 0 || pagingFields.Field(i).Name == "Offset" {
				query := "paging." + toSnakeCase(pagingFields.Field(i).Name) + "=" + strconv.FormatInt(pagingValue.Field(i).Int(), 10)
				pagingQuery += query
				if i != lengthStruct-1 {
					pagingQuery += "&"
				}
			}
		}
	}

	if pagingQuery != "" {
		filterQuery += "&"
	}

	return filterQuery + pagingQuery
}

// RequestAPI ...
func (a *API) RequestAPI() ([]byte, error) {
	jsonBody, err := json.Marshal(a.Body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	url := a.makeAPIURL()
	fmt.Printf("Preparing request: %v - %v \n", a.Method, url)
	req, err := http.NewRequest(a.Method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	if a.Token != "" {
		req.Header.Add("Authorization", "Bearer "+a.Token)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	byteRes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return byteRes, nil
}
