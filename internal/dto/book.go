package dto

type (
	CreateBookRequest struct {
		Title  *string `json:"title"`
		Isbn   *string `json:"isbn"`
		Writer *string `json:"writer"`
	}
	UpdateBookRequest struct {
		Title  *string `json:"title"`
		Isbn   *string `json:"isbn"`
		Writer *string `json:"writer"`
	}
	BookResponse struct {
		ID     uint   `json:"id"`
		Title  string `json:"title"`
		Isbn   string `json:"isbn"`
		Writer string `json:"writer"`
	}
)
