package button

import (
	"errors"
	"github.com/jurun/workflow/organize"
	"github.com/jurun/workflow/utils"
)

type ButtonType string

const (
	SUBMIT_BUTTON   ButtonType = "submit"
	STAGING_BUTTON  ButtonType = "staging"
	ROLLBACK_BUTTON ButtonType = "rollback"
	TRANSFER_BUTTON ButtonType = "transfer"
	KILL_BUTTON     ButtonType = "kill"
	PAUSE_BUTTON    ButtonType = "pause"
	AUTO_BUTTON     ButtonType = "auto"
)

func (static ButtonType) String() string {
	return string(static)
}

func (static ButtonType) Label() string {
	mp := map[ButtonType]string{
		SUBMIT_BUTTON:   "提交",
		STAGING_BUTTON:  "暂存",
		ROLLBACK_BUTTON: "回退",
		TRANSFER_BUTTON: "转交",
		KILL_BUTTON:     "结束流程",
		PAUSE_BUTTON:    "暂停流程",
		AUTO_BUTTON:     "自动触发",
	}

	return mp[static]
}

func (static ButtonType) Init() *Button {
	enabled := false
	if static == SUBMIT_BUTTON {
		enabled = true
	}

	btn := Button{
		Label:   static.Label(),
		Type:    static,
		Enabled: enabled,
	}

	setting := make(map[string]interface{})
	switch static {
	case SUBMIT_BUTTON:
		utils.NewConvert(new(SubmitAction)).Bind(&setting)
	case STAGING_BUTTON:
		utils.NewConvert(new(StagingAction)).Bind(&setting)
	case ROLLBACK_BUTTON:
		a := RollbackAction{
			Range: []int{},
		}
		utils.NewConvert(&a).Bind(&setting)
	case TRANSFER_BUTTON:
		a := TransferAction{
			Callers: organize.ProcessOrganize{}.Init(),
		}
		utils.NewConvert(&a).Bind(&setting)
	case KILL_BUTTON:
		utils.NewConvert(new(KillAction)).Bind(&setting)
	case PAUSE_BUTTON:
		utils.NewConvert(new(PauseAction)).Bind(&setting)
	default:
	}

	btn.Setting = setting

	return &btn
}

type Button struct {
	Label   string                 `json:"label"`
	Type    ButtonType             `json:"type"`
	Enabled bool                   `json:"enabled"`
	Setting map[string]interface{} `json:"setting"`
}

func (this *Button) GetAction() (Action, error) {
	switch this.Type {
	case SUBMIT_BUTTON:
		var at SubmitAction
		err := utils.NewConvert(this.Setting).Bind(&at)
		return &at, err
	case STAGING_BUTTON:
		var at StagingAction
		err := utils.NewConvert(this.Setting).Bind(&at)
		return &at, err
	case ROLLBACK_BUTTON:
		var at RollbackAction
		err := utils.NewConvert(this.Setting).Bind(&at)
		return &at, err
	case TRANSFER_BUTTON:
		var at TransferAction
		err := utils.NewConvert(this.Setting).Bind(&at)
		return &at, err
	case KILL_BUTTON:
		var at KillAction
		err := utils.NewConvert(this.Setting).Bind(&at)
		return &at, err
	case PAUSE_BUTTON:
		var at PauseAction
		err := utils.NewConvert(this.Setting).Bind(&at)
		return &at, err
	default:
		return nil, errors.New("非法的按钮类型")
	}
}

type Action interface {
	GetLabel() string
}
