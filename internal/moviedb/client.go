package moviedb

import (
	"fmt"
	"net/http"
)

type Api struct {
	token   string
	session string
}

const baseUri = "https://api.themoviedb.org/3/"

func Init(token string) *Api {
	return &Api{token: token}
}

func (a *Api) AddMovieToList(movieId, listId int) error {
	payload := fmt.Sprintf(`{"media_id": %v}`, movieId)
	addListEndpoint := fmt.Sprintf("%s/list/%v/add_item", baseUri, listId)
	req, err := generateHttpRequest(http.MethodPost, addListEndpoint, a.token, []byte(payload))
	if err != nil {
		return err
	}

	resp, err := makeHttpRequest(req)
	if err != nil {
		return err
	}

	fmt.Printf("response: %v\n", string(resp))
	return nil
}
