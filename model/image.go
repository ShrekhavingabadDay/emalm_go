package model

type Image struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Uuid        string `json:"uuid"`
	User        uint64 `json:"user_id"`
}

type DisplayedImage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Uuid        string `json:"uuid"`
	User        string `json:"username"`
}

func UploadImage(i Image) error {

	stmt, err := db.Prepare("insert into images (title, descript, uuid, user) values (?, ?, ?, ?);")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(i.Title, i.Description, i.Uuid, i.User)

	return err
}

func GetAllImages() ([]DisplayedImage, error) {

	var images []DisplayedImage

	query := "select images.title, images.descript, images.uuid, users.username from images inner join users on users.id = images.user;"

	rows, err := db.Query(query)

	if err != nil {
		return images, err
	}

	for rows.Next() {
		var title, description, uuid, username string

		err := rows.Scan(&title, &description, &uuid, &description)

		if err != nil {
			return images, err
		}

		image := DisplayedImage{
			Title:       title,
			Description: description,
			Uuid:        uuid,
			User:        username,
		}

		images = append(images, image)
	}

	defer rows.Close()

	return images, nil
}
