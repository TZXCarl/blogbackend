package port

import (
	_ "github.com/go-sql-driver/mysql"
	"file/utils"
)

func InsertFile(path, fileName, type, createDate string) bool {
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()

	stmt, _ := db.Prepare(`INSERT  admin.files (path, fileName, type, createDate) VALUES (?,?,?,?)`)
	row, err := stmt.Exec(path, fileName, type, createDate)
	if err != nil {
		utils.HandleServerError(w, err);
		return false;
	}

	return true;
}