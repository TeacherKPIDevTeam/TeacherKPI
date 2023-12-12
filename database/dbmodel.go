package database

func GetUserById(id uint64) (uint64, string) {
	response := RequestQuery("SELECT * FROM users WHERE id=?", id)
	response.Next()

	var uid uint64
	var username string
	response.Scan(&uid, &username)
	return uid, username
}
