package db

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func InitDB() {
	dbHost := "localhost"
	dbName := "mydatabase"
	dbUser := "myuser"
	dbPass := "mypassword"
	dbPort := "5444"
	sslmode := "disable"

	// 1. GORM + SQL connection
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)
	sqlDB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Run migrations (если хочешь использовать migrate)
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://internal/db/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	// 3. GORM instance
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = gormDB

	// 4. Добавить колонку `role`, если её нет
	err = DB.Exec(`ALTER TABLE users ADD COLUMN IF NOT EXISTS role VARCHAR(20) NOT NULL DEFAULT 'student';`).Error
	if err != nil {
		log.Println("❌ Не удалось добавить колонку role:", err)
	} else {
		log.Println("✅ Колонка role успешно добавлена (если её не было)")
	}

	err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS courses (
            id SERIAL PRIMARY KEY,
            title TEXT NOT NULL,
            description TEXT
        );
    `).Error
	if err != nil {
		log.Fatal("Не удалось создать таблицу courses:", err)
	}

	// ✅ Вот сюда перемести:
	err = DB.Exec(`ALTER TABLE courses ADD COLUMN IF NOT EXISTS teacher_id INTEGER REFERENCES users(id);`).Error
	if err != nil {
		log.Fatal("❌ Не удалось добавить колонку teacher_id:", err)
	} else {
		log.Println("✅ Колонка teacher_id добавлена (если её не было)")
	}
}
