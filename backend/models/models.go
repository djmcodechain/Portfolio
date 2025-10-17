package models

// Path: backend/models/models.go
// Author: Daniel J. Manning
// GitHub: https://github.com/djmcodechain/Portfolio

type Project struct {
	Name     string
	Stack    string
	Branding *Branding
}

type Branding struct {
	Colours    *Colours
	Typography *Typography
}

type Colours struct {
	Primary   string
	Secondary string
	Accent    string
	Links     string
}

type Typography struct {
	Headings string
	Body     string
	Code     string
	Links    string
}

type Metadata struct {
	// The pages metadata
	Title       string
	Description string
	Canonical   string
	Index       string
	CSSlink     string
	JSlink      string
}

type OpenGraphTags struct {
	// The pages Open Graph metadata tags
	Locale         string
	OGtype         string
	Title          string
	Description    string
	URL            string
	siteName       string
	Image          string
	imageSecureURL string
}

// GNU Public
// Copyright (c) 2025 Daniel J. Manning
// Created: Sun, 05 Oct 2025
//
// License: GNU Public (See LICENSE file in repository)
