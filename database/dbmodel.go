package database

func GetUserById(id uint64) map[string]interface{} {
	response := RequestQuery("SELECT * FROM users WHERE id=?", id)
	var ret map[string]interface{}
	if response == nil {
		return ret
	}

	response.Next()

	var uid uint64
	var username string
	response.Scan(&uid, &username)
	return map[string]interface{}{
		"id":       uid,
		"username": username,
	}
}

func GetTaskById(id uint64) map[string]interface{} {
	response := RequestQuery("SELECT * FROM tasks WHERE id=?", id)
	var ret map[string]interface{}
	if response == nil {
		return ret
	}

	response.Next()

	var uid uint64
	var userid uint64
	var tasktypeid uint64
	response.Scan(&uid, &userid, &tasktypeid)
	return map[string]interface{}{
		"id":       uid,
		"owner_id": userid,
		"type_id":  tasktypeid,
	}
}

func GetTasksByUserId(userId uint64) []interface{} {
	response := RequestQuery("SELECT * FROM tasks WHERE user_id=?", userId)
	var ret []interface{} = nil
	if response == nil {
		return ret
	}

	var uid uint64
	var ownerId uint64
	var taskId uint64
	for response.Next() {
		response.Scan(&uid, &ownerId, &taskId)
		ret = append(ret, map[string]interface{}{
			"id":       uid,
			"owner_id": ownerId,
			"type_id":  taskId,
		})
	}

	return ret
}
