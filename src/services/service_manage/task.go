package service_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/repositories/repository_manage"
	"github.com/anden007/afocus-godf/src/services/service_core"
	"github.com/google/uuid"
)

type TaskService struct {
	lruCache *service_core.LruCache
	repo *repository_manage.TaskRepository
}

func NewTaskService(lruCache *service_core.LruCache, repo *repository_manage.TaskRepository) *TaskService {
	instance := &TaskService{
	    lruCache: lruCache,
		repo: repo,
	}
	return instance
}

func (m *TaskService) GetById(id uuid.UUID) (result model_manage.Task, err error) {
	result, err = m.repo.GetById(id)
	return
}

func (m *TaskService) GetAll() (result []model_manage.Task, err error) {
	result, err = m.repo.GetAll()
	return
}

func (m *TaskService) Add(entity model_manage.Task) (err error) {
	err = m.repo.Add(entity)
	return
}

func (m *TaskService) Edit(entity model_manage.Task) (err error) {
	err = m.repo.Edit(entity)
	return
}

func (m *TaskService) DelByIds(ids []uuid.UUID) (err error) {
	err = m.repo.DelByIds(ids)
	return
}


func (m *TaskService) Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error) {
	err = m.repo.Updates(id, fieldValues)
	return
}


func (m *TaskService) GetByCondition(condition map[string]string) (result []model_manage.Task, total int64, err error) {
	result, total, err = m.repo.GetByCondition(condition)
	return
}