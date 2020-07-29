package main

import (
	"errors"
	//"math"
	"reflect"
	"sort"
	"testing"
)

func TestSortByName(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		/*{
			"basic case. nil slice",
			nil,
			nil,
		},*/
		{
			"basic case. empty slice",
			[]string{},
			[]string{},
		},
		{
			"basic case. single element slice",
			[]string{"a"},
			[]string{"a"},
		},
		{
			"single duplicate element. same price",
			[]string{"a", "a"},
			[]string{"a", "a"},
		},
		{
			"single duplicate element with one in the middle. different price",
			[]string{"a", "b", "a"},
			[]string{"a", "a", "b"},
		},
		{
			"five elements with duplicates",
			[]string{"a", "b", "c", "a", "c"},
			[]string{"a", "a", "b", "c", "c"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := PrintSortedNames2(tc.input, 0)

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf("got\t\t%v\nwant\t%v\n%s", res, tc.expected, tc.name)
			}
		})
	}
}

/*
func TestHashOrder(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected uint64
	}{
		{
			"basic case. nil slice",
			nil,
			14695981039346656037,
		},
		{
			"basic case. empty slice",
			[]string{},
			14695981039346656037,
		},
		{
			"basic case. single element slice",
			[]string{"a"},
			12638187200555641996,
		},
		{
			"single duplicate element",
			[]string{"a", "a"},
			620444549055354551,
		},
		{
			"single duplicate element with one in the middle. different price",
			[]string{"a", "b", "a"},
			16653391238245862383,
		},
		{
			"five elements with duplicates",
			[]string{"a", "b", "c", "a", "c"},
			1040503207626252133,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res := hashOrder(tc.input)

			if !reflect.DeepEqual(res, tc.expected) {
				t.Fatalf("got\t\t%v\nwant\t%v", res, tc.expected)
			}
		})
	}
}
*/

func TestCalculateOrder(t *testing.T) {
	testCases := []struct {
		name               string
		shop               map[string]float64
		order              UserCart
		expectedTotalPrice float64
		expectedError      error
	}{
		{
			"basic case. nil slice",
			map[string]float64{},
			nil,
			0,
			nil,
		},
		{
			"basic case. empty slice",
			map[string]float64{},
			UserCart{},
			0,
			nil,
		},
		{
			"basic case. single element slice",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			UserCart{"a": 1},
			1,
			nil,
		},
		{
			"basic case. two element slice",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			UserCart{"b": 1, "a": 1},
			11,
			nil,
		},
		/*{
			"basic case. single unknown item",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			UserCart{"xxx": 1},
			0,
			errors.New("errItemNoFound"),
		},*/
		/*{
			"basic case. single unknown item inbetween ",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			UserCart{"a": 1, "xxx": 1, "b": 1},
			0,
			errors.New("errItemNoFound"),
		},*/
		/*{
			"partial match",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			UserCart{"aa": 1},
			0,
			errors.New("errItemNoFound"),
		},*/
	}

	// переменная магазина
	var strBase = &StorageBase{}
	var usrCrt Carter

	for _, tc := range testCases {
		strBase = &StorageBase{make(StorageIds), make(Products), make(Orders)}

		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for name, price := range tc.shop {
				(strBase).AppendPr(name, price)
			}
			//
			usrCrt = &tc.order

			res := usrCrt.GetTotal(strBase)

			// ошибки не поддерживаются функцией
			// отсутствующие товары имеют нулевую цену
			//var err error = nil
			//res, err := CalculateOrder(tc.shop, tc.order)

			/*if !errors.Is(err, tc.expectedError) {
				t.Fatalf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			}*/
			if !reflect.DeepEqual(res, tc.expectedTotalPrice) {
				t.Fatalf("got\t\t%v\nwant\t%v\nCase name:%s", res, tc.expectedTotalPrice, tc.name)
			}
		})
	}
}

