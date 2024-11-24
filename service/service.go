package service

import (
	"goproject_port/datastruct"
)

type repo interface {
	GetIn(nIn int) (int, error)
	GetChIn(nIn int) (chan datastruct.Reqest, error)
	PostOut(nOut int, newValue int) error
	PostChOut(nOut int) (chan datastruct.Reqest, error)

	GetAll()
}

type service struct {
	Repo repo
}

func New(repo repo) *service {
	return &service{Repo: repo}
}

func (s *service) GetIn(nIn int) (int, error) {
	return s.Repo.GetIn(nIn)
}

func (s *service) GetChIn(nIn int) (chan datastruct.Reqest, error) {
	return s.Repo.GetChIn(nIn)
}

func (s *service) PostOut(nOut int, newValue int) error {
	return s.Repo.PostOut(nOut, newValue)
}

func (s *service) PostChOut(nOut int) (chan datastruct.Reqest, error) {
	return s.Repo.PostChOut(nOut)
}
//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func (s *service) GetAll() {
	s.Repo.GetAll()
}
