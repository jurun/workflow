package organize

import (
	"testing"
)

func TestProcessOrganize_AllUids(t *testing.T) {
	organize := ProcessOrganize{
		Deps: []int{53, 54},
		Uids: []int{118},
		Dynamic: []Dynamic{
			CREATOR,
			CALLER,
			CREATOR_DEP,
			CALLER_DEP,
			CREATOR_SUPERIOR_LEADER,
			CALLER_SUPERIOR_LEADER,
		},
	}

	organize.SetCaller(148)
	organize.SetCreator(183)

	uids, err := organize.AllUids()
	if err != nil {
		t.Error(err)
		return
	}

	organize.AllUids()

	t.Log(uids)
}
