package main

/*

解题思路：
利用二分法，用快慢指针法，快指针走两步，慢指针走一步，快指针越界时，慢指针正好到达中点，只需记录慢指针的前一个指针，就能断成两链。
merge 函数做的事是合并两个有序的左右链

    设置虚拟头结点，用 prev 指针去“穿针引线”，prev 初始指向 dummy
    每次都确定 prev.Next 的指向，并注意 l1，l2指针的推进，和 prev 指针的推进
    最后返回 dummy.Next ，即合并后的链
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	//在这里将head ==nil 写成了head !=nil一直运行不对
	if head == nil || head.Next == nil { // 递归的出口，不用排序 直接返回
		return head
	}
	slow, fast := head, head // 快慢指针
	var preSlow *ListNode    // 保存slow的前一个结点
	//fast !=nil这种是判断偶数个节点的情况，fast.Next !=nil判断的是奇数个节点的时候。
	//（偶数的时候，slow指针是偏左的。例如：1,2,3,4 first=nil，slow指向2）
	//注意是且不是或
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next      // 慢指针走一步
		fast = fast.Next.Next // 快指针走两步
	}
	preSlow.Next = nil     // 断开，分成两链
	l := sortList(head)    // 已排序的左链
	r := sortList(slow)    // 已排序的右链
	return mergeList(l, r) // 合并已排序的左右链，一层层向上返回
}

func mergeList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}   // 虚拟头结点
	prev := dummy                // 用prev去扫，先指向dummy
	for l1 != nil && l2 != nil { // l1 l2 都存在
		if l1.Val < l2.Val { // l1值较小
			prev.Next = l1 // prev.Next指向l1
			l1 = l1.Next   // 考察l1的下一个结点
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next // prev.Next确定了，prev指针推进
	}
	if l1 != nil { // l1存在，l2不存在，让prev.Next指向l1
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}
	return dummy.Next // 真实头结点
}
