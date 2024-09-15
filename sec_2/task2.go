package main

import (
	"fmt"
)

type ListNode struct{
	Val int 
	Next *ListNode
}

type Pop struct{
	val int
}

func createList(nums []int) *ListNode{
	if len(nums) == 0{
		return nil
	}
	
	head := &ListNode{Val: nums[0]}
	current := head

	for _, num := range nums[1:]{
		current.Next = &ListNode{Val: num}
		current = current.Next
	}

	return head
}

func printList(head *ListNode) {
	for head != nil {
	 fmt.Print(head, " ",head.Val, " ", head.Next, "\n")
	 head = head.Next
	}
	fmt.Println()
}

func (head *ListNode)deleteDublicate() *ListNode{
	if head == nil{
		return nil
	}

	current := head
	for current != nil && current.Next != nil{
		if current.Val == current.Next.Val{

			current.Next = current.Next.Next
		}else{

			current = current.Next
		}
	}

	return head
}

func main(){
	s := createList([]int{1, 2, 3, 3, 4,5 })
	printList(s)
	fmt.Println()
	s.deleteDublicate()
	printList(s)
}