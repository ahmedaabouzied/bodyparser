// Parses json reqeust body
package bodyParser

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Parse Jsob Request Body
//
// Returns a Map of String keys and Interface values
//
func JsonParse(r http.Response) (map[string]interface{}, error) {
	// Read the r.body into a byte array
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// Make a map of String keys and Interface Values
	b := make(map[string]interface{})
	// Unmarshal the body array into the map
	err = json.Unmarshal(body, &b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
