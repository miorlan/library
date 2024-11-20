package http

import (
	"Projects/internal/models"
	usecase "Projects/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type SongHandler struct {
	usecase *usecase.SongUsecase
}

func NewSongHandler(usecase *usecase.SongUsecase) *SongHandler {
	return &SongHandler{usecase: usecase}
}

// swagger:operation GET /songs Song GetSongs
//
// ---
// summary: Получить песни с пагинацией и фильтрацией
// description: Возвращает список песен с опциональной пагинацией и фильтрацией.
// parameters:
//   - name: limit
//     in: query
//     description: Количество песен на странице
//     required: false
//     type: integer
//     format: int32
//   - name: offset
//     in: query
//     description: Смещение для пагинации
//     required: false
//     type: integer
//     format: int32
//   - name: group
//     in: query
//     description: Название группы
//     required: false
//     type: string
//   - name: song
//     in: query
//     description: Название песни
//     required: false
//     type: string
//
// responses:
//
//	'200':
//	  description: Успешный ответ
//	  schema:
//	    type: array
//	    items:
//	      $ref: '#/definitions/Song'
//	'500':
//	  description: Внутренняя ошибка сервера
func (h *SongHandler) GetSongs(w http.ResponseWriter, r *http.Request) {
	filter := make(map[string]string)
	for k, v := range r.URL.Query() {
		filter[k] = v[0]
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	songs, err := h.usecase.GetSongs(filter, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(songs)
}

// swagger:operation GET /song/text Song GetSongText
//
// ---
// summary: Получить текст песни с пагинацией
// description: Возвращает текст песни с опциональной пагинацией.
// parameters:
//   - name: id
//     in: query
//     description: ID песни
//     required: true
//     type: integer
//     format: int32
//   - name: limit
//     in: query
//     description: Количество куплетов на странице
//     required: false
//     type: integer
//     format: int32
//   - name: offset
//     in: query
//     description: Смещение для пагинации
//     required: false
//     type: integer
//     format: int32
//
// responses:
//
//	'200':
//	  description: Успешный ответ
//	  schema:
//	    type: string
//	'500':
//	  description: Внутренняя ошибка сервера
func (h *SongHandler) GetSongText(w http.ResponseWriter, r *http.Request) {
	songID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	text, err := h.usecase.GetSongText(songID, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"text": text})
}

// swagger:operation DELETE /song Song DeleteSong
//
// ---
// summary: Удалить песню по ID
// description: Удаляет песню по заданному ID.
// parameters:
//   - name: id
//     in: query
//     description: ID песни
//     required: true
//     type: integer
//     format: int32
//
// responses:
//
//	'200':
//	  description: Успешный ответ
//	'500':
//	  description: Внутренняя ошибка сервера
func (h *SongHandler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	songID, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := h.usecase.DeleteSong(songID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// swagger:operation PUT /song Song UpdateSong
//
// ---
// summary: Обновить песню по ID
// description: Обновляет информацию о песне по заданному ID.
// parameters:
//   - name: song
//     in: body
//     description: Данные для обновления песни
//     required: true
//     schema:
//     $ref: '#/definitions/Song'
//
// responses:
//
//	'200':
//	  description: Успешный ответ
//	'400':
//	  description: Неверный запрос
//	'500':
//	  description: Внутренняя ошибка сервера
func (h *SongHandler) UpdateSong(w http.ResponseWriter, r *http.Request) {
	var song models.Song
	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.usecase.UpdateSong(song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// swagger:operation POST /song Song AddSong
//
// ---
// summary: Добавить новую песню
// description: Добавляет новую песню и получает дополнительные данные из внешнего API.
// parameters:
//   - name: song
//     in: body
//     description: Данные новой песни
//     required: true
//     schema:
//     $ref: '#/definitions/SongRequest'
//
// responses:
//
//	'200':
//	  description: Успешный ответ
//	  schema:
//	    type: integer
//	    description: ID новой песни
//	'500':
//	  description: Внутренняя ошибка сервера
func (h *SongHandler) AddSong(w http.ResponseWriter, r *http.Request) {
	var song models.Song
	err := json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.usecase.AddSong(song)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"id": id})
}
