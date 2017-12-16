package proofs

import (
	"bytes"
	"errors"

	lc "github.com/tendermint/light-client"
	"github.com/tendermint/tendermint/types"
)

func ValidateBlockMeta(meta *types.BlockMeta, check lc.Checkpoint) error {
	// TODO: check the BlockID??
	return ValidateHeader(meta.Header, check)
}

func ValidateBlock(meta *types.Block, check lc.Checkpoint) error {
	err := ValidateHeader(meta.Header, check)
	if err != nil {
		return err
	}
	if !bytes.Equal(meta.Data.Hash(), meta.Header.DataHash) {
		return errors.New("Data hash doesn't match header")
	}
	return nil
}

func ValidateHeader(head *types.Header, check lc.Checkpoint) error {
	// make sure they are for the same height (obvious fail)
	if head.Height != check.Height() {
		return lc.ErrHeightMismatch(head.Height, check.Height())
	}
	// check if they are equal by using hashes
	chead := check.Header
	if !bytes.Equal(head.Hash(), chead.Hash()) {
		return errors.New("Headers don't match")
	}
	return nil
}
