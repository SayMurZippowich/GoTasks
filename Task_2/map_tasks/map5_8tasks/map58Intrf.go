package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

type Carter interface {
	GetTotal(*StorageBase) float64
	GetSotedKeys() []string
	MakeSrtByCart(*StorageBase, *[]string) string
	GetTotalMulti(*StorageBase) float64
	IsAllowedFor(*StorageBase, User, string) bool
}

type Keeper interface {
	GetPriceByName(string) float64
	AppendPr(string, float64) error
	UpdatePr(string, float64) error
}

type IUser interface {
	String()
	PrSortedNames(int)
	PrByBalance()
}

// типы для реализации магазина
type UserCart map[string]int    // корзина[товар]количество
type StorageIds map[string]uint // [товар]id
type Products map[uint]float64  // [id_товара]цена
// строка заказа вида ["|id:кол-во|id:кол-во..."]суммаЗаказа
type Orders map[string]float64
type User map[string]float64 // [имя]счёт

type StorageBase struct {
	Ids StorageIds
	Prd Products
	Ord Orders
}

type AppendErr string

func (err AppendErr) Error() string {
	return fmt.Sprintf("%s", err)
}

// User methods

func (usr *User) String() string {
	var str strings.Builder
	for k, v := range *usr {
		str.WriteString(
			fmt.Sprintf("%s : %f\n", k, v))
	}
	return str.String()
}

