package handler

import (
	"YTStreamGoApi/models"
	"time"

	"github.com/kkdai/youtube/v2"
)

type userResponse struct {
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func newUserResponse(u *models.User) *userResponse {
	r := new(userResponse)
	r.Name = u.Name
	r.CreatedAt = u.CreatedAt
	r.UpdatedAt = u.UpdatedAt
	return r
}

type videoDetailsResponse struct {
	VideoId string `json:"videoId"`
	Title   string `json:"title"`
}

func newVideoDetailsResponse(v *youtube.Video) *videoDetailsResponse {
	r := new(videoDetailsResponse)
	r.VideoId = v.ID
	r.Title = v.Title
	return r
}

type videoDataResponse struct {
	VideoId     string               `json:"videoId"`
	Title       string               `json:"title"`
	Author      string               `json:"author"`
	Duration    time.Duration        `json:"duration"`
	PublishDate time.Time            `json:"publishDate"`
	Streams     []*streamResponse    `json:"streams"`
	Thumbnails  []*thumbnailResponse `json:"thumbnails"`
}

type thumbnailResponse struct {
	Url    string `json:"url"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
}

type streamResponse struct {
	Url      string `json:"url"`
	MimeType string `json:"mimeType"`
	Quality  string `json:"quality"`
	Bitrate  int    `json:"bitrate"`
}

func newVideoDataResponse(v *youtube.Video) *videoDataResponse {
	r := new(videoDataResponse)
	r.VideoId = v.ID
	r.Title = v.Title
	r.Author = v.Author
	r.Duration = v.Duration
	r.PublishDate = v.PublishDate

	if formats := v.Formats; formats != nil {
		streams := make([]*streamResponse, len(formats))

		for i, f := range formats {
			s := new(streamResponse)
			s.Url = f.URL
			s.MimeType = f.MimeType
			s.Quality = f.Quality
			s.Bitrate = f.Bitrate
			streams[i] = s
		}

		r.Streams = streams
	}

	if thumbnails := v.Thumbnails; thumbnails != nil {
		thumbs := make([]*thumbnailResponse, len(thumbnails))

		for i, t := range thumbnails {
			thumbnail := new(thumbnailResponse)
			thumbnail.Url = t.URL
			thumbnail.Height = t.Height
			thumbnail.Width = t.Width
			thumbs[i] = thumbnail
		}

		r.Thumbnails = thumbs
	}

	return r
}
