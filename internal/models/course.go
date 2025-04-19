package models

type Course struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TeacherID   *uint  `json:"teacher_id"`
}
