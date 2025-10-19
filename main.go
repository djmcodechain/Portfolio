package main

import (
	"net/http"

	"github.com/djmcodechain/Portfolio/backend/logging"
)

// Path: backend/cmd/main.go
// Author: Daniel J. Manning
// GitHub: https://github.com/djmcodechain/Portfolio
func main() {
	defer logging.Logger.Info("main() ~ called")
	// Serve Static Files
	fs := http.FileServer(http.Dir("../../frontend/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// GNU Public
// Copyright (c) 2025 Daniel J. Manning
// Created: Sun, 05 Oct 2025
//
// License: GNU Public (See LICENSE file in repository)
