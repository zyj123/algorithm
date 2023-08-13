package main

import p "wode/alternatePrint"

func main() {
	p.Print()

}

//上机编程题
//题目描述：扁平数组转Tree
//
//给定一个扁平数组，数组内每个对象的id属性是唯一的。每个对象具有pid属性，
//pid属性为0表示为根节点（根节点只有一个），其它表示自己的父节点id。
//编写一段程序，输入为给定的扁平数组，输出要求为一个树结构，
//为其中每个对象增加children数组属性（里面存放child对象）。
//
//
//给定输入：
//[
//{"id": 1, "name": "部门1", "pid": 0},
//{"id": 2, "name": "部门2", "pid": 1},
//{"id": 3, "name": "部门3", "pid": 1},
//{"id": 4, "name": "部门4", "pid": 3},
//{"id": 5, "name": "部门5", "pid": 4}
//]
//
//给定输出：
//{
//"id": 1,
//"name": "部门1",
//"pid": 0,
//"children": [
//{
//"id": 2,
//"name": "部门2",
//"pid": 1,
//"children": []
//},
//{
//"id": 3,
//"name": "部门3",
//"pid": 1,
//"children": [
//{
//"id": 4,
//"name": "部门4",
//"pid": 3,
//"children": [
//{
//"id": 5,
//"name": "部门5",
//"pid": 4,
//"children": []
//}
//]
//}
//]
//}
//]
//}

type Num struct {
	ID   int
	Name string
	PID  int
}

type Tree struct {
	ID       int
	Name     string
	Children []*Tree
}

func toTree(nums []Num) *Tree {
	if len(nums) == 0 {
		return nil
	}
	mp := make(map[int][]Num, 0)
	for _, num := range nums {
		mp[num.PID] = append(mp[num.PID], num)
	}

	var head *Tree
	if numArr, ok := mp[0]; ok {
		num := numArr[0]
		head = &Tree{
			ID:   num.ID,
			Name: num.Name,
		}
	} else {
		return nil
	}

	q := make([]*Tree, 0)
	q = append(q, head)
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			node := q[i]
			if numArr, ok := mp[node.ID]; ok {
				children := make([]*Tree, 0)
				for _, num := range numArr {
					child := &Tree{
						ID:   num.ID,
						Name: num.Name,
					}
					children = append(children, child)
					q = append(q, child)
				}
				node.Children = children
			}
		}
		q = q[n:]
	}

	return head
}
