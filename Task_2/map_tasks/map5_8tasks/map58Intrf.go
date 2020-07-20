package main

import (
	"fmt"
)

// Carter methods
func main() {

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

	//
	// база магазина
	var strBase = &StorageBase{storeBaseIds, storeBase, ordersBase}

	// корзина пользователя
	var usrCrt Carter = &UserCart{"keyboard": 1, "phone": 1, "lamp": 2}

	fmt.Println("5. Пользователь даёт список заказа, программа должна по map с наименованиями товаров и ценами, посчитать сумму заказа. И сделать метод добавления новых товаров в map, и метод обновления цены уже существующего товара")
	fmt.Println(usrCrt.GetTotal(strBase))
	fmt.Println("Добавим новый товар:")
	(strBase).AppendPr("microphone", 3.25)

	fmt.Println((strBase).Ids)
	fmt.Println("Обновим его цену:")
	(strBase).UpdatePr("microphone", 999.0)
	fmt.Println((strBase).Prd)

	fmt.Println("6. Сделать 1е, но у нас приходит несколько сотен таких списков заказов и мы хотим запоминать уже посчитанные заказы, чтобы если встречается такой же, то сразу говорить его цену без расчёта")
	fmt.Println(usrCrt.GetTotalMulti(strBase))
	fmt.Println(usrCrt.GetTotalMulti(strBase))

	fmt.Println("7. К 2 добавить, чтобы хранились пользовательские аккаунты со счетом типа \"вася: 300р, петя: 30000000р\". И перед оформлением заказа, но после его расчёта мы проверяли, а есть ли деньги у пользователя, и если есть, то списывали сумму заказа.")
	user := "Admin"
	fmt.Println("Счёт пользователя", usersBase[user])
	fmt.Println("Достаточно ли средств:", usrCrt.IsAllowedFor(strBase, usersBase, user))
	fmt.Println("Счёт пользователя после сделки", usersBase[user])

	fmt.Println("8. Есть map аккаунтов и счетов, как описано в 3. Надо вывести ее в отсортированном виде с сортировкой: по имени в алфавитном порядке, по имени в обратном порядке, по количеству денег по убыванию")
	fmt.Println()
	usersBase.PrintSortedNames(0)
	fmt.Println("\nОбратный порядок:")
	usersBase.PrintSortedNames(1)
	fmt.Println("\nПо убыванию средств пользователей:")
	usersBase.PrintByBalance()
}
