package users

import (
	"io"
	"log"
	"net/http"
)

// Add user to pool for search match
func (i *Implementation) SearchMatch(writer http.ResponseWriter, request *http.Request) {

	ctx := request.Context()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("read body request is failed")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = i.userService.SearchMatch(ctx, body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)

}
