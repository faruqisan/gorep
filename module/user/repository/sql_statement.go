package repository

var (
	get    = "SELECT * FROM users"
	find   = "SELECT * FROM users WHERE id = ?"
	store  = "INSERT INTO users (name, email, password) VALUES (?,?,?)"
	update = "UPDATE users SET name = ?, password = ?"
	delete = "DELETE FROM users WHERE id = ?"
)
