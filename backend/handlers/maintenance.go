package handlers

import (
	"danieljmanning/backend/models"
	"net/http"
)

func MaintainenceHandler(w http.ResponseWriter, r *http.Request) {
	meta := models.Metadata{
		Title:       "djmcodechain [Under Maintenance]",
		Description: "djmcodechain is currently under some development.",
		Canonical:   "djmcodechain.dev/",
		Index:       "index, follow",
		CSSlink:     "frontend/assets/css/style.css",
	}

	ogTags := models.OpenGraphTags{
		Locale:      "GB-en",
		OGtype:      "website",
		Title:       meta.Title,
		Description: meta.Description,
		URL:         meta.Canonical,
	}
}
