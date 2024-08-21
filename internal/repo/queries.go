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

const selectStatusQuery = `
	SELECT table_name, read_only FROM table_lock_status;
`
