package main

import (
	"fmt"
	"sort"
)

func DelElement(sl *[]int, num int) int {
	// сохранение значения
	val := (*sl)[num]
	/*
		элементы справа от удаляемого
		копируются влево на один элемент
		тем самым замещая собой удаляемый элемент
	*/
	copy((*sl)[num:], (*sl)[num+1:])
	// возникает дубликат последнего в списке значения
	// берётся срез без дубликата
	*sl = (*sl)[:len(*sl)-1]
	return val
}

func Difference(slA *[]int, slB *[]int) []int {

	m := make(map[int]bool, len(*slA)) // map[val]existsInB
	// map значений slA
	for _, val := range *slA {
		m[val] = false
	}

	for _, val := range *slB {
		/* если элемент из slB есть в m
		(т.е. присутствует в slА)
		то он помечается как true */
		if _, exists := m[val]; exists {
			m[val] = true
		}
	}

	var diffSl []int

	for val, inB := range m {
		if !inB {
			diffSl = append(diffSl, val)
		}
	}
	return diffSl
}

func ShiftLeftOn(shiftSl *[]int, step int) {

	if lenSl := len(*shiftSl); step > lenSl {
		step -= lenSl
	}

	*shiftSl = append((*shiftSl)[step:], (*shiftSl)[:step]...)
}

func ShiftRightOn(shiftSl *[]int, step int) {

	if lenSl := len(*shiftSl); step > lenSl {
		step -= lenSl
	}
	shiftBorder := len(*shiftSl) - step
	*shiftSl = append((*shiftSl)[shiftBorder:], (*shiftSl)[:shiftBorder]...)

}

func CopySl(base *[]int) []int {
	slCopy := make([]int, len(*base))
	copy(slCopy, *base)
	return slCopy
}

func SliceSwap(sl *[]int) {
	slLen := len(*sl)

	if mod := slLen % 2; mod == 1 {
		// если количество элементов нечетное
		// то последний элемент игнорируется
		slLen--
	}

	for i := 0; i < slLen; i += 2 {
		(*sl)[i], (*sl)[i+1] = (*sl)[i+1], (*sl)[i]
	}
}

