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
)

func setupTestCustOrderDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, err
	}

	createTable := `
	CREATE TABLE customerorder (
		orderID INTEGER PRIMARY KEY,
		cellnumber TEXT NOT NULL,
		catalogueID TEXT NOT NULL,
		orderitems TEXT NOT NULL,
		orderTotal INTEGER DEFAULT 0,
		ispaid BOOLEAN DEFAULT 0,
		datetimedelivered DATETIME,
		isclosed BOOLEAN DEFAULT 0
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func setupTestCatalogueItemDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return nil, err
	}

	createTable := `
	CREATE TABLE catalogueitem (
		catalogueID varchar(255) NOT NULL,
		catalogueitemID int4 NOT NULL,
		"selection" varchar(255) NULL,
		"item" varchar(255) NULL,
		"options" _text NULL,
		pricingType pricingTypeEnum,
		CONSTRAINT catalogueitem_pk PRIMARY KEY (catalogueID, catalogueitemID)
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}

var GardeningSelection = mb.CatalogueSelection{
	Preamble: grdngSlctnPreamble,
	Items: []mb.CatalogueItem{
		{
			CatalogueItemID: 1,
			Item:            "Denitrified fertilizer",
			Options: []string{
				"5g @ R110 p.g.",
				"10g @ R90 p.g.",
			},
			PricingType: mb.WeightItem,
		},
		{
			CatalogueItemID: 2,
			Item:            "Dehydrogenated water",
			Options: []string{
				"5g @ R140 p.g.",
				"10g @ R120 p.g.",
			},
			PricingType: mb.WeightItem,
		},
		{
			CatalogueItemID: 3,
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
			CatalogueItemID: 7,
			Item:            "DIY ready cake mix",
			Options: []string{
				"5g @ R180 p.g.",
				"10g @ R160 p.g.",
			},
			PricingType: mb.WeightItem,
		},
		{
			CatalogueItemID: 8,
			Item:            "Sugarless Sugar",
			Options: []string{
				"5g @ R210 p.g.",
				"10g @ R190 p.g.",
			},
			PricingType: mb.WeightItem,
		},
		{
			CatalogueItemID: 9,
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
			CatalogueItemID: 10,
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
			CatalogueItemID: 11,
			Item:            "Macless Apple @ R100 each",
			Options:         []string{},
			PricingType:     mb.SingleItem,
		},
		{
			CatalogueItemID: 12,
			Item:            "Unchargeable cellphone @ R150 each",
			Options:         []string{},
			PricingType:     mb.SingleItem,
		},
	},
}

var EdiblesSelection = mb.CatalogueSelection{
	Preamble: edblsSlctnPreamble,
	Items: []mb.CatalogueItem{
		{
			CatalogueItemID: 13,
			Item:            "Fruit toffees - 400mg",
			Options: []string{
				"10-Pack @ R200",
			},
			PricingType: mb.SingleItem,
		},
		{
			CatalogueItemID: 14,
			Item:            "Sour space strips - 400mg",
			Options: []string{
				"10-Pack @ R180",
			},
			PricingType: mb.SingleItem,
		},
		{
			CatalogueItemID: 15,
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
