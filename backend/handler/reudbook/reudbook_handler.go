package reudbook

import (
	"backend/gen/restapi/operations/bookshelf"
)

func AddReudBookHandler() bookshelf.AddReudBookHandler {
	return bookshelf.AddReudBookHandlerFunc(AddReudBook)
}

func DeleteReudBookHandler() bookshelf.DeleteReudBookHandler {
	return bookshelf.DeleteReudBookHandlerFunc(DeleteReudBook)
}

func GetReudBookHandler() bookshelf.GetReudBookHandler {
	return bookshelf.GetReudBookHandlerFunc(GetReudBook)
}
