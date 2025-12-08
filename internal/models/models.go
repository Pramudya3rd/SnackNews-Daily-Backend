package models

import "time"

type News struct {
    ID        string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
    Title          string    `json:"title"`
    Content        string    `json:"content"`
    Author         string    `json:"author"`
    Category       string    `json:"category" gorm:"default:'General'"`
    Image          string    `json:"image"`
    SourceURL      string    `json:"sourceUrl"`
    DisplaySection string    `json:"displaySection"`
    Archived       bool      `json:"archived" gorm:"default:false"`
    CreatedAt      time.Time `json:"createdAt" gorm:"autoCreateTime"`
    UpdatedAt      time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type Category struct {
    ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
    Name string `json:"name" gorm:"unique;not null"`
}

type User struct {
    ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
    Username string `json:"username" gorm:"unique;not null"`
    Password string `json:"-" gorm:"not null"`
}