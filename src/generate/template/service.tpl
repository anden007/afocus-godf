<%: func Service(pkgCode string, model string, buffer *bytes.Buffer) %>
package service_<%==s pkgCode%>

import (
	"github.com/anden007/afocus-godf/src/model/model_<%==s pkgCode%>"
	"github.com/anden007/afocus-godf/src/repositories/repository_<%==s pkgCode%>"
	"github.com/anden007/afocus-godf/src/services/service_core"
	"github.com/google/uuid"
)

type <%==s model%>Service struct {
	lruCache *service_core.LruCache
	repo *repository_<%==s pkgCode%>.<%==s model%>Repository
}

func New<%==s model%>Service(lruCache *service_core.LruCache, repo *repository_<%==s pkgCode%>.<%==s model%>Repository) *<%==s model%>Service {
	instance := &<%==s model%>Service{
	    lruCache: lruCache,
		repo: repo,
	}
	return instance
}

func (m *<%==s model%>Service) GetById(id uuid.UUID) (result model_<%==s pkgCode%>.<%==s model%>, err error) {
	result, err = m.repo.GetById(id)
	return
}

func (m *<%==s model%>Service) GetAll() (result []model_<%==s pkgCode%>.<%==s model%>, err error) {
	result, err = m.repo.GetAll()
	return
}

func (m *<%==s model%>Service) Add(entity model_<%==s pkgCode%>.<%==s model%>) (err error) {
	err = m.repo.Add(entity)
	return
}

func (m *<%==s model%>Service) Edit(entity model_<%==s pkgCode%>.<%==s model%>) (err error) {
	err = m.repo.Edit(entity)
	return
}

func (m *<%==s model%>Service) DelByIds(ids []uuid.UUID) (err error) {
	err = m.repo.DelByIds(ids)
	return
}


func (m *<%==s model%>Service) Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error) {
	err = m.repo.Updates(id, fieldValues)
	return
}


func (m *<%==s model%>Service) GetByCondition(condition map[string]string) (result []model_<%==s pkgCode%>.<%==s model%>, total int64, err error) {
	result, total, err = m.repo.GetByCondition(condition)
	return
}