package model

// User defines model for User.
type User struct {
	Age       int    `db:"age" json:"age"`
	CreatedAt string `db:"created_at" json:"createdAt"`
	Faculty   int    `db:"faculty" json:"faculty"`
	Gender    int    `db:"gender" json:"gender"`
	Id        int    `db:"id" json:"id"`
	Password  string `db:"password" json:"password"`
	Mail      string `db:"mail" json:"mail"`
	Name      string `db:"name" json:"name"`
	UpdatedAt string `db:"updated_at" json:"updatedAt"`
}