func TestCalculateOrderWithCache(t *testing.T) {
	testCases := []struct {
		name               string
		shop               map[string]float64
		order              UserCart
		expectedTotalPrice float64
		expectedError      error
	}{
		{
			"basic case. nil slice",
			map[string]float64{},
			nil,
			0,
			nil,
		},
		{
			"basic case. empty slice",
			map[string]float64{},
			UserCart{},
			0,
			nil,
		},
		{
			"basic case. single element slice",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			UserCart{"a": 1},
			1,
			nil,
		},
		{
			"basic case. two element slice",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			UserCart{"b": 1, "a": 1},
			11,
			nil,
		},
		/*{
			"basic case. single unknown item",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			UserCart{"xxx": 1},
			0,
			errors.New("errItemNoFound"),
		},
		{
			"basic case. single unknown item inbetween ",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			UserCart{"a": 1, "xxx": 1, "b": 1},
			0,
			errors.New("errItemNoFound"),
		},
		{
			"partial match",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			UserCart{"aa": 1},
			0,
			errors.New("errItemNoFound"),
		},*/
	}
	// переменная магазина
	var strBase = &StorageBase{}
	var usrCrt Carter

	for _, tc := range testCases {
		strBase = &StorageBase{make(StorageIds), make(Products), make(Orders)}
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for name, price := range tc.shop {
				(strBase).AppendPr(name, price)
			}
			//
			usrCrt = &tc.order

			for i := 0; i < 5; i++ {
				res := usrCrt.GetTotalMulti(strBase)

				/*if !errors.Is(err, tc.expectedError) {
					t.Fatalf("got\t\t%v\nwant\t%v", err, tc.expectedError)
				}*/
				if tc.expectedError == nil && i > 0 {
					sortedKeys := usrCrt.GetSortedKeys()
					crtStr := usrCrt.MakeSrtByCart(strBase, &sortedKeys)

					totalFromCache, ok := strBase.Ord[crtStr]

					if !ok {
						t.Fatalf("cache was not used: %v", tc.order)
					}
					if totalFromCache != tc.expectedTotalPrice {
						t.Fatalf("got in cache\t\t%v\nwant\t%v", totalFromCache, tc.expectedTotalPrice)
					}
				}
				if !reflect.DeepEqual(res, tc.expectedTotalPrice) {
					t.Fatalf("got\t\t%v\nwant\t%v", res, tc.expectedTotalPrice)
				}
			}
		})
	}
}

func TestAddItem(t *testing.T) {
	type itemStruct struct {
		nameItem  string
		priceItem float64
	}
	testCases := []struct {
		name          string
		shop          map[string]float64
		item          itemStruct
		expectedShop  map[uint]float64
		expectedError error
	}{
		{
			"basic case. empty Item",
			map[string]float64{},
			itemStruct{},
			map[uint]float64{},
			errors.New("errEmptyItem"),
		}, /*
			{
				"basic case. empty item name",
				map[string]float64{},
				itemStruct{"", 10},
				map[uint]float64{},
				errors.New("errEmptyItem"),
			},*/
		/*{
			"basic case. already exists",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			itemStruct{"a", 10},
			map[uint]float64{
				1: 1,
				2: 10,
			},
			errors.New("errItemAlreadyExists"),
		},*/

		{
			"basic case. correct item",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			itemStruct{"xxx", 100},
			map[uint]float64{
				1: 1,
				2: 10,
				3: 100,
			},
			nil,
		},
	}

	// переменная магазина
	var strBase = &StorageBase{}
	//var usrCrt Carter

	for _, tc := range testCases {
		strBase = &StorageBase{make(StorageIds), make(Products), make(Orders)}
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for name, price := range tc.shop {
				(strBase).AppendPr(name, price)
			}

			(strBase).AppendPr(tc.item.nameItem, tc.item.priceItem)
			//err := AddItem(tc.shop, tc.item)

			//if !errors.Is(err, tc.expectedError) {
			//	t.Errorf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			//}
			var a []float64
			var b []float64

			for _, v := range strBase.Prd {
				a = append(a, v)
			}
			for _, v := range tc.expectedShop {
				b = append(b, v)
			}
			sort.Float64s(a)
			sort.Float64s(b)
			if !reflect.DeepEqual(a, b) {
				t.Fatalf("got\t\t%v\nwant\t%v\n%s", strBase.Prd, tc.expectedShop, tc.name)
			}
		})
	}
}

