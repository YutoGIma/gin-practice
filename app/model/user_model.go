package model

type User struct {
	BaseModel        // 共通フィールドを埋め込み
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"uniqueIndex"`
	Password  string `json:"password"`
}
