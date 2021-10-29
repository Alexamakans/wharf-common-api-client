package apiclient

import (
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
func DoGetUnmarshal(resultPtr interface{}, c Client, baseURL string, path string, queryParams ...string) error {
	v := reflect.ValueOf(resultPtr)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("not ptr; is %T", resultPtr)
	}

	bodyBytes, err := DoGetBytes(c, baseURL, path, queryParams...)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, resultPtr)
	if err != nil {
		return fmt.Errorf("get request failed: %w", err)
	}

	return nil
}

// DoGetBytes is a convenience function that is equivalent to calling NewGet,
// executing the resulting request using http.DefaultClient.Do, checking the
// status code and then reading the response body, returning the resulting
// byte slice.
func DoGetBytes(c Client, baseURL string, path string, queryParams ...string) ([]byte, error) {
	req, err := NewGet(c, baseURL, path, queryParams...)
	if err != nil {
		return []byte{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("get request failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Warn().
			WithInt("code", resp.StatusCode).
			WithString("host", req.Host).
			WithString("path", req.URL.Path).
			WithString("query", req.URL.RawQuery).
			Message("Failed GET request")

		if problem.IsHTTPResponse(resp) {
			prob, err := problem.ParseHTTPResponse(resp)
			if err == nil {
				return []byte{}, prob
			}
		}

		return []byte{}, fmt.Errorf("get request failed: non-2xx HTTP status: %s", resp.Status)
	}

	log.Debug().
		WithInt("code", resp.StatusCode).
		WithString("host", req.Host).
		WithString("path", req.URL.Path).
		WithString("query", req.URL.RawQuery).
		Message("Successful GET request")

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed reading response body: %w", err)
	}

	return bodyBytes, nil
}

// NewGet creates a new HTTP GET request to the set provider URL's endpoint, specified with
// path, using Basic Auth.
//
// Query parameters are interleaved like ( "name_1", "value_1", "name_2", "value_2" ).
func NewGet(c Client, baseURL string, path string, queryParams ...string) (*http.Request, error) {
	// if err := validateClient(c); err != nil {
	// 	return nil, err
	// }

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

func RequireProjectIdentifierString(c *gin.Context) (ProjectIdentifier, bool) {
	piStr, ok := ginutil.RequireParamString(c, "projectIdentifier")
	if !ok {
		return ProjectIdentifier{}, false
	}
	pi, err := projectIdentifierFromPathEscapedString(piStr)
	if err != nil {
		ginutil.WriteProblemError(c, err, problem.Response{
			Title:  "Failed parsing project identifier",
			Detail: "Unable to parse project identifier from path parameters.",
			Type:   "/prob/api/parse-project-identifier",
			Status: http.StatusBadRequest,
		})
		return ProjectIdentifier{}, false
	}

	return pi, true
}

// func SetupClientFromContext(client Client, c *gin.Context) bool {
// 	remoteProviderURL, ok := ginutil.RequireQueryString(c, "remoteProviderUrl")
// 	if !ok {
// 		return false
// 	}
// 	_, token, _ := c.Request.BasicAuth()
// 	setupClient(client, c, remoteProviderURL, token)
// 	return true
// }

// func setupClient(client Client, c *gin.Context, baseURL, token string) {
// 	client.SetContext(c)
// 	client.SetBaseURL(baseURL)
// 	client.SetToken(token)
// }