func TestChangePrice(t *testing.T) {
	type itemStruct struct {
		nameItem  string
		priceItem float64
	}

	errStr := "Unable to append! Possible problems: cost < 0, empty product str, prd. already in base"

	testCases := []struct {
		name          string
		shop          map[string]float64
		item          itemStruct
		expectedShop  map[uint]float64
		expectedError error
	}{
		{
			"basic case. empty Item",
			map[string]float64{},
			itemStruct{},
			map[uint]float64{},
			errors.New(errStr),
		},
		/*{
			"basic case. empty item name",
			map[string]float64{},
			itemStruct{"", 10},
			map[uint]float64{},
			errors.New(errStr),
		},*/
		{
			"basic case. correct price change",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			itemStruct{"a", 10},
			map[uint]float64{
				1: 10,
				2: 10,
			},
			nil,
		},
		{
			"basic case. correct price change",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			itemStruct{"xxx", 10},
			map[uint]float64{
				1: 1,
				2: 10,
			},
			errors.New(errStr),
		},
	}

	var strBase = &StorageBase{}
	for _, tc := range testCases {
		strBase = &StorageBase{make(StorageIds), make(Products), make(Orders)}
		//tc := tc
		t.Run(tc.name, func(t *testing.T) {
			//t.Parallel()

			for name, price := range tc.shop {
				(strBase).AppendPr(name, price)
			}

			err := strBase.UpdatePr(tc.item.nameItem, tc.item.priceItem)

			var a []float64
			var b []float64

			for _, v := range strBase.Prd {
				a = append(a, v)
			}
			for _, v := range tc.expectedShop {
				b = append(b, v)
			}
			sort.Float64s(a)
			sort.Float64s(b)
			if tc.expectedError != nil && err == nil {
				t.Errorf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			}
			if !reflect.DeepEqual(a, b) {
				t.Fatalf("got\t\t%v\nwant\t%v\n%s", strBase.Prd, tc.expectedShop, tc.name)
			}
		})
	}
}

