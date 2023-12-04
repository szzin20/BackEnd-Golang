package constanta

const (
	ErrInvalidBody        = "invalid request body"
	ErrImageFileRequired  = "image file required"
	ErrInvalidImageFormat = "invalid image file format. only jpg, jpeg, png allowed"
	ErrInvalidIDParam     = "invalid id param"
	ErrInvalidParam       = " param not valid"
	ErrQueryParamRequired = " query param required"
	ErrNotFound           = "not found"
	ErrActionGet          = "failed to get "
	ErrActionCreated      = "failed to create "
	ErrActionUpdated      = "failed to update "
	ErrActionDeleted      = "failed to delete "
)

const (
	SuccessActionGet     = "successfully get data "
	SuccessActionCreated = "successfully created "
	SuccessActionUpdated = "successfully updated "
	SuccessActionDeleted = "successfully deleted "
)
