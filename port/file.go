package port

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "fmt"
	"file/domain"
)

func InsertFile(id string, path string, fileName string, _type string, createDate int64) (bool, error) {
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()

	stmt, _ := db.Prepare(`INSERT admin.files (id, path, fileName, type, createDate) VALUES (?,?,?,?,?)`)
	_, err = stmt.Exec(id, path, fileName, _type, createDate,)
	if err != nil {
		return false, err;
	}

	return true, nil;
}


func GetFiles(offset, end int) ([]*domain.File, int, error) {
	db, err := sql.Open("mysql", "abelce:Tzx_301214@tcp(111.231.192.70:3306)/admin?parseTime=true")
	defer db.Close()

	files := []*domain.File{}
	var count int = 0;
	stmt, err := db.Prepare(`SELECT * FROM admin.files ORDER BY createDate desc LIMIT ?,?`)
	rows, err := stmt.Query(offset, end)

	defer rows.Close();  // 防止for循环报错时，不会自动关闭连接
	if err != nil {
		return files, 0, err
	}
	for rows.Next() {
		_file := domain.File{}
		rows.Scan(
			&_file.ID,
			&_file.Path,
			&_file.FileName,
			&_file.Type,
			&_file.CreateDate,
		)
		files = append(files, &_file)
	}
	stmt, err = db.Prepare(`SELECT count(*) FROM admin.files`)
	row := stmt.QueryRow();
	row.Scan(&count)

	return files, count, nil
}