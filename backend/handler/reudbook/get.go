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
		return bookshelf.NewGetReudBookOK().WithPayload([]*models.StoredBook{})
	}

	var ret []*models.StoredBook
	for _, v := range *books {
		ret = append(ret,&v)
	}

	return bookshelf.NewGetReudBookOK().WithPayload(ret)
}
