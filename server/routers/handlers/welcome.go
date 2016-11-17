package handlers

import (
	"net/http"

	"github.com/sstutz/pastestuff/server/helpers"
	"github.com/sstutz/pastestuff/server/templates"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := templates.Load("elm.html")
	release := helpers.ReleaseInformation()
	t.Execute(w, map[string]interface{}{
		"Version": release.Version,
		"Build":   release.Build,
		"AppName": helpers.AppName,
	})
}
