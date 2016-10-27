package swagger2

import (
	"fmt"
	"github.com/emicklei/swagger2/model"
)

func Example() {
	prototype := Annoted{}
	ref, definitions := NewSchemaBuilder().Build(prototype)

	op := model.Operation{}
	op.Produces = []string{"application/json"}
	op.Consumes = []string{"application/json"}
	op.Parameters = []model.Parameter{}
	op.Responses = map[string]*model.Response{"200": &model.Response{Description: "test", Schema: ref}}
	pathItem := model.PathItem{Get: &op}

	swagger := model.NewSwagger()
	modelBuilder := model.NewInfo().Title("test")
	swagger.Definitions(definitions).Host("localhost").Path("/api/v1", &pathItem).Info(modelBuilder)
	fmt.Println(doc(swagger.Build()))
}
