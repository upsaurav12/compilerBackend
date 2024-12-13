package handlers

import (
	"encoding/json"
	"net/http"
	"online-compiler/executor"
)

type RequestPayLoad struct {
	Language string `json:"language"`
	Code     string `json:"code"`
	Input    string `json:"output"`
}

type ResponsePayLoad struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}

func HandleExecute(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello from HandleExecute!")

	var payload RequestPayLoad

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid Request payload", http.StatusBadRequest)
		return
	}

	output, execErr := executor.Execute(payload.Language, []byte(payload.Code), payload.Input)

	response := ResponsePayLoad{
		Output: output,
		Error:  execErr,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}
