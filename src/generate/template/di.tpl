<%: func DI(pkgCode string, model string, buffer *bytes.Buffer) %>
// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/anden007/afocus-godf/src/interfaces"
	"github.com/anden007/afocus-godf/src/interfaces/interface_<%==s pkgCode%>"
	"github.com/anden007/afocus-godf/src/repositories/repository_<%==s pkgCode%>"
	"github.com/anden007/afocus-godf/src/services/service_<%==s pkgCode%>"
	"github.com/google/wire"
)

var <%==s model%>Service_Set = wire.NewSet(service_<%==s pkgCode%>.New<%==s model%>Service, LruCache_SingleSet, repository_<%==s pkgCode%>.New<%==s model%>Repository, wire.Bind(new(interface_<%==s pkgCode%>.I<%==s model%>Service), new(*service_<%==s pkgCode%>.<%==s model%>Service)))

func Get<%==s model%>Service() interface_<%==s pkgCode%>.I<%==s model%>Service {
	panic(wire.Build(<%==s model%>Service_Set))
}

func init() {
	container := interfaces.DI()
	container.Get<%==s model%>Service = Get<%==s model%>Service
}
