package menubotlib_test

import (
	"database/sql"

	mb "github.com/JeremyJalpha/MenuBotLib"
	_ "modernc.org/sqlite"
)

const (
	catalogueID string = "Pig"

	//pricelistPreamble = "All fertilizer quoted per gram."

	grdngSlctnPreamble = "Gardening:"
	ktcnSlctnPreamble  = "Kitchen:"
	diySlctnPreamble   = "DIY:"
	tchSlctnPreamble   = "Tech:"
	edblsSlctnPreamble = "Edibles:"

	crtCustomerOrderTbl = `
	CREATE TABLE customerorder (
		orderID INTEGER PRIMARY KEY,
		cellnumber varchar(15) NOT NULL,
		catalogueID varchar(30) NOT NULL,
		orderitems varchar(255) NOT NULL,
		orderTotal INTEGER DEFAULT 0,
		ispaid BOOLEAN DEFAULT 0,
		datetimedelivered DATETIME,
		isclosed BOOLEAN DEFAULT 0
	);`

	crtCatalogueItemTbl = `
	CREATE TABLE catalogueitem (
		catalogueID varchar(255) NOT NULL,
		catalogueitemID INTEGER NOT NULL,
		"selection" varchar(255) NULL,
		"item" varchar(255) NULL,
		"options" varchar(255) NULL,
		pricingType pricingTypeEnum,
		CONSTRAINT catalogueitem_pk PRIMARY KEY (catalogueID, catalogueitemID)
	);`
)

func setupTestDBInstance() (*sql.DB, error) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, err
	}

	return db, nil
}

var GardeningSelection = mb.CatalogueSelection{
	Preamble: grdngSlctnPreamble,
	Items: []mb.CatalogueItem{
		{
			CatalogueID:     catalogueID,
			CatalogueItemID: 1,
			Selection:       grdngSlctnPreamble,
			Item:            "Denitrified fertilizer",
			Options: []string{
				"5g @ R110 p.g.",
				"10g @ R90 p.g.",
			},
			PricingType: mb.WeightItem,
		},
		{
			CatalogueID:     catalogueID,
			CatalogueItemID: 2,
			Selection:       grdngSlctnPreamble,
			Item:            "Dehydrogenated water",
			Options: []string{
				"5g @ R140 p.g.",
				"10g @ R120 p.g.",
			},
			PricingType: mb.WeightItem,
		},
		{
			CatalogueID:     catalogueID,
			CatalogueItemID: 3,
			Selection:       grdngSlctnPreamble,
			Item:            "Decarbonized soil",
			Options: []string{
				"5g @ R150 p.g.",
				"10g @ R130 p.g.",
			},
			PricingType: mb.WeightItem,
		},
	},
}

var KitchenSelection = mb.CatalogueSelection{
	Preamble: ktcnSlctnPreamble,
	Items: []mb.CatalogueItem{
		{
			CatalogueID:     catalogueID,
			CatalogueItemID: 4,
			Item:            "DIY ready cake mix",
			Selection:       ktcnSlctnPreamble,
			Options: []string{
				"5g @ R180 p.g.",
				"10g @ R160 p.g.",
			},
			PricingType: mb.WeightItem,
		},
		{
			CatalogueID:     catalogueID,
			CatalogueItemID: 5,
			Selection:       ktcnSlctnPreamble,
			Item:            "Sugarless Sugar",
			Options: []string{
				"5g @ R210 p.g.",
				"10g @ R190 p.g.",
			},
			PricingType: mb.WeightItem,
		},
		{
			CatalogueID:     catalogueID,
			CatalogueItemID: 6,
			Selection:       ktcnSlctnPreamble,
			Item:            "Burnt bread crumbs",
			Options: []string{
				"5g @ R250 p.g.",
				"10g @ R230 p.g.",
			},
			PricingType: mb.WeightItem,
		},
	},
}

var DIYSelection = mb.CatalogueSelection{
	Preamble: diySlctnPreamble,
	Items: []mb.CatalogueItem{
		{
			CatalogueID:     catalogueID,
			CatalogueItemID: 7,
			Selection:       diySlctnPreamble,
			Item:            "Bristleless Broom",
			Options: []string{
				"Vacuumless roomba version @ R650",
				"Bristled handleless version @ R650",
				"Floppy handled kinetic version @ R650",
			},
			PricingType: mb.SingleItem,
		},
	},
}

var TechSelection = mb.CatalogueSelection{
	Preamble: tchSlctnPreamble,
	Items: []mb.CatalogueItem{
		{
			CatalogueID:     catalogueID,
			CatalogueItemID: 8,
			Selection:       tchSlctnPreamble,
			Item:            "Macless Apple @ R100 each",
			Options:         nil,
			PricingType:     mb.SingleItem,
		},
		{
			CatalogueID:     catalogueID,
			CatalogueItemID: 9,
			Selection:       tchSlctnPreamble,
			Item:            "Unchargeable cellphone @ R150 each",
			Options:         nil,
			PricingType:     mb.SingleItem,
		},
	},
}

var EdiblesSelection = mb.CatalogueSelection{
	Preamble: edblsSlctnPreamble,
	Items: []mb.CatalogueItem{
		{
			CatalogueID:     catalogueID,
			CatalogueItemID: 10,
			Selection:       edblsSlctnPreamble,
			Item:            "Fruit toffees - 400mg",
			Options: []string{
				"10-Pack @ R200",
			},
			PricingType: mb.SingleItem,
		},
		{
			CatalogueID:     catalogueID,
			CatalogueItemID: 11,
			Selection:       edblsSlctnPreamble,
			Item:            "Sour space strips - 400mg",
			Options: []string{
				"10-Pack @ R180",
			},
			PricingType: mb.SingleItem,
		},
		{
			CatalogueID:     catalogueID,
			CatalogueItemID: 12,
			Selection:       edblsSlctnPreamble,
			Item:            "Space bud treats - 240mg",
			Options: []string{
				"3-Pack @ R200",
			},
			PricingType: mb.SingleItem,
		},
	},
}

var selections = []mb.CatalogueSelection{
	GardeningSelection,
	KitchenSelection,
	DIYSelection,
	TechSelection,
	EdiblesSelection,
}
