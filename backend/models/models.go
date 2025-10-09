package models

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
