package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func RegisterCommand() *cli.Command {
	return &cli.Command{
		Name:  "register",
		Usage: "Register a new operator",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "keystore-path",
				Usage:   "Path to the keystore file",
				EnvVars: []string{"KEYSTORE_PATH"},
			},
		},
		Action: registerOperator,
	}
}

func DeregisterCommand() *cli.Command {
	return &cli.Command{
		Name:   "deregister",
		Usage:  "Deregister an operator",
		Action: deregisterOperator,
	}
}

func registerOperator(c *cli.Context) error {
	fmt.Println(
		`Register a new operator to the TriggerX AVS. It is assumed that the operator has already been registered to the EigenLayer protocol.
If not, please register to EigenLayer first. Follow the instructions here: 
https://github.com/Layr-Labs/eigenlayer-cli/blob/master/README.md`)
	return nil
}

func deregisterOperator(c *cli.Context) error {
	return nil
}

// func (o *Operator) registerOperatorOnStartup(
// 	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
// 	mockTokenStrategyAddr common.Address,
// ) {
// 	err := o.RegisterOperatorWithEigenlayer()
// 	if err != nil {
// 		// This error might only be that the operator was already registered with eigenlayer, so we don't want to fatal
// 		o.logger.Error("Error registering operator with eigenlayer", "err", err)
// 	} else {
// 		o.logger.Infof("Registered operator with eigenlayer")
// 	}

// 	// TODO(samlaf): shouldn't hardcode number here
// 	amount := big.NewInt(1000)
// 	err = o.DepositIntoStrategy(mockTokenStrategyAddr, amount)
// 	if err != nil {
// 		o.logger.Fatal("Error depositing into strategy", "err", err)
// 	}
// 	o.logger.Infof("Deposited %s into strategy %s", amount, mockTokenStrategyAddr)

// 	err = o.RegisterOperatorWithAvs(operatorEcdsaPrivateKey)
// 	if err != nil {
// 		o.logger.Fatal("Error registering operator with avs", "err", err)
// 	}
// 	o.logger.Infof("Registered operator with avs")
// }

// func (o *Operator) RegisterOperatorWithEigenlayer() error {
// 	op := eigenSdkTypes.Operator{
// 		Address:                   o.operatorAddr.String(),
// 		DelegationApproverAddress: o.operatorAddr.String(),
// 	}
// 	_, err := o.eigenlayerWriter.RegisterAsOperator(context.Background(), op, true)
// 	if err != nil {
// 		o.logger.Error("Error registering operator with eigenlayer", "err", err)
// 		return err
// 	}
// 	return nil
// }

// func (o *Operator) DepositIntoStrategy(strategyAddr common.Address, amount *big.Int) error {
// 	_, tokenAddr, err := o.eigenlayerReader.GetStrategyAndUnderlyingToken(&bind.CallOpts{}, strategyAddr)
// 	if err != nil {
// 		o.logger.Error("Failed to fetch strategy contract", "err", err)
// 		return err
// 	}
// 	contractErc20Mock, err := o.avsReader.GetErc20Mock(context.Background(), tokenAddr)
// 	if err != nil {
// 		o.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
// 		return err
// 	}
// 	txOpts, err := o.avsWriter.TxMgr.GetNoSendTxOpts()
// 	tx, err := contractErc20Mock.Mint(txOpts, o.operatorAddr, amount)
// 	if err != nil {
// 		o.logger.Errorf("Error assembling Mint tx")
// 		return err
// 	}
// 	_, err = o.avsWriter.TxMgr.Send(context.Background(), tx, true)
// 	if err != nil {
// 		o.logger.Errorf("Error submitting Mint tx")
// 		return err
// 	}

// 	_, err = o.eigenlayerWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount, true)
// 	if err != nil {
// 		o.logger.Errorf("Error depositing into strategy", "err", err)
// 		return err
// 	}
// 	return nil
// }

