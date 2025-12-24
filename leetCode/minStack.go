package leetcode

type MinStack struct {
    stack []int
    mins []int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    this.stack = append(this.stack, val)

    if len(this.mins) == 0 {
        this.mins = append(this.mins, val)
        return
    }

    currentMin := this.mins[len(this.mins)-1]
    if val < currentMin {
        this.mins = append(this.mins, val)
    } else {
        this.mins = append(this.mins, currentMin)
    }
}


func (this *MinStack) Pop()  {
    if len(this.stack) == 0 {
        return
    }
    this.stack = this.stack[:len(this.stack)-1]

    this.mins = this.mins[:len(this.mins)-1]
}


func (this *MinStack) Top() int {
    if len(this.stack) == 0 {
        return 0
    }
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.mins) == 0 {
        return 0
    }
    return this.mins[len(this.mins)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */