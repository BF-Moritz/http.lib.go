package httplibgo

import (
	"io/ioutil"
	"net/http"
)

func (c *Client) Get(url string) (status int, respBody []byte, err error) {
	return c.GetWithHeaders(url, map[string]string{})
}

func (c *Client) GetWithHeaders(url string, headers map[string]string) (status int, respBody []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, respBody, nil
}
