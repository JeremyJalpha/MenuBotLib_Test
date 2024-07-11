package menubotlib

import (
	"fmt"
)

// CatalogueSelection represents a section of the catalogue with a specific pricing regime
type CatalogueSelection struct {
	Preamble string
	Items    []CatalogueItem
}

// Generate a string for a single question and answer
func CatalogueItemAsAString(item CatalogueItem) string {
	optionsText := ""
	for i, option := range item.Options {
		optionsText += fmt.Sprintf("   %d. %s\n", i+1, option)
	}

	qA := fmt.Sprintf("%d: %s\n%s\n", item.CatalogueItemID, item.Item, optionsText)

	return qA
}

// Iterate over Questions array and populate questions array
func SingleSelectionAsAString(section CatalogueSelection) string {
	allItems := section.Preamble + "\n"

	for _, item := range section.Items {
		allItems += CatalogueItemAsAString(item)
	}

	return allItems
}

// Updated function to use the new CatalogueSection structure
func AssembleCatalogueSelections(pricelistpreamble string, ctlgselections []CatalogueSelection) string {
	selectionString := pricelistpreamble + "\n\n"

	for i, selection := range ctlgselections {
		// Outdoor Selection
		selectionString += SingleSelectionAsAString(selection)
		if i < len(ctlgselections)-1 {
			selectionString += "\n"
		}
	}

	return selectionString
}
