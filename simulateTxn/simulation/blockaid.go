package simulation

import (
	"context"
	"fmt"

	"github.com/blockaid-official/blockaid-client-go"
	"github.com/blockaid-official/blockaid-client-go/option"
	"github.com/blockaid-official/blockaid-client-go/shared"
)

type blockaidSimulationService struct {
	client *blockaidclientgo.Client
}

func NewBlockaidSimulationService(apiKey string) SimulationService {
	return blockaidSimulationService{
		client: blockaidclientgo.NewClient(
			option.WithAPIKey(apiKey),
		),
	}
}

func (bs blockaidSimulationService) SimulateTransaction(ctx context.Context, callData string, fromAddr string, toAddr string, value string, block string) (string, error) {
	metadata := blockaidclientgo.F(blockaidclientgo.MetadataParam{
		Domain: blockaidclientgo.F("ORIGIN"),
	})
	params := blockaidclientgo.EvmJsonRpcScanParams{
		Chain: blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChain(blockaidclientgo.TransactionScanSupportedChainImmutableZkevm)),
		Block: blockaidclientgo.F[blockaidclientgo.EvmJsonRpcScanParamsBlockUnion](shared.UnionString(block)),
		Data: blockaidclientgo.F(blockaidclientgo.EvmJsonRpcScanParamsData{
			Method: blockaidclientgo.F("eth_sendTransaction"),
			Params: blockaidclientgo.F([]interface{}{
				map[string]interface{}{
					"data":  callData,
					"to":    toAddr,
					"from":  fromAddr,
					"value": value,
				},
			}),
		}),
		Metadata:       metadata,
		AccountAddress: blockaidclientgo.F(fromAddr),
		Options:        blockaidclientgo.F([]blockaidclientgo.EvmJsonRpcScanParamsOption{blockaidclientgo.EvmJsonRpcScanParamsOptionSimulation, blockaidclientgo.EvmJsonRpcScanParamsOptionValidation}),
	}

	transactionScanResponse, err := bs.client.Evm.JsonRpc.Scan(
		ctx,
		params,
	)
	if err != nil {
		return "", err
	}

	simulation := transactionScanResponse.Simulation.JSON.RawJSON()

	fmt.Printf("%+v\n", simulation)
	fmt.Print("\n", transactionScanResponse.Block)
	fmt.Printf("%+v\n", transactionScanResponse.Validation.JSON.RawJSON())

	return simulation, nil
}

func (bs blockaidSimulationService) SimulateMessage(ctx context.Context, msg string, account string, block string) (string, error) {
	transactionScanResponse, err := bs.client.Evm.JsonRpc.Scan(context.TODO(), blockaidclientgo.EvmJsonRpcScanParams{
		Chain: blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChain(blockaidclientgo.TransactionScanSupportedChainEthereum)),
		Block: blockaidclientgo.F[blockaidclientgo.EvmJsonRpcScanParamsBlockUnion](shared.UnionString(block)),
		Data: blockaidclientgo.F(blockaidclientgo.EvmJsonRpcScanParamsData{
			Method: blockaidclientgo.F("eth_signTypedData_v4"),
			Params: blockaidclientgo.F([]interface{}{account, msg}),
		}),
		Metadata: blockaidclientgo.F(blockaidclientgo.MetadataParam{
			Domain: blockaidclientgo.F("https://boredapeyartclub.com"),
		}),
		Options: blockaidclientgo.F([]blockaidclientgo.EvmJsonRpcScanParamsOption{blockaidclientgo.EvmJsonRpcScanParamsOptionSimulation, blockaidclientgo.EvmJsonRpcScanParamsOptionValidation}),
	})
	if err != nil {
		return "", err
	}

	simulation := transactionScanResponse.JSON.RawJSON()
	fmt.Printf("%+v\n", simulation)
	return simulation, nil
}

func (bs blockaidSimulationService) SimulateBulkTransactions(ctx context.Context, callData string, fromAddr string, toAddr string, value string) (string, error) {
	metadata := blockaidclientgo.F(blockaidclientgo.MetadataParam{
		Domain: blockaidclientgo.F("https://boredapeyartclub.com"),
	})
	params := blockaidclientgo.EvmTransactionBulkScanParams{
		Chain: blockaidclientgo.F(blockaidclientgo.TransactionScanSupportedChain(blockaidclientgo.TransactionScanSupportedChainEthereum)),
		Data: blockaidclientgo.F([]blockaidclientgo.EvmTransactionBulkScanParamsData{
			{
				From:     blockaidclientgo.F(fromAddr),
				Data:     blockaidclientgo.F(callData),
				Gas:      blockaidclientgo.F("0x0"),
				GasPrice: blockaidclientgo.F("0x0"),
				To:       blockaidclientgo.F(toAddr),
				Value:    blockaidclientgo.F(value),
			},
			{
				From:     blockaidclientgo.F(fromAddr),
				Data:     blockaidclientgo.F(callData),
				Gas:      blockaidclientgo.F("0x0"),
				GasPrice: blockaidclientgo.F("0x0"),
				To:       blockaidclientgo.F(toAddr),
				Value:    blockaidclientgo.F(value),
			},
		}),
		Metadata: metadata,
		Options:  blockaidclientgo.F([]blockaidclientgo.EvmTransactionBulkScanParamsOption{blockaidclientgo.EvmTransactionBulkScanParamsOptionSimulation}),
	}

	transactionScanResponse, err := bs.client.Evm.TransactionBulk.Scan(
		ctx,
		params,
	)
	if err != nil {
		return "", err
	}

	var simulationRes string

	for _, txn := range *transactionScanResponse {
		fmt.Printf("%+v\n", txn.JSON.RawJSON())
		simulationRes += txn.JSON.RawJSON()
	}

	return simulationRes, nil
}
