package main

import (
	"bytes"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"strings"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"message": "Hello, this is compressor head",
	}
	writeJsonResponse(w, res, http.StatusOK)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {

	imgUrl, err := getImageUrl(r)
	if err != nil {
		log.Println("INFO: image_url not specifies")
		res := map[string]string{
			"message": err.Error(),
		}
		writeJsonResponse(w, res, http.StatusBadRequest)
		return
	}

	height, err := getRequestedHeight(r)
	if err != nil {
		log.Println("INFO: invalid height specified")
		res := map[string]string{
			"message": err.Error(),
		}
		writeJsonResponse(w, res, http.StatusBadRequest)
		return
	}

	width, err := getRequestedWidth(r)
	if err != nil {
		log.Println("INFO: invalid width specified")
		res := map[string]string{
			"message": err.Error(),
		}
		writeJsonResponse(w, res, http.StatusBadRequest)
		return
	}

	if height == 0 && width == 0 {
		log.Println("INFO: Both height and width can't be zero")
		res := map[string]string{
			"message": "Specify at least one of the dimension",
		}
		writeJsonResponse(w, res, http.StatusBadRequest)
		return
	}

	format, err := getRequestedFormat(r)
	if err != nil {
		log.Printf("INFO: invalid image format specifies")
		res := map[string]string{
			"message": err.Error(),
		}
		writeJsonResponse(w, res, http.StatusBadRequest)
		return
	}

	cacheKey := fmt.Sprintf("%s,%d,%d,%s", imgUrl, height, width, format)
	validCacheKey := len([]byte(cacheKey)) <= 250 // Memcache key must be at most 250 bytes in length.
	if !validCacheKey {
		log.Printf("WARN: Not looking for image in cache, key size exceeds 250 byte limit - %s", cacheKey)
	}

	if validCacheKey {
		item, err := memcache.Get(appengine.NewContext(r), cacheKey)
		if err != nil {
			if err == memcache.ErrCacheMiss {
				log.Printf("WARN: Image with key(%s) not found in cache", cacheKey)
			} else {
				log.Printf("WARN: Failed to get image from key: %v", err)
			}
		} else {
			log.Printf("INFO: Serving image from memcache")
			writeImageResponse(w, item.Value, format)
			return
		}
	}

	client := http.Client{
		Timeout: time.Second * 10,
	}
	imageRes, err := client.Get(imgUrl)
	if err != nil {
		log.Printf("ERROR: Failed to get image from url %s: %v", imgUrl, err)
		res := map[string]string{
			"message": "Failed to download image, check image url",
		}
		writeJsonResponse(w, res, http.StatusBadRequest)
		return
	}
	defer imageRes.Body.Close()


	rawImage, rawFormat, err := image.Decode(imageRes.Body)
	if err != nil {
		log.Printf("ERROR: Failed to decode image: %v", err)
		res := map[string]string{
			"message": "Failed to decode image",
		}
		writeJsonResponse(w, res, http.StatusBadRequest)
		return
	}

	img := resize.Resize(width, height, rawImage, resize.Lanczos3)

	var buf bytes.Buffer
	switch strings.ToUpper(format) {
	case "JPEG", "JPG":
		err = jpeg.Encode(&buf, img, nil)
	case "PNG":
		err = png.Encode(&buf, img)
	case "GIF":
		err = gif.Encode(&buf, img, nil)
	case "BMP":
		err = bmp.Encode(&buf, img)
	case "TIFF":
		err = tiff.Encode(&buf, img, nil)
	case "WEBP":
		err = webp.Encode(&buf, img, nil)
	default:
		log.Printf("ERROR: Unknown format - %v", rawFormat)
		res := map[string]string{
			"message": fmt.Sprintf("Only following formats are supported - %v", imageFormats),
		}
		writeJsonResponse(w, res, http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Printf("ERROR: Failed to encode %s image: %v", format, err)
		res := map[string]string{
			"message": "Failed to encode image",
		}
		writeJsonResponse(w, res, http.StatusInternalServerError)
		return
	}

	if !validCacheKey {
		log.Printf("WARN: Not saving image in cache, key size exceeds 250 byte limit - %s", cacheKey)
	} else {
		log.Println("INFO: Saving image to cache")
		cacheImage := &memcache.Item{
			Key:   cacheKey,
			Value: buf.Bytes(),
		}
		err = memcache.Add(appengine.NewContext(r), cacheImage)
		if err != nil {
			if err == memcache.ErrNotStored {
				log.Println("INFO: Image already exists in cache")
			} else {
				log.Printf("WARN: Failed to save image in cache key (%s)", cacheKey)
			}
		}
	}

	writeImageResponse(w, buf.Bytes(), format)
}
