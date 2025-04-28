package todo

type Todo struct {
	ID     int    `json:"id"`
	UserID int    `json:"-"`
	Title  string `json:"title"`
	Done   string `json:"done"`
}
