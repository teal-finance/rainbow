// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package zeta

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitializeZetaMarket is the `initializeZetaMarket` instruction.
type InitializeZetaMarket struct {
	Args *InitializeMarketArgs

	// [0] = [] state
	//
	// [1] = [] marketIndexes
	//
	// [2] = [WRITE] zetaGroup
	//
	// [3] = [SIGNER] admin
	//
	// [4] = [WRITE] market
	//
	// [5] = [WRITE] requestQueue
	//
	// [6] = [WRITE] eventQueue
	//
	// [7] = [WRITE] bids
	//
	// [8] = [WRITE] asks
	//
	// [9] = [WRITE] baseMint
	//
	// [10] = [WRITE] quoteMint
	//
	// [11] = [WRITE] zetaBaseVault
	//
	// [12] = [WRITE] zetaQuoteVault
	//
	// [13] = [WRITE] dexBaseVault
	//
	// [14] = [WRITE] dexQuoteVault
	//
	// [15] = [] vaultOwner
	//
	// [16] = [] mintAuthority
	//
	// [17] = [] serumAuthority
	//
	// [18] = [] dexProgram
	//
	// [19] = [] systemProgram
	//
	// [20] = [] tokenProgram
	//
	// [21] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeZetaMarketInstructionBuilder creates a new `InitializeZetaMarket` instruction builder.
func NewInitializeZetaMarketInstructionBuilder() *InitializeZetaMarket {
	nd := &InitializeZetaMarket{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 22),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *InitializeZetaMarket) SetArgs(args InitializeMarketArgs) *InitializeZetaMarket {
	inst.Args = &args
	return inst
}

// SetStateAccount sets the "state" account.
func (inst *InitializeZetaMarket) SetStateAccount(state ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(state)
	return inst
}

// GetStateAccount gets the "state" account.
func (inst *InitializeZetaMarket) GetStateAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMarketIndexesAccount sets the "marketIndexes" account.
func (inst *InitializeZetaMarket) SetMarketIndexesAccount(marketIndexes ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(marketIndexes)
	return inst
}

// GetMarketIndexesAccount gets the "marketIndexes" account.
func (inst *InitializeZetaMarket) GetMarketIndexesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetZetaGroupAccount sets the "zetaGroup" account.
func (inst *InitializeZetaMarket) SetZetaGroupAccount(zetaGroup ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(zetaGroup).WRITE()
	return inst
}

// GetZetaGroupAccount gets the "zetaGroup" account.
func (inst *InitializeZetaMarket) GetZetaGroupAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAdminAccount sets the "admin" account.
func (inst *InitializeZetaMarket) SetAdminAccount(admin ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *InitializeZetaMarket) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMarketAccount sets the "market" account.
func (inst *InitializeZetaMarket) SetMarketAccount(market ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(market).WRITE()
	return inst
}

// GetMarketAccount gets the "market" account.
func (inst *InitializeZetaMarket) GetMarketAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetRequestQueueAccount sets the "requestQueue" account.
func (inst *InitializeZetaMarket) SetRequestQueueAccount(requestQueue ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(requestQueue).WRITE()
	return inst
}

// GetRequestQueueAccount gets the "requestQueue" account.
func (inst *InitializeZetaMarket) GetRequestQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetEventQueueAccount sets the "eventQueue" account.
func (inst *InitializeZetaMarket) SetEventQueueAccount(eventQueue ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(eventQueue).WRITE()
	return inst
}

// GetEventQueueAccount gets the "eventQueue" account.
func (inst *InitializeZetaMarket) GetEventQueueAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetBidsAccount sets the "bids" account.
func (inst *InitializeZetaMarket) SetBidsAccount(bids ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
func (inst *InitializeZetaMarket) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetAsksAccount sets the "asks" account.
func (inst *InitializeZetaMarket) SetAsksAccount(asks ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(asks).WRITE()
	return inst
}

// GetAsksAccount gets the "asks" account.
func (inst *InitializeZetaMarket) GetAsksAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetBaseMintAccount sets the "baseMint" account.
func (inst *InitializeZetaMarket) SetBaseMintAccount(baseMint ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(baseMint).WRITE()
	return inst
}

// GetBaseMintAccount gets the "baseMint" account.
func (inst *InitializeZetaMarket) GetBaseMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetQuoteMintAccount sets the "quoteMint" account.
func (inst *InitializeZetaMarket) SetQuoteMintAccount(quoteMint ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(quoteMint).WRITE()
	return inst
}

// GetQuoteMintAccount gets the "quoteMint" account.
func (inst *InitializeZetaMarket) GetQuoteMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetZetaBaseVaultAccount sets the "zetaBaseVault" account.
func (inst *InitializeZetaMarket) SetZetaBaseVaultAccount(zetaBaseVault ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(zetaBaseVault).WRITE()
	return inst
}

// GetZetaBaseVaultAccount gets the "zetaBaseVault" account.
func (inst *InitializeZetaMarket) GetZetaBaseVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetZetaQuoteVaultAccount sets the "zetaQuoteVault" account.
func (inst *InitializeZetaMarket) SetZetaQuoteVaultAccount(zetaQuoteVault ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(zetaQuoteVault).WRITE()
	return inst
}

// GetZetaQuoteVaultAccount gets the "zetaQuoteVault" account.
func (inst *InitializeZetaMarket) GetZetaQuoteVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetDexBaseVaultAccount sets the "dexBaseVault" account.
func (inst *InitializeZetaMarket) SetDexBaseVaultAccount(dexBaseVault ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(dexBaseVault).WRITE()
	return inst
}

// GetDexBaseVaultAccount gets the "dexBaseVault" account.
func (inst *InitializeZetaMarket) GetDexBaseVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetDexQuoteVaultAccount sets the "dexQuoteVault" account.
func (inst *InitializeZetaMarket) SetDexQuoteVaultAccount(dexQuoteVault ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(dexQuoteVault).WRITE()
	return inst
}

// GetDexQuoteVaultAccount gets the "dexQuoteVault" account.
func (inst *InitializeZetaMarket) GetDexQuoteVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetVaultOwnerAccount sets the "vaultOwner" account.
func (inst *InitializeZetaMarket) SetVaultOwnerAccount(vaultOwner ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(vaultOwner)
	return inst
}

// GetVaultOwnerAccount gets the "vaultOwner" account.
func (inst *InitializeZetaMarket) GetVaultOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetMintAuthorityAccount sets the "mintAuthority" account.
func (inst *InitializeZetaMarket) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(mintAuthority)
	return inst
}

// GetMintAuthorityAccount gets the "mintAuthority" account.
func (inst *InitializeZetaMarket) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetSerumAuthorityAccount sets the "serumAuthority" account.
func (inst *InitializeZetaMarket) SetSerumAuthorityAccount(serumAuthority ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(serumAuthority)
	return inst
}

// GetSerumAuthorityAccount gets the "serumAuthority" account.
func (inst *InitializeZetaMarket) GetSerumAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

// SetDexProgramAccount sets the "dexProgram" account.
func (inst *InitializeZetaMarket) SetDexProgramAccount(dexProgram ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(dexProgram)
	return inst
}

// GetDexProgramAccount gets the "dexProgram" account.
func (inst *InitializeZetaMarket) GetDexProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(18)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *InitializeZetaMarket) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *InitializeZetaMarket) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(19)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *InitializeZetaMarket) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *InitializeZetaMarket) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(20)
}

