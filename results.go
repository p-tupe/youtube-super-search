package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func queryYoutube(query string) (*Results, error) {
  url := baseUrl + "&q=" + url.QueryEscape(query) + "&part=snippet&type=video&key=" // TODO: Add your key here

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
