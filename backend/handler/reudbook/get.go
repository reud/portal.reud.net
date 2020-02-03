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
		ret = append(ret, &models.Book{
			ID:              v.ID,
			Href:            v.Href,
			IrjpImageSource: v.IrjpImageSource,
			WsfeImageSource: v.WsfeImageSource,
			Title:           v.Title,
			Tag1:            v.Tag1,
			Tag2:            v.Tag2,
			Tag3:            v.Tag3,
		})
	}
	db.Close()

	return bookshelf.NewGetReudBookOK().WithPayload(ret)
}
