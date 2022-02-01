package model

import "time"

type Comment struct {
	User_id      uint64    `json:"user_id"`
	Content_id   uint64    `json:"content_id"`
	Content_type string    `json:"content_type"`
	Posted_at    time.Time `json:"posted_at"`
	Content      string    `json:"content"`
	Reply_to     uint64    `json:"reply_to"`
}

type DisplayedComment struct {
	Username  string    `json:"username"`
	Posted_at time.Time `json:"posted_at"`
	Content   string    `json:"content"`
}

func PostComment(c Comment) error {

	stmt, err := db.Prepare("insert into comments (user_id, content_id, content_type, posted_at, content, reply_to) values (?, ?, ?, ?, ?, ?);")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(c.User_id, c.Content_id, c.Content_type, c.Posted_at, c.Content, c.Reply_to)

	return err
}

func GetAllComments() ([]DisplayedComment, error) {
	var comments []DisplayedComment

	query := "SELECT users.username, comments.posted_at, comments.content from comments inner join users on comments.user_id = users.ID;"

	rows, err := db.Query(query)

	if err != nil {
		return comments, err
	}

	for rows.Next() {
		var username, content string
		var posted_at time.Time

		err := rows.Scan(&username, &posted_at, &content)

		if err != nil {
			return comments, err
		}

		comment := DisplayedComment{
			Username:  username,
			Posted_at: posted_at,
			Content:   content,
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