// SetRentAccount sets the "rent" account.
func (inst *InitializeZetaMarket) SetRentAccount(rent ag_solanago.PublicKey) *InitializeZetaMarket {
	inst.AccountMetaSlice[21] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *InitializeZetaMarket) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(21)
}

func (inst InitializeZetaMarket) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitializeZetaMarket,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitializeZetaMarket) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitializeZetaMarket) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.State is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MarketIndexes is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ZetaGroup is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Market is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.RequestQueue is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.EventQueue is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Asks is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.BaseMint is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.QuoteMint is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.ZetaBaseVault is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.ZetaQuoteVault is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.DexBaseVault is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.DexQuoteVault is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.VaultOwner is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.SerumAuthority is not set")
		}
		if inst.AccountMetaSlice[18] == nil {
			return errors.New("accounts.DexProgram is not set")
		}
		if inst.AccountMetaSlice[19] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[21] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *InitializeZetaMarket) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitializeZetaMarket")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=22]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         state", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" marketIndexes", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     zetaGroup", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         admin", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        market", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("  requestQueue", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("    eventQueue", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("          bids", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("          asks", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("      baseMint", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("     quoteMint", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta(" zetaBaseVault", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("zetaQuoteVault", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("  dexBaseVault", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta(" dexQuoteVault", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("    vaultOwner", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta(" mintAuthority", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("serumAuthority", inst.AccountMetaSlice.Get(17)))
						accountsBranch.Child(ag_format.Meta("    dexProgram", inst.AccountMetaSlice.Get(18)))
						accountsBranch.Child(ag_format.Meta(" systemProgram", inst.AccountMetaSlice.Get(19)))
						accountsBranch.Child(ag_format.Meta("  tokenProgram", inst.AccountMetaSlice.Get(20)))
						accountsBranch.Child(ag_format.Meta("          rent", inst.AccountMetaSlice.Get(21)))
					})
				})
		})
}

func (obj InitializeZetaMarket) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitializeZetaMarket) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeZetaMarketInstruction declares a new InitializeZetaMarket instruction with the provided parameters and accounts.
func NewInitializeZetaMarketInstruction(
	// Parameters:
	args InitializeMarketArgs,
	// Accounts:
	state ag_solanago.PublicKey,
	marketIndexes ag_solanago.PublicKey,
	zetaGroup ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	market ag_solanago.PublicKey,
	requestQueue ag_solanago.PublicKey,
	eventQueue ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	asks ag_solanago.PublicKey,
	baseMint ag_solanago.PublicKey,
	quoteMint ag_solanago.PublicKey,
	zetaBaseVault ag_solanago.PublicKey,
	zetaQuoteVault ag_solanago.PublicKey,
	dexBaseVault ag_solanago.PublicKey,
	dexQuoteVault ag_solanago.PublicKey,
	vaultOwner ag_solanago.PublicKey,
	mintAuthority ag_solanago.PublicKey,
	serumAuthority ag_solanago.PublicKey,
	dexProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *InitializeZetaMarket {
	return NewInitializeZetaMarketInstructionBuilder().
		SetArgs(args).
		SetStateAccount(state).
		SetMarketIndexesAccount(marketIndexes).
		SetZetaGroupAccount(zetaGroup).
		SetAdminAccount(admin).
		SetMarketAccount(market).
		SetRequestQueueAccount(requestQueue).
		SetEventQueueAccount(eventQueue).
		SetBidsAccount(bids).
		SetAsksAccount(asks).
		SetBaseMintAccount(baseMint).
		SetQuoteMintAccount(quoteMint).
		SetZetaBaseVaultAccount(zetaBaseVault).
		SetZetaQuoteVaultAccount(zetaQuoteVault).
		SetDexBaseVaultAccount(dexBaseVault).
		SetDexQuoteVaultAccount(dexQuoteVault).
		SetVaultOwnerAccount(vaultOwner).
		SetMintAuthorityAccount(mintAuthority).
		SetSerumAuthorityAccount(serumAuthority).
		SetDexProgramAccount(dexProgram).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram).
		SetRentAccount(rent)
}
