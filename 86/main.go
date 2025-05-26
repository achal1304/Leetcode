package main

import "fmt"

// ListNode defines the structure for a singly linked list node.
type ListNode struct {
    Val  int
    Next *ListNode
}

// Function to partition the linked list based on value x
func partition(head *ListNode, x int) *ListNode {
    // Create two dummy nodes for the two partitions
    lessHead := &ListNode{}
    greaterHead := &ListNode{}
    
    // Create two pointers to build the two partitions
    less := lessHead
    greater := greaterHead
    
    // Traverse the original list
    current := head
    for current != nil {
        if current.Val < x {
            less.Next = current
            less = less.Next
        } else {
            greater.Next = current
            greater = greater.Next
        }
        current = current.Next
    }
    
    // End the greater list
    greater.Next = nil
    
    // Connect the two lists
    less.Next = greaterHead.Next
    
    return lessHead.Next
}

// Helper function to create a linked list from a slice
func createList(arr []int) *ListNode {
    if len(arr) == 0 {
        return nil
    }
    head := &ListNode{Val: arr[0]}
    current := head
    for i := 1; i < len(arr); i++ {
        current.Next = &ListNode{Val: arr[i]}
        current = current.Next
    }
    return head
}

// Helper function to print the linked list
func printList(head *ListNode) {
    current := head
    for current != nil {
        fmt.Print(current.Val, " ")
        current = current.Next
    }
    fmt.Println()
}

func main() {
    arr := []int{1, 4, 3, 2, 5, 2}
    head := createList(arr)
    fmt.Print("Original List: ")
    printList(head)

    partitioned := partition(head, 3)
    fmt.Print("Partitioned List: ")
    printList(partitioned)
}
