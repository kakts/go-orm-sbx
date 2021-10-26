package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

/* gorm.Modelの定義
type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
  }
*/

// usersテーブル用のデータモデル
type Users struct {
	gorm.Model
	Name string
	Age  uint8
}

// usersテーブル用のデータモデル
type UserTest struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

// GORMの規約に沿わない形でテーブル名を定義したい場合
// Tablerインタフェースを実装することで変更できる
// type Tabler interface {
// 	TableName() string
// }

func (UserTest) TableName() string {
	return "testuser"
}
