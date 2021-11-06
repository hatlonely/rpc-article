package article

import (
	"time"
)

type Author struct {
	ID   string `gorm:"type:char(32);primary_key" json:"id"`
	Name string `gorm:"type:char(32);not null" json:"name"`
}

type Article struct {
	ID       string    `gorm:"type:char(32);primary_key" json:"id"`
	AuthorID int       `gorm:"type:bigint(20);column:authorID;unique_index:author_title_idx;default:0" json:"authorID,omitempty"`
	Title    string    `gorm:"type:varchar(128);unique_index:author_title_idx;not null" json:"title,omitempty"`
	Tags     string    `gorm:"type:varchar(256);default:''" json:"tags,omitempty"`
	Brief    string    `gorm:"type:varchar(256);default:''" json:"brief,omitempty"`
	Content  string    `gorm:"type:longtext COLLATE utf8mb4_unicode_520_ci;not null" json:"content,omitempty"`
	CreateAt time.Time `gorm:"type:timestamp;column:createAt;not null;default:CURRENT_TIMESTAMP;index:create_at_idx" json:"createAt,omitempty"`
	UpdateAt time.Time `gorm:"type:timestamp;column:updateAt;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index:update_at_idx" json:"updateAt,omitempty"`
}
