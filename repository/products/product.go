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
			JOIN checkout_items ci ON pr.checkout_item_id = ci.id
			WHERE ci.product_id = p.id
				AND ci.deleted_at IS NULL
		), 0), 1) AS rating,
		CAST((
			SELECT COUNT(pr.id)
			FROM previews pr
			JOIN checkout_items ci ON pr.checkout_item_id = ci.id
			WHERE ci.product_id = p.id
				AND ci.deleted_at IS NULL
		) AS INTEGER) AS total_reviewers,
		CAST((
			SELECT SUM(ci.qty)
			FROM checkout_items ci 
			WHERE ci.product_id = p.id
				AND ci.deleted_at IS NULL
		) AS INTEGER) AS total_sold,
		
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

		if err := rows.Scan(&product.ID, &product.Name, &product.ImageUrl, &product.Price, &product.Discount, &product.Description, &product.Timestamps.Created_at, &product.Timestamps.Updated_at, &product.Categories.Name, &product.Previews.Rating, &product.Previews.TotalReviewers, &product.Checkouts.TotalSold); err != nil {
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

func (p *ProductRepository) GetProductByID(id int) (*model.Products, error) {

	product := model.Products{
		Categories: &model.Categories{},
		Previews:   &model.Previews{},
	}

	query := `SELECT p.id, p.name, p.image_url, p.price, p.discount, p.description, c.name,
			ROUND(COALESCE((
				SELECT AVG(pr.rating)::numeric
				FROM previews pr
				JOIN checkout_items ci ON pr.checkout_item_id = ci.id
				WHERE ci.product_id = p.id
					AND ci.deleted_at IS NULL
			), 0), 1) AS rating,
			CAST((
				SELECT COUNT(pr.id)
				FROM previews pr
				JOIN checkout_items ci ON pr.checkout_item_id = ci.id
				WHERE ci.product_id = p.id
					AND ci.deleted_at IS NULL
			) AS INTEGER) AS total_reviewers
			FROM products p
			LEFT JOIN categories c ON p.category_id = c.id
			WHERE p.deleted_at IS NULL
			AND p.id = $1`
	if err := p.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.ImageUrl, &product.Price, &product.Discount, &product.Description, &product.Categories.Name, &product.Previews.Rating, &product.Previews.TotalReviewers); err != nil {
		p.Logger.Error("Error from GetProductByID repository: " + err.Error())
		return nil, err
	}

	variant, err := p.GetVariantProducts(int(product.ID))
	if err != nil {
		p.Logger.Error("Error from Get Product ID: " + err.Error())
		return nil, err
	}

	product.Variants = variant
	product.PriceAfterDiscount = product.Price - (product.Price * product.Discount / 100)
	return &product, nil
}

func (p *ProductRepository) GetVariantProducts(productID int) (*[]model.Variant, error) {

	variants := []model.Variant{}
	query := `SELECT pv.size, pv.color, pv.stock, pv.image_url
		FROM product_varians pv
		JOIN products p ON pv.product_id = p.id
		WHERE p.id = $1`

	row, err := p.DB.Query(query, productID)
	if err != nil {
		p.Logger.Error("Error from query GetVariantProducts repository: " + err.Error())
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		variant := model.Variant{}
		if err := row.Scan(&variant.Size, &variant.Color, &variant.Stocks, &variant.ImageUrl); err != nil {
			p.Logger.Error("Error from scan GetVariantProducts repository: " + err.Error())
			return nil, err
		}

		variants = append(variants, variant)
	}

	return &variants, nil
}
