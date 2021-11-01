package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iver-wharf/wharf-core/pkg/ginutil"
	"github.com/iver-wharf/wharf-core/pkg/logger"
	"github.com/iver-wharf/wharf-core/pkg/problem"
)

var log = logger.NewScoped("API-CLIENT")

// DoGetUnmarshal is a convenience function that is equivalent to calling DoGetBytes and then
// unmarshalling the resulting byte array into the passed pointer.
func DoGetUnmarshal(resultPtr interface{}, c Client, baseURL, path string, queryParams ...string) error {
	req, err := NewGet(c, baseURL, path, queryParams...)
	if err != nil {
		return err
	}

	return DoRespAsUnmarshalled(c, resultPtr, req, nil, queryParams...)
}

// DoGetBytes is a convenience function that is equivalent to calling NewGet,
// executing the resulting request using http.DefaultClient.Do, checking the
// status code and then reading the response body, returning the resulting
// byte slice.
func DoGetBytes(c Client, baseURL, path string, queryParams ...string) ([]byte, error) {
	req, err := NewGet(c, baseURL, path, queryParams...)
	if err != nil {
		return []byte{}, err
	}

	return DoRespAsBytes(c, req, nil, queryParams...)
}

// NewGet creates a new HTTP GET request to the set provider URL's endpoint, specified with
// path, using Basic Auth.
//
// Query parameters are interleaved like ( "name_1", "value_1", "name_2", "value_2" ).
func NewGet(c Client, baseURL, path string, queryParams ...string) (*http.Request, error) {
	queryParamsLen := len(queryParams)
	if queryParamsLen%2 != 0 {
		return nil, fmt.Errorf("invalid number of query parameter entries (%d, should be divisible by 2)", queryParamsLen)
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse base URL: %q", baseURL)
	}
	u.Path = path
	if queryParamsLen > 0 {
		q := url.Values{}
		for i := 0; i < queryParamsLen; i += 2 {
			q.Add(queryParams[i], queryParams[i+1])
		}
		u.RawQuery = strings.ReplaceAll(q.Encode(), ".", "%2E")
	}

	req, err := http.NewRequestWithContext(c.GetContext(), http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("creating get request failed: %w", err)
	}

	req.SetBasicAuth("", c.GetToken())

	return req, nil
}

func DoPostUnmarshal(resultPtr interface{}, c Client, baseURL, path string, bodyPtr interface{}, queryParams ...string) error {
	req, err := NewPost(c, baseURL, path, bodyPtr, queryParams...)
	if err != nil {
		return err
	}

	return DoRespAsUnmarshalled(c, resultPtr, req, bodyPtr, queryParams...)
}

func DoPostBytes(c Client, baseURL, path string, bodyPtr interface{}, queryParams ...string) ([]byte, error) {
	req, err := NewPost(c, baseURL, path, bodyPtr, queryParams...)
	if err != nil {
		return []byte{}, err
	}

	return DoRespAsBytes(c, req, bodyPtr, queryParams...)
}

// NewPost creates a new HTTP POST request to the set provider URL's endpoint, specified with
// path, using Basic Auth.
//
// Query parameters are interleaved like ( "name_1", "value_1", "name_2", "value_2" ).
func NewPost(c Client, baseURL string, path string, bodyPtr interface{}, queryParams ...string) (*http.Request, error) {
	postBody, err := json.Marshal(bodyPtr)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal object of type %T", bodyPtr)
	}

	bodyBytes := bytes.NewBuffer(postBody)

	queryParamsLen := len(queryParams)
	if queryParamsLen%2 != 0 {
		return nil, fmt.Errorf("invalid number of query parameter entries (%d, should be divisible by 2)", queryParamsLen)
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse base URL: %q", baseURL)
	}
	u.Path = path
	if queryParamsLen > 0 {
		q := url.Values{}
		for i := 0; i < queryParamsLen; i += 2 {
			q.Add(queryParams[i], queryParams[i+1])
		}
		u.RawQuery = strings.ReplaceAll(q.Encode(), ".", "%2E")
	}

	req, err := http.NewRequestWithContext(c.GetContext(), http.MethodPost, u.String(), bodyBytes)
	if err != nil {
		return nil, fmt.Errorf("creating post request failed: %w", err)
	}

	req.SetBasicAuth("", c.GetToken())

	return req, nil
}

func DoRespAsUnmarshalled(c Client, resultPtr interface{}, req *http.Request, bodyPtr interface{}, queryParams ...string) error {
	v := reflect.ValueOf(resultPtr)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("not ptr; is %T", resultPtr)
	}

	bodyBytes, err := DoRespAsBytes(c, req, nil, queryParams...)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, resultPtr)
	if err != nil {
		return fmt.Errorf("%s request failed: %w", req.Method, err)
	}

	return nil
}

func DoRespAsBytes(c Client, req *http.Request, bodyPtr interface{}, queryParams ...string) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("%s request failed: %w", req.Method, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Warn().
			WithInt("code", resp.StatusCode).
			WithString("host", req.Host).
			WithString("path", req.URL.Path).
			WithString("query", req.URL.RawQuery).
			Messagef("Failed %s request", req.Method)

		if problem.IsHTTPResponse(resp) {
			prob, err := problem.ParseHTTPResponse(resp)
			if err == nil {
				return []byte{}, prob
			}
		}

		return []byte{}, fmt.Errorf("%s request failed: non-2xx HTTP status: %s", req.Method, resp.Status)
	}

	log.Debug().
		WithInt("code", resp.StatusCode).
		WithString("host", req.Host).
		WithString("path", req.URL.Path).
		WithString("query", req.URL.RawQuery).
		Messagef("Successful %s request", req.Method)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed reading response body: %w", err)
	}

	return bodyBytes, nil
}

func RequireProject(c *gin.Context) (WharfProject, bool) {
	var project WharfProject
	err := c.ShouldBindJSON(&project)
	if err != nil {
		ginutil.WriteInvalidBindError(c, err, "Unable to bind WharfProject JSON.")
		return WharfProject{}, false
	}
	return project, true
}
