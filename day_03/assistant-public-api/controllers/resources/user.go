package resources

type CreateUserRequest struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	ManagerID int64  `json:"manager_id"`
}

type UpdateUserRequest struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	ManagerID int64  `json:"manager_id"`
}

type ListUserResponse struct {
	Users []*User `json:"users"`
}
