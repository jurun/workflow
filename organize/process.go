package organize

import (
	"errors"
	"github.com/json-iterator/go"
	"math/rand"
)

type Dynamic int

const (
	CREATOR Dynamic = iota + 1
	CREATOR_LEADER
	CREATOR_SUPERIOR_LEADER
	CREATOR_DEP
	CREATOR_ROLE
	CALLER
	CALLER_LEADER
	CALLER_SUPERIOR_LEADER
	CALLER_DEP
	CALLER_ROLE
)

func (static Dynamic) String() string {
	mp := map[Dynamic]string{
		CREATOR:                 "发起人",        // 1
		CREATOR_LEADER:          "发起人上级",      // 2
		CREATOR_SUPERIOR_LEADER: "发起人上级的上级",   // 3
		CREATOR_DEP:             "发起人同部门",     // 4
		CREATOR_ROLE:            "发起人同角色",     // 5
		CALLER:                  "当前负责人",      // 6
		CALLER_LEADER:           "当前负责人上级",    // 7
		CALLER_SUPERIOR_LEADER:  "当前负责人上级的上级", // 8
		CALLER_DEP:              "负责人同部门",     // 9
		CALLER_ROLE:             "负责人同角色",     //  10
	}
	return mp[static]
}

func (static Dynamic) Int() int {
	return int(static)
}

type ProcessOrganize struct {
	Deps       []int     `json:"deps"`
	Uids       []int     `json:"uids"`
	Roles      []int     `json:"roles"`
	Dynamic    []Dynamic `json:"dynamic"`
	Field      string    `json:"field"`
	finalUsers []int     `json:"-"`
	hasFetch   bool      `json:"-"`

	caller  int `json:"-"`
	creator int `json:"-"`

	//callerInfo  uims.User `json:"caller_info"`
	//creatorInfo uims.User `json:"creator_info"`

	_getLeaderIds []int `json:"-"`
}

func (static ProcessOrganize) Init() *ProcessOrganize {
	return &ProcessOrganize{
		Deps:          []int{},
		Uids:          []int{},
		Roles:         []int{},
		Dynamic:       []Dynamic{},
		_getLeaderIds: []int{},
		finalUsers:    []int{},
	}
}

func (this *ProcessOrganize) SetCaller(uid int) {
	this.caller = uid
}

func (this *ProcessOrganize) SetCreator(uid int) {
	this.creator = uid
}

func (this *ProcessOrganize) HasPermission() (bool, error) {
	if this.hasFetch == false {
		if err := this.fetch(); err != nil {
			return false, err
		}
	}

	for _, u := range this.finalUsers {
		if this.caller == u {
			return true, nil
		}
	}
	return false, nil
}

