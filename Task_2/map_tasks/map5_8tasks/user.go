package main

import (
	"fmt"
	"sort"
	"strings"
)

type IUser interface {
	String()
	PrintSortedNames(int)
	PrintByBalance()
}

type User map[string]float64 // [имя]счёт

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
func (usr *User) PrintSortedNames(reverse int) {
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
func (usr *User) PrintByBalance() {

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
