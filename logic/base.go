package logic

import (
	"github.com/kudoochui/kudos/log"
	. "mahjong-server/util/intlist"
)

func getCardType(card int) int {
	return card/9
}

func getCardValue(card int) int {
	card = (card + 1)%9
	if card == 0 {
		card = 9
	}
	return card-1
}

func getCardSlice(cards *[]int) *[27]int {
	s := [27]int{}
	for _,card := range *cards {
		s[card]++
	}
	return &s
}

func checkLack(cards *[]int,lack int) bool {
	for i:=len(*cards)-1;i>=0;i-- {
		if getCardType((*cards)[i]) == lack {
			return false
		}
	}
	return true
}

// 巧七对
func isSevenPair(cards *[]int,array *[27]int) bool {
	if array == nil {
		array = getCardSlice(cards)
	}

	if len(*cards) != 14 {
		return false
	}

	// 要保证ar数据正常
	for _, cnt := range array {
		if !(cnt == 2 || cnt == 0) {
			return false
		}
	}
	return true
}

// 龙七对
func isBigSevenPair(cards *[]int,array *[27]int) bool {
	if array == nil {
		array = getCardSlice(cards)
	}

	if len(*cards) != 14 {
		return false
	}

	// 要保证ar数据正常
	big := false
	for _, cnt := range array {
		if !(cnt == 2 || cnt == 4 || cnt == 0) {
			return false
		}
		if cnt == 4 {
			big = true
		}
	}
	return big
}

// 清一色
func isSameColor(cards *[]int) bool {
	if len(*cards) == 0 {
		return false
	}
	color := getCardValue((*cards)[0])
	for _,card := range *cards {
		if getCardType(card) != color {
			return false
		}
	}
	return true
}

// 大单吊
func isSinglePair(cards *[]int) bool {
	if len(*cards) == 2 && (*cards)[0] == (*cards)[1] {
		return true
	}
	return false
}

// 大对子
func isBigPair(cards *[]int,array *[27]int) bool {
	if array == nil {
		array = getCardSlice(cards)
	}
	if len(*cards) != 5 {
		return false
	}

	// 要保证ar数据正常
	for _, cnt := range array {
		if !(cnt == 3 || cnt == 2 || cnt == 0) {
			return false
		}
	}
	return true
}

func isSomeGroup(markList IntList,array [27]int) bool {
	checkGroup := func(array [27]int) bool {
		var index int
		for i:=0;i<3;i++{
			for j:=0;j<7;j++{
				index = i*9+j
				if array[index] > 0 {
					array[index+1] -= array[index]
					array[index+2] -= array[index]
					array[index] = 0

					if array[index+1] < 0 || array[index+2] < 0 {
						return false
					}
				}
			}
			if array[i*9+7] > 0 || array[i*9+8] > 0 {
				return false
			}
		}
		return true
	}

	getMark := func(list IntList,array *[27]int) int {
		for card,cnt := range array {
			if cnt >= 3 && list.IndexOf(card) == 0 {
				return card
			}
		}
		return -1
	}

	if mark := getMark(markList,&array); mark >= 0 {
		markList.Insert(mark)
		if isSomeGroup(markList.Clone(),array) {
			return true
		}else{
			array[mark] = array[mark] - 3
			return  isSomeGroup(markList,array)
		}
	}else{
		return checkGroup(array)
	}

	return false
}

func isNormalH(cards *[]int,array [27]int) bool {
	//一个对子+N坎牌（三个相同的牌、三个数值连续牌)
	if (len(*cards)-2)%3 != 0 {
		//打印错误，牌数错误
		return false
	}

	for i:=0;i<len(array);i++ {
		if array[i] >= 2 {
			array[i] -= 2
			if isSomeGroup(IntList{},array) {
				log.Info("true-2")
				return true
			}
			array[i] += 2
		}
	}

	return false
}

func checkHu(cards []int,lack int,args...int) bool {
	log.Info(cards)
	log.Info(lack)

	if !checkLack(&cards,lack) {
		return false
	}

	array := getCardSlice(&cards)
	if isBigSevenPair(&cards,array) || isSevenPair(&cards,array) {
		log.Info("true-1")
		return true
	}

	return isNormalH(&cards,*array)
}

func arrangeCards(cards *[]int) map[int][]int {
	array := make(map[int][]int)
	for _,c := range *cards {
		typ := getCardType(c)
		val := getCardValue(c)
		if array[typ] == nil {
			array[typ] = make([]int,9)
		}
		array[typ][val]++
	}
	return array
}
