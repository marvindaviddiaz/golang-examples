package dbMysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // _ run init method from package
	"log"
)

type User struct {
	id int
	username string
	email string
	enabled string
}

func Main() {
	db, err := sql.Open("mysql", "root:root@/sandbox?parseTime=true")
	if err != nil {
		log.Println("Error", err)
	}
	defer db.Close()

	query(db)

	update(db)


}
func update(db *sql.DB) {
	res, err := db.Exec("update user set enabled = ? ", "Y")
	if err != nil {
		log.Fatal(err)
	}
	ra, _ := res.RowsAffected()
	id, _ := res.LastInsertId()

	log.Println("Rows affected:", ra, "Las inserted id:", id)

}

func query(db *sql.DB) {

	rows, err := db.Query("SELECT id, username, email, enabled from user where enabled = ?", "Y")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	log.Println("Columns detected:", cols)
	for rows.Next() {
		member := User{}
		err = rows.Scan(&member.id, &member.username, &member.email, &member.enabled)
		if err != nil {
			log.Fatal("Error scanning row", err)
		}
		log.Println(member)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}