// вывести пользователей и их баланс
// имена пользователей идут в алфавитном порядке
// если reverse == 0
// при любом ином значении
// значений выводятся в обратном алф. порядке
func (usr *User) PrSortedNames(reverse int) {
	keys := make([]string, 0, len(*usr))
	for k := range *usr {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	if reverse == 0 {
		for _, name := range keys {
			fmt.Println(name, " : ", (*usr)[name])
		}
	} else {
		for i := len(keys) - 1; i >= 0; i-- {
			fmt.Println(keys[i], " : ", (*usr)[keys[i]])
		}
	}
}

// По убыванию средств пользователей
func (usr *User) PrByBalance() {

	type kVal struct {
		Key   string
		Value float64
	}

	sl := make([]kVal, 0, len(*usr))
	for k, v := range *usr {
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

// Keeper methods
func (strBase *StorageBase) GetPriceByName(prdName string) float64 {
	return strBase.Prd[strBase.Ids[prdName]]
}

func (strBase *StorageBase) AppendPr(product string, cost float64) error {
	if _, ok := strBase.Ids[product]; product == "" || cost <= 0 || ok {
		return AppendErr("Unable to append! Possible problems: cost < 0, empty product str, prd. already in base")
	}

	// получить новое id
	idNew := (strBase.Ids).GetUniqId()
	// записать новые продукт и id в карту [продукты]id
	strBase.Ids[product] = idNew
	// записать новые id и центу в карту [id]цена
	strBase.Prd[idNew] = cost

	return nil
}

// получить id на единицу большее максимального
func (ids *StorageIds) GetUniqId() uint {
	var max uint = 0
	for _, val := range *ids {
		if val > max {
			max = val
		}
	}
	return max + 1
}

// обновить товар
func (strBase *StorageBase) UpdatePr(product string, cost float64) error {
	if _, ok := strBase.Ids[product]; product == "" || cost <= 0 || !ok {
		return AppendErr("Unable to append! Possible problems: cost < 0, empty product str, prd. not in base")
	}
	id := strBase.Ids[product]
	strBase.Prd[id] = cost
	return nil
}

// Carter methods
// получить сумму заказа
func (crt *UserCart) GetTotal(strBase *StorageBase) float64 {
	total := 0.0
	for name, number := range *crt {
		total += strBase.GetPriceByName(name) * float64(number)
	}
	return total
}

// получить из map заказа стоку вида "|0:2|1:5|2:5..."
// где |id товара:кол-во_товара
func (crt *UserCart) MakeSrtByCart(strBase *StorageBase, sortedKeys *[]string) string {
	b := new(bytes.Buffer)
	for _, val := range *sortedKeys {
		fmt.Fprintf(b, "|%d:%d", strBase.Ids[val], (*crt)[val])
	}
	return b.String()
}

// получить срез отсортированных ключей
func (crt *UserCart) GetSotedKeys() []string {
	keys := make([]string, 0, len(*crt))
	for k := range *crt {
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
func (crt *UserCart) GetTotalMulti(strBase *StorageBase) float64 {
	var total float64
	// получить срез отсортированных ключей
	sortedKeys := crt.GetSotedKeys()
	// на основе упорядоченной карты
	// создать ключ-строку
	crtStr := crt.MakeSrtByCart(strBase, &sortedKeys)

	// проверить наличие значения для ключ-строки в ordersBase
	if val, ok := strBase.Ord[crtStr]; ok {
		fmt.Println("The order has already been done by someone - the total is taken from the database")
		// в случае успеха вернуть найденное значение
		return val
	} else {
		// если строка не найдена то посчитать сумму заказа
		fmt.Println("No such order in base - counting total...")
		total = crt.GetTotal(strBase)
		strBase.Ord[crtStr] = total
	}
	return total
}

// считается сумма заказа
// и списывается со счёта если
// у пользователя достаточно средств
func (crt *UserCart) IsAllowedFor(strBase *StorageBase, usersBase User, user string) bool {
	total := crt.GetTotalMulti(strBase)
	if balance := usersBase[user]; balance < total {
		return false
	}
	usersBase[user] -= total
	return true
}

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
	var strBase Keeper = &StorageBase{storeBaseIds, storeBase, ordersBase}

	// корзина пользователя
	var usrCrt Carter = &UserCart{"keyboard": 1, "phone": 1, "lamp": 2}

	fmt.Println("5. Пользователь даёт список заказа, программа должна по map с наименованиями товаров и ценами, посчитать сумму заказа. И сделать метод добавления новых товаров в map, и метод обновления цены уже существующего товара")
	fmt.Println(usrCrt.GetTotal(strBase.(*StorageBase)))
	fmt.Println("Добавим новый товар:")
	(strBase.(*StorageBase)).AppendPr("microphone", 3.25)

	fmt.Println((strBase.(*StorageBase)).Ids)
	fmt.Println("Обновим его цену:")
	(strBase.(*StorageBase)).UpdatePr("microphone", 999.0)
	fmt.Println((strBase.(*StorageBase)).Prd)

	fmt.Println("6. Сделать 1е, но у нас приходит несколько сотен таких списков заказов и мы хотим запоминать уже посчитанные заказы, чтобы если встречается такой же, то сразу говорить его цену без расчёта")
	fmt.Println(usrCrt.GetTotalMulti(strBase.(*StorageBase)))
	fmt.Println(usrCrt.GetTotalMulti(strBase.(*StorageBase)))

	fmt.Println("7. К 2 добавить, чтобы хранились пользовательские аккаунты со счетом типа \"вася: 300р, петя: 30000000р\". И перед оформлением заказа, но после его расчёта мы проверяли, а есть ли деньги у пользователя, и если есть, то списывали сумму заказа.")
	user := "Admin"
	fmt.Println("Счёт пользователя", usersBase[user])
	fmt.Println("Достаточно ли средств:", usrCrt.IsAllowedFor(strBase.(*StorageBase), usersBase, user))
	fmt.Println("Счёт пользователя после сделки", usersBase[user])

	fmt.Println("8. Есть map аккаунтов и счетов, как описано в 3. Надо вывести ее в отсортированном виде с сортировкой: по имени в алфавитном порядке, по имени в обратном порядке, по количеству денег по убыванию")
	fmt.Println()
	usersBase.PrSortedNames(0)
	fmt.Println("\nОбратный порядок:")
	usersBase.PrSortedNames(1)
	fmt.Println("\nПо убыванию средств пользователей:")
	usersBase.PrByBalance()
}