func main() {
	// 1. Прибавить 1 к каждому элементу []int
	intSl := []int{-1, 0, 1, 2, 3}
	fmt.Println("1. К каждому элементу []int была прибавлена 1")
	fmt.Println("Изначальный срез:", intSl)
	for i := range intSl {
		intSl[i]++
	}
	fmt.Println("Полученный срез:", intSl, "\n")

	fmt.Println("2. Добавим число 5 в конец slice:")
	fmt.Printf("%d -> ", intSl)
	intSl = append(intSl, 5)
	fmt.Println(intSl, "\n")

	fmt.Println("3. Добавим число 5 в начало slice:")
	fmt.Printf("%d -> ", intSl)
	intSl = append([]int{5}, intSl...)
	fmt.Println(intSl, "\n")

	fmt.Println("4. Взять последнее число slice, вернуть его пользователю, а из slice этот элемент удалить:")
	fmt.Printf("%d -> ", intSl)
	elm, intSl := intSl[len(intSl)-1], intSl[:len(intSl)-1]
	fmt.Printf("%d полученный элемент %d\n\n", intSl, elm)

	fmt.Println("5. Взять первое число slice, вернуть его пользователю, а из slice этот элемент удалить:")
	fmt.Printf("%d -> ", intSl)
	elm, intSl = intSl[0], intSl[1:]
	fmt.Printf("%d полученный элемент %d\n\n", intSl, elm)

	fmt.Println("6. Взять i-е число slice, вернуть его пользователю, а из slice этот элемент удалить. Число i передает пользователь в функцию:")
	fmt.Printf("%d -> ", intSl)
	elm = DelElement(&intSl, 2)
	fmt.Printf("%d полученный элемент %d\n\n", intSl, elm)

	fmt.Println("7. Объединить два slice и вернуть новый со всеми элементами первого и второго:")
	unionArr := []int{9, 9, 9}
	fmt.Printf("%d + %d = ", intSl, unionArr)
	unionSl := append(intSl, unionArr...)
	fmt.Println(unionSl, "\n")

	fmt.Println("8. Из первого slice удалить все числа, которые есть во втором:")
	minusSl := []int{0, 3, 9, 101}
	fmt.Printf("%d - %d = ", unionSl, minusSl)
	fmt.Println(Difference(&unionSl, &minusSl), "\n")

	fmt.Println("9. Сдвинуть все элементы slice на 1 влево. Нулевой становится последним, первый - нулевым, последний - предпоследним:")
	shiftArr := []int{9, 1, 2, 3, 4}
	fmt.Printf("%d <-1 \n", shiftArr)
	shiftArr = append(shiftArr[1:], shiftArr[0])
	fmt.Println(shiftArr, "//shifted one step left\n")

	fmt.Println("10. Тоже, но сдвиг на заданное пользователем i:")
	shiftLeftArr := []int{4, 5, 6, 7, 8, 9}
	fmt.Printf("%d <--- \n", shiftLeftArr)
	ShiftLeftOn(&shiftLeftArr, 9)
	fmt.Println(shiftLeftArr, "//shifted multiple steps left\n")

	fmt.Println("11. Cдвиг вправо на 1:")
	fmt.Printf("%d 1-> \n", shiftArr)
	shiftArrLen := len(shiftArr) - 1
	shiftArr = append([]int{shiftArr[shiftArrLen]}, shiftArr[:shiftArrLen]...)
	fmt.Println(shiftArr, "//shifted one step right\n")

	fmt.Println("12. Сдвиг вправо на задaнное количество шагов:")
	shiftRightArr := []int{3, 4, 5, 6, 1, 2}
	fmt.Printf("%d ---> \n", shiftRightArr)
	ShiftRightOn(&shiftRightArr, 2)
	fmt.Println(shiftRightArr, "//shifted multiple steps right\n")

	fmt.Println("13. Вернуть пользователю копию переданного slice:")
	baseSlice := []int{9, 9, 9}
	fmt.Printf("Base slice: %d\n", baseSlice)
	copiedSlice := CopySl(&baseSlice)
	fmt.Printf("Copy: %d\n", copiedSlice)
	baseSlice[0]++
	fmt.Printf("Modified base slice: %d\n", baseSlice)
	fmt.Printf("Copy: %d\n", copiedSlice)

	fmt.Println("14. В slice поменять все четные с ближайшими нечетными индексами. 0 и 1, 2 и 3, 4 и 5...")
	twistedSlice := []int{2, 1, 4, 3, 6, 5, 7}
	fmt.Printf("Original slice: %d\n", twistedSlice)
	SliceSwap(&twistedSlice)
	fmt.Printf("Result: %d\n\n", twistedSlice)

	fmt.Println("15. Упорядочить slice в порядке: прямом, обратном, лексикографическом:")
	sliceToSort := []int{2, 1, 4, 3, 9, 5, 7}
	fmt.Printf("Sorted slice: %d\n", sliceToSort)
	sort.Ints(sliceToSort)
	fmt.Printf("Sorted in increasing order: %d\n", sliceToSort)

	sliceToSort = []int{2, 1, 4, 3, 9, 5, 7}
	sort.Sort(sort.Reverse(sort.IntSlice(sliceToSort)))
	fmt.Printf("Sorted in decreasing order: %d\n", sliceToSort)

	sliceToSortStr := []string{"zip", "app", "acb", "abc"}
	fmt.Printf("Unsorted strings slice: %s\n", sliceToSortStr)
	sort.Strings(sliceToSortStr)
	fmt.Printf("Sorted strings slice: %s\n", sliceToSortStr)

}
