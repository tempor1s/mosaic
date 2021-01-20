package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tempor1s/mosaic/auth"
	"github.com/tempor1s/mosaic/models"
)

// Upload will upload a file to the server and return the URL to the image
func (h *Handlers) Upload(c echo.Context) error {
	// get the user id from context (jwt)
	userID, err := auth.GetUserIDFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// TODO: create the image in google cloud objects to get URL information etc

	newImg := models.Image{
		FullURL:  "https://storage.googleapis.com/mosaic-image-bucket/629bbec443dae461f7fa746093e1a122.png",
		ShortURL: "https://i.benl.dev/629bbec443dae461f7fa746093e1a122.png",
		UserID:   userID,      // set the user ID in the DB to be that from auth
		Name:     "file name", // TODO: pull this from the uploaded files header
	}

	// TODO: create the image in the database
	store := models.NewStore(h.Database)
	err = store.CreateImage(newImg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create image in the database")
	}

	return c.JSON(http.StatusOK, newImg)
}
