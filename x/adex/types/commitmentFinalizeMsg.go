package types

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CommitmentFinalizeMsg struct {
	Commitment Commitment `json:"commitment"`
	Signatures [][]byte `json:"signatures"`
	Vote []byte `json:"vote"`
	Submitter sdk.AccAddress `json:"submitter"`
}

func (msg CommitmentFinalizeMsg) Name() string {
	return "CommitmentFinalizeMsg"
}

func (msg CommitmentFinalizeMsg) Type() string {
	return "adex"
}

func (msg CommitmentFinalizeMsg) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg CommitmentFinalizeMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Submitter}
}

func (msg CommitmentFinalizeMsg) ValidateBasic() sdk.Error {
	if !msg.Commitment.IsValid() {
		// @TODO: proper error
		return sdk.ErrUnknownRequest("invalid commitment")
	}

	// @TODO: validate signatures here, require 2/3rds or more

	return nil
}

func (msg CommitmentFinalizeMsg) String() string {
	return fmt.Sprintf("CommitmentFinalizeMsg{%v}", msg.Commitment)
}
