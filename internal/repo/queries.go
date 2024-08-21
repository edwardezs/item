package repo

const selectItemsQuery = `
	SELECT id, title, description, done FROM items;
`

const selectItemQuery = `
	SELECT id, title, description, done FROM items WHERE id = $1;
`

const createItemQuery = `
	INSERT INTO items (title, description, done) VALUES ($1, $2, $3) RETURNING id;
`

const deleteItemQuery = `
    DELETE FROM items WHERE id = $1;
`

const selectUsersQuery = `
	SELECT id, name, job FROM users;
`

const selectUserQuery = `
	SELECT id, name, job FROM users WHERE id = $1;
`

const createUserQuery = `
	INSERT INTO users (name, job) VALUES ($1, $2) RETURNING id;
`

const deleteUserQuery = `
    DELETE FROM users WHERE id = $1;
`

const selectStatusQuery = `
	SELECT table_name, read_only FROM table_lock_status;
`
