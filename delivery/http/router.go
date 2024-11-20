package http

import (
	usecase "Projects/internal/usecase"
	"github.com/go-chi/chi"
)

func NewRouter(usecase *usecase.SongUsecase) *chi.Mux {
	router := chi.NewRouter()
	handler := NewSongHandler(usecase)

	router.Get("/songs", handler.GetSongs)
	router.Get("/song/text", handler.GetSongText)
	router.Delete("/song", handler.DeleteSong)
	router.Put("/song", handler.UpdateSong)
	router.Post("/song", handler.AddSong)

	return router
}
