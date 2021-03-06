package accounts

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Username string
	Password string
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	account := new(object.Account)
	account.Username = req.Username
	if err := account.SetPassword(req.Password); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	// 既存のusernameを弾く、ここでやらないほうがいい気がするが、BadRequestを返すとなると、ここ？
	alreadyAccount, err := h.app.Dao.Account().FindByUsername(ctx, req.Username)
	log.Printf("already account: %v", alreadyAccount)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	if alreadyAccount != nil {
		httperror.BadRequest(w, fmt.Errorf("%v", "Validation Error: This name exists."))
		return
	}

	account, err = h.app.Dao.Account().Create(ctx, *account)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
