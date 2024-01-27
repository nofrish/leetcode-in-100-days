package _170

type TwoSum struct {
	exists map[int]int
}

func Constructor() TwoSum {
	return TwoSum{exists: make(map[int]int, 0)}
}

func (this *TwoSum) Add(number int) {
	this.exists[number] += 1
}

func (this *TwoSum) Find(value int) bool {
	for one := range this.exists {
		other := value - one
		if v, ok := this.exists[other]; ok {
			if other == one && v == 1 {
				continue
			}
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
