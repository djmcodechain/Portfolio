package handlers

// Path: backend/handlers/maintenance.go
// Author: Daniel J. Manning
// GitHub: https://github.com/djmcodechain/Portfolio

import (
	"net/http"

	"github.com/djmcodechain/Portfolio/backend/models"
	"github.com/djmcodechain/Portfolio/backend/templates"
	"github.com/djmcodechain/Portfolio/backend/templates/views"
)

func MaintainenceHandler(w http.ResponseWriter, r *http.Request) {
	meta := models.Metadata{
		Title:       "djmcodechain [Under Maintenance]",
		Description: "djmcodechain is currently under some development.",
		Canonical:   "djmcodechain.dev/",
		Index:       "index, follow",
		CSSlink:     "frontend/assets/css/style.css",
	}
	ogTags := models.OpenGraphTags{
		Locale:      "en_GB",
		OGtype:      "website",
		Title:       meta.Title,
		Description: meta.Description,
		URL:         meta.Canonical,
	}
	body := views.Maintenance()
	layout := templates.Layout(&meta, &ogTags, body)

	if err := layout.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GNU Public
// Copyright (c) 2025 Daniel J. Manning
// Created: Sun, 18 Oct 2025
//
// License: GNU Public (See LICENSE file in repository)
