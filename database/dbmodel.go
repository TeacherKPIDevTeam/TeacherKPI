package database

func GetUserById(id uint64) map[string]interface{} {
	response := RequestQuery("SELECT * FROM users WHERE id=?", id)
	response.Next()

	var uid uint64
	var username string
	if err := response.Scan(&uid, &username); err != nil {
		return nil
	}
	return map[string]interface{}{
		"id":       uid,
		"username": username,
	}
}

func GetTaskById(id uint64) map[string]interface{} {
	response := RequestQuery("SELECT * FROM tasks WHERE id=?", id)
	response.Next()

	var uid uint64
	var userid uint64
	var tasktypeid uint64
	if err := response.Scan(&uid, &userid, &tasktypeid); err != nil {
		return nil
	}
	return map[string]interface{}{
		"id":       uid,
		"owner_id": userid,
		"type_id":  tasktypeid,
	}
}

func GetTasksByUserId(userId uint64) []interface{} {
	response := RequestQuery("SELECT * FROM tasks WHERE user_id=?", userId)
	var ret []interface{}

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
