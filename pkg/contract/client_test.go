package contract

import (
	"math/big"
	"strconv"
	"strings"
	"testing"

	"github.com/Ankr-network/uscan/pkg/rpcclient"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/stretchr/testify/assert"
)

var (
	testClient            = NewClient(rpcclient.NewRpcClient([]string{"wss://testnet.ankr.com/ws"}))
	testContract20        = "0x6a92f2e354228e866c44419860233cc23bec0d8a"
	testContract721       = "0xB502432eD49b7c1AD5Cdca7C133F4334DD09e8cd"
	testContract721Token  = "0x1348c63"
	testContract1155      = "0xD9ED70B30Fe75978093c0Cd7da422Fb24BC27b60"
	testContract1155Token = "0x1348c63"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestContract_isEip(t *testing.T) {
	testClient.eip1155Cache.Add(testContract1155, true)
	testClient.eip721Cache.Add(testContract1155, true)
	value, ok := testClient.eip721Cache.Get(testContract1155)
	assert.True(t, ok)
	assert.True(t, value.(bool))

	value, ok = testClient.eip1155Cache.Get(testContract1155)
	assert.True(t, ok)
	assert.True(t, value.(bool))

	testClient.eip721Cache.Remove(testContract1155)
	testClient.eip1155Cache.Remove(testContract1155)
}

func TestContract_data(t *testing.T) {
	var byte4 [4]byte
	copy(byte4[:], hexutil.MustDecode(eip721Inf)[0:4])
	data, err := testClient.eip1155Abi.Pack("supportsInterface", byte4)
	assert.NoError(t, err)
	t.Log(hexutil.Bytes(data).String())
}

func TestContract_IsEIP721(t *testing.T) {
	assert.True(t, testClient.IsEIP721(testContract721))
	assert.False(t, testClient.IsEIP721(testContract1155))
}

func TestContract_Bytes(t *testing.T) {
	t.Log(hexutil.MustDecode(eip721Inf))
}

func TestContract_IsEIP1155(t *testing.T) {
	assert.False(t, testClient.IsEIP1155(testContract721))
	assert.True(t, testClient.IsEIP1155(testContract1155))
}

func TestContract_hashToAddr(t *testing.T) {
	hash := common.HexToHash("0x00000000000000000000000066acaf662822dbfd66a5240788ec528b743885e6")
	t.Log(common.BytesToAddress(hash[:]).Hex())

	nonHash := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
	assert.True(t, common.BytesToAddress(nonHash[:]) == common.Address{})
}

func TestContract_bigInt(t *testing.T) {
	bi1 := big.NewInt(110)
	bi2 := big.NewInt(22)
	t.Log(bi1.Sub(bi1, bi2).Cmp(big.NewInt(0))) // 1:>;    0:=;     -1:<
}

func TestContact_TokenID(t *testing.T) {
	tokenID, err := hexutil.DecodeBig("0x3a04186db96039c6497fbb15ce84cf1aa568978fe1803cb8e5cd3a9dfce62561")
	assert.NoError(t, err)
	t.Log(tokenID.String())

	data, err := testClient.eip721Abi.Pack("ownerOf", tokenID)
	assert.NoError(t, err)
	t.Log(hexutil.Bytes(data).String())

	bi, ok := big.NewInt(0).SetString("100034502830371529892874507663703497903", 10)
	assert.True(t, ok)

	data, err = testClient.eip721Abi.Pack("ownerOf", bi)
	assert.NoError(t, err)
	t.Log(hexutil.Bytes(data).String())

	id, _ := hexutil.DecodeBig("0x11e68")
	res, err := testClient.eip721MetaAbi.Pack("tokenURI", id)
	assert.NoError(t, err)
	t.Log(hexutil.Bytes(res).String())
}

func TestContract_Pack(t *testing.T) {
	// arg := abi.NonIndexed()
	id, _ := hexutil.DecodeBig("0x11e68")
	t.Log(common.Bytes2Hex(math.U256Bytes(id)))

	t.Log(common.Bytes2Hex(id.Bytes()))

}

func TestContract_Meta721(t *testing.T) {
	uri, err := testClient.GetEIP721Meta(testContract721, testContract721Token)
	assert.NoError(t, err)
	t.Log(uri)
}

func TestContract_Meta1155(t *testing.T) {
	uri, err := testClient.GetEIP1155Meta(testContract1155, testContract1155Token)
	assert.NoError(t, err)
	t.Log(uri)
}

func TestGetContactName(t *testing.T) {
	name, err := testClient.GetContractName(testContract721)
	assert.NoError(t, err)
	t.Log(name)
}

func TestGetContactSymbol(t *testing.T) {
	symbol, err := testClient.GetContractSymbol(testContract721)
	assert.NoError(t, err)
	t.Log(symbol)
}

func TestGetContactDecimals(t *testing.T) {
	decimals, err := testClient.GetContractDecimals(testContract20)
	assert.NoError(t, err)
	t.Log(decimals)
}

func TestGetContractTotalSupply(t *testing.T) {
	totalSupply, err := testClient.GetContractTotalSupply(testContract20)
	assert.NoError(t, err)
	t.Log(totalSupply)
}

func TestErc20Transfer(t *testing.T) {
	blockNumber, _ := strconv.ParseInt(strings.TrimPrefix("0xa2aec6", "0x"), 16, 64)
	txIndex, _ := strconv.ParseInt(strings.TrimPrefix("0x16", "0x"), 16, 64)
	index, _ := strconv.ParseInt(strings.TrimPrefix("0x53", "0x"), 16, 64)
	log := types.Log{
		Address: common.HexToAddress("0x77b0faf252b1c7d28728af93ef28f73783c47ed2"),
		Topics: []common.Hash{
			common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
			common.HexToHash("0x0000000000000000000000002ffa266bc99566999fcf7a19941ab6426aa25d61"),
		},
		Data:        hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000012888a1cb"),
		BlockNumber: uint64(blockNumber),
		TxHash:      common.HexToHash("0x0df31368a16cb7e31f6b450cce62e6a7f36c58d7012afa58376bcdd5a6b7a9e2"),
		TxIndex:     uint(txIndex),
		BlockHash:   common.HexToHash("0xc42511b7b3f32c979e6542bfa6c8ea359249e889022c5f12122dde979677367d"),
		Index:       uint(index),
		Removed:     false,
	}
	res, err := testClient.Erc20Transfer("0x8f3ed97839f0606101ef94f3db110fba9b649c07", log)
	assert.NoError(t, err)
	t.Log(res.To)
	t.Log(res.From)
	t.Log(res.Value)
}

func TestErc721Transfer(t *testing.T) {
	blockNumber, _ := strconv.ParseInt(strings.TrimLeft("0x936d32", "0x"), 16, 64)
	txIndex, _ := strconv.ParseInt(strings.TrimLeft("0xa", "0x"), 16, 64)
	index, _ := strconv.ParseInt(strings.TrimLeft("0x18", "0x"), 16, 64)
	log := types.Log{
		Address: common.HexToAddress("0x4ca28dfb93a5f0c137f39561b8dcf93db485870d"),
		Topics: []common.Hash{
			common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
			common.HexToHash("0x0000000000000000000000002ffa266bc99566999fcf7a19941ab6426aa25d61"),
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000547"),
		},
		Data:        hexutil.MustDecode("0x"),
		BlockNumber: uint64(blockNumber),
		TxHash:      common.HexToHash("0xce7af7f3d4675cb8e9029abf4562a292c882eb2ae6d08cf380425edeede7bece"),
		TxIndex:     uint(txIndex),
		BlockHash:   common.HexToHash("0x47234edaad6efa9fd251adc779226575eafb8b34aa621ddf884af08d75597e9c"),
		Index:       uint(index),
		Removed:     false,
	}
	res, err := testClient.Erc721Transfer("0xef7d805cda3bccb3516d582f6ef9554d3e1edee5", log)
	assert.NoError(t, err)
	t.Log(res.To)
	t.Log(res.From)
	t.Log(res.TokenId)
}

func TestErc1155TransferSingle(t *testing.T) {
	blockNumber, _ := strconv.ParseInt(strings.TrimLeft("0x9d646e", "0x"), 16, 64)
	txIndex, _ := strconv.ParseInt(strings.TrimLeft("0x9", "0x"), 16, 64)
	index, _ := strconv.ParseInt(strings.TrimLeft("0xf", "0x"), 16, 64)
	log := types.Log{
		Address: common.HexToAddress("0xcf83c9def5e880574860f0f47872a18f48b191af"),
		Topics: []common.Hash{
			common.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"),
			common.HexToHash("0x0000000000000000000000006a82fa5cf82ef724d1f1955fecc27ddd0758132e"),
			common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
			common.HexToHash("0x0000000000000000000000006a82fa5cf82ef724d1f1955fecc27ddd0758132e"),
		},
		Data:        hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000005"),
		BlockNumber: uint64(blockNumber),
		TxHash:      common.HexToHash("0x97ad07c77e75e77a24e9fb726e7d1787f7a9f7c3810600bfbeefeeb64ab48e7b"),
		TxIndex:     uint(txIndex),
		BlockHash:   common.HexToHash("0x88d6c961d6a6b6443cf94ee704e31b34665b46066ef39df203240a4d140fbc0f"),
		Index:       uint(index),
		Removed:     false,
	}
	res, err := testClient.Erc1155TransferSingle("0xCF83c9deF5E880574860F0F47872A18F48b191aF", log)
	assert.NoError(t, err)
	t.Log(res.To)
	t.Log(res.From)
	t.Log(res.Id)
	t.Log(res.Value)
}

func TestErc1155TransferBatch(t *testing.T) {
	blockNumber, _ := strconv.ParseInt(strings.TrimLeft("0xa332e1", "0x"), 16, 64)
	txIndex, _ := strconv.ParseInt(strings.TrimLeft("0xf", "0x"), 16, 64)
	index, _ := strconv.ParseInt(strings.TrimLeft("0x12", "0x"), 16, 64)
	log := types.Log{
		Address: common.HexToAddress("0x359b72bb7302855f28c01ba8d0ee72b5701b8b03"),
		Topics: []common.Hash{
			common.HexToHash("0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"),
			common.HexToHash("0x000000000000000000000000f6564501ad364fd20cb8853b7fa0aaffcd469734"),
			common.HexToHash("0x000000000000000000000000b37a5ba4060d6bfd00a3bfcb235bb596f13932bd"),
			common.HexToHash("0x000000000000000000000000401d24076331c077099dea08fe898ea7ff7254cd"),
		},
		Data:        hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000090000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001"),
		BlockNumber: uint64(blockNumber),
		TxHash:      common.HexToHash("0x297d8dc0257e280b231aadf38622752a565d6f1d73a29a198da111892e345284"),
		TxIndex:     uint(txIndex),
		BlockHash:   common.HexToHash("0x3c0a37fcacff410e60f53d74beb87970f1099cc59eb34ce375dd0b0bad8bbd9b"),
		Index:       uint(index),
		Removed:     false,
	}
	res, err := testClient.Erc1155TransferBatch("0x359b72bb7302855f28c01ba8d0ee72b5701b8b03", log)
	assert.NoError(t, err)
	t.Log(res.To)
	t.Log(res.From)
	t.Log(res.Ids)
	t.Log(res.Values)
	t.Log(res.Operator)
}

/*
eip721  TransferEventTopic
https://rinkeby.etherscan.io/tx/0x29b622f0b541121356428f173256ac6d0826ab48bcd546503ad1532061f1940c#eventlog




eip1155 TransferSingleEventTopic
https://rinkeby.etherscan.io/tx/0x97ad07c77e75e77a24e9fb726e7d1787f7a9f7c3810600bfbeefeeb64ab48e7b#eventlog

TransferBatchEventTopic
https://bscscan.com/tx/0x2dcfe8650d6deee75dbd55d748a5d5b604a25edfe98588e1a096bfd365281a8e#eventlog
*/
