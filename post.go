package httplibgo

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func (c *Client) Post(url string, body []byte) (status int, respBody []byte, err error) {
	return c.PostWithHeaders(url, body, map[string]string{})
}

func (c *Client) PostWithHeaders(url string, body []byte, headers map[string]string) (status int, respBody []byte, err error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
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
