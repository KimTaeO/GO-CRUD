package presentation

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}