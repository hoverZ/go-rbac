package permission

import "strings"

type TreeNode struct {
	Id       int         `json:"id"`       // ID 标志
	Level    int         `json:"level"`    // 第几层
	Status   int         `json:"status"`   // 状态，是否需要验证
	Key      string      `json:"key"`      // 匹配字段
	Children []*TreeNode `json:"children"` // 子节点
}

type Tree struct {
	Size            int       `json:"size"`
	RootNode        *TreeNode `json:"root_node"`
	CurrentMaxLevel int       `json:"current_max_level"`
	ValidSize       int       `json:"valid_size"` // 有 ID 的节点
}

type InsertNodeItem struct {
	Id     int    `json:"id"`
	Path   string `json:"path"`
	Status int    `json:"status"`
}

const find_leaf_mode_init = 1 // 初始化模式
const find_leaf_mode_auth = 2 // 验证模式

// 插入节点
func InsertTreeNode(tree *Tree, node *InsertNodeItem) *TreeNode {
	// 根据 '/' 将 path 切分为多个字符串
	keys := strings.Split(node.Path, "/")
	leafNode := FindLeafNode(tree.RootNode, keys, find_leaf_mode_init)
	if leafNode.Level+1 == len(keys) {
		// 长度相等，表示已经完整匹配了
		tree.ValidSize++
		leafNode.Id = node.Id
		leafNode.Status = node.Status
		return leafNode
	}
	// 若没有完整匹配所有 key，需要创建节点
	for begin := leafNode.Level + 1; begin < len(keys); begin++ {
		treeNode := &TreeNode{
			Id:     0,
			Status: 0,
			Level:  begin,
			Key:    keys[begin],
		}
		// 查到最后一个元素，需要让 treeNode 设置其ID 和 status
		if begin == len(keys)-1 {
			treeNode.Id = node.Id
			treeNode.Status = node.Status
			tree.ValidSize++
		}
		leafNode.Children = append(leafNode.Children, treeNode)
		leafNode = treeNode
		tree.Size++
	}
	if leafNode.Level > tree.CurrentMaxLevel {
		tree.CurrentMaxLevel = tree.Size
	}
	return leafNode
}

// mode 为模式，目前存在两种模式：初始化模式（init）、验证模式（auth）
func FindLeafNode(rootNode *TreeNode, keys []string, mode int) (targetItem *TreeNode) {
	var traversingItems []*TreeNode
	var likeItem *TreeNode
	// 将根节点添加到需要遍历的节点数组中
	traversingItems = append(traversingItems, rootNode)
	for _, key := range keys {
		var hadMatch bool
		var like bool
		// 遍历数组中的节点，取出匹配的节点
		// 重置 traversingItems = targetItem.Children
		// 退出 traversingItems 的遍历
		for index, node := range traversingItems {

			if key == node.Key {
				// 完全匹配
				targetItem = node
				traversingItems = targetItem.Children
				hadMatch = true
				break
			} else if mode == find_leaf_mode_init && node.Key != "" && node.Key[0:1] == ":" {
				// 初始化时的模糊匹配，将带有 : 前缀的 key 判定为同一个 key
				// 由于完全匹配的优先于模糊匹配，所以需要确保没有完全匹配后才能退出
				like = true
				likeItem = node
			}
			if index+1 == len(traversingItems) && like &&
				((mode == find_leaf_mode_init && key[0:1] == ":") || mode == find_leaf_mode_auth) {
				// 最后一个元素已经匹配完了 且 存在模糊匹配成功 且 ((初始化模式 且 key带':'前缀) 或者 为验证模式)
				targetItem = likeItem
				traversingItems = targetItem.Children
				hadMatch = true
				break
			}
		}
		if !hadMatch {
			// 没有匹配成功，退出
			break
		}
	}
	// 若没有匹配成功就设置为根节点
	if targetItem == nil {
		targetItem = rootNode
	}
	return
}
