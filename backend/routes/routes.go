package routes

// Path: backend/routes/routes.go
// Author: Daniel J. Manning
// GitHub: https://github.com/djmcodechain/Portfolio

import (
	// import packages
	"net/http"
	"os"
	"strings"

	"github.com/djmcodechain/Portfolio/backend/handlers"
	"github.com/djmcodechain/Portfolio/backend/logging"
)

func Routes() *http.ServeMux {
	defer logging.Logger.Info("Routes() ~ called")
	// Create a new Mux
	var Mux = http.NewServeMux()

	// The routes to pages
	siteMode := os.Getenv("SITE_MODE")
	switch strings.ToLower(siteMode) {
	case "dev", "devmode", "":
		if siteMode == "" {
			logging.Logger.Warn("SITE_MODE not set, defaulting to maintenance mode")
		}
		Mux.HandleFunc("/", handlers.MaintainenceHandler)
	case "live":
		Mux.HandleFunc("/", handlers.IndexHandler)
	default:
		logging.Logger.Warn("Unknown SITE_MODE, defaulting to maintenance mode", "site_mode", siteMode)
		Mux.HandleFunc("/", handlers.MaintainenceHandler)
	}

	// Return Mux to allow it to be served in main.go
	return Mux
}

// GNU Public
// Copyright (c) 2025 Daniel J. Manning
// Created: Sun, 05 Oct 2025
//
// License: GNU Public (See LICENSE file in repository)
