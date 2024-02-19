package _170

type TwoSum struct {
	exists map[int]int
}

func Constructor() TwoSum {
	return TwoSum{exists: make(map[int]int)}
}

func (this *TwoSum) Add(number int) {
	this.exists[number] += 1
}

func (this *TwoSum) Find(value int) bool {
	for one := range this.exists {
		other := value - one
		if count, ok := this.exists[other]; ok {
			if other == one && count == 1 {
				continue
			}
			return true
		}
	}
	return false
}
