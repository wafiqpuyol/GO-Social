package main

import (
	"net/http"

	"github.com/wafiqpuyol/GO-Social/internal/helper"
	"github.com/wafiqpuyol/GO-Social/internal/store"
)

type PostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

func (a *application) createPost(w http.ResponseWriter, r *http.Request) {
	var payload PostPayload

	// read payload
	if err := helper.ReadJson(w, r, &payload); err != nil {
		helper.WriteJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	// validate payload
	if err := helper.Validate.Struct(payload); err != nil {
		helper.WriteJsonError(w, http.StatusBadRequest, err.Error())
	}

	// create post
	post := store.Post{
		Content: payload.Content,
		UserID:  45,
		Title:   payload.Title,
		Tags:    payload.Tags,
	}

	ctx := r.Context()
	if err := a.store.Post.CreatePost(ctx, &post); err != nil {
		helper.WriteJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// write response
	if err := helper.WriteJson(w, http.StatusCreated, &post); err != nil {
		helper.WriteJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
