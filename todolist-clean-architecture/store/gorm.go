package store

import (
	"github.com/montheankul-k/todolist-clean-architecture/todo"
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{db: db}
}

func (s *GormStore) New(todo *todo.Todo) error {
	return s.db.Create(todo).Error
}

func (s *GormStore) List(todo *[]todo.Todo) error {
	return s.db.Find(todo).Error
}

func (s *GormStore) Remove(todo *todo.Todo, id int) error {
	return s.db.Delete(todo, id).Error
}
