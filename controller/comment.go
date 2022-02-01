package controller

import (
	"database/sql"
	"emalm/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// /comments/{content_type}/{uuid}

func PostComment(w http.ResponseWriter, r *http.Request) {

	var content_id uint64
	var err error

	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()

	content_type := mux.Vars(r)["content_type"]

	uuid := mux.Vars(r)["uuid"]

	if content_type == "image" {
		content_id, err = model.GetImageIDByUuid(uuid)
		if err != nil {
			WriteError(w, err)
			return
		}
	}

	var reply_to sql.NullInt64

	comment := model.Comment{
		User_id:    1,
		Content_id: content_id,
		Posted_at:  time.Now(),
		Content:    r.FormValue("content"),
		Reply_to:   uint64(reply_to.Int64),
	}

	err = model.PostComment(comment)

	if err != nil {
		WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	comments, err := model.GetAllComments()

	if err != nil {
		WriteError(w, err)
		return
	}

	json.NewEncoder(w).Encode(comments)

}
