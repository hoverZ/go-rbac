package permission

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// InsertTreeNode 测试用例：
// 1. 插入后的树结构
// 2. 插入根节点 path=""
// 3. 插入多级节点 path="/a", path="/a/b"
// 4. 插入模糊节点 path="/:t", path="/:m/abc"
// 5. 插入模糊节点 path="/:t", path="/a"
// FindLeafNode 测试用例：
// 1. 初始化节点
// 根节点
// 2. 授权时查询
// 域名
func TestInsertTreeNode(t *testing.T) {
	// InsertTreeNode 测试用例：
	insertNodeItems := []*InsertNodeItem{
		&InsertNodeItem{
			Id:     1,
			Path:   "",
			Status: 1,
		},
		&InsertNodeItem{
			Id:     3,
			Path:   "/",
			Status: 1,
		},
		&InsertNodeItem{
			Id:     2,
			Path:   "/a/b",
			Status: 2,
		},
		&InsertNodeItem{
			Id:     4,
			Path:   "/a",
			Status: 1,
		},
		&InsertNodeItem{
			Id:     5,
			Path:   "/:t",
			Status: 1,
		},
		&InsertNodeItem{
			Id:     6,
			Path:   "/:m/abc",
			Status: 0,
		},
	}
	var tree = &Tree{
		Size:            1,
		CurrentMaxLevel: 0,
		RootNode: &TreeNode{
			Id:     0,
			Level:  0,
			Status: 0,
			Key:    "",
		},
	}
	for i := 0; i < len(insertNodeItems); i++ {
		InsertTreeNode(tree, insertNodeItems[i])
	}
	Convey("check tree struct", t, func() {
		So(tree, ShouldNotBeNil)
		So(tree.RootNode, ShouldNotBeNil)
		So(tree.RootNode.Id, ShouldEqual, 1)
		So(tree.ValidSize, ShouldEqual, 6)
		So(tree.CurrentMaxLevel, ShouldEqual, 2)
	})
	Convey("check treeNode struct", t, func() {
		// 首节点
		So(tree.RootNode.Id, ShouldEqual, 1)
		// 首节点的子节点
		So(len(tree.RootNode.Children), ShouldEqual, 3)
		So(tree.RootNode.Children[0].Id, ShouldEqual, 3)
		So(tree.RootNode.Children[1].Id, ShouldEqual, 4)
		So(tree.RootNode.Children[2].Id, ShouldEqual, 5)
		// 第二层, 子节点数和子节点对应的id 是否正确
		So(len(tree.RootNode.Children[0].Children), ShouldEqual, 0)
		So(len(tree.RootNode.Children[1].Children), ShouldEqual, 1)
		So(tree.RootNode.Children[1].Children[0].Id, ShouldEqual, 2)
		So(len(tree.RootNode.Children[2].Children), ShouldEqual, 1)
		So(tree.RootNode.Children[2].Children[0].Id, ShouldEqual, 6)
	})
	// FindLeafNode 测试用例：
}
