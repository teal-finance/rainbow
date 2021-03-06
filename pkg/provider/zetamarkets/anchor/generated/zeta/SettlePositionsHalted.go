// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package zeta

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SettlePositionsHalted is the `settlePositionsHalted` instruction.
type SettlePositionsHalted struct {

	// [0] = [] state
	//
	// [1] = [] zetaGroup
	//
	// [2] = [] greeks
	//
	// [3] = [SIGNER] admin
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSettlePositionsHaltedInstructionBuilder creates a new `SettlePositionsHalted` instruction builder.
func NewSettlePositionsHaltedInstructionBuilder() *SettlePositionsHalted {
	nd := &SettlePositionsHalted{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 4),
	}
	return nd
}

// SetStateAccount sets the "state" account.
func (inst *SettlePositionsHalted) SetStateAccount(state ag_solanago.PublicKey) *SettlePositionsHalted {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *SettlePositionsHalted) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetZetaGroupAccount sets the "zetaGroup" account.
func (inst *SettlePositionsHalted) SetZetaGroupAccount(zetaGroup ag_solanago.PublicKey) *SettlePositionsHalted {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(zetaGroup)
	return inst
}

// GetZetaGroupAccount gets the "zetaGroup" account.
func (inst *SettlePositionsHalted) GetZetaGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetGreeksAccount sets the "greeks" account.
func (inst *SettlePositionsHalted) SetGreeksAccount(greeks ag_solanago.PublicKey) *SettlePositionsHalted {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(greeks)
	return inst
}

// GetGreeksAccount gets the "greeks" account.
func (inst *SettlePositionsHalted) GetGreeksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAdminAccount sets the "admin" account.
func (inst *SettlePositionsHalted) SetAdminAccount(admin ag_solanago.PublicKey) *SettlePositionsHalted {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *SettlePositionsHalted) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

func (inst SettlePositionsHalted) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SettlePositionsHalted,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SettlePositionsHalted) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SettlePositionsHalted) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ZetaGroup is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Greeks is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *SettlePositionsHalted) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SettlePositionsHalted")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("zetaGroup", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   greeks", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("    admin", inst.AccountMetaSlice.Get(3)))
					})
				})
		})
}

func (obj SettlePositionsHalted) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SettlePositionsHalted) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSettlePositionsHaltedInstruction declares a new SettlePositionsHalted instruction with the provided parameters and accounts.
func NewSettlePositionsHaltedInstruction(
	// Accounts:
	state ag_solanago.PublicKey,
	zetaGroup ag_solanago.PublicKey,
	greeks ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *SettlePositionsHalted {
	return NewSettlePositionsHaltedInstructionBuilder().
		SetStateAccount(state).
		SetZetaGroupAccount(zetaGroup).
		SetGreeksAccount(greeks).
		SetAdminAccount(admin)
}
