package requests

type InsertManagerRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateManagerRequest struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DeleteManagerRequest struct {
	Id int `json:"id" binding:"required"`
}
