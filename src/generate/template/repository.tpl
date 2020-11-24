<%: func Repository(pkgCode string, model string, hasCreateTime bool, buffer *bytes.Buffer) %>
package repository_<%==s pkgCode%>

import (
	"github.com/anden007/afocus-godf/src/interfaces"
	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_<%==s pkgCode%>"
	"github.com/google/uuid"
<%if hasCreateTime{%>
	"time"
<%}%>
)

type <%==s model%>Repository struct{}

func New<%==s model%>Repository() *<%==s model%>Repository {
	instance := new(<%==s model%>Repository)
	return instance
}

func (m *<%==s model%>Repository) Add(entity model_<%==s pkgCode%>.<%==s model%>) (err error) {
	db := interfaces.DI().GetDataBase()
<%if hasCreateTime{%>
    if entity.CreateTime.IsZero(){
        entity.CreateTime = time.Now()
    }
<%}%>
	err = db.GetDB().Create(&entity).Error
	return
}

func (m *<%==s model%>Repository) Edit(entity model_<%==s pkgCode%>.<%==s model%>) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Save(&entity).Error
	return
}

func (m *<%==s model%>Repository) DelByIds(ids []uuid.UUID ) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Delete(model_<%==s pkgCode%>.<%==s model%>{}, "id in (?)", ids).Error
	return
}

func (m *<%==s model%>Repository) GetById(id uuid.UUID) (result model_<%==s pkgCode%>.<%==s model%>, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_<%==s pkgCode%>.<%==s model%>{}).Where(&model_<%==s pkgCode%>.<%==s model%>{Id:id}).First(&result).Error
	return
}

func (m *<%==s model%>Repository) GetAll() (result []model_<%==s pkgCode%>.<%==s model%>, err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_<%==s pkgCode%>.<%==s model%>{}).Find(&result).Error
	return
}

func (m *<%==s model%>Repository) Updates(id uuid.UUID , fieldValues map[string]interface{}) (err error) {
	db := interfaces.DI().GetDataBase()
	err = db.GetDB().Model(&model_<%==s pkgCode%>.<%==s model%>{}).Where(model_<%==s pkgCode%>.<%==s model%>{Id: id}).Updates(fieldValues).Error
	return
}

func (m *<%==s model%>Repository) GetByCondition(condition map[string]string) (result []model_<%==s pkgCode%>.<%==s model%>, total int64, err error) {
	db := interfaces.DI().GetDataBase()
	Query := db.GetDB().Model(model_<%==s pkgCode%>.<%==s model%>{})
	CountQuery := db.GetDB().Model(model_<%==s pkgCode%>.<%==s model%>{})
	err = lib.NewQueryCondition().GetQuery(model_<%==s pkgCode%>.<%==s model%>{}.TableName(), condition, Query, CountQuery)
	if err == nil {
		err = CountQuery.Count(&total).Error
		err = Query.Find(&result).Error
	}
	return
}