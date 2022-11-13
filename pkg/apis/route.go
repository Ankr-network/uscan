package apis

import (
	"github.com/Ankr-network/uscan/pkg/response"
	"github.com/Ankr-network/uscan/pkg/service"
	"github.com/Ankr-network/uscan/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func SetupRouter(g fiber.Router) {
	g.Get("/search", search)
	g.Get("/home", getHome)

	g.Get("/blocks", listBlocks)
	g.Get("/blocks/:blockNum", getBlock)

	g.Get("/txs", listTxs)
	g.Get("/txs/:txHash", getTx)
	//g.Get("/internal-txs", getInternalTxs)
	//g.Get("/txs/:txHash/internal", getInternalTx)

	g.Get("/txs/:txHash/base", getTxBase)
	//g.Get("/txs/:txHash/event-logs", getTxEventLogs) // 合并到/txs/:txHas

	//g.Get("/txs/:txHash/tracetx", getTraceTx)
	//g.Get("/txs/:txHash/tracetx2", getTraceTx2)

	//g.Get("/accounts", listAccounts)
	g.Get("/accounts/:address", getAccountInfo)
	g.Get("/accounts/:address/txns", getAccountTxns)
	//g.Get("/accounts/:address/txns/download", downloadAccountTxns)
	g.Get("/accounts/:address/txns-erc20", getAccountErc20Txns)
	//g.Get("/accounts/:address/txns-erc20/download", downloadAccountErc20Txns)
	g.Get("/accounts/:address/txns-erc721", getAccountErc721Txns)
	//g.Get("/accounts/:address/txns-erc721/download", downloadAccountErc721Txns)
	g.Get("/accounts/:address/txns-erc1155", getAccountErc1155Txns)
	//g.Get("/accounts/:address/txns-erc1155/download", downloadAccountErc1155Txns)
	g.Get("/accounts/:address/txns-internal", getAccountInternalTxns)
	//g.Get("/accounts/:address/txns-internal/download", downloadAccountInternalTxns)
	//g.Get("/tokens/txns/erc20", listTokenTxnsErc20)
	//g.Get("/tokens/txns/erc721", listTokenTxnsErc721)
	//g.Get("/tokens/txns/erc1155", listTokenTxnsErc1155)
	//
	//g.Get("/tokens/:address/type", getTokenType)
	//g.Get("/tokens/:address/transfers", listTokenTransfers)
	//g.Get("/tokens/:address/transfers/download", downloadTokenTransfers)
	//g.Get("/tokens/:address/holders", listTokenHolders)
	//g.Get("/tokens/:address/inventory", listInventory)
	//g.Get("/nfts/:address/:tokenID", getNft)
	//g.Get("/contracts/:address/verify", validateContract)
	//g.Get("/contracts-verify/:id/status", getValidateContractStatus)
	//g.Get("/contracts/metadata", getValidateContractMetadata)
	//g.Get("/contracts/:address/content", getValidateContract)
}

func search(c *fiber.Ctx) error {
	f := &types.SearchFilter{}
	if err := c.QueryParser(&f); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.Err(response.ErrInvalidParameter))
	}
	resp, err := service.Search(f)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err))
	}
	return c.Status(http.StatusOK).JSON(response.Ok(resp))
}

func getHome(c *fiber.Ctx) error {
	resp, err := service.Home()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err))
	}
	return c.Status(http.StatusOK).JSON(response.Ok(resp))
}

func listBlocks(c *fiber.Ctx) error {
	f := &types.Pager{}
	if err := c.QueryParser(&f); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.Err(response.ErrInvalidParameter))
	}
	f.Complete()
	resp, total, err := service.ListFullFieldBlocks(f)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err))
	}
	return c.Status(http.StatusOK).JSON(response.Ok(map[string]interface{}{"items": resp, "total": total}))
}

func getBlock(c *fiber.Ctx) error {
	blockNum := c.Params("blockNum")
	block, err := service.GetBlock(blockNum)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err))
	}
	return c.Status(http.StatusOK).JSON(response.Ok(block))
}

func listTxs(c *fiber.Ctx) error {
	f := &types.Pager{}
	if err := c.QueryParser(&f); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.Err(response.ErrInvalidParameter))
	}
	f.Complete()
	resp, total, err := service.ListTxs(f)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err))
	}
	return c.Status(http.StatusOK).JSON(response.Ok(map[string]interface{}{"items": resp, "total": total}))
}

func getInternalTxs(c *fiber.Ctx) error {
	//return c.Status(http.StatusBadRequest).JSON(response.Err(err))
	return c.Status(http.StatusOK).JSON(response.Ok(nil))
}

func getTx(c *fiber.Ctx) error {
	txHash := c.Params("txHash")
	if txHash == "" {
		return c.Status(http.StatusBadRequest).JSON(response.ErrInvalidParameter)
	}
	resp, err := service.GetTx(common.HexToHash(txHash).Hex())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err))
	}
	return c.Status(http.StatusOK).JSON(response.Ok(resp))
}

func getInternalTx(c *fiber.Ctx) error {
	//return c.Status(http.StatusBadRequest).JSON(response.Err(err))
	return c.Status(http.StatusOK).JSON(response.Ok(nil))
}

func getTxBase(c *fiber.Ctx) error {
	txHash := c.Params("txHash")
	if txHash == "" {
		return c.Status(http.StatusBadRequest).JSON(response.ErrInvalidParameter)
	}
	resp, err := service.GetTx(common.HexToHash(txHash).Hex())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err))
	}
	return c.Status(http.StatusOK).JSON(response.Ok(resp))
}

func listAccounts(c *fiber.Ctx) error {
	f := &types.Pager{}
	if err := c.QueryParser(&f); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.Err(response.ErrInvalidParameter))
	}
	f.Complete()
	resp, total, err := service.ListAccount(f)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err))
	}
	return c.Status(http.StatusOK).JSON(response.Ok(map[string]interface{}{"items": resp, "total": total}))
}

func getAccountInfo(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(http.StatusBadRequest).JSON(response.Err(response.ErrInvalidParameter))
	}
	resp, err := service.GetAccountInfo(address)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err))
	}
	return c.Status(http.StatusOK).JSON(response.Ok(resp))
}

func getAccountTxns(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(http.StatusBadRequest).JSON(response.Err(response.ErrInvalidParameter))
	}
	f := &types.Pager{}
	if err := c.QueryParser(&f); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.Err(response.ErrInvalidParameter))
	}
	resp, err := service.GetAccountTxs(f, address)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err))
	}
	return c.Status(http.StatusOK).JSON(response.Ok(resp))
}

func getAccountErc20Txns(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(http.StatusBadRequest).JSON(response.Err(response.ErrInvalidParameter))
	}
	f := &types.Pager{}
	if err := c.QueryParser(&f); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.Err(response.ErrInvalidParameter))
	}
	resp, err := service.GetAccountErc20Txns(f, address)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.Err(err))
	}
	return c.Status(http.StatusOK).JSON(response.Ok(resp))
}
