//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
//
// Related Topics 链表 分治 堆（优先队列） 归并排序 👍 2489 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	// 不要过度解读，这里数组每一个有序的链表都是有序衔接的
	// 直接顺序合并是可以的，每一个都是一次合并两个有序链表
	// 这里采取分治优化一下
	return merge(lists, 0, len(lists)-1)
}

// 这里分治
func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r)>>1
	return mergeTwoList(merge(lists, l, mid), merge(lists, mid+1, r))
}
func mergeTwoList(a, b *ListNode) *ListNode {
	if a == nil && b != nil {
		return b
	}
	if a != nil && b == nil {
		return a
	}
	pre := &ListNode{}
	head, aPtr, bPtr := pre, a, b
	for aPtr != nil && bPtr != nil {
		if aPtr.Val < bPtr.Val {
			head.Next = aPtr
			aPtr = aPtr.Next
		} else {
			head.Next = bPtr
			bPtr = bPtr.Next
		}
		head = head.Next
	}
	if aPtr != nil {
		head.Next = aPtr
	}
	if bPtr != nil {
		head.Next = bPtr
	}
	return pre.Next
}

//leetcode submit region end(Prohibit modification and deletion)
