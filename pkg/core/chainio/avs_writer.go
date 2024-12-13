package chainio

// import (
// 	"context"
// 	"math/big"

// 	sdkcommon "github.com/trigg3rX/triggerx-keeper/pkg/common"
// 	gethcommon "github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"

// 	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
// 	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
// 	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
// 	logging "github.com/Layr-Labs/eigensdk-go/logging"
// 	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

// 	txtaskmanager "github.com/trigg3rX/go-backend/pkg/avsinterface/bindings/TriggerXTaskManager"
// 	"github.com/trigg3rX/triggerx-keeper/pkg/core/config"
// )

// type AvsWriterer interface {
// 	//avsregistry.ChainWriter

// 	SendNewTaskNumberToSquare(
// 		ctx context.Context,
// 		numToSquare *big.Int,
// 		quorumThresholdPercentage sdktypes.QuorumThresholdPercentage,
// 		quorumNumbers sdktypes.QuorumNums,
// 	) (txtaskmanager.ITriggerXTaskManagerTask, uint32, error)
// 	RaiseChallenge(
// 		ctx context.Context,
// 		task txtaskmanager.ITriggerXTaskManagerTask,
// 		taskResponse txtaskmanager.ITriggerXTaskManagerTaskResponse,
// 		// taskResponseMetadata txtaskmanager.ITriggerXTaskManagerTaskResponseMetadata,
// 		pubkeysOfNonSigningOperators []txtaskmanager.BN254G1Point,
// 	) (*types.Receipt, error)
// 	SendAggregatedResponse(ctx context.Context,
// 		task txtaskmanager.ITriggerXTaskManagerTask,
// 		taskResponse txtaskmanager.ITriggerXTaskManagerTaskResponse,
// 		nonSignerStakesAndSignature txtaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
// 	) (*types.Receipt, error)
// }

// type AvsWriter struct {
// 	avsregistry.ChainWriter
// 	AvsContractBindings *AvsManagersBindings
// 	logger              logging.Logger
// 	TxMgr               txmgr.TxManager
// 	client              eth.HttpBackend
// }

// var _ AvsWriterer = (*AvsWriter)(nil)

// func BuildAvsWriterFromConfig(c *config.Config) (*AvsWriter, error) {
// 	return BuildAvsWriter(c.TxMgr, c.TriggerXServiceManagerAddr, c.OperatorStateRetrieverAddr, &c.EthHttpClient, c.Logger)
// }

// func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethHttpClient sdkcommon.EthClientInterface, logger logging.Logger) (*AvsWriter, error) {
// 	avsServiceBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
// 	if err != nil {
// 		logger.Error("Failed to create contract bindings", "err", err)
// 		return nil, err
// 	}
// 	avsRegistryWriter, err := avsregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return NewAvsWriter(*avsRegistryWriter, avsServiceBindings, logger, txMgr), nil
// }
// func NewAvsWriter(avsRegistryWriter avsregistry.ChainWriter, avsServiceBindings *AvsManagersBindings, logger logging.Logger, txMgr txmgr.TxManager) *AvsWriter {
// 	return &AvsWriter{
// 		ChainWriter:         avsRegistryWriter,
// 		AvsContractBindings: avsServiceBindings,
// 		logger:              logger,
// 		TxMgr:               txMgr,
// 	}
// }

// // returns the tx receipt, as well as the task index (which it gets from parsing the tx receipt logs)
// func (w *AvsWriter) SendNewTaskNumberToSquare(ctx context.Context, numToSquare *big.Int, quorumThresholdPercentage sdktypes.QuorumThresholdPercentage, quorumNumbers sdktypes.QuorumNums) (cstaskmanager.IIncredibleSquaringTaskManagerTask, uint32, error) {
// 	txOpts, err := w.TxMgr.GetNoSendTxOpts()
// 	if err != nil {
// 		w.logger.Errorf("Error getting tx opts")
// 		return txtaskmanager.ITriggerXTaskManagerTask{}, 0, err
// 	}
// 	tx, err := w.AvsContractBindings.TaskManager.CreateNewTask(txOpts, jobId, uint32(quorumThresholdPercentage), quorumNumbers.UnderlyingType())
// 	if err != nil {
// 		w.logger.Errorf("Error assembling CreateNewTask tx")
// 		return txtaskmanager.ITriggerXTaskManagerTask{}, 0, err
// 	}
// 	receipt, err := w.TxMgr.Send(ctx, tx, true)
// 	if err != nil {
// 		w.logger.Errorf("Error submitting CreateNewTask tx")
// 		return txtaskmanager.ITriggerXTaskManagerTask{}, 0, err
// 	}
// 	taskCreatedEvent, err := w.AvsContractBindings.TaskManager.ContractTriggerXTaskManagerFilterer.ParseTaskCreated(*receipt.Logs[0])
// 	if err != nil {
// 		w.logger.Error("Aggregator failed to parse new task created event", "err", err)
// 		return txtaskmanager.ITriggerXTaskManagerTask{}, 0, err
// 	}
// 	return taskCreatedEvent.TaskId, taskCreatedEvent.TaskResponseHash, nil
// }

// func (w *AvsWriter) SendAggregatedResponse(
// 	ctx context.Context, task cstaskmanager.IIncredibleSquaringTaskManagerTask,
// 	taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
// 	nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
// ) (*types.Receipt, error) {
// 	txOpts, err := w.TxMgr.GetNoSendTxOpts()
// 	if err != nil {
// 		w.logger.Errorf("Error getting tx opts")
// 		return nil, err
// 	}
// 	tx, err := w.AvsContractBindings.TaskManager.RespondToTask(txOpts, task, taskResponse, nonSignerStakesAndSignature)
// 	if err != nil {
// 		w.logger.Error("Error submitting SubmitTaskResponse tx while calling respondToTask", "err", err)
// 		return nil, err
// 	}
// 	receipt, err := w.TxMgr.Send(ctx, tx, true)
// 	if err != nil {
// 		w.logger.Errorf("Error submitting respondToTask tx")
// 		return nil, err
// 	}
// 	return receipt, nil
// }

// func (w *AvsWriter) RaiseChallenge(
// 	ctx context.Context,
// 	task cstaskmanager.IIncredibleSquaringTaskManagerTask,
// 	taskResponse cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse,
// 	taskResponseMetadata cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata,
// 	pubkeysOfNonSigningOperators []cstaskmanager.BN254G1Point,
// ) (*types.Receipt, error) {
// 	txOpts, err := w.TxMgr.GetNoSendTxOpts()
// 	if err != nil {
// 		w.logger.Errorf("Error getting tx opts")
// 		return nil, err
// 	}
// 	tx, err := w.AvsContractBindings.TaskManager.RaiseAndResolveChallenge(txOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
// 	if err != nil {
// 		w.logger.Errorf("Error assembling RaiseChallenge tx")
// 		return nil, err
// 	}
// 	receipt, err := w.TxMgr.Send(ctx, tx, true)
// 	if err != nil {
// 		w.logger.Errorf("Error submitting RaiseChallenge tx")
// 		return nil, err
// 	}
// 	return receipt, nil
// }
