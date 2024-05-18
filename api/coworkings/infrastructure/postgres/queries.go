package postgres

var GET_ALL_COWORKINGS = `SELECT id, name, createdAt FROM coworkings`

var CREATE_COWORKING = `INSERT INTO coworkings (name, createdAt) VALUES ($1, $2) RETURNING id`

var CREATE_WORKSPACE = `INSERT INTO workspaces (name, createdAt, coworking_id) VALUES ($1, $2, $3) RETURNING id`
