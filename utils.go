package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
)

var (
	imageFormats = []string{"JPEG", "PNG", "WEBP", "GIF", "BMP", "TIFF"}
)

func writeJsonResponse(w http.ResponseWriter, res interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(res)
}

func writeImageResponse(w http.ResponseWriter, image []byte, format string) {
	w.Header().Set("Content-Type", fmt.Sprintf("image/%s", format))
	_, _ = w.Write(image)
}

func getImageUrl(r *http.Request) (string, error) {
	imgUrls, ok := r.URL.Query()["image_url"]
	if !ok || len(imgUrls) == 0 {
		return "", errors.New("Specify the image_url parameter")
	}
	return imgUrls[0], nil
}

func getRequestedHeight(r *http.Request) (uint, error) {
	heights, ok := r.URL.Query()["height"]
	if !ok || len(heights) == 0 {
		return 0, nil
	}
	height, err := strconv.Atoi(heights[0])
	if err != nil {
		return 0, errors.New("Specify height as a valid integer")
	}
	if height < 0 {
		return 0, errors.New("Specify a positive height")
	}
	return uint(height), nil
}

func getRequestedWidth(r *http.Request) (uint, error) {
	widths, ok := r.URL.Query()["width"]
	if !ok || len(widths) == 0 {
		return 0, nil
	}
	width, err := strconv.Atoi(widths[0])
	if err != nil {
		return 0, errors.New("Specify width as a valid integer")
	}
	if width < 0 {
		return 0, errors.New("Specify a positive width")
	}
	return uint(width), nil
}

func getRequestedFormat(r *http.Request) (string, error) {
	formats, ok := r.URL.Query()["format"]
	if !ok || len(formats) == 0 {
		return "JPEG", nil
	}
	format := strings.ToUpper(strings.TrimSpace(formats[0]))
	correctFormat := false
	for _, val := range imageFormats {
		if format == val {
			correctFormat = true
			break
		}
	}
	if !correctFormat {
		return "", errors.New(fmt.Sprintf("Specify a valid image format %v", imageFormats))
	}
	return format, nil
}
