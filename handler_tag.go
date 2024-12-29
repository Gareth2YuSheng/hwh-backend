package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CREATE NEW TAG
func (apiCfg *APIConfig) handlerCreateTag(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerCreateTag")
	req := CreateTagRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ErrorParsingJSON(err, w)
		return
	}
	defer r.Body.Close()

	tag, err := NewTag(req.Name)
	if err != nil {
		logError("Error Creating New Standard Tag Template", err)
		respondERROR(w, http.StatusBadRequest, "Failed to Create Tag, Invalid Tag Details")
		return
	}
	fmt.Printf("New Tag: %v\n", tag) //remove later

	err = apiCfg.DB.CreateTag(tag)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			logError("Tag Already Exists", err)
			respondERROR(w, http.StatusBadRequest, "Tag Already Exists")
			return
		}
		logError("Unable to Create Tag", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Create Tag")
		return
	}

	respondOK(w, http.StatusCreated, "Tag Created Successfully", nil)
}

// GET ALL TAGS
func (apiCfg *APIConfig) handlerGetAllTags(w http.ResponseWriter, r *http.Request, user User) {
	logInfo("Running handlerGetAllTags")
	tags, err := apiCfg.DB.GetAllTags()
	if err != nil {
		logError("Unable to Get Tags", err)
		respondERROR(w, http.StatusInternalServerError, "Failed to Get All Tags")
		return
	}

	respondOK(w, http.StatusOK, "", GetTagsResponse{
		Tags: tags,
	})
}
