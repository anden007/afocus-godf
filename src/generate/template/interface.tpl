<%: func InterFace(pkgCode string, model string, buffer *bytes.Buffer) %>
package interface_<%==s pkgCode%>

import (
	"github.com/anden007/afocus-godf/src/model/model_<%==s pkgCode%>"
	"github.com/google/uuid"
)

type I<%==s model%>Service interface {
	Add(entity model_<%==s pkgCode%>.<%==s model%>) (err error)
	DelByIds(ids []uuid.UUID ) (err error)
	Edit(entity model_<%==s pkgCode%>.<%==s model%>) (err error)
	Updates(id uuid.UUID , fieldValues map[string]interface{}) (err error)
	GetAll() (result []model_<%==s pkgCode%>.<%==s model%>, err error)
	GetById(id uuid.UUID) (result model_<%==s pkgCode%>.<%==s model%>, err error)
	GetByCondition(condition map[string]string) (result []model_<%==s pkgCode%>.<%==s model%>, total int64, err error)
}
