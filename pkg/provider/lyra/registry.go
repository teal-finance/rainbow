// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lyra

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LyraRegistryOptionMarketAddresses is an auto generated low-level Go binding around an user-defined struct.
type LyraRegistryOptionMarketAddresses struct {
	LiquidityPool      common.Address
	LiquidityToken     common.Address
	GreekCache         common.Address
	OptionMarket       common.Address
	OptionMarketPricer common.Address
	OptionToken        common.Address
	PoolHedger         common.Address
	ShortCollateral    common.Address
	GwavOracle         common.Address
	QuoteAsset         common.Address
	BaseAsset          common.Address
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"contractName\",\"type\":\"bytes32\"}],\"name\":\"NonExistentGlobalContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"optionMarket\",\"type\":\"address\"}],\"name\":\"NonExistentMarket\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nominatedOwner\",\"type\":\"address\"}],\"name\":\"OnlyNominatedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"thrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"RemovingInvalidMarket\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"GlobalAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"MarketRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractGWAVOracle\",\"name\":\"gwavOracle\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structLyraRegistry.OptionMarketAddresses\",\"name\":\"market\",\"type\":\"tuple\"}],\"name\":\"MarketUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractGWAVOracle\",\"name\":\"gwavOracle\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structLyraRegistry.OptionMarketAddresses\",\"name\":\"newMarketAddresses\",\"type\":\"tuple\"}],\"name\":\"addMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"contractName\",\"type\":\"bytes32\"}],\"name\":\"getGlobalAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"globalContract\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"}],\"name\":\"getMarketAddresses\",\"outputs\":[{\"components\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractGWAVOracle\",\"name\":\"gwavOracle\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"internalType\":\"structLyraRegistry.OptionMarketAddresses\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"marketAddresses\",\"outputs\":[{\"internalType\":\"contractLiquidityPool\",\"name\":\"liquidityPool\",\"type\":\"address\"},{\"internalType\":\"contractLiquidityToken\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"contractOptionGreekCache\",\"name\":\"greekCache\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarket\",\"name\":\"optionMarket\",\"type\":\"address\"},{\"internalType\":\"contractOptionMarketPricer\",\"name\":\"optionMarketPricer\",\"type\":\"address\"},{\"internalType\":\"contractOptionToken\",\"name\":\"optionToken\",\"type\":\"address\"},{\"internalType\":\"contractPoolHedger\",\"name\":\"poolHedger\",\"type\":\"address\"},{\"internalType\":\"contractShortCollateral\",\"name\":\"shortCollateral\",\"type\":\"address\"},{\"internalType\":\"contractGWAVOracle\",\"name\":\"gwavOracle\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"baseAsset\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"optionMarkets\",\"outputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOptionMarket\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"removeMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"names\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"updateGlobalAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}],\"bytecode\":\"0x608060405234801561001057600080fd5b50600080546001600160a01b0319163390811782556040805192835260208301919091527fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c910160405180910390a16113c98061006e6000396000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c80637f2e060011610081578063c4158a511161005b578063c4158a51146101e2578063c71b7e5314610202578063db9132361461030e57600080fd5b80637f2e0600146101795780638acce68f1461018c5780638da5cb5b146101c257600080fd5b806330a12dc4116100b257806330a12dc41461013e57806353a47bb71461015157806379ba50971461017157600080fd5b80631627540c146100d9578063199bc905146100ee5780631cb6684514610101575b600080fd5b6100ec6100e7366004610e1e565b610321565b005b6100ec6100fc366004610ef5565b6103a2565b61011461010f366004610fcf565b6105cf565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100ec61014c366004611080565b610606565b6001546101149073ffffffffffffffffffffffffffffffffffffffff1681565b6100ec6107aa565b610114610187366004610fcf565b6108c7565b61011461019a366004610fcf565b60046020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6000546101149073ffffffffffffffffffffffffffffffffffffffff1681565b6101f56101f0366004610e1e565b610931565b6040516101359190611139565b610296610210366004610e1e565b600360208190526000918252604090912080546001820154600283015493830154600484015460058501546006860154600787015460088801546009890154600a9099015473ffffffffffffffffffffffffffffffffffffffff9889169a978916999789169896871697958716969485169593851694928316939183169291821691168b565b6040805173ffffffffffffffffffffffffffffffffffffffff9c8d1681529a8c1660208c0152988b16988a01989098529589166060890152938816608088015291871660a0870152861660c0860152851660e08501528416610100840152831661012083015290911661014082015261016001610135565b6100ec61031c366004610e1e565b610a9a565b610329610aae565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229060200160405180910390a150565b6103aa610aae565b606081015173ffffffffffffffffffffffffffffffffffffffff9081166000908152600360208190526040909120015416610459576060810151600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790555b60608101805173ffffffffffffffffffffffffffffffffffffffff908116600090815260036020818152604092839020865181547fffffffffffffffffffffffff000000000000000000000000000000000000000090811691871691909117825591870151600182018054841691871691909117905583870151600282018054841691871691909117905594519185018054821692851692831790556080860151600486018054831691861691909117905560a0860151600586018054831691861691909117905560c0860151600686018054831691861691909117905560e0860151600786018054831691861691909117905561010086015160088601805483169186169190911790556101208601516009860180548316918616919091179055610140860151600a90950180549091169490931693909317909155517fa090264792b8766dd953a6d1775028cdd591d04a748e045c19b98135970e9127906105c4908490611139565b60405180910390a250565b600281815481106105df57600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b61060e610aae565b81518151811461067f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f6c656e677468206d69736d61746368000000000000000000000000000000000060448201526064015b60405180910390fd5b60005b818110156107a45782818151811061069c5761069c6112b6565b6020026020010151600460008684815181106106ba576106ba6112b6565b6020026020010151815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838181518110610720576107206112b6565b60200260200101517fdf028fe8b079d8195054d23416d2537debdd0f58695e83cd840d822e6a4c2e8084838151811061075b5761075b6112b6565b602002602001015160405161078c919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b60405180910390a261079d81611314565b9050610682565b50505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610823576001546040517f96cf9ed800000000000000000000000000000000000000000000000000000000815230600482015233602482015273ffffffffffffffffffffffffffffffffffffffff9091166044820152606401610676565b6000546001546040805173ffffffffffffffffffffffffffffffffffffffff93841681529290911660208301527fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c910160405180910390a160018054600080547fffffffffffffffffffffffff000000000000000000000000000000000000000090811673ffffffffffffffffffffffffffffffffffffffff841617909155169055565b60008181526004602052604090205473ffffffffffffffffffffffffffffffffffffffff1680156108f757919050565b6040517f8e5b5d0800000000000000000000000000000000000000000000000000000000815260048101839052602401610676565b919050565b6040805161016081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081019190915273ffffffffffffffffffffffffffffffffffffffff80831660009081526003602081815260409283902083516101608101855281548616815260018201548616928101929092526002810154851693820193909352908201548316606082018190526004830154841660808301526005830154841660a08301526006830154841660c08301526007830154841660e08301526008830154841661010083015260098301548416610120830152600a9092015490921661014083015215610a505792915050565b6040517f228c8d1600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610676565b610aa2610aae565b610aab81610b29565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610b27576000546040517f1abc2f9800000000000000000000000000000000000000000000000000000000815230600482015233602482015273ffffffffffffffffffffffffffffffffffffffff9091166044820152606401610676565b565b600254600080805b83811015610ba7578473ffffffffffffffffffffffffffffffffffffffff1660028281548110610b6357610b636112b6565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff161415610b975780925060019150610ba7565b610ba081611314565b9050610b31565b5080610bfd576040517ff3aface400000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff85166024820152604401610676565b6002610c0a60018561134d565b81548110610c1a57610c1a6112b6565b6000918252602090912001546002805473ffffffffffffffffffffffffffffffffffffffff9092169184908110610c5357610c536112b6565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506002805480610cac57610cac611364565b60008281526020812082017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590910190915560405173ffffffffffffffffffffffffffffffffffffffff8616917f59d7b1e52008dc342c9421dadfc773114b914a65682a4e4b53cf60a970df0d7791a250505073ffffffffffffffffffffffffffffffffffffffff16600090815260036020819052604090912080547fffffffffffffffffffffffff000000000000000000000000000000000000000090811682556001820180548216905560028201805482169055918101805483169055600481018054831690556005810180548316905560068101805483169055600781018054831690556008810180548316905560098101805483169055600a0180549091169055565b73ffffffffffffffffffffffffffffffffffffffff81168114610aab57600080fd5b600060208284031215610e3057600080fd5b8135610e3b81610dfc565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610160810167ffffffffffffffff81118282101715610e9557610e95610e42565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610ee257610ee2610e42565b604052919050565b803561092c81610dfc565b60006101608284031215610f0857600080fd5b610f10610e71565b610f1983610eea565b8152610f2760208401610eea565b6020820152610f3860408401610eea565b6040820152610f4960608401610eea565b6060820152610f5a60808401610eea565b6080820152610f6b60a08401610eea565b60a0820152610f7c60c08401610eea565b60c0820152610f8d60e08401610eea565b60e0820152610100610fa0818501610eea565b90820152610120610fb2848201610eea565b90820152610140610fc4848201610eea565b908201529392505050565b600060208284031215610fe157600080fd5b5035919050565b600067ffffffffffffffff82111561100257611002610e42565b5060051b60200190565b600082601f83011261101d57600080fd5b8135602061103261102d83610fe8565b610e9b565b82815260059290921b8401810191818101908684111561105157600080fd5b8286015b8481101561107557803561106881610dfc565b8352918301918301611055565b509695505050505050565b6000806040838503121561109357600080fd5b823567ffffffffffffffff808211156110ab57600080fd5b818501915085601f8301126110bf57600080fd5b813560206110cf61102d83610fe8565b82815260059290921b840181019181810190898411156110ee57600080fd5b948201945b8386101561110c578535825294820194908201906110f3565b9650508601359250508082111561112257600080fd5b5061112f8582860161100c565b9150509250929050565b815173ffffffffffffffffffffffffffffffffffffffff1681526101608101602083015161117f602084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060408301516111a7604084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060608301516111cf606084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060808301516111f7608084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060a083015161121f60a084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060c083015161124760c084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060e083015161126f60e084018273ffffffffffffffffffffffffffffffffffffffff169052565b506101008381015173ffffffffffffffffffffffffffffffffffffffff90811691840191909152610120808501518216908401526101409384015116929091019190915290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611346576113466112e5565b5060010190565b60008282101561135f5761135f6112e5565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea264697066735822122042d5b51a2967d3b5d324ef0e55f66327cc0bbbd893c48589f4b15d830068e9db64736f6c63430008090033\",\"deployedBytecode\":\"0x608060405234801561001057600080fd5b50600436106100d45760003560e01c80637f2e060011610081578063c4158a511161005b578063c4158a51146101e2578063c71b7e5314610202578063db9132361461030e57600080fd5b80637f2e0600146101795780638acce68f1461018c5780638da5cb5b146101c257600080fd5b806330a12dc4116100b257806330a12dc41461013e57806353a47bb71461015157806379ba50971461017157600080fd5b80631627540c146100d9578063199bc905146100ee5780631cb6684514610101575b600080fd5b6100ec6100e7366004610e1e565b610321565b005b6100ec6100fc366004610ef5565b6103a2565b61011461010f366004610fcf565b6105cf565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100ec61014c366004611080565b610606565b6001546101149073ffffffffffffffffffffffffffffffffffffffff1681565b6100ec6107aa565b610114610187366004610fcf565b6108c7565b61011461019a366004610fcf565b60046020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6000546101149073ffffffffffffffffffffffffffffffffffffffff1681565b6101f56101f0366004610e1e565b610931565b6040516101359190611139565b610296610210366004610e1e565b600360208190526000918252604090912080546001820154600283015493830154600484015460058501546006860154600787015460088801546009890154600a9099015473ffffffffffffffffffffffffffffffffffffffff9889169a978916999789169896871697958716969485169593851694928316939183169291821691168b565b6040805173ffffffffffffffffffffffffffffffffffffffff9c8d1681529a8c1660208c0152988b16988a01989098529589166060890152938816608088015291871660a0870152861660c0860152851660e08501528416610100840152831661012083015290911661014082015261016001610135565b6100ec61031c366004610e1e565b610a9a565b610329610aae565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce229060200160405180910390a150565b6103aa610aae565b606081015173ffffffffffffffffffffffffffffffffffffffff9081166000908152600360208190526040909120015416610459576060810151600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9092169190911790555b60608101805173ffffffffffffffffffffffffffffffffffffffff908116600090815260036020818152604092839020865181547fffffffffffffffffffffffff000000000000000000000000000000000000000090811691871691909117825591870151600182018054841691871691909117905583870151600282018054841691871691909117905594519185018054821692851692831790556080860151600486018054831691861691909117905560a0860151600586018054831691861691909117905560c0860151600686018054831691861691909117905560e0860151600786018054831691861691909117905561010086015160088601805483169186169190911790556101208601516009860180548316918616919091179055610140860151600a90950180549091169490931693909317909155517fa090264792b8766dd953a6d1775028cdd591d04a748e045c19b98135970e9127906105c4908490611139565b60405180910390a250565b600281815481106105df57600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b61060e610aae565b81518151811461067f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f6c656e677468206d69736d61746368000000000000000000000000000000000060448201526064015b60405180910390fd5b60005b818110156107a45782818151811061069c5761069c6112b6565b6020026020010151600460008684815181106106ba576106ba6112b6565b6020026020010151815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838181518110610720576107206112b6565b60200260200101517fdf028fe8b079d8195054d23416d2537debdd0f58695e83cd840d822e6a4c2e8084838151811061075b5761075b6112b6565b602002602001015160405161078c919073ffffffffffffffffffffffffffffffffffffffff91909116815260200190565b60405180910390a261079d81611314565b9050610682565b50505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610823576001546040517f96cf9ed800000000000000000000000000000000000000000000000000000000815230600482015233602482015273ffffffffffffffffffffffffffffffffffffffff9091166044820152606401610676565b6000546001546040805173ffffffffffffffffffffffffffffffffffffffff93841681529290911660208301527fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c910160405180910390a160018054600080547fffffffffffffffffffffffff000000000000000000000000000000000000000090811673ffffffffffffffffffffffffffffffffffffffff841617909155169055565b60008181526004602052604090205473ffffffffffffffffffffffffffffffffffffffff1680156108f757919050565b6040517f8e5b5d0800000000000000000000000000000000000000000000000000000000815260048101839052602401610676565b919050565b6040805161016081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052610120810182905261014081019190915273ffffffffffffffffffffffffffffffffffffffff80831660009081526003602081815260409283902083516101608101855281548616815260018201548616928101929092526002810154851693820193909352908201548316606082018190526004830154841660808301526005830154841660a08301526006830154841660c08301526007830154841660e08301526008830154841661010083015260098301548416610120830152600a9092015490921661014083015215610a505792915050565b6040517f228c8d1600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602401610676565b610aa2610aae565b610aab81610b29565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610b27576000546040517f1abc2f9800000000000000000000000000000000000000000000000000000000815230600482015233602482015273ffffffffffffffffffffffffffffffffffffffff9091166044820152606401610676565b565b600254600080805b83811015610ba7578473ffffffffffffffffffffffffffffffffffffffff1660028281548110610b6357610b636112b6565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff161415610b975780925060019150610ba7565b610ba081611314565b9050610b31565b5080610bfd576040517ff3aface400000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff85166024820152604401610676565b6002610c0a60018561134d565b81548110610c1a57610c1a6112b6565b6000918252602090912001546002805473ffffffffffffffffffffffffffffffffffffffff9092169184908110610c5357610c536112b6565b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506002805480610cac57610cac611364565b60008281526020812082017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590910190915560405173ffffffffffffffffffffffffffffffffffffffff8616917f59d7b1e52008dc342c9421dadfc773114b914a65682a4e4b53cf60a970df0d7791a250505073ffffffffffffffffffffffffffffffffffffffff16600090815260036020819052604090912080547fffffffffffffffffffffffff000000000000000000000000000000000000000090811682556001820180548216905560028201805482169055918101805483169055600481018054831690556005810180548316905560068101805483169055600781018054831690556008810180548316905560098101805483169055600a0180549091169055565b73ffffffffffffffffffffffffffffffffffffffff81168114610aab57600080fd5b600060208284031215610e3057600080fd5b8135610e3b81610dfc565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610160810167ffffffffffffffff81118282101715610e9557610e95610e42565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610ee257610ee2610e42565b604052919050565b803561092c81610dfc565b60006101608284031215610f0857600080fd5b610f10610e71565b610f1983610eea565b8152610f2760208401610eea565b6020820152610f3860408401610eea565b6040820152610f4960608401610eea565b6060820152610f5a60808401610eea565b6080820152610f6b60a08401610eea565b60a0820152610f7c60c08401610eea565b60c0820152610f8d60e08401610eea565b60e0820152610100610fa0818501610eea565b90820152610120610fb2848201610eea565b90820152610140610fc4848201610eea565b908201529392505050565b600060208284031215610fe157600080fd5b5035919050565b600067ffffffffffffffff82111561100257611002610e42565b5060051b60200190565b600082601f83011261101d57600080fd5b8135602061103261102d83610fe8565b610e9b565b82815260059290921b8401810191818101908684111561105157600080fd5b8286015b8481101561107557803561106881610dfc565b8352918301918301611055565b509695505050505050565b6000806040838503121561109357600080fd5b823567ffffffffffffffff808211156110ab57600080fd5b818501915085601f8301126110bf57600080fd5b813560206110cf61102d83610fe8565b82815260059290921b840181019181810190898411156110ee57600080fd5b948201945b8386101561110c578535825294820194908201906110f3565b9650508601359250508082111561112257600080fd5b5061112f8582860161100c565b9150509250929050565b815173ffffffffffffffffffffffffffffffffffffffff1681526101608101602083015161117f602084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060408301516111a7604084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060608301516111cf606084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060808301516111f7608084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060a083015161121f60a084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060c083015161124760c084018273ffffffffffffffffffffffffffffffffffffffff169052565b5060e083015161126f60e084018273ffffffffffffffffffffffffffffffffffffffff169052565b506101008381015173ffffffffffffffffffffffffffffffffffffffff90811691840191909152610120808501518216908401526101409384015116929091019190915290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611346576113466112e5565b5060010190565b60008282101561135f5761135f6112e5565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea264697066735822122042d5b51a2967d3b5d324ef0e55f66327cc0bbbd893c48589f4b15d830068e9db64736f6c63430008090033\",\"linkReferences\":{},\"deployedLinkReferences\":{}}",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// GetGlobalAddress is a free data retrieval call binding the contract method 0x7f2e0600.
//
// Solidity: function getGlobalAddress(bytes32 contractName) view returns(address globalContract)
func (_Registry *RegistryCaller) GetGlobalAddress(opts *bind.CallOpts, contractName [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getGlobalAddress", contractName)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGlobalAddress is a free data retrieval call binding the contract method 0x7f2e0600.
//
// Solidity: function getGlobalAddress(bytes32 contractName) view returns(address globalContract)
func (_Registry *RegistrySession) GetGlobalAddress(contractName [32]byte) (common.Address, error) {
	return _Registry.Contract.GetGlobalAddress(&_Registry.CallOpts, contractName)
}

// GetGlobalAddress is a free data retrieval call binding the contract method 0x7f2e0600.
//
// Solidity: function getGlobalAddress(bytes32 contractName) view returns(address globalContract)
func (_Registry *RegistryCallerSession) GetGlobalAddress(contractName [32]byte) (common.Address, error) {
	return _Registry.Contract.GetGlobalAddress(&_Registry.CallOpts, contractName)
}

// GetMarketAddresses is a free data retrieval call binding the contract method 0xc4158a51.
//
// Solidity: function getMarketAddresses(address optionMarket) view returns((address,address,address,address,address,address,address,address,address,address,address))
func (_Registry *RegistryCaller) GetMarketAddresses(opts *bind.CallOpts, optionMarket common.Address) (LyraRegistryOptionMarketAddresses, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getMarketAddresses", optionMarket)

	if err != nil {
		return *new(LyraRegistryOptionMarketAddresses), err
	}

	out0 := *abi.ConvertType(out[0], new(LyraRegistryOptionMarketAddresses)).(*LyraRegistryOptionMarketAddresses)

	return out0, err

}

// GetMarketAddresses is a free data retrieval call binding the contract method 0xc4158a51.
//
// Solidity: function getMarketAddresses(address optionMarket) view returns((address,address,address,address,address,address,address,address,address,address,address))
func (_Registry *RegistrySession) GetMarketAddresses(optionMarket common.Address) (LyraRegistryOptionMarketAddresses, error) {
	return _Registry.Contract.GetMarketAddresses(&_Registry.CallOpts, optionMarket)
}

// GetMarketAddresses is a free data retrieval call binding the contract method 0xc4158a51.
//
// Solidity: function getMarketAddresses(address optionMarket) view returns((address,address,address,address,address,address,address,address,address,address,address))
func (_Registry *RegistryCallerSession) GetMarketAddresses(optionMarket common.Address) (LyraRegistryOptionMarketAddresses, error) {
	return _Registry.Contract.GetMarketAddresses(&_Registry.CallOpts, optionMarket)
}

// GlobalAddresses is a free data retrieval call binding the contract method 0x8acce68f.
//
// Solidity: function globalAddresses(bytes32 ) view returns(address)
func (_Registry *RegistryCaller) GlobalAddresses(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "globalAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalAddresses is a free data retrieval call binding the contract method 0x8acce68f.
//
// Solidity: function globalAddresses(bytes32 ) view returns(address)
func (_Registry *RegistrySession) GlobalAddresses(arg0 [32]byte) (common.Address, error) {
	return _Registry.Contract.GlobalAddresses(&_Registry.CallOpts, arg0)
}

// GlobalAddresses is a free data retrieval call binding the contract method 0x8acce68f.
//
// Solidity: function globalAddresses(bytes32 ) view returns(address)
func (_Registry *RegistryCallerSession) GlobalAddresses(arg0 [32]byte) (common.Address, error) {
	return _Registry.Contract.GlobalAddresses(&_Registry.CallOpts, arg0)
}

// MarketAddresses is a free data retrieval call binding the contract method 0xc71b7e53.
//
// Solidity: function marketAddresses(address ) view returns(address liquidityPool, address liquidityToken, address greekCache, address optionMarket, address optionMarketPricer, address optionToken, address poolHedger, address shortCollateral, address gwavOracle, address quoteAsset, address baseAsset)
func (_Registry *RegistryCaller) MarketAddresses(opts *bind.CallOpts, arg0 common.Address) (struct {
	LiquidityPool      common.Address
	LiquidityToken     common.Address
	GreekCache         common.Address
	OptionMarket       common.Address
	OptionMarketPricer common.Address
	OptionToken        common.Address
	PoolHedger         common.Address
	ShortCollateral    common.Address
	GwavOracle         common.Address
	QuoteAsset         common.Address
	BaseAsset          common.Address
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "marketAddresses", arg0)

	outstruct := new(struct {
		LiquidityPool      common.Address
		LiquidityToken     common.Address
		GreekCache         common.Address
		OptionMarket       common.Address
		OptionMarketPricer common.Address
		OptionToken        common.Address
		PoolHedger         common.Address
		ShortCollateral    common.Address
		GwavOracle         common.Address
		QuoteAsset         common.Address
		BaseAsset          common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityPool = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LiquidityToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.GreekCache = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.OptionMarket = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.OptionMarketPricer = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.OptionToken = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.PoolHedger = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.ShortCollateral = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.GwavOracle = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.QuoteAsset = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.BaseAsset = *abi.ConvertType(out[10], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// MarketAddresses is a free data retrieval call binding the contract method 0xc71b7e53.
//
// Solidity: function marketAddresses(address ) view returns(address liquidityPool, address liquidityToken, address greekCache, address optionMarket, address optionMarketPricer, address optionToken, address poolHedger, address shortCollateral, address gwavOracle, address quoteAsset, address baseAsset)
func (_Registry *RegistrySession) MarketAddresses(arg0 common.Address) (struct {
	LiquidityPool      common.Address
	LiquidityToken     common.Address
	GreekCache         common.Address
	OptionMarket       common.Address
	OptionMarketPricer common.Address
	OptionToken        common.Address
	PoolHedger         common.Address
	ShortCollateral    common.Address
	GwavOracle         common.Address
	QuoteAsset         common.Address
	BaseAsset          common.Address
}, error) {
	return _Registry.Contract.MarketAddresses(&_Registry.CallOpts, arg0)
}

// MarketAddresses is a free data retrieval call binding the contract method 0xc71b7e53.
//
// Solidity: function marketAddresses(address ) view returns(address liquidityPool, address liquidityToken, address greekCache, address optionMarket, address optionMarketPricer, address optionToken, address poolHedger, address shortCollateral, address gwavOracle, address quoteAsset, address baseAsset)
func (_Registry *RegistryCallerSession) MarketAddresses(arg0 common.Address) (struct {
	LiquidityPool      common.Address
	LiquidityToken     common.Address
	GreekCache         common.Address
	OptionMarket       common.Address
	OptionMarketPricer common.Address
	OptionToken        common.Address
	PoolHedger         common.Address
	ShortCollateral    common.Address
	GwavOracle         common.Address
	QuoteAsset         common.Address
	BaseAsset          common.Address
}, error) {
	return _Registry.Contract.MarketAddresses(&_Registry.CallOpts, arg0)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Registry *RegistryCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Registry *RegistrySession) NominatedOwner() (common.Address, error) {
	return _Registry.Contract.NominatedOwner(&_Registry.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Registry *RegistryCallerSession) NominatedOwner() (common.Address, error) {
	return _Registry.Contract.NominatedOwner(&_Registry.CallOpts)
}

// OptionMarkets is a free data retrieval call binding the contract method 0x1cb66845.
//
// Solidity: function optionMarkets(uint256 ) view returns(address)
func (_Registry *RegistryCaller) OptionMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "optionMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptionMarkets is a free data retrieval call binding the contract method 0x1cb66845.
//
// Solidity: function optionMarkets(uint256 ) view returns(address)
func (_Registry *RegistrySession) OptionMarkets(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.OptionMarkets(&_Registry.CallOpts, arg0)
}

// OptionMarkets is a free data retrieval call binding the contract method 0x1cb66845.
//
// Solidity: function optionMarkets(uint256 ) view returns(address)
func (_Registry *RegistryCallerSession) OptionMarkets(arg0 *big.Int) (common.Address, error) {
	return _Registry.Contract.OptionMarkets(&_Registry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Registry *RegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Registry *RegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _Registry.Contract.AcceptOwnership(&_Registry.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Registry *RegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Registry.Contract.AcceptOwnership(&_Registry.TransactOpts)
}

// AddMarket is a paid mutator transaction binding the contract method 0x199bc905.
//
// Solidity: function addMarket((address,address,address,address,address,address,address,address,address,address,address) newMarketAddresses) returns()
func (_Registry *RegistryTransactor) AddMarket(opts *bind.TransactOpts, newMarketAddresses LyraRegistryOptionMarketAddresses) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addMarket", newMarketAddresses)
}

// AddMarket is a paid mutator transaction binding the contract method 0x199bc905.
//
// Solidity: function addMarket((address,address,address,address,address,address,address,address,address,address,address) newMarketAddresses) returns()
func (_Registry *RegistrySession) AddMarket(newMarketAddresses LyraRegistryOptionMarketAddresses) (*types.Transaction, error) {
	return _Registry.Contract.AddMarket(&_Registry.TransactOpts, newMarketAddresses)
}

// AddMarket is a paid mutator transaction binding the contract method 0x199bc905.
//
// Solidity: function addMarket((address,address,address,address,address,address,address,address,address,address,address) newMarketAddresses) returns()
func (_Registry *RegistryTransactorSession) AddMarket(newMarketAddresses LyraRegistryOptionMarketAddresses) (*types.Transaction, error) {
	return _Registry.Contract.AddMarket(&_Registry.TransactOpts, newMarketAddresses)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Registry *RegistryTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Registry *RegistrySession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.NominateNewOwner(&_Registry.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Registry *RegistryTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.NominateNewOwner(&_Registry.TransactOpts, _owner)
}

// RemoveMarket is a paid mutator transaction binding the contract method 0xdb913236.
//
// Solidity: function removeMarket(address market) returns()
func (_Registry *RegistryTransactor) RemoveMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "removeMarket", market)
}

// RemoveMarket is a paid mutator transaction binding the contract method 0xdb913236.
//
// Solidity: function removeMarket(address market) returns()
func (_Registry *RegistrySession) RemoveMarket(market common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemoveMarket(&_Registry.TransactOpts, market)
}

// RemoveMarket is a paid mutator transaction binding the contract method 0xdb913236.
//
// Solidity: function removeMarket(address market) returns()
func (_Registry *RegistryTransactorSession) RemoveMarket(market common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemoveMarket(&_Registry.TransactOpts, market)
}

// UpdateGlobalAddresses is a paid mutator transaction binding the contract method 0x30a12dc4.
//
// Solidity: function updateGlobalAddresses(bytes32[] names, address[] addresses) returns()
func (_Registry *RegistryTransactor) UpdateGlobalAddresses(opts *bind.TransactOpts, names [][32]byte, addresses []common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateGlobalAddresses", names, addresses)
}

// UpdateGlobalAddresses is a paid mutator transaction binding the contract method 0x30a12dc4.
//
// Solidity: function updateGlobalAddresses(bytes32[] names, address[] addresses) returns()
func (_Registry *RegistrySession) UpdateGlobalAddresses(names [][32]byte, addresses []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpdateGlobalAddresses(&_Registry.TransactOpts, names, addresses)
}

// UpdateGlobalAddresses is a paid mutator transaction binding the contract method 0x30a12dc4.
//
// Solidity: function updateGlobalAddresses(bytes32[] names, address[] addresses) returns()
func (_Registry *RegistryTransactorSession) UpdateGlobalAddresses(names [][32]byte, addresses []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpdateGlobalAddresses(&_Registry.TransactOpts, names, addresses)
}

// RegistryGlobalAddressUpdatedIterator is returned from FilterGlobalAddressUpdated and is used to iterate over the raw logs and unpacked data for GlobalAddressUpdated events raised by the Registry contract.
type RegistryGlobalAddressUpdatedIterator struct {
	Event *RegistryGlobalAddressUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryGlobalAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryGlobalAddressUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryGlobalAddressUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryGlobalAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryGlobalAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryGlobalAddressUpdated represents a GlobalAddressUpdated event raised by the Registry contract.
type RegistryGlobalAddressUpdated struct {
	Name [32]byte
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGlobalAddressUpdated is a free log retrieval operation binding the contract event 0xdf028fe8b079d8195054d23416d2537debdd0f58695e83cd840d822e6a4c2e80.
//
// Solidity: event GlobalAddressUpdated(bytes32 indexed name, address addr)
func (_Registry *RegistryFilterer) FilterGlobalAddressUpdated(opts *bind.FilterOpts, name [][32]byte) (*RegistryGlobalAddressUpdatedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "GlobalAddressUpdated", nameRule)
	if err != nil {
		return nil, err
	}
	return &RegistryGlobalAddressUpdatedIterator{contract: _Registry.contract, event: "GlobalAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalAddressUpdated is a free log subscription operation binding the contract event 0xdf028fe8b079d8195054d23416d2537debdd0f58695e83cd840d822e6a4c2e80.
//
// Solidity: event GlobalAddressUpdated(bytes32 indexed name, address addr)
func (_Registry *RegistryFilterer) WatchGlobalAddressUpdated(opts *bind.WatchOpts, sink chan<- *RegistryGlobalAddressUpdated, name [][32]byte) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "GlobalAddressUpdated", nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryGlobalAddressUpdated)
				if err := _Registry.contract.UnpackLog(event, "GlobalAddressUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGlobalAddressUpdated is a log parse operation binding the contract event 0xdf028fe8b079d8195054d23416d2537debdd0f58695e83cd840d822e6a4c2e80.
//
// Solidity: event GlobalAddressUpdated(bytes32 indexed name, address addr)
func (_Registry *RegistryFilterer) ParseGlobalAddressUpdated(log types.Log) (*RegistryGlobalAddressUpdated, error) {
	event := new(RegistryGlobalAddressUpdated)
	if err := _Registry.contract.UnpackLog(event, "GlobalAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryMarketRemovedIterator is returned from FilterMarketRemoved and is used to iterate over the raw logs and unpacked data for MarketRemoved events raised by the Registry contract.
type RegistryMarketRemovedIterator struct {
	Event *RegistryMarketRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryMarketRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryMarketRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryMarketRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryMarketRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryMarketRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryMarketRemoved represents a MarketRemoved event raised by the Registry contract.
type RegistryMarketRemoved struct {
	Market common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketRemoved is a free log retrieval operation binding the contract event 0x59d7b1e52008dc342c9421dadfc773114b914a65682a4e4b53cf60a970df0d77.
//
// Solidity: event MarketRemoved(address indexed market)
func (_Registry *RegistryFilterer) FilterMarketRemoved(opts *bind.FilterOpts, market []common.Address) (*RegistryMarketRemovedIterator, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "MarketRemoved", marketRule)
	if err != nil {
		return nil, err
	}
	return &RegistryMarketRemovedIterator{contract: _Registry.contract, event: "MarketRemoved", logs: logs, sub: sub}, nil
}

// WatchMarketRemoved is a free log subscription operation binding the contract event 0x59d7b1e52008dc342c9421dadfc773114b914a65682a4e4b53cf60a970df0d77.
//
// Solidity: event MarketRemoved(address indexed market)
func (_Registry *RegistryFilterer) WatchMarketRemoved(opts *bind.WatchOpts, sink chan<- *RegistryMarketRemoved, market []common.Address) (event.Subscription, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "MarketRemoved", marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryMarketRemoved)
				if err := _Registry.contract.UnpackLog(event, "MarketRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketRemoved is a log parse operation binding the contract event 0x59d7b1e52008dc342c9421dadfc773114b914a65682a4e4b53cf60a970df0d77.
//
// Solidity: event MarketRemoved(address indexed market)
func (_Registry *RegistryFilterer) ParseMarketRemoved(log types.Log) (*RegistryMarketRemoved, error) {
	event := new(RegistryMarketRemoved)
	if err := _Registry.contract.UnpackLog(event, "MarketRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryMarketUpdatedIterator is returned from FilterMarketUpdated and is used to iterate over the raw logs and unpacked data for MarketUpdated events raised by the Registry contract.
type RegistryMarketUpdatedIterator struct {
	Event *RegistryMarketUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryMarketUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryMarketUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryMarketUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryMarketUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryMarketUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryMarketUpdated represents a MarketUpdated event raised by the Registry contract.
type RegistryMarketUpdated struct {
	OptionMarket common.Address
	Market       LyraRegistryOptionMarketAddresses
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMarketUpdated is a free log retrieval operation binding the contract event 0xa090264792b8766dd953a6d1775028cdd591d04a748e045c19b98135970e9127.
//
// Solidity: event MarketUpdated(address indexed optionMarket, (address,address,address,address,address,address,address,address,address,address,address) market)
func (_Registry *RegistryFilterer) FilterMarketUpdated(opts *bind.FilterOpts, optionMarket []common.Address) (*RegistryMarketUpdatedIterator, error) {

	var optionMarketRule []interface{}
	for _, optionMarketItem := range optionMarket {
		optionMarketRule = append(optionMarketRule, optionMarketItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "MarketUpdated", optionMarketRule)
	if err != nil {
		return nil, err
	}
	return &RegistryMarketUpdatedIterator{contract: _Registry.contract, event: "MarketUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketUpdated is a free log subscription operation binding the contract event 0xa090264792b8766dd953a6d1775028cdd591d04a748e045c19b98135970e9127.
//
// Solidity: event MarketUpdated(address indexed optionMarket, (address,address,address,address,address,address,address,address,address,address,address) market)
func (_Registry *RegistryFilterer) WatchMarketUpdated(opts *bind.WatchOpts, sink chan<- *RegistryMarketUpdated, optionMarket []common.Address) (event.Subscription, error) {

	var optionMarketRule []interface{}
	for _, optionMarketItem := range optionMarket {
		optionMarketRule = append(optionMarketRule, optionMarketItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "MarketUpdated", optionMarketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryMarketUpdated)
				if err := _Registry.contract.UnpackLog(event, "MarketUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketUpdated is a log parse operation binding the contract event 0xa090264792b8766dd953a6d1775028cdd591d04a748e045c19b98135970e9127.
//
// Solidity: event MarketUpdated(address indexed optionMarket, (address,address,address,address,address,address,address,address,address,address,address) market)
func (_Registry *RegistryFilterer) ParseMarketUpdated(log types.Log) (*RegistryMarketUpdated, error) {
	event := new(RegistryMarketUpdated)
	if err := _Registry.contract.UnpackLog(event, "MarketUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Registry contract.
type RegistryOwnerChangedIterator struct {
	Event *RegistryOwnerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryOwnerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnerChanged represents a OwnerChanged event raised by the Registry contract.
type RegistryOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Registry *RegistryFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*RegistryOwnerChangedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &RegistryOwnerChangedIterator{contract: _Registry.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Registry *RegistryFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *RegistryOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnerChanged)
				if err := _Registry.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Registry *RegistryFilterer) ParseOwnerChanged(log types.Log) (*RegistryOwnerChanged, error) {
	event := new(RegistryOwnerChanged)
	if err := _Registry.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Registry contract.
type RegistryOwnerNominatedIterator struct {
	Event *RegistryOwnerNominated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnerNominated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryOwnerNominated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnerNominated represents a OwnerNominated event raised by the Registry contract.
type RegistryOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Registry *RegistryFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*RegistryOwnerNominatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &RegistryOwnerNominatedIterator{contract: _Registry.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Registry *RegistryFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *RegistryOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnerNominated)
				if err := _Registry.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Registry *RegistryFilterer) ParseOwnerNominated(log types.Log) (*RegistryOwnerNominated, error) {
	event := new(RegistryOwnerNominated)
	if err := _Registry.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
