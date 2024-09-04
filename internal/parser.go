package internal

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetLatestGoVersion(u string) (string, error) {
	resp, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`go\d+\.\d+(\.\d+)?`)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}

	var latestVersion string
	s := doc.Find("div")
	latestVersion = re.FindString(s.Text())
	if latestVersion == "" {
		return "", fmt.Errorf("no Go versions found")
	}

	return latestVersion, nil
}
