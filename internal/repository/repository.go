package repository

import (
	"Projects/internal/models"
	"database/sql"
)

type SongRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) GetSongs(filter map[string]string, limit, offset int) ([]models.Song, error) {
	query := "SELECT * FROM songs"
	var args []interface{}
	var conditions []string

	for k, v := range filter {
		conditions = append(conditions, k+"=$"+string(len(args)+1))
		args = append(args, v)
	}

	if len(conditions) > 0 {
		query += " WHERE " + conditions[0]
		for i := 1; i < len(conditions); i++ {
			query += " AND " + conditions[i]
		}
	}

	query += " LIMIT $" + string(len(args)+1) + " OFFSET $" + string(len(args)+2)
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	return songs, nil
}

func (r *SongRepository) GetSongText(songID, limit, offset int) (string, error) {
	var text string
	err := r.db.QueryRow("SELECT text FROM songs WHERE id=$1 LIMIT $2 OFFSET $3", songID, limit, offset).Scan(&text)
	if err != nil {
		return "", err
	}
	return text, nil
}

func (r *SongRepository) DeleteSong(songID int) error {
	_, err := r.db.Exec("DELETE FROM songs WHERE id=$1", songID)
	return err
}

func (r *SongRepository) UpdateSong(song models.Song) error {
	_, err := r.db.Exec("UPDATE songs SET group=$1, song=$2, release_date=$3, text=$4, link=$5 WHERE id=$6",
		song.Group, song.Song, song.ReleaseDate, song.Text, song.Link, song.ID)
	return err
}

func (r *SongRepository) AddSong(song models.Song) (int, error) {
	var id int
	err := r.db.QueryRow("INSERT INTO songs (group, song, release_date, text, link) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		song.Group, song.Song, song.ReleaseDate, song.Text, song.Link).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
