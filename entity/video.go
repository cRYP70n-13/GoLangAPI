package entity

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Email     string `json:"email" validate:"required,email"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
}

type Video struct {
	Title       string `json:"title" binding:"min=2,max=20"`
	Description string `json:"description" binding:"max=30"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
