// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package zeta

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UnhaltZetaGroup is the `unhaltZetaGroup` instruction.
type UnhaltZetaGroup struct {

	// [0] = [] state
	//
	// [1] = [WRITE] zetaGroup
	//
	// [2] = [SIGNER] admin
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUnhaltZetaGroupInstructionBuilder creates a new `UnhaltZetaGroup` instruction builder.
func NewUnhaltZetaGroupInstructionBuilder() *UnhaltZetaGroup {
	nd := &UnhaltZetaGroup{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetStateAccount sets the "state" account.
func (inst *UnhaltZetaGroup) SetStateAccount(state ag_solanago.PublicKey) *UnhaltZetaGroup {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UnhaltZetaGroup) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetZetaGroupAccount sets the "zetaGroup" account.
func (inst *UnhaltZetaGroup) SetZetaGroupAccount(zetaGroup ag_solanago.PublicKey) *UnhaltZetaGroup {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(zetaGroup).WRITE()
	return inst
}

// GetZetaGroupAccount gets the "zetaGroup" account.
func (inst *UnhaltZetaGroup) GetZetaGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAdminAccount sets the "admin" account.
func (inst *UnhaltZetaGroup) SetAdminAccount(admin ag_solanago.PublicKey) *UnhaltZetaGroup {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UnhaltZetaGroup) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UnhaltZetaGroup) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UnhaltZetaGroup,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UnhaltZetaGroup) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UnhaltZetaGroup) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ZetaGroup is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *UnhaltZetaGroup) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UnhaltZetaGroup")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("zetaGroup", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("    admin", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj UnhaltZetaGroup) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *UnhaltZetaGroup) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewUnhaltZetaGroupInstruction declares a new UnhaltZetaGroup instruction with the provided parameters and accounts.
func NewUnhaltZetaGroupInstruction(
	// Accounts:
	state ag_solanago.PublicKey,
	zetaGroup ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *UnhaltZetaGroup {
	return NewUnhaltZetaGroupInstructionBuilder().
		SetStateAccount(state).
		SetZetaGroupAccount(zetaGroup).
		SetAdminAccount(admin)
}
