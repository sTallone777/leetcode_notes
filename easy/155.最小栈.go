type MinStack struct {
	t, b  int
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		t:     -1,
		b:     -1,
		stack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.t++
	if this.t < len(this.stack) {
		this.stack[this.t] = val
	} else {
		this.stack = append(this.stack, val)
	}

	if this.b == -1 {
		this.b = 0
	}
}

func (this *MinStack) Pop() {
	this.t--
}

func (this *MinStack) Top() int {
	return this.stack[this.t]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt32
	for i := range this.stack[this.b : this.t+1] {
		if this.stack[i] < min {
			min = this.stack[i]
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */