package database

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "fmt"
    "math"
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


// Function to map rating to image file
func GetRatingImage(rating float64) string {
    // Scale and round the rating to the nearest multiple of 5
    scaledRating := int(math.Round(rating * 10 / 5) * 5)
    
    // Ensure the rating stays within the range of 0 to 50
    if scaledRating < 0 {
        scaledRating = 0
    } else if scaledRating > 50 {
        scaledRating = 50
    }
    // Format the image file name
    return fmt.Sprintf("./images/ratings/rating-%d.png", scaledRating)
}


//displaying price in cents into price in dollars (basically a money converter)
func GetPriceCents(price int) string {
    return fmt.Sprintf("%.2f", float64(price)/100.0)
}


//function for searching products via keywords
func SearchProducts(query string) ([]Product, error) {
    query = "%" + query + "%"
    sqlQuery := `
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
        WHERE keywords ILIKE $1
    `
    rows, err := DB.Query(sqlQuery, query)
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
            &product.ProductType,
            &product.SizeChartLink, 
            &product.InstructionsLink, 
            &product.WarrantyLink,
        )
        if err != nil {
            return nil, err
        }
        product.RatingImages = GetRatingImage(product.RatingStars)
        product.PriceDollar = GetPriceCents(product.PriceCents)
        products = append(products, product)
    }

    return products, nil
}



//get all prods from db
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
        // convert rating from numbers to local links
        product.RatingImages = GetRatingImage(product.RatingStars)
        product.PriceDollar = GetPriceCents(product.PriceCents)
        
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

    //customized attributes
    RatingImages    string
    PriceDollar     string
}
