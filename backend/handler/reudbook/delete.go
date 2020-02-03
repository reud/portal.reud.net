package reudbook

import (
	"backend/database"
	"backend/gen/restapi/operations/bookshelf"
	"encoding/json"
	"github.com/go-openapi/runtime/middleware"
)

func DeleteReudBook(params bookshelf.DeleteReudBookParams, jwt *map[string]*json.RawMessage) middleware.Responder {
	db := database.NewDatabase()
	db.Delete(params.BookID)
	db.Close()

	return bookshelf.NewDeleteReudBookNoContent()
}
