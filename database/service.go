package database

import (
	"github.com/DmitriyPrischep/Course-Project-of-Data-Base/models"
)

const (
	getStatusSQL = `
		SELECT 
		(SELECT COUNT(*) FROM users) AS users,
		(SELECT COUNT(*) FROM forums) AS forums,
		(SELECT COUNT(*) FROM posts) AS posts,
		(SELECT COALESCE(SUM(threads), 0) FROM forums WHERE threads > 0) AS threads
	`
	clearSQL = `
		TRUNCATE users, forums, threads, posts, votes;
	`
)

// /service/status Получение инфомарции о базе данных
func GetStatusDB() *models.Status {
	status := &models.Status{}
	DB.pool.QueryRow(
		getStatusSQL,
	).Scan(
		&status.User,
		&status.Forum,
		&status.Post,
		&status.Thread,
	)
	return status
}

// /service/clear Очистка всех данных в базе
func ClearDB() {
	DB.pool.Exec(clearSQL)
}
