package web

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required, max=255, min=1"`
}