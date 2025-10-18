package handlers

import (
	// import packages
	"net/http"
)

// Path: backend/handlers/index.go
// Author: Daniel J. Manning
// GitHub: https://github.com/djmcodechain/Portfolio
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<p> Just testing to see if this does response when it has data </p>
	`))
}

// GNU Public
// Copyright (c) 2025 Daniel J. Manning
// Created: Sun, 05 Oct 2025
//
// License: GNU Public (See LICENSE file in repository)
