package dto

type GetUserRequest struct {
	User_Id   string `json:"user_id"`
	User_Type string `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
}
