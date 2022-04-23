package request

type CategoryUpdateRequest struct {
	Name string `validate:"required"`
}
