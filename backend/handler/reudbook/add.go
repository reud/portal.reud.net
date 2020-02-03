package reudbook

import (
	"backend/database"
	"backend/gen/restapi/operations/bookshelf"
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
)

func AddReudBook(params bookshelf.AddReudBookParams, jwt *map[string]*json.RawMessage) middleware.Responder {
	db := database.NewDatabase()
	db.Create(params.Body)
	db.Close()

	return bookshelf.NewAddReudBookNoContent()
}
