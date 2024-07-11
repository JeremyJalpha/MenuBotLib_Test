package menubotlib

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
