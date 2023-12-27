package entity

type Client struct {
	Id        int    `db:"id"`
	ManagerId int    `db:"manager_id"`
	Name      string `db:"name"`
}
