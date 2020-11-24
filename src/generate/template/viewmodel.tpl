<%: func ViewModel(pkgCode string, model string, buffer *bytes.Buffer) %>
package view_model_<%==s pkgCode%>

import (
	"github.com/anden007/afocus-godf/src/model/model_<%==s pkgCode%>"
	"github.com/anden007/afocus-godf/src/web/view_model"
)

type VM_<%==s model%> struct {
	view_model.BaseViewModel
	model_<%==s pkgCode%>.<%==s model%>
}
