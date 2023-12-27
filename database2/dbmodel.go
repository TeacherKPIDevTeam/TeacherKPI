package database2

import "database/sql"

func GetUserDataById(id uint64) (map[string]interface{}, error) {
	response, err := RequestQuery("SELECT * FROM users WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	response.Next()

	var uid uint64
	var username string
	if err := response.Scan(&uid, &username); err != nil {
		return nil, err
	}
	response.Close()
	return map[string]interface{}{
		"id":       uid,
		"username": username,
	}, nil
}

func CreateUser(userData map[string]interface{}) (sql.Result, error) {
	username := userData["username"].(string)

	result, err := RequestNonQuery("INSERT INTO users VALUES(NULL, ?)", username)
	if err != nil {
		return nil, err
	}
	return result, err
}

func SaveUser(userData map[string]interface{}) error {
	id := userData["id"].(uint64)
	username := userData["username"].(string)

	result, err := RequestNonQuery("UPDATE users SET username=? WHERE id=?", username, id)
	_ = result
	return err
}

func GetTaskDataById(id uint64) (map[string]interface{}, error) {
	response, err := RequestQuery("SELECT * FROM tasks WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	response.Next()

	var uid uint64
	var userid uint64
	var tasktypeid uint64
	if err := response.Scan(&uid, &userid, &tasktypeid); err != nil {
		return nil, err
	}

	response.Close()
	return map[string]interface{}{
		"id":       uid,
		"owner_id": userid,
		"type_id":  tasktypeid,
	}, nil
}

func GetTaskIdsByUserId(userId uint64) ([]uint64, error) {
	response, err := RequestQuery("SELECT id FROM tasks WHERE user_id=?", userId)
	if err != nil {
		return nil, err
	}

	var ret []uint64

	var id uint64
	for response.Next() {
		if err := response.Scan(&id); err != nil {
			return nil, err
		}
		ret = append(ret, id)
	}

	response.Close()
	return ret, nil
}

func GetStageDataById(id uint64) (map[string]interface{}, error) {
	request_text := "SELECT t1.`id`, t1.`task_id`, t1.`type_id`, t1.`status`, t2.`queue_position`, t2.`name`, t3.`name` " +
		"FROM stages AS t1, stage_types AS t2, stage_statuses AS t3 " +
		"WHERE t1.`id` = ? AND t2.`id` = t1.`status` AND t3.`id` = t1.`status`"
	response, err := RequestQuery(request_text, id)
	if err != nil {
		return nil, err
	}
	response.Next()

	var baseId, taskId, typeId uint64
	var status, position int
	var statusName, typeName string
	if err := response.Scan(&baseId, &taskId, &typeId, &status, &position, &typeName, &statusName); err != nil {
		return nil, err
	}

	response.Close()
	return map[string]interface{}{
		"id":          baseId,
		"task_id":     taskId,
		"type_id":     typeId,
		"type_name":   typeName,
		"status":      status,
		"queue_pos":   position,
		"status_name": statusName,
	}, nil
}

func GetStageIdsByTaskId(taskId uint64) ([]uint64, error) {
	response, err := RequestQuery("SELECT id FROM stages WHERE task_id=?", taskId)
	if err != nil {
		return nil, err
	}
	var ret []uint64

	var id uint64
	for response.Next() {
		if err := response.Scan(&id); err != nil {
			return nil, err
		}
		ret = append(ret, id)
	}

	response.Close()
	return ret, nil
}
