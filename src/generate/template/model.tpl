<%: func Model(pkgCode string, model string, fieldTags map[string]string, buffer *bytes.Buffer) %>
package model_<%==s pkgCode%>

import (
	"github.com/anden007/afocus-godf/src/model"
)

func init() {
	tableName := <%==s model%>{}.TableName()

	model.ModelDBFields[tableName] = make(map[string]string)
<%for tag, dbFieldName := range fieldTags { %>
	model.ModelDBFields[tableName]["<%==s tag%>"] = "<%==s dbFieldName%>"
<% } %>
}