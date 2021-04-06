package common

// ITreeNode 要使用树形结构，对应结构体需要实现以下方法
type ITreeNode interface {
	GetPrimKey() int       // 主键
	GetParentPrimKey() int // 父主键
	GetName() string       // 节点名称
	GetData() interface{}  // 数据
	Root() bool            // 是否根节点
}

// TreeNode 树形结构的节点数据
type TreeNode struct {
	PrimKey       int         `json:"primKey"`
	ParentPrimKey int         `json:"parentPrimKey"`
	Name          string      `json:"name"`
	Data          interface{} `json:"data"`
	Children      []TreeNode  `json:"children"`
}

// GenerateTree 生成树形结构
func GenerateTree(nodes []ITreeNode) (trees []TreeNode) {
	if len(nodes) <= 0 {
		return
	}
	var roots, childs []ITreeNode

	// 分类根节点和子节点
	for _, v := range nodes {
		if v.Root() {
			roots = append(roots, v)
		} else {
			childs = append(childs, v)
		}
	}

	for _, root := range roots {
		treeNode := &TreeNode{
			Name:          root.GetName(),
			PrimKey:       root.GetPrimKey(),
			ParentPrimKey: root.GetParentPrimKey(),
			Data:          root.GetData(),
		}

		// 递归寻找子节点
		recursiveTree(treeNode, childs)

		trees = append(trees, *treeNode)
	}
	return
}

func recursiveTree(node *TreeNode, childs []ITreeNode) {
	iNode := node.Data.(ITreeNode)

	for _, v := range childs {
		if v.GetParentPrimKey() == iNode.GetPrimKey() {
			childTreeNode := &TreeNode{
				Name:          v.GetName(),
				PrimKey:       v.GetPrimKey(),
				ParentPrimKey: v.GetParentPrimKey(),
				Data:          v.GetData(),
			}

			// 递归寻找子节点
			recursiveTree(childTreeNode, childs)

			node.Children = append(node.Children, *childTreeNode)
		}
	}
}
