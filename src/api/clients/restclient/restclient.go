package restclient

import (
	"encoding/json"
	"net/http"
)

func Post(url string, body interface{}) error {
	client := http.Client{}

	bytes, err := json.Marshal(body)

	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, url)
}
