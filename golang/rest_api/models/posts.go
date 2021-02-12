package models

// swagger:model post
type Post struct {
	// the ID of the post
	//
	// required: false
	ID int64 `gorm:"unique_index;not null" json:"id"`
	// the title of the post.
	//
	// required: true
	// max length: 40
	Title string `gorm:"type:varchar(40)" json:"title"`
	// the content of the post.
	//
	// required: true
	// max length: 255
	Content string `gorm:"type:varchar(255)" json:"content"`
}
