package handler

import (
	"YTStreamGoApi/models"
	"github.com/nvvi9/YTStreamGo/model"
	"time"
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

func newVideoDetailsResponse(videoDetails *model.VideoDetails) *videoDetailsResponse {
	r := new(videoDetailsResponse)
	r.VideoId = videoDetails.Id
	r.Title = videoDetails.Title
	return r
}

type videoDataResponse struct {
	VideoId string   `json:"videoId"`
	Title   string   `json:"title"`
	Streams []string `json:"streams"`
}

func newVideoDataResponse(videoData *model.VideoData) *videoDataResponse {
	r := new(videoDataResponse)
	r.VideoId = videoData.VideoDetails.Id
	r.Title = videoData.VideoDetails.Title
	streams := make([]string, len(videoData.Streams))

	for i, s := range videoData.Streams {
		streams[i] = s.Url
	}

	r.Streams = streams

	return r
}
