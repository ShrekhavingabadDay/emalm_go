package model

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
	Pw   string `json:"password"`
}

func GetAllUsers() ([]User, error) {
	var users []User

	query := "select id, username from users;"

	rows, err := db.Query(query)

	if err != nil {
		return users, err
	}

	for rows.Next() {
		var id uint64
		var name string

		err := rows.Scan(&id, &name)

		if err != nil {
			return users, err
		}

		user := User{
			ID:   id,
			Name: name,
			Pw:   "",
		}

		users = append(users, user)
	}

	defer rows.Close()

	return users, nil
}
