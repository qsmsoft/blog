package repository

const (
	createUserQuery = `INSERT INTO users (first_name, last_name, email, password, role, avatar)
							VALUES ($1, $2, $3, $4, COALESCE(NULLIF($5, ''), 'user'), $6)
							RETURNING *`

	updateUserQuery = `UPDATE users  SET first_name = COALESCE(NULLIF($1, ''), first_name),
                  			last_name = COALESCE(NULLIF($2, ''), last_name),
                  			email = COALESCE(NULLIF($3, ''), email),
                  			role = COALESCE(NULLIF($4, ''), role),
                  			avatar = COALESCE(NULLIF($5, ''), avatar),
						WHERE id = $6
						RETURNING *`

	deleteUserQuery = `DELETE FROM users WHERE id = $1`

	getUserQuery = `SELECT id, first_name, last_name, email, role, avatar, created_at, updated_at, login_date  
					 FROM users 
					 WHERE id = $1`

	getTotalCount = `SELECT COUNT(id) FROM users WHERE first_name ILIKE '%' || $1 || '%' or last_name ILIKE '%' || $1 || '%'`

	findUsers = `SELECT id, first_name, last_name, email, role, avatar, created_at, updated_at, login_date 
				  FROM users 
				  WHERE first_name ILIKE '%' || $1 || '%' or last_name ILIKE '%' || $1 || '%'
				  ORDER BY first_name, last_name
				  OFFSET $2 LIMIT $3`

	getTotal = `SELECT COUNT(id) FROM users`

	getUsers = `SELECT id, first_name, last_name, email, role, avatar, created_at, updated_at, login_date
				 FROM users 
				 ORDER BY COALESCE(NULLIF($1, ''), first_name) OFFSET $2 LIMIT $3`

	findUserByEmail = `SELECT id, first_name, last_name, email, role, avatar,created_at, updated_at, login_date, password
				 		FROM users 
				 		WHERE email = $1`
)