// // Registration specific functions
// func (o *Operator) RegisterOperatorWithAvs(
// 	operatorEcdsaKeyPair *ecdsa.PrivateKey,
// ) error {
// 	// hardcode these things for now
// 	quorumNumbers := eigenSdkTypes.QuorumNums{eigenSdkTypes.QuorumNum(0)}
// 	socket := "Not Needed"
// 	operatorToAvsRegistrationSigSalt := [32]byte{123}
// 	curBlockNum, err := o.ethClient.BlockNumber(context.Background())
// 	if err != nil {
// 		o.logger.Errorf("Unable to get current block number")
// 		return err
// 	}
// 	curBlock, err := o.ethClient.HeaderByNumber(context.Background(), big.NewInt(int64(curBlockNum)))
// 	if err != nil {
// 		o.logger.Errorf("Unable to get current block")
// 		return err
// 	}
// 	sigValidForSeconds := int64(1_000_000)
// 	operatorToAvsRegistrationSigExpiry := big.NewInt(int64(curBlock.Time) + sigValidForSeconds)
// 	_, err = o.avsWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
// 		context.Background(),
// 		operatorEcdsaKeyPair, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry,
// 		o.blsKeypair, quorumNumbers, socket, true,
// 	)

// 	if err != nil {
// 		o.logger.Errorf("Unable to register operator with avs registry coordinator")
// 		return err
// 	}
// 	o.logger.Infof("Registered operator with avs registry coordinator.")

// 	return nil
// }

// type OperatorStatus struct {
// 	EcdsaAddress string
// 	// pubkey compendium related
// 	PubkeysRegistered bool
// 	G1Pubkey          string
// 	G2Pubkey          string
// 	// avs related
// 	RegisteredWithAvs bool
// 	OperatorId        string
// }

// func (o *Operator) PrintOperatorStatus() error {
// 	fmt.Println("Printing operator status")
// 	operatorId, err := o.avsReader.GetOperatorId(&bind.CallOpts{}, o.operatorAddr)
// 	if err != nil {
// 		return err
// 	}
// 	pubkeysRegistered := operatorId != [32]byte{}
// 	registeredWithAvs := o.operatorId != [32]byte{}
// 	operatorStatus := OperatorStatus{
// 		EcdsaAddress:      o.operatorAddr.String(),
// 		PubkeysRegistered: pubkeysRegistered,
// 		G1Pubkey:          o.blsKeypair.GetPubKeyG1().String(),
// 		G2Pubkey:          o.blsKeypair.GetPubKeyG2().String(),
// 		RegisteredWithAvs: registeredWithAvs,
// 		OperatorId:        hex.EncodeToString(o.operatorId[:]),
// 	}
// 	operatorStatusJson, err := json.MarshalIndent(operatorStatus, "", " ")
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(string(operatorStatusJson))
// 	return nil
// }

// func pubKeyG1ToBN254G1Point(p *bls.G1Point) regcoord.BN254G1Point {
// 	return regcoord.BN254G1Point{
// 		X: p.X.BigInt(new(big.Int)),
// 		Y: p.Y.BigInt(new(big.Int)),
// 	}
// }

