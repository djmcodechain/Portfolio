package handlers

import (
	"net/http"

	"github.com/djmcodechain/Portfolio/backend/models"
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
		Locale:      "en_GB",
		OGtype:      "website",
		Title:       meta.Title,
		Description: meta.Description,
		URL:         meta.Canonical,
	}
}
