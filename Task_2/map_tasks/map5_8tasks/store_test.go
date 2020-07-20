package main

import (
	"testing"
)

var storeBase = Products{
	0: 2.35,
	1: 5.25,
	2: 152.0,
	3: 3.0,
}

// товары Id
var storeBaseIds = StorageIds{
	"lamp":      0,
	"keyboard":  1,
	"phone":     2,
	"usb_cable": 3,
}

// заказы
var ordersBase = make(Orders)

// баланс пользователей
var usersBase = User{
	"Admin":    454.00,
	"User":     55.45,
	"RichUser": 1150,
}

// Ожидается: сумма данного заказа равнa 161.95
func TestGetTotal_161_95(t *testing.T) {
	// база магазина
	var strBase = &StorageBase{storeBaseIds, storeBase, ordersBase}

	var usrCrt Carter = &UserCart{"keyboard": 1, "phone": 1, "lamp": 2}

	v := usrCrt.GetTotal(strBase)

	if v != 161.95 {
		t.Fatal()
	}
}

// Ожидается: новое id = 4
func TestGetCorrId_4(t *testing.T) {
	// база магазина
	var strBase = &StorageBase{storeBaseIds, storeBase, ordersBase}

	vPre := (strBase).Ids
	v := vPre.GetUniqId()

	if v != 4 {
		t.Fatal()
	}
}

// Ожидается: отсутствие ошибки
func TestAppendPr_Nil(t *testing.T) {
	// база магазина
	var strBase = &StorageBase{storeBaseIds, storeBase, ordersBase}

	v := (strBase).AppendPr("microphone", 3.25)

	if v != nil {
		t.Fatal()
	}
}

// Ожидается: отсутствие ошибки
func TestUpdatePr_Nil(t *testing.T) {
	// база магазина
	var strBase = &StorageBase{storeBaseIds, storeBase, ordersBase}

	v := (strBase).UpdatePr("phone", 3.25)

	if v != nil {
		t.Fatal()
	}
}

// Ожидается: будут получены два значения
func TestGetTotalMulti_161_95(t *testing.T) {
	// база магазина
	var strBase = &StorageBase{storeBaseIds, storeBase, ordersBase}

	var usrCrt Carter = &UserCart{"keyboard": 1, "phone": 1, "lamp": 2}

	v := usrCrt.GetTotalMulti(strBase)
	v2 := usrCrt.GetTotalMulti(strBase)

	if v != 161.95 && v2 != 161.95 {
		t.Fatal()
	}
}

// Ожидается, что Admin может купить товары из заданной корзины, а User нет
func TestIsAllowedFor_True(t *testing.T) {
	// база магазина
	var strBase = &StorageBase{storeBaseIds, storeBase, ordersBase}

	var usrCrt Carter = &UserCart{"keyboard": 1, "phone": 1, "lamp": 2}

	v := usrCrt.IsAllowedFor(strBase, usersBase, "Admin")
	v2 := usrCrt.IsAllowedFor(strBase, usersBase, "User")

	if !v && v2 {
		t.Fatal()
	}
}
