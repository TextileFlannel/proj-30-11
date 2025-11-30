package storage

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"proj/internal/models"
	"sync"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Storage struct {
	data   []map[string]string
	length int
}

func NewStorage() *Storage {
	return &Storage{
		data:   make([]map[string]string, 0),
		length: 0,
	}
}

func (s *Storage) CheckLinks(links models.LinksRequest) (models.LinksResponse, error) {
	res := make(map[string]string)
	var wg sync.WaitGroup
	ch := make(chan map[string]string)

	for _, link := range links.Links {
		wg.Add(1)
		go checkLink(link, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		for k, v := range i {
			res[k] = v
		}
	}

	s.data = append(s.data, res)
	s.length += 1

	return models.LinksResponse{
		Links:    res,
		LinkNums: s.length,
	}, nil
}

func checkLink(link string, ch chan map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()

	if !(link[:7] == "http://" || link[:8] == "https://") {
		link = "http://" + link
	}

	parsedUrl, err := url.Parse(link)
	if err != nil {
		ch <- map[string]string{
			link: "not available",
		}
		return
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(parsedUrl.String())
	if err != nil {
		ch <- map[string]string{
			link: "not available",
		}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		ch <- map[string]string{
			link: "available",
		}
	} else {
		ch <- map[string]string{
			link: "not available",
		}
	}
}

func (s *Storage) GetAllLinks() ([]map[string]string, error) {
	return s.data, nil
}

func (s *Storage) ReportLinks(linksList models.ReportLinksRequest) (bytes.Buffer, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 24)
	pdf.Cell(0, 20, "List links")
	pdf.Ln(25)

	pdf.SetFont("Arial", "", 18)
	for _, num := range linksList.LinksList {
		if num <= s.length {
			pdf.Cell(0, 15, fmt.Sprintf("[%d]", num))
			pdf.Ln(10)
			elem := s.data[num-1]
			for k, v := range elem {
				pdf.Cell(0, 15, k+"   --->   "+v)
				pdf.Ln(10)
			}
		} else {
			pdf.Cell(0, 15, fmt.Sprintf("[%d]", num))
			pdf.Ln(10)
			pdf.Cell(0, 15, "Not found")
			pdf.Ln(10)
		}
	}

	var buf bytes.Buffer
	_ = pdf.Output(&buf)

	return buf, nil
}
