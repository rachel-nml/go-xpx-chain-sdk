package integration

import (
	"fmt"
	"github.com/proximax-storage/go-xpx-chain-sdk/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSuperContractFlowTransaction(t *testing.T) {
	driveAccount, err := client.NewAccount()
	assert.Nil(t, err)
	fmt.Println(driveAccount)

	replicatorAccount, err := client.NewAccount()
	assert.Nil(t, err)
	fmt.Println(replicatorAccount)

	var storageSize uint64 = 10000
	var billingPrice uint64 = 50
	var billingPeriod = 10

	driveTx, err := client.NewPrepareDriveTransaction(
		sdk.NewDeadline(time.Hour),
		defaultAccount.PublicAccount,
		sdk.Duration(billingPeriod),
		sdk.Duration(billingPeriod),
		sdk.Amount(billingPrice),
		sdk.StorageSize(storageSize),
		1,
		1,
		1,
	);
	driveTx.ToAggregate(driveAccount.PublicAccount)
	assert.Nil(t, err)

	transferStorageToReplicator, err := client.NewTransferTransaction(
		sdk.NewDeadline(time.Hour),
		replicatorAccount.Address,
		[]*sdk.Mosaic{sdk.Storage(storageSize)},
		sdk.NewPlainMessage(""),
	);
	transferStorageToReplicator.ToAggregate(defaultAccount.PublicAccount)
	assert.Nil(t, err)

	transferXpxToReplicator, err := client.NewTransferTransaction(
		sdk.NewDeadline(time.Hour),
		driveAccount.Address,
		[]*sdk.Mosaic{sdk.Xpx(10000000)},
		sdk.NewPlainMessage(""),
	);
	transferXpxToReplicator.ToAggregate(defaultAccount.PublicAccount)
	assert.Nil(t, err)

	result := sendTransaction(t, func() (sdk.Transaction, error) {
		return client.NewCompleteAggregateTransaction(
			sdk.NewDeadline(time.Hour),
			[]sdk.Transaction{driveTx, transferStorageToReplicator, transferXpxToReplicator},
		)
	}, defaultAccount, driveAccount)
	assert.Nil(t, result.error)

	result = sendTransaction(t, func() (sdk.Transaction, error) {
		return client.NewJoinToDriveTransaction(
			sdk.NewDeadline(time.Hour),
			driveAccount.PublicAccount,
		)
	}, replicatorAccount)
	assert.Nil(t, result.error)

	var fileSize uint64 = 147
	fileHash, err := sdk.StringToHash("AA2D2427E105A9B60DF634553849135DF629F1408A018D02B07A70CAFFB43093")
	assert.Nil(t, err)

	fsTx, err := client.NewDriveFileSystemTransaction(
		sdk.NewDeadline(time.Hour),
		driveAccount.PublicAccount,
		&sdk.Hash{1},
		&sdk.Hash{},
		[]*sdk.Action{
			{
				FileHash: fileHash,
				FileSize: sdk.StorageSize(fileSize),
			},
		},
		[]*sdk.Action{},
	)
	fsTx.ToAggregate(defaultAccount.PublicAccount)
	assert.Nil(t,err)

	transferStreamingToReplicator, err := client.NewTransferTransaction(
		sdk.NewDeadline(time.Hour),
		replicatorAccount.Address,
		[]*sdk.Mosaic{sdk.Streaming(fileSize)},
		sdk.NewPlainMessage(""),
	);
	transferStreamingToReplicator.ToAggregate(defaultAccount.PublicAccount)
	assert.Nil(t, err)

	result = sendTransaction(t, func() (sdk.Transaction, error) {
		return client.NewCompleteAggregateTransaction(
			sdk.NewDeadline(time.Hour),
			[]sdk.Transaction{fsTx, transferStreamingToReplicator},
		)
	}, defaultAccount)
	assert.Nil(t, result.error)

	superContract, err := client.NewAccount()
	assert.Nil(t, err)
	deploy, err := client.NewDeployTransaction(
		sdk.NewDeadline(time.Hour),
		driveAccount.PublicAccount,
		defaultAccount.PublicAccount,
		fileHash,
		0,
	);
	deploy.ToAggregate(superContract.PublicAccount)
	assert.Nil(t, err)

	result = sendTransaction(t, func() (sdk.Transaction, error) {
		return client.NewCompleteAggregateTransaction(
			sdk.NewDeadline(time.Hour),
			[]sdk.Transaction{transferStreamingToReplicator, deploy},
		)
	}, defaultAccount, superContract)
	assert.Nil(t, result.error)

	initiator, err := client.NewAccount()
	assert.Nil(t, err)
	transferSCToInitiator, err := client.NewTransferTransaction(
		sdk.NewDeadline(time.Hour),
		initiator.Address,
		[]*sdk.Mosaic{sdk.SuperContractMosaic(1000)},
		sdk.NewPlainMessage(""),
	);
	transferSCToInitiator.ToAggregate(defaultAccount.PublicAccount)

	assert.Nil(t, err)
	execute, err := client.NewStartExecuteTransaction(
		sdk.NewDeadline(time.Hour),
		superContract.PublicAccount,
		[]*sdk.Mosaic{sdk.SuperContractMosaic(1000)},
		"GoGoGo",
		[]int64{
			123,
			228,
		},
	);
	execute.ToAggregate(initiator.PublicAccount)
	assert.Nil(t, err)

	result = sendTransaction(t, func() (sdk.Transaction, error) {
		return client.NewCompleteAggregateTransaction(
			sdk.NewDeadline(time.Hour),
			[]sdk.Transaction{transferSCToInitiator, execute},
		)
	}, defaultAccount, initiator, replicatorAccount)
}
