package model

import (
	"time"

	domain "github.com/sakaguchi-0725/echo-onion-arch/domain/model"
)

type Book struct {
	ID         string `gorm:"primaryKey;not null"`
	Title      string `gorm:"not null"`
	Author     string
	CategoryID uint      `gorm:"not null"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
	Status     string    `gorm:"type:enum('available', 'loaned', 'reserved') not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func ToModelBook(b domain.Book) Book {
	return Book{
		ID:         b.ID.String(),
		Title:      b.Title,
		Author:     b.Author,
		CategoryID: b.CategoryID,
		Status:     b.Status.String(),
	}
}

func ToDomainBook(b Book) domain.Book {
	return domain.RecreateBook(domain.BookID(b.ID), b.Title, b.Author, b.CategoryID, domain.BookStatus(b.Status))
}
