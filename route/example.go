package route

import (
	"encoding/json"
	"net/http"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, _ := json.Marshal([]string{"uno", "due", "tre"})
	w.Write(data)
}
