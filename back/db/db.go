package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env 로드 실패")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB 연결 실패:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB Ping 실패:", err)
	}

	log.Println("✅ DB 연결 완료")
}

/*

※ 반드시 규칙!! string 만 null 을 허용!! (datetime도 string으로 처리) 그 외 int / boolean 은 반드시 default 값 주기!

CREATE DATABASE testdb default CHARACTER SET UTF8;

CREATE TABLE IF NOT EXISTS test_table(
    id              INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name            VARCHAR(250) UNIQUE,
	memo            VARCHAR(250),
    age             INT,
    confirm_type    BOOLEAN DEFAULT FALSE,
    created_at      DATETIME DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE test_table ADD CONSTRAINT unique_name UNIQUE (name);
ALTER TABLE test_table MODIFY name VARCHAR(250) UNIQUE;
*/
