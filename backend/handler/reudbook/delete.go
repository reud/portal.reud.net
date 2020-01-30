package reudbook

import (
	"backend/gen/restapi/operations/bookshelf"
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
)

func DeleteReudBook(params bookshelf.DeleteReudBookParams, jwt *map[string]*json.RawMessage) middleware.Responder {
	return middleware.NotImplemented("operation bookshelf.DeleteReudBook has not yet been implemented")
}
