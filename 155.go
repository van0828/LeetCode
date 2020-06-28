package main

func main() {
	
}

//设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
//
//push(x) -- 将元素 x 推入栈中。
//pop() -- 删除栈顶的元素。
//top() -- 获取栈顶元素。
//getMin() -- 检索栈中的最小元素。
//示例:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//min-stack
type MinStack struct {
	ValueAndMin []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{ValueAndMin:make([]int,0)}
}


func (this *MinStack) Push(x int)  {
	if len(this.ValueAndMin) <= 0 {
		this.ValueAndMin = append(this.ValueAndMin,x)
		this.ValueAndMin = append(this.ValueAndMin,x)
		return
	}
	min := this.GetMin()
	this.ValueAndMin = append(this.ValueAndMin,x)
	if x < min {
		this.ValueAndMin = append(this.ValueAndMin,x)
	} else {
		this.ValueAndMin = append(this.ValueAndMin,min)
	}
}

func (this *MinStack) Pop()  {
	this.ValueAndMin = this.ValueAndMin[:len(this.ValueAndMin)-2]
}

func (this *MinStack) Top() int {
	return this.ValueAndMin[len(this.ValueAndMin)-2]
}

func (this *MinStack) GetMin() int {
	return this.ValueAndMin[len(this.ValueAndMin)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
