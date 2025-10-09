package models

// Path: backend/cmd/main.go
// Author: Daniel J. Manning
// GitHub: https://github.com/djmcodechain/Go-Portfolio
type Metadata struct {
	// The pages metadata
	title       string
	description string
	canonical   string
	index       string
	cssLink     string
	jsLink      string
}

type OpenGraphTags struct {
	// The pages Open Graph metadata tags
	locale         string
	ogType         string
	title          string
	description    string
	url            string
	siteName       string
	image          string
	imageSecureURL string
}

// GNU Public
// Copyright (c) 2025 Daniel J. Manning
// Created: Sun, 05 Oct 2025
//
// License: GNU Public (See LICENSE file in repository)
