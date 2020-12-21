package service

import "time"

type Store struct {
}

func (s *Store) List() ([]Service, error) {
	return []Service{{ID: 1, Name: "Antispam Daemon", UpdatedAt: time.Now()}}, nil
}

func (s *Store) Get(id int64) (*Service, error) {
	return &Service{ID: 1, Name: "Antispam Daemon", UpdatedAt: time.Now()}, nil
}
