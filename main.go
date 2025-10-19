package main

import (
	"log"
	"net/http"

	"github.com/djmcodechain/Portfolio/backend/logging"
	"github.com/djmcodechain/Portfolio/backend/routes"
)

// Path: backend/cmd/main.go
// Author: Daniel J. Manning
// GitHub: https://github.com/djmcodechain/Portfolio
func main() {
	defer logging.Logger.Info("main() ~ called")
	mux := routes.Routes()

	fs := http.FileServer(http.Dir("frontend/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	addr := ":8080"
	logging.Logger.Info("HTTP server starting", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}

// GNU Public
// Copyright (c) 2025 Daniel J. Manning
// Created: Sun, 05 Oct 2025
//
// License: GNU Public (See LICENSE file in repository)
