// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package zeta

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateAdmin is the `updateAdmin` instruction.
type UpdateAdmin struct {

	// [0] = [WRITE] state
	//
	// [1] = [SIGNER] admin
	//
	// [2] = [WRITE, SIGNER] newAdmin
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateAdminInstructionBuilder creates a new `UpdateAdmin` instruction builder.
func NewUpdateAdminInstructionBuilder() *UpdateAdmin {
	nd := &UpdateAdmin{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetStateAccount sets the "state" account.
func (inst *UpdateAdmin) SetStateAccount(state ag_solanago.PublicKey) *UpdateAdmin {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state).WRITE()
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *UpdateAdmin) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdateAdmin) SetAdminAccount(admin ag_solanago.PublicKey) *UpdateAdmin {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdateAdmin) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetNewAdminAccount sets the "newAdmin" account.
func (inst *UpdateAdmin) SetNewAdminAccount(newAdmin ag_solanago.PublicKey) *UpdateAdmin {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(newAdmin).WRITE().SIGNER()
	return inst
}

// GetNewAdminAccount gets the "newAdmin" account.
func (inst *UpdateAdmin) GetNewAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdateAdmin) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateAdmin,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateAdmin) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateAdmin) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.NewAdmin is not set")
		}
	}
	return nil
}

func (inst *UpdateAdmin) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateAdmin")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("   admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("newAdmin", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj UpdateAdmin) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *UpdateAdmin) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewUpdateAdminInstruction declares a new UpdateAdmin instruction with the provided parameters and accounts.
func NewUpdateAdminInstruction(
	// Accounts:
	state ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	newAdmin ag_solanago.PublicKey) *UpdateAdmin {
	return NewUpdateAdminInstructionBuilder().
		SetStateAccount(state).
		SetAdminAccount(admin).
		SetNewAdminAccount(newAdmin)
}