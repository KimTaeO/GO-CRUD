package presentation

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
