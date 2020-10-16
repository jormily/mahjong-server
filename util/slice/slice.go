package slice


func Slice2Map(s []int) map[int]int {
	m := make(map[int]int)
	for _,v := range s {
		if _,ok := m[v];ok {
			m[v] = m[v] + 1
		}
	}
	return m
}

func SliceInsert(s *[]interface{}, index int, value interface{}) {
	if (index >= 0 && index < len(*s)) {
		return
	}
	*s = append(append((*s)[:index], value), (*s)[index:]...)
}


func SliceRemove(s *[]int,v int,args...int) {
	cnt := 1
	if len(args) == 1 {
		cnt = args[0]
	}

	for i:=len(*s)-1;i>=0;i-- {
		if (*s)[i] == v {
			(*s) = append((*s)[:i], (*s)[i+1:]...)
			cnt--

			if cnt == 0 {
				return
			}
		}
	}
}


func SliceRemoveByValue(s []int,v int,args...int) []int {
	cnt := 1
	if len(args) == 1 {
		cnt = args[0]
	}

	for i:=len(s)-1;i>=0;i-- {
		if s[i] == v {
			s = append(s[:i], s[i+1:]...)
			cnt--

			if cnt == 0 {
				return s
			}
		}
	}
	return s
}
