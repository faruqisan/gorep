package model

type (
	// Post struct describing data model
	Post struct {
		ID     int64  `json:"id"`
		UserID int64  `json:"user-id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	// UserPosts describe post owned by user
	UserPosts struct {
		UserID int64  `json:"user-id"`
		Posts  []Post `json:"posts"`
	}
)
