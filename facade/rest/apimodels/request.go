package apimodels

type UpdateGameServersRequest struct {
	Filter    string `json:"filter"`
	JsonPatch []byte `json:"jsonPatch"`
}
