package requests

type InsertClientRequest struct {
	ManagerId int    `json:"manager_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

type UpdateClientRequest struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DeleteClientRequest struct {
	Id int `json:"id" binding:"required"`
}
