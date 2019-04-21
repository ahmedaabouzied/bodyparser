package bodyParser

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func JsonParse(r http.Response) (map[string]interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	b := make(map[string]interface{})
	err = json.Unmarshal(body, &b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
