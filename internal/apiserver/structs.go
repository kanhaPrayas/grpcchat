package apiserver

type PostParams struct {
	Name     string `json:"name"`
	RoomName string `json:"name"`
	Message  string `json:"message"`
}
