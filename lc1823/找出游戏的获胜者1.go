package lc1823

/*
分析：模拟+queue；约瑟夫环
模拟游戏过程，使用queue存储圈子中的小伙伴编号，
初始时将1到n的所有编号依次加入queue，队首元素即为第1名小伙伴的编号，
每一轮游戏中，从当前小伙伴开始数k名小伙伴，数到的第k名小伙伴离开圈子。
模拟游戏过程的做法是，将队首元素取出并将该元素在队尾处重新加入队列，
重复该操作k−1次，则在k−1次操作之后，队首元素即为这一轮中数到的第k名小伙伴的编号，
将队首元素取出，即为数到的第k名小伙伴离开圈子。
上述操作之后，新的队首元素即为下一轮游戏的起始小伙伴的编号。
每一轮游戏之后，圈子中减少一名小伙伴，队列中减少一个元素。
重复上述过程，直到队列中只剩下1个元素，该元素即为获胜的小伙伴的编号。
时间复杂度：O(nk), 其中n是做游戏的小伙伴数量，k是每一轮离开圈子的小伙伴的计数。
初始时需要将n个元素加入队列，每一轮需要将k个元素从队列中取出，将k−1个元素加入队列，
一共有n−1轮，因此时间复杂度是O(nk)
空间复杂度：O(n), 其中n是做游戏的小伙伴数量。空间复杂度主要取决于队列，队列中最多有n个元素。
 */


func findTheWinner(n, k int) int {
	q := make([]int, n)
	for i := range q {
		q[i] = i + 1
	}
	for len(q) > 1 {
		// 前k-1个重新入队
		for i := 1; i < k; i++ {
			q = append(q, q[0])[1:]
		}
		// 第k个出队
		q = q[1:]
	}
	return q[0]
}



