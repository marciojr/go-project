package request

type UserRequest struct {
	ID       string
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=30,containsany=!@#$%¨&*0123456789"`
	Name     string `json:"name" binding:"required,min=4,max=50"`
	Age      int8   `json:"age" binding:"required,min=1,max=100"`
}
