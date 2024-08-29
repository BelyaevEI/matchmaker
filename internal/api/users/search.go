package users

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// AddUserToPool Add user to pool for search match
func (i *Implementation) AddUserToPool(writer http.ResponseWriter, request *http.Request) {

	ctx := request.Context()
	fmt.Print("Hello world")
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("read body request is failed")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = i.userService.AddUserToPool(ctx, body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)

}
