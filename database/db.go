package database

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

// Accept the connection string as a parameter
func Connect(connStr string) {
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
}

func GetProducts() ([]Product, error) {
    query := `
        SELECT 
            id, 
            image, 
            name, 
            ratingstars, 
            ratingcount, 
            pricecents, 
            keywords, 
            type, 
            sizechartlink, 
            instructionslink, 
            warrantylink 
        FROM products
    `
    rows, err := DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []Product
    for rows.Next() {
        var product Product
        err := rows.Scan(
            &product.ID, 
            &product.Image, 
            &product.Name, 
            &product.RatingStars, 
            &product.RatingCount, 
            &product.PriceCents, 
            &product.Keywords, 
            &product.ProductType,  // Use ProductType here to avoid 'type' conflict
            &product.SizeChartLink, 
            &product.InstructionsLink, 
            &product.WarrantyLink,
        )
        if err != nil {
            return nil, err
        }
        products = append(products, product)
    }

    return products, nil
}

type Product struct {
    ID              string
    Image           string
    Name            string
    RatingStars     float64
    RatingCount     int
    PriceCents      int
    Keywords        string
    ProductType     string  // Rename the field to avoid using 'type'
    SizeChartLink   sql.NullString
    InstructionsLink sql.NullString
    WarrantyLink    sql.NullString
}
