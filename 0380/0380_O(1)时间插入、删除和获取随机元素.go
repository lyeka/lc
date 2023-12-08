package insertDeleteGetrandomO1

import (
	"math/rand"
)

//leetcode submit region begin(Prohibit modification and deletion)
type RandomizedSet struct {
	index map[int]int
	data  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{index: map[int]int{}, data: []int{}}

}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.index[val]
	if ok {
		return false
	}
	this.data = append(this.data, val)
	this.index[val] = len(this.data) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.index[val]
	if !ok {
		return false
	}

	last := len(this.data) - 1
	this.data[index] = this.data[last]
	this.index[this.data[index]] = index
	this.data = this.data[:last]
	delete(this.index, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.Intn(len(this.data))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
//leetcode submit region end(Prohibit modification and deletion)
