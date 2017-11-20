package curl

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// JSON a raw url
func JSON(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// https://stackoverflow.com/a/31399046/145435
	buf = bytes.TrimPrefix(buf, []byte("\xef\xbb\xbf"))

	return string(buf), nil
}

// XML a raw url
func XML(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// https://stackoverflow.com/a/31399046/145435
	buf = bytes.TrimPrefix(buf, []byte("\xef\xbb\xbf"))

	return string(buf), nil
}
