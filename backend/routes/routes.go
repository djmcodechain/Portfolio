package routes

// Path: backend/cmd/main.go
// Author: Daniel J. Manning
// GitHub: https://github.com/djmcodechain/Go-Portfolio

import (
	// import packages
	"net/http"

	"github.com/djmcodechain/Portfolio/backend/handlers"
)

func Routes() *http.ServeMux {
	// Create a new Mux
	var Mux = http.NewServeMux()

	// The routes to pages
	Mux.HandleFunc("/", handlers.IndexHandler)

	// Return Mux to allow it to be served in main.go
	return Mux
}

// GNU Public
// Copyright (c) 2025 Daniel J. Manning
// Created: Sun, 05 Oct 2025
//
// License: GNU Public (See LICENSE file in repository)
