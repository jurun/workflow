package model

import (
	"errors"
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/json-iterator/go"
)

type Position struct {
	Node       *Node
	List       *Nodes
	Index      int
	ParentNode *Position
}

func (Position) Init() *Position {
	return &Position{
		Node:       Node{}.Init(EMPTY_NODE),
		List:       Nodes{}.Init(),
		ParentNode: new(Position),
	}
}

func (this *Position) Exits() bool {
	if this.Node != nil && this.Node.Name != "" {
		return true
	}
	return false
}

func (this *Position) Next() (*Position, bool) {
	if this.Exits() == false {
		return Position{}.Init(), false
	}

	val, found := this.List._list.Get(this.Index + 1)
	if found {
		return &Position{
			Node:       val.(*Node),
			List:       this.List,
			Index:      this.Index + 1,
			ParentNode: this.ParentNode,
		}, true
	}
	return Position{}.Init(), false
}

func (this *Position) Prev() (*Position, bool) {
	if this.Exits() == false {
		return Position{}.Init(), false
	}

	val, found := this.List._list.Get(this.Index - 1)
	if found {
		return &Position{
			Node:       val.(*Node),
			List:       this.List,
			Index:      this.Index - 1,
			ParentNode: this.ParentNode,
		}, true
	}
	return Position{}.Init(), false
}

func (this *Position) Subs() [][]*Position {
	subs := make([][]*Position, 0)
	if len(this.Node.Branches) > 0 {
		for _, list := range this.Node.Branches {
			_sub := make([]*Position, 0)
			list.Each(func(index int, node *Node) error {
				_sub = append(_sub, &Position{
					Node:       node,
					List:       list,
					Index:      index,
					ParentNode: this,
				})
				return nil
			})
			subs = append(subs, _sub)
		}
	}
	return subs
}

func (Nodes) Init() *Nodes {
	nds := &Nodes{
		_list: arraylist.New(),
	}
	return nds
}

type Nodes struct {
	_list *arraylist.List
}

func (this *Nodes) AddBranch(parentId string) error {
	position, found := this.Find(parentId)
	if !found {
		return fmt.Errorf("父级节点：%s 不存在", parentId)
	}

	nds := Nodes{}.Init()
	nds.AddNode(CONDITION_NODE)
	position.Node.Branches = append(position.Node.Branches, nds)
	return nil
}

func (this *Nodes) AddNode(nodeType NodeType, prevId ...string) error {
	node := Node{}.Init(nodeType)
	if len(prevId) == 0 {
		this._list.Insert(this._list.Size(), node)
		return nil
	}

	prev, found := this.Find(prevId[0])
	if !found {
		return fmt.Errorf("前置节点：%s 不存在", prevId[0])
	}

	prev.List._list.Insert(prev.Index+1, node)
	return nil
}

func (this *Nodes) Update(node *Node) error {
	position, found := this.Find(node.Id)
	if !found {
		return errors.New("节点不存在")
	}

	position.List._list.Set(position.Index, node)
	return nil
}

func (this *Nodes) remove(nodeId string) error {

	position, found := this.Find(nodeId)
	if found {
		switch position.Node.Type {
		case START_NODE:
			return fmt.Errorf("%s不能删除", position.Node.Type.Text())
		}

		if position.Node.Type == CONDITION_NODE {
			position.List._list.Clear()
		} else {
			position.List._list.Remove(position.Index)
			if position.List.Size() == 1 {
				last, _ := position.List._list.Get(0)
				if last.(*Node).Type == CONDITION_NODE {
					position.List._list.Clear()
				}
			}
		}
	}
	return nil
}

func (this *Nodes) compensate() {
	if this._list.Size() == 0 {
		return
	}

	newList := arraylist.New()

	this._list.Each(func(index int, value interface{}) {
		node := value.(*Node)
		if node.Type == BRANCH_NODE {
			newBranches := make([]*Nodes, 0)
			for i := 0; i < len(node.Branches); i++ {
				node.Branches[i].compensate()

				if node.Branches[i]._list.Size() == 0 {
					continue
				}

				newBranches = append(newBranches, node.Branches[i])
			}

			node.Branches = newBranches
			if len(node.Branches) == 1 {
				node.Branches[0].Each(func(i int, nd *Node) error {
					if nd.Type == CONDITION_NODE {
						return nil
					}
					newList.Add(nd)
					return nil
				})
			} else if len(node.Branches) == 0 {
				return
			} else {
				node.Branches = newBranches
				newList.Add(value)
			}
		} else if node.Type == CONDITION_NODE {
			newList.Add(value)
		} else {
			newList.Add(value)
		}

		return
	})

	this._list = newList
}

func (this *Nodes) Delete(nodeId string) error {
	err := this.remove(nodeId)
	if err != nil {
		return err
	}
	this.compensate()
	return nil
}

func (this *Nodes) Find(nodeId string) (position *Position, found bool) {
	if nodeId == "" {
		return &Position{}, false
	}

	this._list.Find(func(index int, value interface{}) bool {

		parentNode := value.(*Node)
		parentPosition := &Position{
			Node:       parentNode,
			Index:      index,
			List:       this,
			ParentNode: new(Position),
		}

		if parentPosition.Node.Id == nodeId {
			found = true
			position = parentPosition
			return true
		}

		//if value.(*Node).Id == nodeId {
		//	position.node = value.(*Node)
		//	position.index = index
		//	position.list = this
		//	found = true
		//	return true
		//}

		if parentNode.Type == BRANCH_NODE && len(parentNode.Branches) > 0 {
			for _, nds := range parentNode.Branches {
				if subPos, fd := nds.Find(nodeId); fd {
					subPos.ParentNode = parentPosition
					position = subPos
					found = true
					return true
				}
			}
		}

		//if value.(*Node).Type == BRANCH_NODE && len(value.(*Node).Branches) > 0 {
		//	for _, nds := range value.(*Node).Branches {
		//		if pos, fd := nds.Find(nodeId); fd {
		//			position = pos
		//			found = true
		//			return true
		//		}
		//	}
		//	return false
		//}
		return false
	})

	return
}

func (this *Nodes) Each(fn func(index int, node *Node) error) error {
	values := this._list.Values()
	for i := 0; i < this._list.Size(); i++ {
		err := fn(i, values[i].(*Node))
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *Nodes) Size() int {
	return this._list.Size()
}

func (this *Nodes) hasConditionNode() bool {
	index, _ := this._list.Find(func(index int, value interface{}) bool {
		if value.(*Node).Type == CONDITION_NODE {
			return true
		}
		return false
	})

	if index != -1 {
		return true
	}
	return false
}

func (this *Nodes) MarshalJSON() ([]byte, error) {
	if this._list == nil {
		return nil, nil
	}
	return this._list.ToJSON()
}

func (this *Nodes) UnmarshalJSON(b []byte) error {
	nodes := make([]*Node, 0)
	err := jsoniter.Unmarshal(b, &nodes)
	if err != nil {
		return err
	}

	this._list = arraylist.New()
	for _, n := range nodes {
		this._list.Add(n)
	}
	return nil
}

func (this *Nodes) FromDB(b []byte) error {
	return this.UnmarshalJSON(b)
}

func (this *Nodes) ToDB() ([]byte, error) {
	return this.MarshalJSON()
}

func (this *Nodes) List() *arraylist.List {
	return this._list
}

func (this *Nodes) Get(index int) (*Node, bool) {
	val, found := this._list.Get(index)
	if !found {
		return nil, false
	}
	return val.(*Node), true
}
