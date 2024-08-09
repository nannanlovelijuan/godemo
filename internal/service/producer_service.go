package service

import "gitlab.ezrpro.in/godemo/internal/repo"

type producerService struct {
	producerRepo repo.IProducerRepo
}

func NewProducerService(producerRepo repo.IProducerRepo) IProducerService {
	return &producerService{
		producerRepo: producerRepo,
	}
}

func (p *producerService) Send(topic, msg string) (partition int32, offset int64, err error) {
	return p.producerRepo.Send(topic, msg)
}
