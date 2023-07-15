package service

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/raman-vhd/video-converter/internal/lib"
	"github.com/raman-vhd/video-converter/internal/repository"
)

type IVideoService interface {
	ConvertVideo(videoID string, quality string, ext string) error
}

type videoService struct {
	repo repository.IVideoRepository
	env  lib.Env
}

func NewVideo(
	repo repository.IVideoRepository,
	env lib.Env,
) IVideoService {
	return videoService{
		repo: repo,
		env:  env,
	}
}

func (s videoService) ConvertVideo(videoID string, quality string, ext string) error {
	err := s.repo.UpdateVideoState(videoID, quality, "converting", 0)
	if err != nil {
		log.Printf("err: %s\n", err)
	}

	input := s.env.VideoDir + videoID + ext
	output := s.env.VideoDir + videoID + "-" + quality + ext

	err = startConverting(input, output, quality)
	if err != nil {
		log.Printf("err: %s\n", err)
		return err
	}

	info, err := os.Stat(output)
	if err != nil {
		log.Printf("err: %s\n", err)
		return err
	}
	size := info.Size()
	err = s.repo.UpdateVideoState(videoID, quality, "done", int(size))
	if err != nil {
		log.Printf("err: %s\n", err)
		return err
	}
	return nil
}

func startConverting(input string, output string, quality string) error {
	args := fmt.Sprintf("-i %v -vf scale=-2:%v %v -y", input, quality, output)
	cmd := exec.Command("ffmpeg", strings.Split(args, " ")...)

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Printf("err: %s\n", err)
		return nil
	}
	return nil
}

func info(data string) map[string]string {
	myExp := regexp.MustCompile(`frame=(?P<frame>.*)fps=(?P<fps>.*)q=(?P<q>.*)size=(?P<size>.*)` +
		`time=(?P<time>.*)bitrate=(?P<bitrate>.*)speed=(?P<speed>.*)`)

	result := map[string]string{}
	if !myExp.Match([]byte(data)) {
		return result
	}

	matches := myExp.FindStringSubmatch(data)
	names := myExp.SubexpNames()
	for i, m := range matches {
		if names[i] == "" {
			continue
		}
		result[names[i]] = strings.TrimSpace(m)
	}
	return result
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func rawScan(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	return len(data), dropCR(data), nil
}
