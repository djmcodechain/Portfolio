package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/djmcodechain/Portfolio/backend/db"
	"github.com/djmcodechain/Portfolio/backend/logging"
	"github.com/djmcodechain/Portfolio/backend/routes"
)

// Path: backend/cmd/main.go
// Author: Daniel J. Manning
// GitHub: https://github.com/djmcodechain/Portfolio
func main() {
	defer logging.Logger.Info("main() ~ called")

	// Load the registry DB + run migrations
	registry := db.InitDB("backend/db/core/djmcodechain.db")

	// Load the domain DBs listed in the registry
	domainDBs, err := db.LoadAllDatabases(registry)
	if err != nil {
		panic(err)
	}

	fmt.Println("Registry DB loaded:", registry)
	fmt.Println("Domain DBs loaded:", domainDBs)

	mux := routes.Routes()

	fs := http.FileServer(http.Dir("frontend/assets/"))
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
