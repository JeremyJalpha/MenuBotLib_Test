package menubotlib

import (
	"database/sql"
)

// Define a custom type for PricingType
type PricingType string

// Define constants for the PricingType values
const (
	WeightItem PricingType = "WeightItem"
	SingleItem PricingType = "SingleItem"
)

type CatalogueItem struct {
	CatalogueID     string
	Selection       string
	CatalogueItemID int
	Item            string
	Options         []string
	PricingType     PricingType
}

// GetCatalogueItemsFromDB retrieves catalogue items from the database based on catalogueID.
func GetCatalogueItemsFromDB(db *sql.DB, catalogueID string) ([]CatalogueItem, error) {
	queryString := `SELECT catalogueitemID, "item", "options", pricingType FROM catalogueitem WHERE catalogueID = $1`
	rows, err := db.Query(queryString, catalogueID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []CatalogueItem
	for rows.Next() {
		var item CatalogueItem
		err := rows.Scan(&item.CatalogueItemID, &item.Item, &item.Options, &item.PricingType)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
