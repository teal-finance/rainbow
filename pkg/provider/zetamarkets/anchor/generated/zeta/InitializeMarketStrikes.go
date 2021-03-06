// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package zeta

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeMarketStrikes is the `initializeMarketStrikes` instruction.
type InitializeMarketStrikes struct {

	// [0] = [] state
	//
	// [1] = [WRITE] zetaGroup
	//
	// [2] = [] oracle
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeMarketStrikesInstructionBuilder creates a new `InitializeMarketStrikes` instruction builder.
func NewInitializeMarketStrikesInstructionBuilder() *InitializeMarketStrikes {
	nd := &InitializeMarketStrikes{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 3),
	}
	return nd
}

// SetStateAccount sets the "state" account.
func (inst *InitializeMarketStrikes) SetStateAccount(state ag_solanago.PublicKey) *InitializeMarketStrikes {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *InitializeMarketStrikes) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetZetaGroupAccount sets the "zetaGroup" account.
func (inst *InitializeMarketStrikes) SetZetaGroupAccount(zetaGroup ag_solanago.PublicKey) *InitializeMarketStrikes {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(zetaGroup).WRITE()
	return inst
}

// GetZetaGroupAccount gets the "zetaGroup" account.
func (inst *InitializeMarketStrikes) GetZetaGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetOracleAccount sets the "oracle" account.
func (inst *InitializeMarketStrikes) SetOracleAccount(oracle ag_solanago.PublicKey) *InitializeMarketStrikes {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(oracle)
	return inst
}

// GetOracleAccount gets the "oracle" account.
func (inst *InitializeMarketStrikes) GetOracleAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst InitializeMarketStrikes) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeMarketStrikes,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeMarketStrikes) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeMarketStrikes) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ZetaGroup is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Oracle is not set")
		}
	}
	return nil
}

func (inst *InitializeMarketStrikes) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeMarketStrikes")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("    state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("zetaGroup", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("   oracle", inst.AccountMetaSlice.Get(2)))
					})
				})
		})
}

func (obj InitializeMarketStrikes) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *InitializeMarketStrikes) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewInitializeMarketStrikesInstruction declares a new InitializeMarketStrikes instruction with the provided parameters and accounts.
func NewInitializeMarketStrikesInstruction(
	// Accounts:
	state ag_solanago.PublicKey,
	zetaGroup ag_solanago.PublicKey,
	oracle ag_solanago.PublicKey) *InitializeMarketStrikes {
	return NewInitializeMarketStrikesInstructionBuilder().
		SetStateAccount(state).
		SetZetaGroupAccount(zetaGroup).
		SetOracleAccount(oracle)
}
