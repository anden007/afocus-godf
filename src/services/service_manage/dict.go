package service_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/repositories/repository_manage"

	"github.com/google/uuid"
)

type DictService struct {
	repo *repository_manage.DictRepository
}

func NewDictService(repo *repository_manage.DictRepository) *DictService {
	instance := &DictService{
		repo: repo,
	}
	return instance
}

func (m *DictService) GetAll() (result []model_manage.Dict, err error) {
	result, err = m.repo.GetAll()
	return
}
func (m *DictService) FindByType(dictType string) (result model_manage.Dict, err error) {
	result, err = m.repo.FindByType(dictType)
	return
}
func (m *DictService) FindByTitleOrTypeLike(key string) (result []model_manage.Dict, err error) {
	result, err = m.repo.FindByTitleOrTypeLike(key)
	return
}

func (m *DictService) GetDataByType(dictType string) (result []*model_manage.DictData, err error) {
	result, err = m.repo.GetDataByType(dictType)
	return
}

func (m *DictService) GetDictDataByCondition(condition map[string]string) (result []model_manage.DictData, total int64, err error) {
	result, total, err = m.repo.GetDictDataByCondition(condition)
	return
}

func (m *DictService) AddOrEditDictData(targetModel model_manage.DictData) error {
	return m.repo.AddOrEditDictData(targetModel)
}
func (m *DictService) DelDictDataByIds(ids []uuid.UUID) error {
	return m.repo.DelDictDataByIds(ids)
}
func (m *DictService) AddOrEditDict(targetModel model_manage.Dict) error {
	return m.repo.AddOrEditDict(targetModel)
}
func (m *DictService) DelDictByIds(ids []uuid.UUID) error {
	return m.repo.DelDictByIds(ids)
}
