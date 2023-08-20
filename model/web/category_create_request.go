package web

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required,max=200,min=1"`
}