/*func TestChangeName(t *testing.T) {
	type itemStruct struct {
		nameItem  string
		priceItem float64
	}
	testCases := []struct {
		name          string
		shop          map[string]float64
		item          itemStruct
		newItemName   itemStruct
		expectedShop  map[string]uint
		expectedError error
	}{
		{
			"basic case. empty Item",
			map[string]float64{},
			itemStruct{"a", 1},
			itemStruct{"", 1},
			map[string]uint{},
			errEmptyItem,
		},
		{
			"basic case. empty item name",
			map[string]float64{
				"a": 10,
			},
			itemStruct{"a", 1},
			itemStruct{"", 1},
			map[string]uint{
				"a": 0,
			},
			errEmptyItem,
		},
		{
			"basic case. already exists",
			map[string]float64{
				"a": 1,
				"b": 10,
			},
			itemStruct{"a", 1},
			itemStruct{"aa", 1},
			map[string]float64{
				"aa": 1,
				"b":  10,
			},
			nil,
		},
		{
			"basic case. item not found",
			map[string]itemStruct{
				"a": 1,
				"b": 10,
			},
			itemStruct{"xxx", 1},
			itemStruct{"aa", 1},
			map[string]itemStruct{
				"a": 1,
				"b": 10,
			},
			errItemNoFound,
		},
	}

	var strBase = &StorageBase{}
	for _, tc := range testCases {
		strBase = &StorageBase{make(StorageIds), make(Products), make(Orders)}
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for name, price := range tc.shop {
				(strBase).AppendPr(name, price)
			}
			err := strBase.UpdatePr(tc.item.nameItem, tc.item.priceItem)

			var a []float64
			var b []float64

			for _, v := range strBase.Prd {
				a = append(a, v)
			}
			for _, v := range tc.expectedShop {
				b = append(b, v)
			}
			sort.Float64s(a)
			sort.Float64s(b)
			if tc.expectedError != nil && err == nil {
				t.Errorf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			}
			if !reflect.DeepEqual(a, b) {
				t.Fatalf("got\t\t%v\nwant\t%v\n%s", strBase.Prd, tc.expectedShop, tc.name)
			}
		})
	}
}
*/
/*
func TestAddAccount(t *testing.T) {
	testCases := []struct {
		name             string
		accounts         map[string]Account
		account          Account
		expectedAccounts map[string]Account
		expectedError    error
	}{
		{
			"basic case. empty Account",
			map[string]Account{},
			Account{},
			map[string]Account{},
			errEmptyAccount,
		},
		{
			"basic case. empty Account name",
			map[string]Account{},
			Account{"", 10},
			map[string]Account{},
			errEmptyAccount,
		},
		{
			"basic case. already exists",
			map[string]Account{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			Account{"a", 10},
			map[string]Account{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			errAccountAlreadyExists,
		},
		{
			"basic case. correct item",
			map[string]Account{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			Account{"xxx", 100},
			map[string]Account{
				"a":   {"a", 1},
				"b":   {"b", 10},
				"xxx": {"xxx", 100},
			},
			nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := AddAccount(tc.accounts, tc.account)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			}
			if !reflect.DeepEqual(tc.accounts, tc.expectedAccounts) {
				t.Fatalf("got\t\t%v\nwant\t%v", tc.account, tc.expectedAccounts)
			}
		})
	}
}
*/
/*
func TestChangeBalance(t *testing.T) {
	testCases := []struct {
		name             string
		accounts         map[string]Account
		account          Account
		expectedAccounts map[string]Account
		expectedError    error
	}{
		{
			"basic case. empty Account",
			map[string]Account{},
			Account{},
			map[string]Account{},
			errEmptyAccount,
		},
		{
			"basic case. empty Account name",
			map[string]Account{},
			Account{"", 10},
			map[string]Account{},
			errEmptyAccount,
		},
		{
			"basic case. correct balance change",
			map[string]Account{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			Account{"a", 10},
			map[string]Account{
				"a": {"a", 10},
				"b": {"b", 10},
			},
			nil,
		},
		{
			"basic case. correct balance change. the same balance",
			map[string]Account{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			Account{"a", 1},
			map[string]Account{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			nil,
		},
		{
			"unknown",
			map[string]Account{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			Account{"xxx", 1},
			map[string]Account{
				"a": {"a", 1},
				"b": {"b", 10},
			},
			errAccountNoFound,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := ChangeBalance(tc.accounts, tc.account)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			}
			if !reflect.DeepEqual(tc.accounts, tc.expectedAccounts) {
				t.Fatalf("got\t\t%v\nwant\t%v", tc.accounts, tc.expectedAccounts)
			}
		})
	}
}

func TestSortAccounts(t *testing.T) {
	testCases := []struct {
		name             string
		accounts         map[string]Account
		sortBy           SortBy
		expectedAccounts []Account
		expectedError    error
	}{
		// name
		{
			"SortByNameAsc. empty Accounts",
			map[string]Account{},
			SortByNameAsc,
			[]Account{},
			nil,
		},
		{
			"SortByNameAsc. single Account ",
			map[string]Account{
				"a": {"a", 10},
			},
			SortByNameAsc,
			[]Account{{"a", 10}},
			nil,
		},
		{
			"SortByNameAsc. already sorted",
			map[string]Account{
				"a":     {"a", 1},
				"b":     {"b", 10},
				"d":     {"d", 12},
				"d10":   {"d10", 11},
				"xxx_1": {"xxx_1", 22},
			},
			SortByNameAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d", 12},
				{"d10", 11},
				{"xxx_1", 22},
			},
			nil,
		},
		{
			"SortByNameAsc. already sorted in reverse order",
			map[string]Account{
				"xxx_1": {"xxx_1", 22},
				"d10":   {"d10", 11},
				"d":     {"d", 12},
				"b":     {"b", 10},
				"a":     {"a", 1},
			},
			SortByNameAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d", 12},
				{"d10", 11},
				{"xxx_1", 22},
			},
			nil,
		},
		{
			"SortByNameAsc. random order",
			map[string]Account{
				"d10":   {"d10", 11},
				"a":     {"a", 1},
				"xxx_1": {"xxx_1", 22},
				"b":     {"b", 10},
				"d":     {"d", 12},
			},
			SortByNameAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d", 12},
				{"d10", 11},
				{"xxx_1", 22},
			},
			nil,
		},

		{
			"SortByNameDesc. empty Accounts",
			map[string]Account{},
			SortByNameDesc,
			[]Account{},
			nil,
		},
		{
			"SortByNameDesc. single Account ",
			map[string]Account{
				"a": {"a", 10},
			},
			SortByNameDesc,
			[]Account{{"a", 10}},
			nil,
		},
		{
			"SortByNameDesc. already sorted",
			map[string]Account{
				"xxx_1": {"xxx_1", 22},
				"d10":   {"d10", 11},
				"d":     {"d", 12},
				"b":     {"b", 10},
				"a":     {"a", 1},
			},
			SortByNameDesc,
			[]Account{
				{"xxx_1", 22},
				{"d10", 11},
				{"d", 12},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByNameDesc. already sorted in reverse order",
			map[string]Account{
				"a":     {"a", 1},
				"b":     {"b", 10},
				"d":     {"d", 12},
				"d10":   {"d10", 11},
				"xxx_1": {"xxx_1", 22},
			},
			SortByNameDesc,
			[]Account{
				{"xxx_1", 22},
				{"d10", 11},
				{"d", 12},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByNameDesc. random order",
			map[string]Account{
				"d10":   {"d10", 11},
				"a":     {"a", 1},
				"xxx_1": {"xxx_1", 22},
				"b":     {"b", 10},
				"d":     {"d", 12},
			},
			SortByNameDesc,
			[]Account{
				{"xxx_1", 22},
				{"d10", 11},
				{"d", 12},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},

		// balance
		{
			"SortByBalanceAsc. empty Accounts",
			map[string]Account{},
			SortByBalanceAsc,
			[]Account{},
			nil,
		},
		{
			"SortByBalanceAsc. single Account",
			map[string]Account{
				"a": {"a", 10},
			},
			SortByBalanceAsc,
			[]Account{{"a", 10}},
			nil,
		},
		{
			"SortByBalanceAsc. already sorted",
			map[string]Account{
				"a":     {"a", 1},
				"b":     {"b", 10},
				"d10":   {"d10", 11},
				"d":     {"d", 12},
				"xxx_1": {"xxx_1", 22},
			},
			SortByBalanceAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d10", 11},
				{"d", 12},
				{"xxx_1", 22},
			},
			nil,
		},
		{
			"SortByBalanceAsc. already sorted with duplicates",
			map[string]Account{
				"a":     {"a", 1},
				"b":     {"b", 10},
				"d10":   {"d10", 11},
				"d11":   {"d11", 11},
				"d":     {"d", 12},
				"xxx_1": {"xxx_1", 22},
				"xxx_2": {"xxx_2", 22},
			},
			SortByBalanceAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d10", 11},
				{"d11", 11},
				{"d", 12},
				{"xxx_1", 22},
				{"xxx_2", 22},
			},
			nil,
		},
		{
			"SortByBalanceAsc. already sorted in reverse order",
			map[string]Account{
				"xxx_1": {"xxx_1", 22},
				"d":     {"d", 12},
				"d10":   {"d10", 11},
				"b":     {"b", 10},
				"a":     {"a", 1},
			},
			SortByBalanceAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d10", 11},
				{"d", 12},
				{"xxx_1", 22},
			},
			nil,
		},
		{
			"SortByBalanceAsc. already sorted in reverse order with duplicated",
			map[string]Account{
				"xxx_2": {"xxx_2", 22},
				"xxx_1": {"xxx_1", 22},
				"d":     {"d", 12},
				"d11":   {"d11", 11},
				"d10":   {"d10", 11},
				"b":     {"b", 10},
				"a":     {"a", 1},
			},
			SortByBalanceAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d10", 11},
				{"d11", 11},
				{"d", 12},
				{"xxx_1", 22},
				{"xxx_2", 22},
			},
			nil,
		},
		{
			"SortByBalanceAsc. random order",
			map[string]Account{
				"d10":   {"d10", 11},
				"a":     {"a", 1},
				"xxx_1": {"xxx_1", 22},
				"b":     {"b", 10},
				"d":     {"d", 12},
			},
			SortByBalanceAsc,
			[]Account{
				{"a", 1},
				{"b", 10},
				{"d10", 11},
				{"d", 12},
				{"xxx_1", 22},
			},
			nil,
		},
		{
			"SortByBalanceAsc. random order with duplicates",
			map[string]Account{
				"d10":   {"d10", 11},
				"a":     {"a", 1},
				"a1":    {"a1", 1},
				"xxx_1": {"xxx_1", 22},
				"xxx_2": {"xxx_2", 22},
				"b":     {"b", 10},
				"d":     {"d", 12},
			},
			SortByBalanceAsc,
			[]Account{
				{"a", 1},
				{"a1", 1},
				{"b", 10},
				{"d10", 11},
				{"d", 12},
				{"xxx_1", 22},
				{"xxx_2", 22},
			},
			nil,
		},

		{
			"SortByBalanceDesc. empty Accounts",
			map[string]Account{},
			SortByBalanceDesc,
			[]Account{},
			nil,
		},
		{
			"SortByBalanceDesc. single Account",
			map[string]Account{
				"a": {"a", 10},
			},
			SortByBalanceDesc,
			[]Account{{"a", 10}},
			nil,
		},
		{
			"SortByBalanceDesc. already sorted",
			map[string]Account{
				"xxx_1": {"xxx_1", 22},
				"d":     {"d", 12},
				"d10":   {"d10", 11},
				"b":     {"b", 10},
				"a":     {"a", 1},
			},
			SortByBalanceDesc,
			[]Account{
				{"xxx_1", 22},
				{"d", 12},
				{"d10", 11},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByBalanceDesc. already sorted with duplicates",
			map[string]Account{
				"xxx_1": {"xxx_1", 22},
				"xxx_2": {"xxx_2", 22},
				"d":     {"d", 12},
				"d11":   {"d11", 11},
				"d10":   {"d10", 11},
				"b":     {"b", 10},
				"a":     {"a", 1},
			},
			SortByBalanceDesc,
			[]Account{
				{"xxx_1", 22},
				{"xxx_2", 22},
				{"d", 12},
				{"d10", 11},
				{"d11", 11},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByBalanceDesc. already sorted in reverse order",
			map[string]Account{
				"a":     {"a", 1},
				"b":     {"b", 10},
				"d10":   {"d10", 11},
				"d":     {"d", 12},
				"xxx_1": {"xxx_1", 22},
			},
			SortByBalanceDesc,
			[]Account{
				{"xxx_1", 22},
				{"d", 12},
				{"d10", 11},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByBalanceDesc. already sorted in reverse order with duplicated",
			map[string]Account{
				"a":     {"a", 1},
				"b":     {"b", 10},
				"d10":   {"d10", 11},
				"d11":   {"d11", 11},
				"d":     {"d", 12},
				"xxx_1": {"xxx_1", 22},
				"xxx_2": {"xxx_2", 22},
			},
			SortByBalanceDesc,
			[]Account{
				{"xxx_1", 22},
				{"xxx_2", 22},
				{"d", 12},
				{"d10", 11},
				{"d11", 11},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByBalanceDesc. random order",
			map[string]Account{
				"d10":   {"d10", 11},
				"a":     {"a", 1},
				"xxx_1": {"xxx_1", 22},
				"b":     {"b", 10},
				"d":     {"d", 12},
			},
			SortByBalanceDesc,
			[]Account{
				{"xxx_1", 22},
				{"d", 12},
				{"d10", 11},
				{"b", 10},
				{"a", 1},
			},
			nil,
		},
		{
			"SortByBalanceDesc. random order with duplicates",
			map[string]Account{
				"d10":   {"d10", 11},
				"a":     {"a", 1},
				"a1":    {"a1", 1},
				"xxx_1": {"xxx_1", 22},
				"xxx_2": {"xxx_2", 22},
				"b":     {"b", 10},
				"d":     {"d", 12},
			},
			SortByBalanceDesc,
			[]Account{
				{"xxx_1", 22},
				{"xxx_2", 22},
				{"d", 12},
				{"d10", 11},
				{"b", 10},
				{"a", 1},
				{"a1", 1},
			},
			nil,
		},

		// unknown
		{
			"Unknown order on empty accounts",
			map[string]Account{},
			SortBy(math.MaxUint8),
			[]Account{},
			nil,
		},
		{
			"Unknown order error.",
			map[string]Account{
				"a": {"a", 1},
			},
			SortBy(math.MaxUint8),
			nil,
			errAccountUnknownSort,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, err := SortAccounts(tc.accounts, tc.sortBy)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("got\t\t%v\nwant\t%v", err, tc.expectedError)
			}
			if !reflect.DeepEqual(res, tc.expectedAccounts) {
				t.Fatalf("got\t\t%v\nwant\t%v", res, tc.expectedAccounts)
			}
		})
	}
}
*/
