package stock

const (
	CreateStockQuery = `
	INSERT INTO stock (
		id,
		name,
		price,
		availability,
		is_active
	) VALUES (?,?,?,?,?);`

	GetStockQueryByID = `
	SELECT 
		id,
		name,
		price,
		availability,
		is_active
	FROM stock WHERE id = ?`
)
