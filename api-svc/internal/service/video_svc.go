package service

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/raman-vhd/video-converter/internal/lib"
	"github.com/raman-vhd/video-converter/internal/model"
	"github.com/raman-vhd/video-converter/internal/repository"
	"github.com/raman-vhd/video-converter/internal/util"
)

var qualityList = []string{"144", "240", "360", "480", "720", "1080"}

type IVideoService interface{
    CreateVideo(ctx context.Context, file io.Reader, ext string, formats []string) (string, error)
    GetVideo(ctx context.Context, videoID string) (model.Video, error)
}

type videoService struct {
    repo repository.IVideoRepository
    amqp *lib.AMQP
    env lib.Env
}

func NewVideo(
    repo repository.IVideoRepository,
    amqp *lib.AMQP,
    env lib.Env,
) IVideoService{
    return videoService{
        repo: repo,
        amqp: amqp,
        env: env,
    }
}

func (s videoService) CreateVideo(ctx context.Context, reader io.Reader, ext string, quality []string) (string, error) {
    videoID := util.GenerateLink(10)
    filePath := s.env.VideoDir + videoID + ext
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

    n, err := io.Copy(file, reader)
	if err != nil {
        err := os.Remove(filePath)
        if err != nil {
            return "", err
        }
		return "", err
	}

    err = s.repo.CreateVideo(ctx, videoID, ext, int(n))
    if err != nil {
        return "", err
    }

    for _, q := range quality {
        var ok bool
        for _, i := range qualityList {
            if i == q {
                ok = true
                break
            }
        }
        if !ok {
            continue
        }
        msg := model.AMQPMsg{
            VideoID: videoID,
            Quality: q,
            Ext: ext,
        }
        msgRaw, err := json.Marshal(msg)
        if err != nil {
            return "", err
        }
        
        err = s.amqp.Publish(msgRaw)
        if err != nil {
            return "", err
        }
        
        err = s.repo.AddVersion(ctx, videoID, q)
        if err != nil {
            return "", err
        }
    }
    return videoID, nil
}

func (s videoService) GetVideo(ctx context.Context, videoID string) (model.Video, error) {
    v, err := s.repo.GetVideo(ctx, videoID)
    if err != nil {
        return model.Video{}, err
    }
    return v, nil
}

