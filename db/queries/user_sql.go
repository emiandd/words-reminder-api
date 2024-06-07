package queries

const (
	SQLCreateNewUser = `
	INSERT INTO user (email, password, created)
	VALUES (?, ?, ?)
	`

	SQLFetchUsers = `
	SELECT 
		user_id,
		email,
		password,
		created
	FROM user WHERE 1 = 1 `

	SQLFindUser = `
	SELECT 
		user_id,
		email,
		password,
		created
	FROM user
	WHERE email = ?`

	SQLUserCount = "SELECT COUNT(*) AS count FROM user WHERE 1 = 1 "

	SQLLimitOffset = "LIMIT ? OFFSET ?"
)
