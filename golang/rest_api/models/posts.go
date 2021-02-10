package models

type Post struct {
	ID      int64  `gorm:"unique_index;not null" json:"id"`
	Title   string `gorm:"type:varchar(40)" json:"title"`
	Content string `gorm:"type:varchar(255)" json:"content"`
}
