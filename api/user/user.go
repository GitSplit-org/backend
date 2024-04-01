package user

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"net/http"
	"os"

	"github.com/GitSplit-org/backend/generated/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ApplyRoutes(r *gin.RouterGroup) {
	api := r.Group("/user")
	{
		api.POST("/assign", AssignWalletToUsername)
	}
}

func AssignWalletToUsername(c *gin.Context) {
	var request request
	err := c.BindJSON(&request)
	if err != nil {
		logrus.Error("payload invalid")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	nodeUrl := os.Getenv("POLYGON_MUMBAI_RPC_HTTP")
	client, err := ethclient.Dial(nodeUrl)
	if err != nil {
		logrus.Errorf("failed to dial client at url %v, error: %v", nodeUrl, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "bad request",
		})
		return
	}
	privateKey, err := crypto.HexToECDSA(os.Getenv("WALLET_PRIVATE_KEY"))
	if err != nil {
		logrus.Errorf("failed to parse privatekey: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to parse privatekey",
		})
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logrus.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot assert type: publicKey is not of type *ecdsa.PublicKey",
		})

	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logrus.Errorf("cannot generate nonce")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot generate nonce",
		})
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logrus.Errorf("cannot fetch gas price")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot fetch gas price",
		})
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(80001))
	if err != nil {
		logrus.Errorf("cannot get TransactOpts")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot get TransactOpts",
		})
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := contract.NewContract(address, client)
	if err != nil {
		logrus.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "bad request",
		})
	}
	tx, err := instance.AssignAddressToUsername(auth, request.Username, common.HexToAddress(request.WalletAddress))
	if err != nil {
		logrus.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "bad request",
		})
	}
	logrus.Info("transaction send", tx.Hash().String())
	c.JSON(http.StatusOK, gin.H{"data": tx.Hash().String()})
}