func (this *ProcessOrganize) fetch() error {

	if this.hasFetch == true {
		return nil
	}

	if this.creator == 0 {
		return errors.New("Organize 缺少 creator")
	}

	if this.caller == 0 {
		return errors.New("Organize 缺少 caller")
	}

	this.hasFetch = true

	// @todo 缓存钉钉SDK
	//rsp, err := uims.GetUsers([]int{this.caller, this.creator}, false)
	//if err != nil {
	//	return err
	//}

	//for _, u := range rsp.Attachment.Users {
	//	if u.Id == this.creator {
	//		this.creatorInfo = u
	//	} else if u.Id == this.caller {
	//		this.callerInfo = u
	//	}
	//}

	//if len(this.Dynamic) > 0 {
	//	for _, d := range this.Dynamic {
	//		switch d {
	//		case CREATOR:
	//			this.Uids = append(this.Uids, this.creator)
	//		case CREATOR_LEADER:
	//			for _, id := range this.creatorInfo.Leaders {
	//				this.Uids = append(this.Uids, id)
	//			}
	//		case CREATOR_SUPERIOR_LEADER:
	//			this._getLeaderIds = append(this._getLeaderIds, this.creatorInfo.Leaders...)
	//		case CREATOR_DEP:
	//			for _, d := range this.creatorInfo.Depinfos {
	//				this.Deps = append(this.Deps, d.Id)
	//			}
	//		case CREATOR_ROLE:
	//			for _, r := range this.creatorInfo.Roleinfos {
	//				this.Roles = append(this.Roles, r.Id)
	//			}
	//		case CALLER:
	//			this.Uids = append(this.Uids, this.caller)
	//		case CALLER_LEADER:
	//			for _, id := range this.callerInfo.Leaders {
	//				this.Uids = append(this.Uids, id)
	//			}
	//		case CALLER_SUPERIOR_LEADER:
	//			this._getLeaderIds = append(this._getLeaderIds, this.callerInfo.Leaders...)
	//		case CALLER_DEP:
	//			for _, d := range this.callerInfo.Depinfos {
	//				this.Deps = append(this.Deps, d.Id)
	//			}
	//		case CALLER_ROLE:
	//			for _, r := range this.callerInfo.Roleinfos {
	//				this.Roles = append(this.Roles, r.Id)
	//			}
	//		}
	//	}
	//}

	//if len(this.Deps) > 0 {
	//	depRsp, err := uims.GetUsersByDeps(this.Deps, false)
	//	if err != nil {
	//		return err
	//	}
	//	for _, u := range depRsp.Attachment.Users {
	//		this.Uids = append(this.Uids, u.Id)
	//	}
	//}

	//if len(this.Roles) > 0 {
	//	roleRsp, err := uims.GetUsersByRoles(this.Roles)
	//	if err != nil {
	//		return err
	//	}
	//	for _, u := range roleRsp.Attachment.Users {
	//		this.Uids = append(this.Uids, u.Id)
	//	}
	//}

	//if len(this._getLeaderIds) > 0 {
	//	usersRsp, err := uims.GetUsers(this._getLeaderIds, false)
	//	if err != nil {
	//		return err
	//	}
	//	for _, u := range usersRsp.Attachment.Users {
	//		this.Uids = append(this.Uids, u.Leaders...)
	//	}
	//}

	mp := make(map[int]bool)

	for _, id := range this.Uids {
		if id == 0 {
			continue
		}

		if _, found := mp[id]; found {
			continue
		}
		mp[id] = true
		this.finalUsers = append(this.finalUsers, id)
	}

	return nil
}

func (this *ProcessOrganize) AllUids() ([]int, error) {
	if err := this.fetch(); err != nil {
		return nil, err
	}

	return this.finalUsers, nil
}

func (this *ProcessOrganize) GetRandUid() (int, error) {
	if err := this.fetch(); err != nil {
		return 0, err
	}

	if len(this.finalUsers) == 0 {
		return 0, nil
	}

	if len(this.finalUsers) == 1 {
		return this.finalUsers[0], nil
	}

	sixah := this.finalUsers
	rand.Shuffle(len(sixah), func(i, j int) { //调用算法
		sixah[i], sixah[j] = sixah[j], sixah[i]
	})

	return sixah[0], nil
}

func (this *ProcessOrganize) ToDB() ([]byte, error) {
	return jsoniter.Marshal(this)
}

func (this *ProcessOrganize) FromDB(data []byte) error {
	this.Deps = make([]int, 0)
	this.Uids = make([]int, 0)
	this.Roles = make([]int, 0)
	this.Dynamic = make([]Dynamic, 0)
	this._getLeaderIds = make([]int, 0)
	this.finalUsers = make([]int, 0)

	err := jsoniter.Unmarshal(data, this)
	return err
}

func (this *ProcessOrganize) Reset() {
	this.Deps = make([]int, 0)
	this.Uids = make([]int, 0)
	this.Roles = make([]int, 0)
	this.Dynamic = make([]Dynamic, 0)
	this._getLeaderIds = make([]int, 0)
	this.finalUsers = make([]int, 0)
}
