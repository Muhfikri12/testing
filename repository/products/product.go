package products

import (
	"database/sql"
	"ecommers/model"

	"go.uber.org/zap"
)

type ProductRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewProductRepository(db *sql.DB, Log *zap.Logger) ProductRepository {
	return ProductRepository{
		DB:     db,
		Logger: Log,
	}
}

func (p *ProductRepository) ShowAllProducts(limit, page int, category, name string) (*[]model.Products, int, error) {
	offset := (page - 1) * limit

	query := `SELECT p.id, p.name, p.image_url, p.price, p.discount, p.description, p.created_at, p.updated_at, ca.name,
		ROUND(COALESCE((
			SELECT AVG(pr.rating)::numeric
			FROM previews pr
			JOIN checkouts c ON pr.checkout_id = c.id
			JOIN users u ON c.user_id = u.id
			JOIN shopping_carts sc ON u.id = sc.user_id
			JOIN product_varians pv ON sc.product_variant_id = pv.id
			WHERE pv.product_id = p.id
			AND pv.deleted_at IS NULL
		), 0), 1) AS rating,
		CAST((
			SELECT COUNT(pr.id)
			FROM previews pr
			JOIN checkouts c ON pr.checkout_id = c.id
			JOIN users u ON c.user_id = u.id
			JOIN shopping_carts sc ON u.id = sc.user_id
			JOIN product_varians pv ON sc.product_variant_id = pv.id
			WHERE pv.product_id = p.id
			AND pv.deleted_at IS NULL
		) AS INTEGER) AS total_reviewers,
		CAST((
			SELECT COUNT(c.id)
			FROM checkouts c 
			JOIN users u ON c.user_id = u.id
			JOIN shopping_carts sc ON u.id = sc.user_id
			JOIN product_varians pv ON sc.product_variant_id = pv.id
			WHERE pv.product_id = p.id
			AND pv.deleted_at IS NULL
		) AS INTEGER) AS total_buyers
		FROM products p
		JOIN categories ca ON p.category_id = ca.id
		WHERE p.deleted_at IS NULL 
			AND (ca.name = $1 OR $1 = '') 
			AND (p.name ILIKE '%' || $2 || '%' OR $2 = '')
		LIMIT $3 OFFSET $4`

	rows, err := p.DB.Query(query, category, name, limit, offset)
	if err != nil {
		p.Logger.Error("Error from repository: " + err.Error())
		return nil, 0, err
	}
	defer rows.Close()

	productsarr := []model.Products{}

	for rows.Next() {
		product := model.Products{
			Timestamps: &model.Basic{},
			Previews:   &model.Previews{},
			Checkouts:  &model.Checkouts{},
			Categories: &model.Categories{},
		}

		if err := rows.Scan(&product.ID, &product.Name, &product.ImageUrl, &product.Price, &product.Discount, &product.Description, &product.Timestamps.Created_at, &product.Timestamps.Updated_at, &product.Categories.Name, &product.Previews.Rating, &product.Previews.TotalReviewers, &product.Checkouts.TotalBuyers); err != nil {
			p.Logger.Error("Error from repository: " + err.Error())
			return nil, 0, err
		}

		productsarr = append(productsarr, product)
	}

	var totalData int

	countQuery := `SELECT COUNT(*) 
				FROM products p
				JOIN categories ca ON p.category_id = ca.id
				WHERE p.deleted_at IS NULL 
                 AND (ca.name = $1 OR $1 = '') 
                 AND (p.name ILIKE '%' || $2 || '%' OR $2 = '')`

	err = p.DB.QueryRow(countQuery, category, name).Scan(&totalData)
	if err != nil {
		p.Logger.Error("event repository: failed to fetch total count", zap.Error(err))
		return nil, 0, err
	}

	return &productsarr, totalData, nil
}
