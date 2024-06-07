package github

import (
	"encoding/json"
	"fmt"
	"os"
	"rustdesk-api-server-pro/util"
)

type Release struct {
	ID              int     `json:"id"`
	NodeID          string  `json:"node_id"`
	TagName         string  `json:"tag_name"`
	TargetCommitish string  `json:"target_commitish"`
	Name            string  `json:"name"`
	Draft           bool    `json:"draft"`
	Prerelease      bool    `json:"prerelease"`
	CreatedAt       string  `json:"created_at"`
	PublishedAt     string  `json:"published_at"`
	Assets          []Asset `json:"assets"`
}

type Asset struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Label              string `json:"label"`
	ContentType        string `json:"content_type"`
	State              string `json:"state"`
	Size               int    `json:"size"`
	DownloadCount      int    `json:"download_count"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func GetReleases(repo string) *[]Release {
	repo = fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	resp, err := util.HttpGetString(repo)
	if err != nil {
		fmt.Println("GetReleases:: Request error:", err)
		os.Exit(0)
	}
	releases := &[]Release{}
	err = json.Unmarshal([]byte(resp), releases)
	if err != nil {
		fmt.Println("GetReleases:: json.Unmarshal:", err)
		os.Exit(0)
	}
	return releases
}

func GetLatestRelease(repo string) *Release {
	repo = fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repo)
	resp, err := util.HttpGetString(repo)
	if err != nil {
		fmt.Println("GetLatestRelease:: Request error:", err)
		os.Exit(0)
	}
	release := &Release{}
	err = json.Unmarshal([]byte(resp), release)
	if err != nil {
		fmt.Println("GetLatestRelease:: json.Unmarshal:", err)
		os.Exit(0)
	}
	return release
}

func GetReleaseByTag(repo, tag string) *Release {
	repo = fmt.Sprintf("https://api.github.com/repos/%s/releases/tags/%s", repo, tag)
	resp, err := util.HttpGetString(repo)
	if err != nil {
		fmt.Println("GetReleaseByTag:: Request error:", err)
		os.Exit(0)
	}
	release := &Release{}
	err = json.Unmarshal([]byte(resp), release)
	if err != nil {
		fmt.Println("GetReleaseByTag:: json.Unmarshal:", err)
		os.Exit(0)
	}
	return release
}
