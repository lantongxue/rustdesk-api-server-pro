package util

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
)

var _httpProxy = ""

func SetHttpProxy(proxy string) {
	_httpProxy = proxy
}

func HttpClient() (*http.Client, error) {
	client := http.DefaultClient
	if _httpProxy != "" {
		proxyUrl, err := url.Parse(_httpProxy)
		if err != nil {
			return nil, err
		}
		scheme := strings.ToLower(proxyUrl.Scheme)
		allow_schemes := []string{"http", "https", "socks5"}
		if !InArray(allow_schemes, scheme) {
			return nil, errors.New("only support http, https, socks5 proxy protocols")
		}

		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	}
	return client, nil
}

func HttpGetString(remoteAddr string) (string, error) {
	req, err := http.NewRequest("GET", remoteAddr, nil)
	if err != nil {
		return "", err
	}
	client, _ := HttpClient()
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

func DownloadFile(remoteAddr, filename string, showConsoleProgressBar bool) error {
	req, err := http.NewRequest("GET", remoteAddr, nil)
	if err != nil {
		return err
	}
	httpClient, _ := HttpClient()
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	if showConsoleProgressBar {
		bar := progressbar.DefaultBytes(
			resp.ContentLength,
			"Downloading",
		)
		io.Copy(io.MultiWriter(f, bar), resp.Body)
	} else {
		io.Copy(f, resp.Body)
	}

	return nil
}
