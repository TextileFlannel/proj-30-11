package checker

import (
	"net/http"
	"net/url"
	"proj/internal/models"
	"time"
)

type Checker struct {
	client *http.Client
}

func NewChecker() *Checker {
	return &Checker{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Checker) Check(link string) models.Result {
	if len(link) < 7 || !(link[:7] == "http://" || link[:8] == "https://") {
		link = "http://" + link
	}

	parsedUrl, err := url.Parse(link)
	if err != nil {
		return models.Result{
			URL:    link,
			Status: "not available",
		}
	}

	resp, err := c.client.Get(parsedUrl.String())
	if err != nil {
		return models.Result{
			URL:    link,
			Status: "not available",
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return models.Result{
			URL:    link,
			Status: "available",
		}
	}
	return models.Result{
		URL:    link,
		Status: "not available",
	}
}
