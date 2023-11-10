package middleware

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadImage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the uploaded file from the request
		file, err := c.FormFile("thumbnail")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		// Open the uploaded file for reading
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer src.Close()
		// Create a temporary file to save the uploaded file to
		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer tempFile.Close()

		// Copy the uploaded file to the temporary file
		if _, err = io.Copy(tempFile, src); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// Get the filename of the saved file
		data := tempFile.Name()
		// filename := data[8:] // split uploads/

		// Set the filename as a context variable
		c.Set("imageFile", data)
		return next(c)
	}
}