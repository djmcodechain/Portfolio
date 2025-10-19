package handlers

import (
	// import packages
	"net/http"

	"github.com/djmcodechain/Portfolio/backend/logging"
)

// Path: backend/handlers/index.go
// Author: Daniel J. Manning
// GitHub: https://github.com/djmcodechain/Portfolio
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer logging.Logger.InfoContext(ctx, "IndexHandler() ~ called")
	w.Write([]byte(`
	<p> Just testing to see if this does response when it has data </p>
	`))
}

// GNU Public
// Copyright (c) 2025 Daniel J. Manning
// Created: Sun, 05 Oct 2025
//
// License: GNU Public (See LICENSE file in repository)
