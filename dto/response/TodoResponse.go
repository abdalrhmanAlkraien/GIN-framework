package response

type TodoResponse struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
