package extractor

import (
	"net/http"
	"sync"

	"github.com/kkdai/youtube/v2"
)

type Extractor struct {
	client *youtube.Client
}

func NewExtractor() *Extractor {
	return &Extractor{
		client: &youtube.Client{
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyFromEnvironment,
				},
			},
		},
	}
}

func (e *Extractor) GetVideo(id string) (*youtube.Video, error) {
	return e.GetVideo(id)
}

func (e *Extractor) GetVideos(ids []string) []*youtube.Video {
	videos := make([]*youtube.Video, len(ids))
	var wg sync.WaitGroup

	for i, id := range ids {
		wg.Add(1)
		go func(index int, videoId string) {
			defer wg.Done()
			video, err := e.client.GetVideo(videoId)
			if err != nil {
				return
			}

			videos[index] = video
		}(i, id)
	}

	wg.Wait()

	return videos
}
