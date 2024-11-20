package usecases

import (
	"Projects/config"
	"Projects/internal/models"
	"Projects/internal/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SongUsecase struct {
	repo   *repository.SongRepository
	config *config.Config
}

func NewSongUsecase(repo *repository.SongRepository, config *config.Config) *SongUsecase {
	return &SongUsecase{repo: repo, config: config}
}

func (u *SongUsecase) GetSongs(filter map[string]string, limit, offset int) ([]models.Song, error) {
	return u.repo.GetSongs(filter, limit, offset)
}

func (u *SongUsecase) GetSongText(songID, limit, offset int) (string, error) {
	return u.repo.GetSongText(songID, limit, offset)
}

func (u *SongUsecase) DeleteSong(songID int) error {
	return u.repo.DeleteSong(songID)
}

func (u *SongUsecase) UpdateSong(song models.Song) error {
	return u.repo.UpdateSong(song)
}

func (u *SongUsecase) AddSong(song models.Song) (int, error) {
	url := fmt.Sprintf("%s?group=%s&song=%s", u.config.APIUrl, song.Group, song.Song)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var songDetail models.SongDetails
	err = json.Unmarshal(body, &songDetail)
	if err != nil {
		return 0, err
	}

	song.ReleaseDate = songDetail.ReleaseDate
	song.Text = songDetail.Text
	song.Link = songDetail.Link

	return u.repo.AddSong(song)
}
