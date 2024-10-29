package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		minNode := root.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		root.Data = minNode.Data
		root.Right = BTreeDeleteNode(root.Right, minNode)
	}
	return root
}
