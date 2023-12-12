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

func DownloadFile(remoteAddr, filename, proxyServer string, showConsoleProgressBar bool) error {
	req, err := http.NewRequest("GET", remoteAddr, nil)
	if err != nil {
		return err
	}
	httpClient := http.DefaultClient
	if proxyServer != "" {
		proxyUrl, err := url.Parse(proxyServer)
		if err != nil {
			return err
		}
		scheme := strings.ToLower(proxyUrl.Scheme)
		allow_schemes := []string{"http", "https", "socks5"}
		if !InStringArray(allow_schemes, scheme) {
			return errors.New("only support http, https, socks5 proxy protocols")
		}

		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}
	}
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
