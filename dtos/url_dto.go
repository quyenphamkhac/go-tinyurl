package dtos

type CreateURLDto struct {
	OriginalURL string `json:"original_url" binding:"required,url"`
}
