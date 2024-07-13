package dao

import "log"

func GetUserNameById(id int) string {
	var userName string
	row := DB.QueryRow("select * from blog_user where uid = ?", id)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	_ = row.Scan(&userName)
	return userName
}
