// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package zeta

import (
	"bytes"
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
	ag_spew "github.com/spewerspew/spew"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "Zeta"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_InitializeZetaGroup = ag_binary.TypeID([8]byte{6, 135, 36, 232, 35, 39, 250, 71})

	Instruction_OverrideExpiry = ag_binary.TypeID([8]byte{129, 197, 117, 114, 108, 119, 207, 136})

	Instruction_InitializeMarginAccount = ag_binary.TypeID([8]byte{67, 235, 66, 102, 167, 171, 120, 197})

	Instruction_InitializeMarketIndexes = ag_binary.TypeID([8]byte{91, 63, 205, 144, 20, 83, 177, 120})

	Instruction_InitializeMarketNode = ag_binary.TypeID([8]byte{50, 118, 21, 21, 179, 248, 23, 128})

	Instruction_HaltZetaGroup = ag_binary.TypeID([8]byte{27, 67, 114, 244, 196, 125, 146, 117})

	Instruction_UnhaltZetaGroup = ag_binary.TypeID([8]byte{36, 108, 97, 51, 134, 6, 143, 253})

	Instruction_UpdateHaltState = ag_binary.TypeID([8]byte{215, 45, 53, 162, 149, 138, 5, 63})

	Instruction_UpdateVolatility = ag_binary.TypeID([8]byte{190, 105, 116, 221, 229, 198, 208, 83})

	Instruction_UpdateInterestRate = ag_binary.TypeID([8]byte{75, 8, 255, 41, 123, 59, 135, 238})

	Instruction_AddMarketIndexes = ag_binary.TypeID([8]byte{94, 246, 144, 175, 4, 164, 233, 252})

	Instruction_InitializeZetaState = ag_binary.TypeID([8]byte{68, 39, 75, 142, 191, 146, 94, 222})

	Instruction_UpdateAdmin = ag_binary.TypeID([8]byte{161, 176, 40, 213, 60, 184, 179, 228})

	Instruction_UpdateZetaState = ag_binary.TypeID([8]byte{104, 182, 20, 187, 3, 164, 60, 3})

	Instruction_UpdatePricingParameters = ag_binary.TypeID([8]byte{105, 127, 208, 134, 61, 61, 113, 247})

	Instruction_UpdateMarginParameters = ag_binary.TypeID([8]byte{69, 50, 174, 197, 123, 196, 72, 236})

	Instruction_CleanZetaMarkets = ag_binary.TypeID([8]byte{122, 127, 49, 89, 68, 228, 85, 157})

	Instruction_CleanZetaMarketsHalted = ag_binary.TypeID([8]byte{127, 228, 16, 244, 83, 68, 150, 94})

	Instruction_SettlePositions = ag_binary.TypeID([8]byte{43, 66, 200, 216, 219, 186, 44, 87})

	Instruction_SettlePositionsHalted = ag_binary.TypeID([8]byte{170, 147, 139, 163, 19, 104, 167, 77})

	Instruction_InitializeMarketStrikes = ag_binary.TypeID([8]byte{189, 46, 255, 33, 126, 133, 43, 171})

	Instruction_ExpireSeriesOverride = ag_binary.TypeID([8]byte{104, 22, 34, 123, 86, 224, 130, 70})

	Instruction_ExpireSeries = ag_binary.TypeID([8]byte{45, 162, 105, 98, 44, 21, 171, 127})

	Instruction_InitializeZetaMarket = ag_binary.TypeID([8]byte{116, 239, 226, 149, 46, 163, 221, 3})

	Instruction_RetreatMarketNodes = ag_binary.TypeID([8]byte{10, 124, 88, 233, 87, 190, 22, 179})

	Instruction_CleanMarketNodes = ag_binary.TypeID([8]byte{57, 35, 134, 162, 20, 214, 55, 227})

	Instruction_UpdateVolatilityNodes = ag_binary.TypeID([8]byte{100, 170, 196, 34, 72, 228, 219, 236})

	Instruction_UpdatePricing = ag_binary.TypeID([8]byte{157, 225, 208, 150, 23, 153, 253, 18})

	Instruction_UpdatePricingHalted = ag_binary.TypeID([8]byte{3, 71, 109, 241, 165, 14, 38, 19})

	Instruction_Deposit = ag_binary.TypeID([8]byte{242, 35, 198, 137, 82, 225, 242, 182})

	Instruction_DepositInsuranceVault = ag_binary.TypeID([8]byte{47, 53, 25, 47, 109, 122, 22, 22})

	Instruction_Withdraw = ag_binary.TypeID([8]byte{183, 18, 70, 156, 148, 109, 161, 34})

	Instruction_WithdrawInsuranceVault = ag_binary.TypeID([8]byte{17, 250, 213, 45, 172, 117, 81, 225})

	Instruction_InitializeOpenOrders = ag_binary.TypeID([8]byte{55, 234, 16, 82, 100, 42, 126, 192})

	Instruction_InitializeWhitelistDepositAccount = ag_binary.TypeID([8]byte{61, 231, 115, 219, 81, 243, 158, 138})

	Instruction_InitializeWhitelistInsuranceAccount = ag_binary.TypeID([8]byte{43, 46, 240, 155, 80, 4, 86, 102})

	Instruction_InitializeWhitelistTradingFeesAccount = ag_binary.TypeID([8]byte{198, 129, 216, 185, 247, 29, 105, 190})

	Instruction_InitializeInsuranceDepositAccount = ag_binary.TypeID([8]byte{85, 163, 114, 121, 139, 167, 41, 37})

	Instruction_PlaceOrder = ag_binary.TypeID([8]byte{51, 194, 155, 175, 109, 130, 96, 106})

	Instruction_CancelOrder = ag_binary.TypeID([8]byte{95, 129, 237, 240, 8, 49, 223, 132})

	Instruction_CancelOrderHalted = ag_binary.TypeID([8]byte{0, 192, 233, 2, 252, 251, 130, 169})

	Instruction_CancelOrderByClientOrderId = ag_binary.TypeID([8]byte{115, 178, 201, 8, 175, 183, 123, 119})

	Instruction_CancelExpiredOrder = ag_binary.TypeID([8]byte{216, 120, 64, 235, 155, 19, 229, 99})

	Instruction_ForceCancelOrders = ag_binary.TypeID([8]byte{64, 181, 196, 63, 222, 72, 64, 232})

	Instruction_CrankEventQueue = ag_binary.TypeID([8]byte{67, 133, 97, 223, 178, 188, 235, 181})

	Instruction_RebalanceInsuranceVault = ag_binary.TypeID([8]byte{11, 196, 66, 235, 59, 237, 223, 111})

	Instruction_Liquidate = ag_binary.TypeID([8]byte{223, 179, 226, 125, 48, 46, 39, 74})

	Instruction_BurnVaultTokens = ag_binary.TypeID([8]byte{233, 203, 165, 201, 175, 43, 188, 159})

	Instruction_SettleDexFunds = ag_binary.TypeID([8]byte{165, 103, 142, 38, 211, 166, 14, 226})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_InitializeZetaGroup:
		return "InitializeZetaGroup"
	case Instruction_OverrideExpiry:
		return "OverrideExpiry"
	case Instruction_InitializeMarginAccount:
		return "InitializeMarginAccount"
	case Instruction_InitializeMarketIndexes:
		return "InitializeMarketIndexes"
	case Instruction_InitializeMarketNode:
		return "InitializeMarketNode"
	case Instruction_HaltZetaGroup:
		return "HaltZetaGroup"
	case Instruction_UnhaltZetaGroup:
		return "UnhaltZetaGroup"
	case Instruction_UpdateHaltState:
		return "UpdateHaltState"
	case Instruction_UpdateVolatility:
		return "UpdateVolatility"
	case Instruction_UpdateInterestRate:
		return "UpdateInterestRate"
	case Instruction_AddMarketIndexes:
		return "AddMarketIndexes"
	case Instruction_InitializeZetaState:
		return "InitializeZetaState"
	case Instruction_UpdateAdmin:
		return "UpdateAdmin"
	case Instruction_UpdateZetaState:
		return "UpdateZetaState"
	case Instruction_UpdatePricingParameters:
		return "UpdatePricingParameters"
	case Instruction_UpdateMarginParameters:
		return "UpdateMarginParameters"
	case Instruction_CleanZetaMarkets:
		return "CleanZetaMarkets"
	case Instruction_CleanZetaMarketsHalted:
		return "CleanZetaMarketsHalted"
	case Instruction_SettlePositions:
		return "SettlePositions"
	case Instruction_SettlePositionsHalted:
		return "SettlePositionsHalted"
	case Instruction_InitializeMarketStrikes:
		return "InitializeMarketStrikes"
	case Instruction_ExpireSeriesOverride:
		return "ExpireSeriesOverride"
	case Instruction_ExpireSeries:
		return "ExpireSeries"
	case Instruction_InitializeZetaMarket:
		return "InitializeZetaMarket"
	case Instruction_RetreatMarketNodes:
		return "RetreatMarketNodes"
	case Instruction_CleanMarketNodes:
		return "CleanMarketNodes"
	case Instruction_UpdateVolatilityNodes:
		return "UpdateVolatilityNodes"
	case Instruction_UpdatePricing:
		return "UpdatePricing"
	case Instruction_UpdatePricingHalted:
		return "UpdatePricingHalted"
	case Instruction_Deposit:
		return "Deposit"
	case Instruction_DepositInsuranceVault:
		return "DepositInsuranceVault"
	case Instruction_Withdraw:
		return "Withdraw"
	case Instruction_WithdrawInsuranceVault:
		return "WithdrawInsuranceVault"
	case Instruction_InitializeOpenOrders:
		return "InitializeOpenOrders"
	case Instruction_InitializeWhitelistDepositAccount:
		return "InitializeWhitelistDepositAccount"
	case Instruction_InitializeWhitelistInsuranceAccount:
		return "InitializeWhitelistInsuranceAccount"
	case Instruction_InitializeWhitelistTradingFeesAccount:
		return "InitializeWhitelistTradingFeesAccount"
	case Instruction_InitializeInsuranceDepositAccount:
		return "InitializeInsuranceDepositAccount"
	case Instruction_PlaceOrder:
		return "PlaceOrder"
	case Instruction_CancelOrder:
		return "CancelOrder"
	case Instruction_CancelOrderHalted:
		return "CancelOrderHalted"
	case Instruction_CancelOrderByClientOrderId:
		return "CancelOrderByClientOrderId"
	case Instruction_CancelExpiredOrder:
		return "CancelExpiredOrder"
	case Instruction_ForceCancelOrders:
		return "ForceCancelOrders"
	case Instruction_CrankEventQueue:
		return "CrankEventQueue"
	case Instruction_RebalanceInsuranceVault:
		return "RebalanceInsuranceVault"
	case Instruction_Liquidate:
		return "Liquidate"
	case Instruction_BurnVaultTokens:
		return "BurnVaultTokens"
	case Instruction_SettleDexFunds:
		return "SettleDexFunds"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"initialize_zeta_group", (*InitializeZetaGroup)(nil),
		},
		{
			"override_expiry", (*OverrideExpiry)(nil),
		},
		{
			"initialize_margin_account", (*InitializeMarginAccount)(nil),
		},
		{
			"initialize_market_indexes", (*InitializeMarketIndexes)(nil),
		},
		{
			"initialize_market_node", (*InitializeMarketNode)(nil),
		},
		{
			"halt_zeta_group", (*HaltZetaGroup)(nil),
		},
		{
			"unhalt_zeta_group", (*UnhaltZetaGroup)(nil),
		},
		{
			"update_halt_state", (*UpdateHaltState)(nil),
		},
		{
			"update_volatility", (*UpdateVolatility)(nil),
		},
		{
			"update_interest_rate", (*UpdateInterestRate)(nil),
		},
		{
			"add_market_indexes", (*AddMarketIndexes)(nil),
		},
		{
			"initialize_zeta_state", (*InitializeZetaState)(nil),
		},
		{
			"update_admin", (*UpdateAdmin)(nil),
		},
		{
			"update_zeta_state", (*UpdateZetaState)(nil),
		},
		{
			"update_pricing_parameters", (*UpdatePricingParameters)(nil),
		},
		{
			"update_margin_parameters", (*UpdateMarginParameters)(nil),
		},
		{
			"clean_zeta_markets", (*CleanZetaMarkets)(nil),
		},
		{
			"clean_zeta_markets_halted", (*CleanZetaMarketsHalted)(nil),
		},
		{
			"settle_positions", (*SettlePositions)(nil),
		},
		{
			"settle_positions_halted", (*SettlePositionsHalted)(nil),
		},
		{
			"initialize_market_strikes", (*InitializeMarketStrikes)(nil),
		},
		{
			"expire_series_override", (*ExpireSeriesOverride)(nil),
		},
		{
			"expire_series", (*ExpireSeries)(nil),
		},
		{
			"initialize_zeta_market", (*InitializeZetaMarket)(nil),
		},
		{
			"retreat_market_nodes", (*RetreatMarketNodes)(nil),
		},
		{
			"clean_market_nodes", (*CleanMarketNodes)(nil),
		},
		{
			"update_volatility_nodes", (*UpdateVolatilityNodes)(nil),
		},
		{
			"update_pricing", (*UpdatePricing)(nil),
		},
		{
			"update_pricing_halted", (*UpdatePricingHalted)(nil),
		},
		{
			"deposit", (*Deposit)(nil),
		},
		{
			"deposit_insurance_vault", (*DepositInsuranceVault)(nil),
		},
		{
			"withdraw", (*Withdraw)(nil),
		},
		{
			"withdraw_insurance_vault", (*WithdrawInsuranceVault)(nil),
		},
		{
			"initialize_open_orders", (*InitializeOpenOrders)(nil),
		},
		{
			"initialize_whitelist_deposit_account", (*InitializeWhitelistDepositAccount)(nil),
		},
		{
			"initialize_whitelist_insurance_account", (*InitializeWhitelistInsuranceAccount)(nil),
		},
		{
			"initialize_whitelist_trading_fees_account", (*InitializeWhitelistTradingFeesAccount)(nil),
		},
		{
			"initialize_insurance_deposit_account", (*InitializeInsuranceDepositAccount)(nil),
		},
		{
			"place_order", (*PlaceOrder)(nil),
		},
		{
			"cancel_order", (*CancelOrder)(nil),
		},
		{
			"cancel_order_halted", (*CancelOrderHalted)(nil),
		},
		{
			"cancel_order_by_client_order_id", (*CancelOrderByClientOrderId)(nil),
		},
		{
			"cancel_expired_order", (*CancelExpiredOrder)(nil),
		},
		{
			"force_cancel_orders", (*ForceCancelOrders)(nil),
		},
		{
			"crank_event_queue", (*CrankEventQueue)(nil),
		},
		{
			"rebalance_insurance_vault", (*RebalanceInsuranceVault)(nil),
		},
		{
			"liquidate", (*Liquidate)(nil),
		},
		{
			"burn_vault_tokens", (*BurnVaultTokens)(nil),
		},
		{
			"settle_dex_funds", (*SettleDexFunds)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
