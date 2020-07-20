package main

import (
	"bytes"
	"fmt"
	"sort"
)

type Carter interface {
	GetTotal(*StorageBase) float64
	GetSortedKeys() []string
	MakeSrtByCart(*StorageBase, *[]string) string
	GetTotalMulti(*StorageBase) float64
	IsAllowedFor(*StorageBase, User, string) bool
}

type UserCart map[string]int // корзина[товар]количество

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
func (crt *UserCart) GetSortedKeys() []string {
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
	sortedKeys := crt.GetSortedKeys()
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
