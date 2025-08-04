package dto

type TaskRequestById struct {
	ID string `json:"id" validate:"required"`
}

type TaskRequestUpdate struct {
	ID          string `json:"id" validate:"required"`
	Title       string `json:"title" validate:"omitempty,min=3,max=100"`
	Description string `json:"description" validate:"omitempty,min=5,max=500"`
	Status      string `json:"status" validate:"omitempty,oneof=pending in_progress completed"`
}

type TaskRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"oneof=new in_progress done"`
}

type TaskStatusUpdateRequest struct {
	ID     string `json:"id" validate:"required"`
	Status string `json:"status" validate:"required,oneof=new in_progress done"`
}
