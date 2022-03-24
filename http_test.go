package httpclient

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Slideshow struct {
		Author string `json:"author"`
		Date   string `json:"date"`
		Slides []struct {
			Title string   `json:"title"`
			Type  string   `json:"type"`
			Items []string `json:"items,omitempty"`
		} `json:"slides"`
		Title string `json:"title"`
	} `json:"slideshow"`
}

type testMap map[string]any

const JsonResponse = `{
	"slideshow": {
		"author": "Yours Truly",
		"date": "date of publication",
		"slides": [
			{
			"title": "Wake up to WonderWidgets!",
			"type": "all"
			},
			{
			"items": [
			"Why <em>WonderWidgets</em> are great",
			"Who <em>buys</em> WonderWidgets"
			],
			"title": "Overview",
			"type": "all"
			}
		],
		"title": "Sample Slide Show"
	}
}`

func TestClient_Get(t *testing.T) {
	url := "https://api.test/json"

	var data testStruct
	err := json.Unmarshal([]byte(JsonResponse), &data)
	require.NoError(t, err)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, data)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		})

	{
		c := Client[testStruct]{}
		got, err := c.Get(url)
		assert.NoError(t, err)
		assert.IsType(t, &testStruct{}, got)
	}

	{
		c := Client[testMap]{}
		got, err := c.Get(url)
		assert.NoError(t, err)
		assert.IsType(t, &testMap{}, got)
	}
}