/*
package keeper

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	eigenSdkTypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/trigg3rX/triggerx-keeper/pkg/types"

	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
)

func (o *Keeper) registerOperatorOnStartup(
	operatorEcdsaPrivateKey *ecdsa.PrivateKey,
	mockTokenStrategyAddr common.Address,
) {
	err := o.RegisterOperatorWithEigenlayer()
	if err != nil {
		o.logger.Error("Error registering operator with eigenlayer", "err", err)
	} else {
		o.logger.Infof("Registered operator with eigenlayer")
	}

	amount := big.NewInt(1000)
	err = o.DepositIntoStrategy(mockTokenStrategyAddr, amount)
	if err != nil {
		o.logger.Fatal("Error depositing into strategy", "err", err)
	}
	o.logger.Infof("Deposited %s into strategy %s", amount, mockTokenStrategyAddr)

	err = o.RegisterOperatorWithAvs(operatorEcdsaPrivateKey)
	if err != nil {
		o.logger.Fatal("Error registering operator with avs", "err", err)
	}
	o.logger.Infof("Registered operator with avs")
}

func (o *Keeper) RegisterOperatorWithEigenlayer() error {
	op := eigenSdkTypes.Operator{
		Address:                   o.keeperAddr.String(),
		DelegationApproverAddress: o.keeperAddr.String(),
	}
	_, err := o.eigenlayerWriter.RegisterAsOperator(context.Background(), op, true)
	if err != nil {
		o.logger.Error("Error registering operator with eigenlayer", "err", err)
		return err
	}
	return nil
}

func (o *Keeper) DepositIntoStrategy(strategyAddr common.Address, amount *big.Int) error {
	_, tokenAddr, err := o.eigenlayerReader.GetStrategyAndUnderlyingToken(&bind.CallOpts{}, strategyAddr)
	if err != nil {
		o.logger.Error("Failed to fetch strategy contract", "err", err)
		return err
	}
	contractErc20Mock, err := o.avsReader.GetErc20Mock(context.Background(), tokenAddr)
	if err != nil {
		o.logger.Error("Failed to fetch ERC20Mock contract", "err", err)
		return err
	}
	txOpts, err := o.avsWriter.TxMgr.GetNoSendTxOpts()
	tx, err := contractErc20Mock.Mint(txOpts, o.operatorAddr, amount)
	if err != nil {
		o.logger.Errorf("Error assembling Mint tx")
		return err
	}
	_, err = o.avsWriter.TxMgr.Send(context.Background(), tx, true)
	if err != nil {
		o.logger.Errorf("Error submitting Mint tx")
		return err
	}

	_, err = o.eigenlayerWriter.DepositERC20IntoStrategy(context.Background(), strategyAddr, amount, true)
	if err != nil {
		o.logger.Errorf("Error depositing into strategy", "err", err)
		return err
	}
	return nil
}

// Registration specific functions
func (o *Keeper) RegisterOperatorWithAvs(
	operatorEcdsaKeyPair *ecdsa.PrivateKey,
) error {
	// hardcode these things for now
	quorumNumbers := eigenSdkTypes.QuorumNums{eigenSdkTypes.QuorumNum(0)}
	socket := "Not Needed"
	operatorToAvsRegistrationSigSalt := [32]byte{123}
	curBlockNum, err := o.ethClient.BlockNumber(context.Background())
	if err != nil {
		o.logger.Errorf("Unable to get current block number")
		return err
	}
	curBlock, err := o.ethClient.HeaderByNumber(context.Background(), big.NewInt(int64(curBlockNum)))
	if err != nil {
		o.logger.Errorf("Unable to get current block")
		return err
	}
	sigValidForSeconds := int64(1_000_000)
	operatorToAvsRegistrationSigExpiry := big.NewInt(int64(curBlock.Time) + sigValidForSeconds)
	_, err = o.avsWriter.RegisterOperatorInQuorumWithAVSRegistryCoordinator(
		context.Background(),
		operatorEcdsaKeyPair, operatorToAvsRegistrationSigSalt, operatorToAvsRegistrationSigExpiry,
		o.blsKeypair, quorumNumbers, socket, true,
	)

	if err != nil {
		o.logger.Errorf("Unable to register operator with avs registry coordinator")
		return err
	}
	o.logger.Infof("Registered operator with avs registry coordinator.")

	return nil
}



func (o *Keeper) PrintOperatorStatus() error {
	fmt.Println("Printing operator status")
	operatorId, err := o.avsReader.GetOperatorId(&bind.CallOpts{}, o.keeperAddr)
	if err != nil {
		return err
	}
	pubkeysRegistered := operatorId != [32]byte{}
	registeredWithAvs := o.keeperId != [32]byte{}
	operatorStatus := types.KeeperStatus{
		EcdsaAddress:      o.keeperAddr.String(),
		PubkeysRegistered: pubkeysRegistered,
		G1Pubkey:          o.blsKeypair.GetPubKeyG1().String(),
		G2Pubkey:          o.blsKeypair.GetPubKeyG2().String(),
		RegisteredWithAvs: registeredWithAvs,
		KeeperId:          hex.EncodeToString(o.keeperId[:]),
	}
	operatorStatusJson, err := json.MarshalIndent(operatorStatus, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(operatorStatusJson))
	return nil
}

func pubKeyG1ToBN254G1Point(p *bls.G1Point) regcoord.BN254G1Point {
	return regcoord.BN254G1Point{
		X: p.X.BigInt(new(big.Int)),
		Y: p.Y.BigInt(new(big.Int)),
	}
}

*/
