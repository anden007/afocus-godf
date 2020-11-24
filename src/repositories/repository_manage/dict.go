package repository_manage

import (
	"fmt"
	"sort"

	"github.com/anden007/afocus-godf/src/interfaces"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_manage"

	"github.com/google/uuid"
)

type DictRepository struct{}

func NewDictRepository() *DictRepository {
	instance := new(DictRepository)
	return instance
}

func (m *DictRepository) GetAll() (result []model_manage.Dict, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Find(&result).Error
	return
}
func (m *DictRepository) FindByType(dictType string) (result model_manage.Dict, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_manage.Dict{}).Where("type = ?", dictType).Find(&result).Error
	return
}
func (m *DictRepository) FindByTitleOrTypeLike(key string) (result []model_manage.Dict, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Where("title like ?", fmt.Sprint("%", key, "%")).Or("type like ?", fmt.Sprint("%", key, "%")).Find(&result).Error
	return
}

func (m *DictRepository) GetDataByType(dictType string) (result []*model_manage.DictData, err error) {
	db := interfaces.DI().GetDataBase()
	var dict model_manage.Dict
	err = db.GetDB().Model(&model_manage.Dict{}).Where("type = ?", dictType).Preload("Data").First(&dict).Error
	if err == nil {
		tmpArray := model_manage.DictDataArray{}
		for _, item := range dict.Data {
			tmpArray = append(tmpArray, item)
		}
		sort.Stable(tmpArray)
		result = tmpArray
	}
	return
}

func (m *DictRepository) GetDictDataByCondition(condition map[string]string) (result []model_manage.DictData, total int64, err error) {
	db := interfaces.DI().GetDataBase()
	Query := db.GetDB().Model(model_manage.DictData{})
	CountQuery := db.GetDB().Model(model_manage.DictData{})
	err = lib.NewQueryCondition().GetQuery(model_manage.DictData{}.TableName(), condition, Query, CountQuery)
	if err == nil {
		err = CountQuery.Count(&total).Error
		err = Query.Find(&result).Error
	}
	return
}

func (m *DictRepository) AddOrEditDictData(targetModel model_manage.DictData) error {
	db := interfaces.DI().GetDataBase()
	return db.GetDB().Save(&targetModel).Error
}

func (m *DictRepository) DelDictDataByIds(ids []uuid.UUID) error {
	db := interfaces.DI().GetDataBase()
	return db.GetDB().Delete(model_manage.DictData{}, "id in (?)", ids).Error
}

func (m *DictRepository) AddOrEditDict(targetModel model_manage.Dict) error {
	db := interfaces.DI().GetDataBase()
	return db.GetDB().Save(&targetModel).Error
}

func (m *DictRepository) DelDictByIds(ids []uuid.UUID) error {
	db := interfaces.DI().GetDataBase()
	return db.GetDB().Delete(model_manage.Dict{}, "id in (?)", ids).Error
}
