package main

import (
	"bytes"
	"fmt"
	"sort"
)

// типы для реализации магазина
type Cart map[string]int        // корзина[товар]количество
type Storage map[string]float64 // [товар]цена
type StorageIds map[string]uint // [товар]id
// строка заказа вида ["|id:кол-во|id:кол-во..."]суммаЗаказа
type Orders map[string]float64
type User map[string]float64 // [имя]счёт

// товары-цена
var storeBase = Storage{
	"lamp":      2.35,
	"keyboard":  5.25,
	"phone":     152.0,
	"usb_cable": 3.0,
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

// получить сумму заказа
func (crt Cart) GetTotal() float64 {
	total := 0.0
	for name, number := range crt {
		total += storeBase[name] * float64(number)
	}
	return total
}

// добавить товар
func (strg Storage) AppendPr(product string, cost float64) bool {
	if product == "" || cost <= 0.0 {
		return false
	}
	strg[product] = cost
	// получить новое id
	new_id := storeBaseIds.GetUniqId()
	// присвоить новое id
	storeBaseIds[product] = new_id
	return true
}

// получить id на единицу большее максимального
func (ids StorageIds) GetUniqId() uint {
	var max uint = 0
	for _, val := range ids {
		if val > max {
			max = val
		}
	}
	return max + 1
}

// обновить товар
func (strg Storage) UpdatePr(product string, cost float64) bool {
	if _, ok := strg[product]; (product == "" || cost <= 0.0) && ok {
		return false
	}
	strg[product] = cost
	return true
}

// получить из map заказа стоку вида "|0:2|1:5|2:5..."
// где |id товара:кол-во_товара
func (crt Cart) MakeSrtByCart(sortedKeys *[]string) string {
	b := new(bytes.Buffer)
	for _, val := range *sortedKeys {
		fmt.Fprintf(b, "|%d:%d", storeBaseIds[val], crt[val])
	}
	return b.String()
}

// получить срез отсортированных ключей
func (crt Cart) GetSotedKeys() []string {
	keys := make([]string, 0, len(crt))
	for k := range crt {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

/* функция преобразует map каждого поступившего заказа
в строку вида "|id_товара:кол-во|id_товара:кол-во..."
в карте ordersBase проверяется наличие такой-же строки-ключа
если по ключ присутствует, то его значение берётся в качестве суммы к оплате
если такого ключа нет то сумма считается при помощи GetTotal()
из записывается в ordersBase, ключом является - полученная ранее строка
*/
func (crt Cart) GetTotalMulti() float64 {
	var total float64
	// получить срез отсортированных ключей
	sortedKeys := crt.GetSotedKeys()
	// на основе упорядоченной карты
	// создать ключ-строку
	crtStr := crt.MakeSrtByCart(&sortedKeys)

	// проверить наличие значения для ключ-строки в ordersBase
	if val, ok := ordersBase[crtStr]; ok {
		fmt.Println("The order has already been done by someone - the total is taken from the database")
		// в случае успеха вернуть найденное значение
		return val
	} else {
		// если строка не найдена то посчитать сумму заказа
		fmt.Println("No such order in base - counting total...")
		total = crt.GetTotal()
		ordersBase[crtStr] = total
	}
	return total
}

// считается сумма заказа
// и списывается со счёта если
// у пользователя достаточно средств
func (crt Cart) IsAllowedFor(user string) bool {
	total := crt.GetTotalMulti()
	if balance := usersBase[user]; balance < total {
		return false
	}
	usersBase[user] -= total
	return true
}

// вывести пользователей и их баланс
// имена пользователей идут в алфавитном порядке
// если reverse == 0
// при любом ином значении
// значений выводятся в обратном алф. порядке
func (usr User) PrSortedNames(reverse int) {
	keys := make([]string, 0, len(usr))
	for k := range usr {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	if reverse == 0 {
		for _, name := range keys {
			fmt.Println(name, " : ", usr[name])
		}
	} else {
		for i := len(keys) - 1; i >= 0; i-- {
			fmt.Println(keys[i], " : ", usr[keys[i]])
		}
	}
}

// По убыванию средств пользователей
func (usr User) PrByBalance() {

	type kVal struct {
		Key   string
		Value float64
	}

	sl := make([]kVal, 0, len(usr))
	for k, v := range usr {
		sl = append(sl, kVal{k, v})
	}

	// передаётся функция задающая используемые значения
	// и порядок сортировки
	sort.Slice(sl, func(i, j int) bool {
		return sl[i].Value > sl[j].Value
	})

	for _, kv := range sl {
		fmt.Printf("%s: %.2f\n", kv.Key, kv.Value)
	}

}

func main() {
	fmt.Println("5. Пользователь даёт список заказа, программа должна по map с наименованиями товаров и ценами, посчитать сумму заказа. И сделать метод добавления новых товаров в map, и метод обновления цены уже существующего товара")
	crt := Cart{"keyboard": 1, "phone": 1, "lamp": 2}
	fmt.Println(crt.GetTotal())
	fmt.Println("Добавим новый товар:")
	storeBase.AppendPr("microphone", 3.25)
	fmt.Println(storeBase)
	fmt.Println("Обновим его цену:")
	storeBase.UpdatePr("microphone", 999.0)
	fmt.Println(storeBase)

	fmt.Println("6. Сделать 1е, но у нас приходит несколько сотен таких списков заказов и мы хотим запоминать уже посчитанные заказы, чтобы если встречается такой же, то сразу говорить его цену без расчёта")
	fmt.Println(crt.GetTotalMulti())
	fmt.Println(crt.GetTotalMulti())
	fmt.Println("7. К 2 добавить, чтобы хранились пользовательские аккаунты со счетом типа \"вася: 300р, петя: 30000000р\". И перед оформлением заказа, но после его расчёта мы проверяли, а есть ли деньги у пользователя, и если есть, то списывали сумму заказа.")
	user := "Admin"
	fmt.Println("Счёт пользователя", usersBase[user])
	fmt.Println("Достаточно ли средств:", crt.IsAllowedFor(user))
	fmt.Println("Счёт пользователя после сделки", usersBase[user])

	fmt.Println("8. Есть map аккаунтов и счетов, как описано в 3. Надо вывести ее в отсортированном виде с сортировкой: по имени в алфавитном порядке, по имени в обратном порядке, по количеству денег по убыванию\n")
	usersBase.PrSortedNames(0)
	fmt.Println("\nОбратный порядок:")
	usersBase.PrSortedNames(1)
	fmt.Println("\nПо убыванию средств пользователей:")
	usersBase.PrByBalance()
}
