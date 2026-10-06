/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res []string

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			// 用 "N" 代表 nil，讓 string 保留 Tree 的具體形狀
			res = append(res, "N")
			return
		}
		// 將 node.Val 從 int 轉 ASCII 存入
		res = append(res, strconv.Itoa(node.Val))

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
    
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	i := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		val := vals[i]
		i++

		if val == "N" {
			return nil
		}

		num, _ := strconv.Atoi(val)
		node := &TreeNode{Val: num}
		// 照左、右的順序遞迴還原 subtree
		node.Left = dfs()
		node.Right = dfs()

		return node
	}

	return dfs()
}
