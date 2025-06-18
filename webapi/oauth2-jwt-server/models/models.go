package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}

type Client struct {
	ID     string `gorm:"primaryKey"`
	Secret string
	Domain string
}
