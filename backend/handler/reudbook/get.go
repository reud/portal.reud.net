package reudbook

import (
	"backend/database"
	"backend/gen/models"
	"backend/gen/restapi/operations/bookshelf"
	"github.com/go-openapi/runtime/middleware"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetReudBook(params bookshelf.GetReudBookParams) middleware.Responder {
	db := database.NewDatabase()
	books := db.FetchAll()
	if books == nil {
		return bookshelf.NewGetReudBookOK().WithPayload([]*models.Book{})
	}

	var ret []*models.Book
	for _, v := range *books {
		ret = append(ret, &v)
	}
	db.Close()

	return bookshelf.NewGetReudBookOK().WithPayload(ret)
}
