package request

type Todo struct {
	Title  string `json:"title"`
	Status string `json:"status"`
	UserId int    `json:"userId"`
}
