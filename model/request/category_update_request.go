package request

type CategoryUpdateRequest struct {
	ID   int
	Name string `validate:"required"`
}
