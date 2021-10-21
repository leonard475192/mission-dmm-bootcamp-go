package timelines

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"yatter-backend-go/app/handler/httperror"
)

func (h *handler) Public(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// FIXME エラーメッセージの共通の関数が良い気がする
	only_media, err := strconv.Atoi(r.URL.Query().Get("only_media"))
	if err != nil || only_media < 0 || 1 < only_media {
		log.Printf("only_media: %v", err)
		httperror.BadRequest(w, fmt.Errorf("you must set only_media to 0 or 1 in in query parameters"))
		return
	}
	only_media_bool := false
	if only_media != 0 {
		only_media_bool = true
	}
	max_id, err := strconv.Atoi(r.URL.Query().Get("max_id"))
	if err != nil {
		log.Printf("max_id: %v", err)
		httperror.BadRequest(w, fmt.Errorf("you must include max_id as an integer in query parameters"))
		return
	}
	since_id, err := strconv.Atoi(r.URL.Query().Get("since_id"))
	if err != nil {
		log.Printf("since_id: %v", err)
		httperror.BadRequest(w, fmt.Errorf("you must include since_id as an integer in query parameters"))
		return
	}
	// FIXME ユーザーがlimitに整数以外を入れた場合にも、Defaultを返してしまう
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		log.Printf("limit: %v", err)
		limit = 40
	}
	// TODO レスポンスがnull そのほか設定
	status, err := h.app.Dao.Status().Select(ctx, only_media_bool, since_id, max_id, limit)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
