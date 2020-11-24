package interface_manage

import (
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/google/uuid"
)

type IDictService interface {
	GetAll() (result []model_manage.Dict, err error)
	FindByType(dictType string) (result model_manage.Dict, err error)
	FindByTitleOrTypeLike(key string) (result []model_manage.Dict, err error)
	GetDataByType(dictType string) (result []*model_manage.DictData, err error)
	GetDictDataByCondition(condition map[string]string) (result []model_manage.DictData, total int64, err error)
	AddOrEditDictData(targetModel model_manage.DictData) error
	DelDictDataByIds(ids []uuid.UUID) error
	AddOrEditDict(targetModel model_manage.Dict) error
	DelDictByIds(ids []uuid.UUID) error
}
