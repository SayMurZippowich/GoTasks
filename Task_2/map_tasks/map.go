package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// типы для реализации set
type void struct{}
type SetInt map[int]void

func CountWordsFileHandler(path string) map[string]int {
	// открытие файла
	file, err := os.Open("word_map.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// словарь[слово]частота
	wMap := make(map[string]int)

	// функция для FieldsFunc
	// позволяет игнорировать знаки препинания в тексте
	rule := func(chr rune) bool {
		return !unicode.IsLetter(chr)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// strings.Fields с функцией rule
		strArr := strings.FieldsFunc(scanner.Text(), rule)
		// функция подсчета слов
		CountStr(wMap, &strArr)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wMap
}

func CountStr(m map[string]int, str *[]string) {
	for _, val := range *str {
		m[val]++
	}
}

func GetSetInt(sl *[]int) SetInt {
	set := make(SetInt)
	var member void

	for _, val := range *sl {
		if _, isMember := set[val]; !isMember {
			set[val] = member
		}
	}
	return set
}

func (st SetInt) ToSlice() []int {
	outSl := make([]int, len(st))
	j := 0

	for i := range st {
		outSl[j] = i
		j++
	}
	return outSl
}

func CrossSlice(slA *[]int, slB *[]int) []int {
	var minor, major *[]int
	// выбор меньшего среза для
	// изначальной записи в map
	if len(*slA) <= len(*slB) {
		minor = slA
		major = slB
	} else {
		minor = slB
		major = slA
	}

	crossMap := make(map[int]bool)
	// добавление в map меньшего среза
	for _, val := range *minor {
		crossMap[val] = false
	}

	// цикл по второму срезу
	// для поиска совпадений
	crossCounter := 0
	for _, mjVal := range *major {
		if val, inMin := crossMap[mjVal]; inMin {
			if !val {
				crossCounter++
			}
			crossMap[mjVal] = true
		}
	}
	if crossCounter == 0 {
		return []int{}
	}
	// создание возвращаемого slice
	outSl := make([]int, crossCounter)
	// заполнение slice
	count := 0
	for key, val := range crossMap {
		if val {
			outSl[count] = key
			count++
		}
	}
	return outSl
}

func FibMem(memMap map[int]int, n int) int {
	switch {
	case n == 0:
		return 0
	case n == 1 || n == 2:
		return 1
	// случай нулевого значения типа int (ключ не задан)
	case memMap[n] == 0:
		memMap[n] = FibMem(memMap, n-1) + FibMem(memMap, n-2)
	}
	return memMap[n]
}

func main() {

	fmt.Println("1. Есть текст, надо посчитать сколько раз каждое слово встречается:")
	w_map := CountWordsFileHandler("word_map.txt")
	for key, val := range w_map {
		fmt.Printf("Word: %s | Count: %d\n", key, val)
	}

	fmt.Println("2. Есть очень большой массив или slice целых чисел, надо сказать какие числа в нем упоминаются хоть по разу:")
	testArr := make([]int, 300000)
	for i := 0; i < 300000; i += 3 {
		testArr[i], testArr[i+1], testArr[i+2] = 1, 2, 3
	}
	set := GetSetInt(&testArr)
	fmt.Println(set.ToSlice())

	fmt.Println("3. Есть два больших массива чисел, надо найти, какие числа упоминаются в обоих:")
	testArr2 := make([]int, 100000)
	testArr2[0] = 2
	for i := 1; i < 100000; i++ {
		testArr2[i] = 1
	}
	//testArr[1, 2, 3, 1, 2..] testArr2 [2, 1, 1, 1, ....]
	crossSl := CrossSlice(&testArr, &testArr2)
	fmt.Println(crossSl)

	fmt.Println("4. Сделать Фибоначчи с мемоизацией:")
	memMap := make(map[int]int)
	for i := 0; i < 10; i++ {
		f := FibMem(memMap, i)
		fmt.Printf("%d ", f)
	}
	fmt.Printf("\n")

}
