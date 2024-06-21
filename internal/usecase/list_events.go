package usecase

import (
	"github.com/medinaalexandre/imersao-fullcycle-desafio02/internal/domain"
)

type ListEventsUseCase struct {
	repo domain.EventRepository
}

func NewListEventsUseCase(repo domain.EventRepository) *ListEventsUseCase {
	return &ListEventsUseCase{repo: repo}
}

func (uc *ListEventsUseCase) Execute() []domain.Event {
	return uc.repo.ListEvents()
}
