package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/danial.idham/benz/pkg/repo"
	"github.com/danial.idham/benz/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"gorm.io/gorm"
)

type DataController struct {
	Database        *gorm.DB
	ContractAddress string
	RPCURL          string
}

type FormData struct {
	NRIC          string `json:"nric"`
	WalletAddress string `json:"wallet_address"`
}

const (
	queryReceipt = "queryReceipt(address)"
	returnType   = "string"
)

// NewDataController initializes a new data controller with given DB and configurations
func NewDataController(db *gorm.DB, contractAddress, rpcURL string) *DataController {
	return &DataController{
		Database:        db,
		ContractAddress: contractAddress,
		RPCURL:          rpcURL,
	}
}

// Create handles the creation of new data
func (dc *DataController) Create(ctx *gin.Context) {
	var data FormData
	if err := ctx.Bind(&data); err != nil {
		sendBadRequest(ctx, err)
		return
	}

	if err := dc.createProfile(data.NRIC, data.WalletAddress); err != nil {
		sendBadRequest(ctx, err)
		return
	}

	receipt, err := dc.queryReceiptFromContract(data.WalletAddress)
	if err != nil {
		sendBadRequest(ctx, err)
		return
	}

	hash := sha256.Sum256([]byte(receipt))
	ctx.JSON(http.StatusOK, gin.H{"receipt": hex.EncodeToString(hash[:])})
}

func (dc *DataController) createProfile(nric, walletAddress string) error {
	profile := &repo.Profile{
		NRIC:          nric,
		WalletAddress: walletAddress,
	}
	return repo.CreateProfile(dc.Database, profile)
}

func (dc *DataController) queryReceiptFromContract(walletAddress string) (string, error) {
	client, err := w3.Dial(dc.RPCURL)
	if err != nil {
		return "", err
	}
	defer client.Close()

	funcQueryReceipt, err := w3.NewFunc(queryReceipt, returnType)
	if err != nil {
		return "", err
	}
	var receipt string
	err = client.Call(
		eth.CallFunc(funcQueryReceipt, w3.A(dc.ContractAddress), w3.A(walletAddress)).Returns(&receipt),
	)
	return receipt, err
}

func sendBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, response.NewError(http.StatusBadRequest, err.Error()))
}
