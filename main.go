package main

import "fmt"
import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
  Val int
  Next *ListNode
}

func recIsPalindrome(firstHalf *ListNode, secondHalf *ListNode) *ListNode {
  if secondHalf == nil {
    return nil
  }

  if secondHalf.Next == nil {
    if firstHalf.Val == secondHalf.Val {
      return firstHalf.Next
    } else {
      return nil
    }
  }

  nextToCheck := recIsPalindrome(firstHalf, secondHalf.Next)
  if nextToCheck != nil && nextToCheck.Val == secondHalf.Val {
    return nextToCheck.Next
  }

  return nil
}

func isPalindrome(head *ListNode) bool {

  if head == nil || head.Next == nil {
    return true
  }

  var length int

  curr := head
  
  for curr != nil {
    length++
    curr = curr.Next
  }

  if length == 2 {
    return head.Val == head.Next.Val
  }

  secondHalf := head

  for i := 0; i <= length/2; i++ {
    secondHalf = secondHalf.Next
  }

  if secondHalf == nil {
    return false
  }

  isPalin := recIsPalindrome(head, secondHalf)

  return isPalin != nil

}

func main() {
  fmt.Println("Hello World")
}