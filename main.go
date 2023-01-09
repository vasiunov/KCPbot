package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

func main() {
	const percOQuota = 10 // доля особой квоты
	const percSQuota = 10 // доля специальной квоты
	var intKCP, percTQuota, count, count2 int

	intKCP = getInput("общий показатель КЦП :")
	percTQuota = getInput("долю целевой квоты :")

	flKCP := float64(intKCP)

	O := int(math.Ceil(flKCP * percOQuota / 100))             // количество мест по особой квоте
	S := int(math.Round(flKCP * percSQuota / 100.0))          // количество мест по специальной квоте
	T := int(math.Round(flKCP * float64(percTQuota) / 100.0)) // количество мест по целевой квоте

	fmt.Println("Общие сведения")
	fmt.Printf("КЦП\tОсобая\tСпец\tЦелевая\tПеребор\n")
	fmt.Printf("%v\t%v\t%v\t%v\t%v\n", intKCP, O, S, T, O+S+T-intKCP)

	getVariants(intKCP, count, count2, O, S, T)

}

func getInput(message string) (result int) {
	fmt.Printf("Введите %s", message)

	reader := bufio.NewReader(os.Stdin)
	sliceInput, _, err := reader.ReadLine() // читаем с консоли в байтовый слайс
	if err != nil {
		fmt.Println(err)
		getInput(message)
	}

	input := string(sliceInput)

	for _, elem := range input {
		if !unicode.IsDigit(elem) || unicode.IsSpace(elem) { // проверяем строку на наличие пробелов и не цифр
			fmt.Println("Пожалуйста, введите число без лишних символов") // если срабатывает, возвращает 0
			getInput(message)
			break
		}
	}

	result, _ = strconv.Atoi(input)
	return
}

func getVariants(intKCP, count, count2, O, S, T int) {
	for o := 0; o < intKCP; o++ { // перебор особой квоты
		for s := 0; s < intKCP; s++ { // перебор специальной квоты
			for t := 0; t < intKCP; t++ { // перебор целевой квоты
				for os := 0; os < intKCP; os++ { // перебор особой+специальной квоты
					for ot := 0; ot < intKCP; ot++ { // перебор особой+целевой квоты
						for st := 0; st < intKCP; st++ { // перебор специальной+целевой квоты
							for ost := 0; ost < intKCP; ost++ { // перебор особой+специальной+целевой квоты
								if os+ot+ost+o == O && os+st+ost+s == S && ot+st+ost+t == T && intKCP == o+s+t+os+ot+st+ost {
									count++
									// TO DO: писать в файл
									fmt.Println("==================================================")
									fmt.Printf("Вариант %v\n", count)
									fmt.Printf("Плохой вариант %v\n", count2)
									fmt.Printf("О+С\tО+Ц\tС+Ц\tО+С+Ц\tО\tС\tЦ\n")
									fmt.Printf(" %v\t %v\t %v\t  %v\t%v\t%v\t%v\n", os, ot, st, ost, o, s, t)
								} else {
									count2++
								}
							}
						}
					}
				}
			}
		}
	}
}
