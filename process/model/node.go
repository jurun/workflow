package model

import (
	"github.com/gofrs/uuid"
	"github.com/jurun/workflow/button"
	"github.com/jurun/workflow/organize"
	"strings"
)

// NodeType 节点类型
type NodeType string

const (
	// 空白节点，用于默认初始值
	EMPTY_NODE NodeType = "empty"
	// StartNode 申请节点
	START_NODE NodeType = "startNode"
	// AuditNode 审核节点
	AUDIT_NODE NodeType = "auditNode"
	// CopyNode 抄送节点
	COPY_NODE NodeType = "copyNode"
	// ConditionNoe 条件节点
	CONDITION_NODE NodeType = "conditionNode"
	// BranchNode 分支节点
	BRANCH_NODE NodeType = "branchNode"
	// CreatorNode Robot.写入数据节点
	CREATE_NODE NodeType = "creatorNode"
	// UpdaterNode Robot.更新数据节点
	UPDATE_NODE NodeType = "updaterNode"
)

func (static NodeType) String() string {
	return string(static)
}

func (static NodeType) IsAuto() bool {
	if static == BRANCH_NODE || static == CONDITION_NODE || static == COPY_NODE {
		return true
	}
	return false
}

func (static NodeType) Text() string {
	m := map[NodeType]string{
		START_NODE:     "申请节点",
		AUDIT_NODE:     "审批节点",
		COPY_NODE:      "抄送节点",
		CONDITION_NODE: "条件节点",
		BRANCH_NODE:    "分支节点",
		CREATE_NODE:    "添加数据",
		UPDATE_NODE:    "更新数据",
	}
	return m[static]
}

func (static NodeType) Buttons() []*button.Button {
	btns := make([]*button.Button, 0)
	{
		btns = append(
			btns,
			button.SUBMIT_BUTTON.Init(),
		)
	}

	switch static {
	case AUDIT_NODE:
		btns = append(
			btns,
			button.ROLLBACK_BUTTON.Init(),
			button.TRANSFER_BUTTON.Init(),
			button.PAUSE_BUTTON.Init(),
			button.KILL_BUTTON.Init(),
		)
	case BRANCH_NODE, CONDITION_NODE, CREATE_NODE, UPDATE_NODE:
		btns = make([]*button.Button, 0)
	case COPY_NODE:
		btns = make([]*button.Button, 0)
	}
	return btns
}

func (static NodeType) Config() interface{} {
	return nil
}

type AuditType string

const (
	AND_AUDIT AuditType = "AND"
	OR_AUDIT  AuditType = "OR"
)

func (static AuditType) String() string {
	return string(static)
}

func (static AuditType) Text() string {
	mp := map[AuditType]string{
		AND_AUDIT: "会签",
		OR_AUDIT:  "或签",
	}
	return mp[static]
}

type OperatorAuth struct {
	View     bool                    `json:"view"`
	Edit     bool                    `json:"edit"`
	Required bool                    `json:"required"`
	Children map[string]OperatorAuth `json:"children"`
}

type FormAuth struct {
	Tabs   map[string]OperatorAuth `json:"tabs"`
	Fields map[string]OperatorAuth `json:"fields"`
}

func (FormAuth) Init() *FormAuth {
	return &FormAuth{
		Tabs:   map[string]OperatorAuth{},
		Fields: map[string]OperatorAuth{},
	}
}

type Validator struct {
	Formula string `json:"formula"`
	Message string `json:"message"`
}

type Condition struct {
	Formula string   `json:"formula"`
	Fields  []string `json:"fields"`
}

type Node struct {
	Id        string                    `json:"id"`           // 唯一ID
	Name      string                    `json:"name"`         // 节点名称
	Caller    *organize.ProcessOrganize `json:"caller"`       // 负责人设置
	FormAuth  *FormAuth                 `json:"form_auth"`    // 表单操作权限设置
	Type      NodeType                  `json:"type"`         // 节点类型
	Opinion   int                       `json:"opinion"`      // 提交任务是否需要填写反馈意见，0：不需要弹出，1：弹出不必填，2：必填
	Signature bool                      `json:"signature"`    // 是否需要手签
	Buttons   []*button.Button          `json:"buttons"`      // 按钮集合
	Branches  []*Nodes                  `json:"branch_child"` // 子节点集合
	Validator Validator                 `json:"validator"`    // 提交验证
	Condition Condition                 `json:"condition"`    // 条件？
	AuditType AuditType                 `json:"audit_type"`   // 审批类型，会签/或签
}

func (Node) Init(nodeType NodeType) *Node {

	if nodeType == EMPTY_NODE {
		return &Node{}
	}

	id, _ := uuid.NewV4()
	nd := &Node{
		Id:       strings.Replace(id.String(), "-", "", -1),
		Name:     nodeType.Text(),
		Caller:   organize.ProcessOrganize{}.Init(),
		FormAuth: FormAuth{}.Init(),
		Type:     nodeType,
		Buttons:  nodeType.Buttons(),
		Branches: make([]*Nodes, 0),
		Condition: Condition{
			Fields: []string{},
		},
	}

	if nodeType == AUDIT_NODE {
		nd.AuditType = AND_AUDIT
	}

	if nodeType == BRANCH_NODE {
		for i := 0; i < 2; i++ {
			nds := Nodes{}.Init()
			nds.AddNode(CONDITION_NODE)
			nd.Branches = append(nd.Branches, nds)
		}
	}

	return nd
}
