package reudbook

import (
	"backend/gen/restapi/operations/bookshelf"
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
)

func AddReudBook(params bookshelf.AddReudBookParams, jwt *map[string]*json.RawMessage) middleware.Responder {
	return middleware.NotImplemented("operation bookshelf.AddReudBook has not yet been implemented")
}
