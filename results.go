package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// https://developers.google.com/youtube/v3/docs/search
type Results struct {
	Items []struct {
		Kind string
		Etag string
		Id   struct {
			Kind       string
			VideoId    string
			ChannelId  string
			PlaylistId string
		}
		Snippet struct {
			PublishedAt string
			ChannelId   string
			Title       string
			Description string
			Thumbnails  struct {
				Default struct {
					Url    string
					Width  uint
					Height uint
				}
				Medium struct {
					Url    string
					Width  uint
					Height uint
				}
				High struct {
					Url    string
					Width  uint
					Height uint
				}
			}
		}
		ChannelTitle         string
		LiveBroadcastContent string
	}
}

const baseUrl = "https://www.googleapis.com/youtube/v3/search?"

func parseKey() string {
	buf, err := os.ReadFile(".env")
	if err != nil {
		log.Fatalln("Could not read .env file")
	}

	parts := strings.SplitN(string(buf), "=", 2)
	if len(parts) != 2 {
		log.Fatalln("Invalid .env file content")
	}

	env := strings.Trim(strings.TrimSpace(parts[1]), "\"'")

	if env == "" {
		log.Fatalln("No API_KEY found in .env")
	}

	return env
}

var API_KEY = parseKey()

func queryYoutube(query string) (*Results, error) {
	url := baseUrl + "&q=" + url.QueryEscape(query) + "&part=snippet" + "&type=video" + "&key=" + API_KEY

	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var results Results
	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}
