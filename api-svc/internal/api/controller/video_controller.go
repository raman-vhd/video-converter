package controller

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/labstack/echo/v4"
	"github.com/raman-vhd/video-converter/internal/service"
)

type IVideoController interface {
	CreateVideo(ctx echo.Context) error
	GetVideo(ctx echo.Context) error
}

type videoController struct {
	svc service.IVideoService
}

func NewVideo(
	svc service.IVideoService,
) IVideoController {
	return videoController{
		svc: svc,
	}
}

func (c videoController) CreateVideo(ctx echo.Context) error {
	file, _, err := ctx.Request().FormFile("file")
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "video is missing",
		})
	}

	ext, err := checkVideoFormat(file)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal error",
		})
	}
	if ext == "" {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{
			"message": "uploaded file format is not supported",
		})
	}

	quality := ctx.FormValue("quality")
	qualityList := strings.Split(quality, ",")

	file.Seek(0, io.SeekStart)
	link, err := c.svc.CreateVideo(ctx.Request().Context(), file, ext, qualityList)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal error",
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":  "video created successfully",
		"versions": qualityList,
		"link":     link,
	})
}

// checks if the file is video and returns video extension with dot
func checkVideoFormat(file io.Reader) (string, error) {
	t, err := mimetype.DetectReader(file)
	if err != nil {
		return "", err
	}
	ok := strings.HasPrefix(t.String(), "video/")
	if !ok {
		return "", nil
	}
	ext := t.Extension()
	return ext, nil
}

func (c videoController) GetVideo(ctx echo.Context) error {
	videoID := ctx.QueryParam("id")
	v, err := c.svc.GetVideo(ctx.Request().Context(), videoID)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": "video not found",
		})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"video": v,
	})
}
