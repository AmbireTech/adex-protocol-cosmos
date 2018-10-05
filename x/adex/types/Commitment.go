package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Commitment struct {
	BidId []byte `json:"bidId"`
	TotalReward sdk.Coins `json:"totalReward"`
	ValidUntil int64 `json:"validUntil"`
	Advertiser sdk.AccAddress `json:"advertiser"`
	Publisher sdk.AccAddress `json:"publisher"`
	Validators []Validator `json:"validators"`
}

// @TODO: GetHash()
// @TODO: IsValid() : should check if the validator reward IsLT the sum of all validator rewards (same as on eth)
// @TODO: FromBid() : last arg would be extraValidatorAddr; test if != nil, also test if .IsValid or smth
