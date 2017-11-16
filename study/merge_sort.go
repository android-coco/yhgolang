package main


type Px struct {
	Name string
	TestTime string
	Vale string
}

type XX struct {
	Name string
	TestTime string
	Vale1 string
	Vale2 string
}
func main() {
	//a := []Px{Px{"1","2017-10-31 12:00:00","2"},Px{"2","2017-10-31 13:00:00","3"}}
	//b := []Px{Px{"1","2017-10-31 12:01:00","2"},Px{"2","2017-10-31 13:01:00","3"}}
}

func MergeSort(value ...[]Px)  {
	total := 0
	for _, num := range value {
		total += len(num)
	}
}

