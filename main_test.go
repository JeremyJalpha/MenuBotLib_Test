package menubotlib_test

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"testing"

	mb "github.com/JeremyJalpha/MenuBotLib"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func Test_ParseUpdateOrderCommand(t *testing.T) {
	tests := []struct {
		commandText string
		expected    []mb.MenuIndication
		expectError bool
	}{
		{
			commandText: "update order 6:0",
			expected: []mb.MenuIndication{
				{ItemMenuNum: 6, ItemAmount: "0"},
			},
			expectError: false,
		},
		{
			commandText: "update order 9:12, 10: 1x3, 3x2, 2x1, 6:5",
			expected: []mb.MenuIndication{
				{ItemMenuNum: 10, ItemAmount: "1x3, 3x2, 2x1"},
				{ItemMenuNum: 9, ItemAmount: "12"},
				{ItemMenuNum: 6, ItemAmount: "5"},
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		result, err := mb.ParseUpdateOrderCommand(test.commandText)
		if (err != nil) != test.expectError {
			t.Errorf("ParseUpdateOrderCommand(%q) error = %v, expectError %v", test.commandText, err, test.expectError)
			continue
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ParseUpdateOrderCommand(%q) = %v, want %v", test.commandText, result, test.expected)
		}
	}
}

// baseline: "Update order 9:12, 10: 1x3, 3x2, 2x1, 6: 5"
func Test_UpdateCustOrdItems(t *testing.T) {
	var err error

	tests := []struct {
		given       mb.CustomerOrder
		update      mb.OrderItems
		expected    mb.OrderItems
		expectError bool
	}{
		{
			given: mb.CustomerOrder{
				OrderItems: mb.OrderItems{
					MenuIndications: []mb.MenuIndication{
						{ItemMenuNum: 10, ItemAmount: "1x3"},
						{ItemMenuNum: 9, ItemAmount: "5"},
						{ItemMenuNum: 8, ItemAmount: "6"},
						{ItemMenuNum: 6, ItemAmount: "5"},
						{ItemMenuNum: 5, ItemAmount: "9"},
					},
				},
			},
			update: mb.OrderItems{
				MenuIndications: []mb.MenuIndication{
					{ItemMenuNum: 10, ItemAmount: "1x3, 3x2, 2x1"},
					{ItemMenuNum: 6, ItemAmount: "0"},
					{ItemMenuNum: 7, ItemAmount: "7"},
				},
			},
			expected: mb.OrderItems{
				MenuIndications: []mb.MenuIndication{
					{ItemMenuNum: 10, ItemAmount: "1x3, 3x2, 2x1"},
					{ItemMenuNum: 9, ItemAmount: "5"},
					{ItemMenuNum: 8, ItemAmount: "6"},
					{ItemMenuNum: 5, ItemAmount: "9"},
					{ItemMenuNum: 7, ItemAmount: "7"},
				},
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		err = test.given.UpdateCustOrdItems(test.update)
		assert.NoError(t, err)

		if (err != nil) != test.expectError {
			t.Errorf("UpdateCustOrdItems(%q) error = %v, expectError %v", test.update, err, test.expectError)
			continue
		}
		if !reflect.DeepEqual(test.given.OrderItems, test.expected) {
			t.Errorf("UpdateCustOrdItems(%q) = %v, want %v", test.update, test.given.OrderItems, test.expected)
		}
	}
}

func Test_NewOrder_UpdateOrInsertCurrentOrder(t *testing.T) {
	db, err := setupTestDBInstance()
	assert.NoError(t, err)
	defer db.Close()

	_, err = db.Exec(crtCustomerOrderTbl)
	assert.NoError(t, err)

	tests := []struct {
		custOrd     mb.CustomerOrder
		expected    mb.OrderItems
		expectError bool
	}{
		{
			custOrd: mb.CustomerOrder{
				OrderID:     12345,
				CellNumber:  "0766140000",
				CatalogueID: catalogueID,
				OrderItems: mb.OrderItems{
					MenuIndications: []mb.MenuIndication{
						{ItemMenuNum: 10, ItemAmount: "1x3, 3x2, 2x1"},
						{ItemMenuNum: 9, ItemAmount: "12"},
						{ItemMenuNum: 6, ItemAmount: "5"},
					},
				},
			},
			expected: mb.OrderItems{
				MenuIndications: []mb.MenuIndication{
					{ItemMenuNum: 10, ItemAmount: "1x3, 3x2, 2x1"},
					{ItemMenuNum: 9, ItemAmount: "12"},
					{ItemMenuNum: 6, ItemAmount: "5"},
				},
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		err = test.custOrd.UpdateOrInsertCurrentOrder(db, test.custOrd.CellNumber, test.expected, true)
		assert.NoError(t, err)

		var readOrderItems string
		query := `SELECT orderitems FROM customerorder WHERE orderID = ?`
		err = db.QueryRow(query, test.custOrd.OrderID).Scan(&readOrderItems)
		assert.NoError(t, err)

		// Unmarshal the JSON string back to []OrderItem
		var actual mb.OrderItems
		err = json.Unmarshal([]byte(readOrderItems), &actual)
		assert.NoError(t, err)

		if (err != nil) != test.expectError {
			t.Errorf("UpdateOrInsertCurrentOrder(%q) error = %v, expectError %v", test.custOrd.OrderItems, err, test.expectError)
			continue
		}
		result := mb.OrderItems{MenuIndications: actual.MenuIndications}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("UpdateOrInsertCurrentOrder(%q) = %v, want %v", test.custOrd.OrderItems, result, test.expected)
		}
	}
}

func Test_CheckoutNow(t *testing.T) {

	db, err := setupTestDBInstance()
	assert.NoError(t, err)
	defer db.Close()

	senderNum := "0000000000"
	pymntRtrnBase := "payment_return"
	pymntCnclBase := "payment_canceled"
	returnBaseURL := "/" + pymntRtrnBase
	cancelBaseURL := "/" + pymntCnclBase
	notifyBaseURL := "/payment_notify"
	ItemNamePrefix := "Order"

	HomebaseURL := "https://yourhomedomain.com"

	MerchantId := "10000100"
	MerchantKey := "*************"
	Passphrase := "*************"
	PfHost := "https://sandbox.payfast.co.za/eng/process"

	checkoutInfo := mb.CheckoutInfo{
		ReturnURL:      HomebaseURL + returnBaseURL,
		CancelURL:      HomebaseURL + cancelBaseURL,
		NotifyURL:      HomebaseURL + notifyBaseURL,
		MerchantId:     MerchantId,
		MerchantKey:    MerchantKey,
		Passphrase:     Passphrase,
		HostURL:        PfHost,
		ItemNamePrefix: ItemNamePrefix,
	}

	tests := []struct {
		ctlgSelections []mb.CatalogueSelection
		userInfo       mb.UserInfo
		custOrd        mb.CustomerOrder
		expected       mb.OrderItems
		expectError    bool
	}{
		{
			//Remember to change the static definitions to match your pricelist defenition.
			ctlgSelections: selections,
			userInfo: mb.UserInfo{
				CellNumber: senderNum,
				NickName:   mb.NullString{NullString: sql.NullString{String: "testSplurge", Valid: true}},
				Email:      mb.NullString{NullString: sql.NullString{String: "sbtu01@payfast.io", Valid: true}},
			},
			custOrd: mb.CustomerOrder{
				OrderID:     12345,
				CellNumber:  senderNum,
				CatalogueID: catalogueID,
				OrderItems: mb.OrderItems{
					MenuIndications: []mb.MenuIndication{
						{ItemMenuNum: 10, ItemAmount: "1x3, 3x2, 2x1"},
						{ItemMenuNum: 9, ItemAmount: "12"},
						{ItemMenuNum: 6, ItemAmount: "5"},
					},
				},
			},
			expected: mb.OrderItems{
				MenuIndications: []mb.MenuIndication{
					{ItemMenuNum: 10, ItemAmount: "1x3, 3x2, 2x1"},
					{ItemMenuNum: 9, ItemAmount: "12"},
					{ItemMenuNum: 6, ItemAmount: "5"},
				},
			},
			expectError: false,
		},
	}

	for _, test := range tests {
		mb.BeginCheckout(db, test.userInfo, test.ctlgSelections, test.custOrd, checkoutInfo, true)
		assert.NoError(t, err)

		// ...

		// if (err != nil) != test.expectError {
		// 	t.Errorf("UpdateOrInsertCurrentOrder(%q) error = %v, expectError %v", test.custOrd.OrderItems, err, test.expectError)
		// 	continue
		// }
	}
}

func extractItemsFromSelections(selections []mb.CatalogueSelection) []mb.CatalogueItem {
	var items []mb.CatalogueItem
	for _, selection := range selections {
		items = append(items, selection.Items...)
	}
	return items
}

func Test_PrintPriceList(t *testing.T) {
	db, err := setupTestDBInstance()
	assert.NoError(t, err)
	defer db.Close()

	_, err = db.Exec(crtCatalogueItemTbl)
	assert.NoError(t, err)

	err = mb.InsertCatalogueItems(db, selections)
	assert.NoError(t, err)

	expectedItems := extractItemsFromSelections(selections)

	ctlgItms, err := mb.GetCatalogueItemsFromDB(db, catalogueID)
	assert.NoError(t, err)

	if !reflect.DeepEqual(ctlgItms, expectedItems) {
		t.Errorf("getCatalogueItemsFromDB(db, %s...) = %v, want %v", catalogueID, ctlgItms, expectedItems)
	}
}

func Test_ComposedCatalogueSelections(t *testing.T) {
	db, err := setupTestDBInstance()
	assert.NoError(t, err)
	defer db.Close()

	_, err = db.Exec(crtCatalogueItemTbl)
	assert.NoError(t, err)

	err = mb.InsertCatalogueItems(db, selections)
	assert.NoError(t, err)

	ctlgItms, err := mb.GetCatalogueItemsFromDB(db, catalogueID)
	assert.NoError(t, err)

	convertedctlgItms := mb.CmpsCtlgSlctnsFromCtlgItms(ctlgItms)

	if !reflect.DeepEqual(convertedctlgItms, selections) {
		t.Errorf("getCatalogueItemsFromDB(db, %s...) = %v, want %v", catalogueID, ctlgItms, selections)
	}
}

func Test_CalculatePriceWithNoOptions(t *testing.T) {
	// Static Menu definition should say:
	// Denitrified fertilizer:
	//  - "5g @ R110 p.g.",
	//  - "10g @ R90 p.g.",
	// Therefore Expected Total: 12 * 90 = 1080
	tests := []struct {
		ordItems      mb.OrderItems
		expctdTotal   int
		expctdSummary string
		expctError    bool
	}{
		{
			ordItems: mb.OrderItems{
				MenuIndications: []mb.MenuIndication{
					{ItemMenuNum: 1, ItemAmount: "12"},
				},
			},
			expctdTotal:   1080,
			expctdSummary: "",
			expctError:    false,
		},
	}

	for _, test := range tests {
		total, summary := test.ordItems.CalculatePrice(selections)
		if summary != test.expctdSummary {
			t.Errorf("CalculatePrice(selections).summary = %v, want %v", summary, test.expctdSummary)
		}
		if total != test.expctdTotal {
			t.Errorf("CalculatePrice(selections).total = %v, want %v", total, test.expctdTotal)
		}
	}
}
