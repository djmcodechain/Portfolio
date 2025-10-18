package routes

// Path: backend/routes/routes.go
// Author: Daniel J. Manning
// GitHub: https://github.com/djmcodechain/Portfolio

import (
	// import packages
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/djmcodechain/Portfolio/backend/handlers"
)

func Routes() *http.ServeMux {
	// Create a new Mux
	var Mux = http.NewServeMux()

	// The routes to pages
	siteMode := os.Getenv("SITE_MODE")
	switch strings.ToLower(siteMode) {
	case "dev", "devMode":
		Mux.HandleFunc("/", handlers.MaintainenceHandler)
	case "live":
		Mux.HandleFunc("/", handlers.IndexHandler)
	default:
		log.Fatalf("Error: Unknown SITE_MODE '%s'. Expected 'dev' or 'live'.", siteMode)
	}

	// Return Mux to allow it to be served in main.go
	return Mux
}

// GNU Public
// Copyright (c) 2025 Daniel J. Manning
// Created: Sun, 05 Oct 2025
//
// License: GNU Public (See LICENSE file in repository)
