package service

import (
	"github.com/raman-vhd/video-converter/internal/lib"
	"github.com/raman-vhd/video-converter/internal/repository"
)

type ITemplateService interface{
    Action()
}

type templateService struct {
    repo repository.ITemplateRepository
    amqp *lib.AMQP
}

func NewTemplate(
    repo repository.ITemplateRepository,
    amqp *lib.AMQP,
) ITemplateService{
    return templateService{
        repo: repo,
        amqp: amqp,
    }
}

func (s templateService) Action() {
    return
}
