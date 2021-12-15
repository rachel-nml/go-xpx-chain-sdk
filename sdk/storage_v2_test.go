// Copyright 2021 ProximaX Limited. All rights reserved.
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.
package sdk

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/proximax-storage/go-xpx-utils/mock"
	"github.com/stretchr/testify/assert"
)

const (
	testBcDriveInfoJson = `{
    "bcdrive": {
      "multisig": "415C7C61822B063F62A4876A6F6BA2DAAE114AB298D7AC7FC56FDBA95872C309",
      "multisigAddress": "9048760066A50F0F65820D3008A79CF73E1034A564BF44AB3E",
      "owner": "CFC31B3080B36BC3D59DF4AB936AC72F4DC15CE3C3E1B1EC5EA41415A4C33FEE",
      "rootHash": "0100000000000000000000000000000000000000000000000000000000000000",
      "size": [
        1000,
        0
      ],
      "usedSize": [
        0,
        0
      ],
      "metaFilesSize": [
        20,
        0
      ],
      "replicatorCount": 5,
      "activeDataModifications": [
        {
          "id": "0100000000000000000000000000000000000000000000000000000000000000",
          "owner": "CFC31B3080B36BC3D59DF4AB936AC72F4DC15CE3C3E1B1EC5EA41415A4C33FEE",
          "downloadDataCdi": "0100000000000000000000000000000000000000000000000000000000000000",
          "uploadSize": [
            100,
            0
          ]
        }
      ],
      "completedDataModifications": [
        {
          "activeDataModifications": [
            {
              "id": "0100000000000000000000000000000000000000000000000000000000000000",
              "owner": "CFC31B3080B36BC3D59DF4AB936AC72F4DC15CE3C3E1B1EC5EA41415A4C33FEE",
              "downloadDataCdi": "0100000000000000000000000000000000000000000000000000000000000000",
              "uploadSize": [
                100,
                0
              ]
            }
          ],
          "state": 0
        }
      ]
    }
}`

	testBcDriveInfoJsonArr = "[" + testBcDriveInfoJson + ", " + testBcDriveInfoJson + "]"
)

const (
	testReplicatorInfoJson = `{
        "replicator": {
            "key": "36E7F50C8B8BC9A4FC6325B2359E0E5DB50C75A914B5292AD726FD5AE3992691",
            "version": 1,
            "capacity": [
                1000,
                0
            ],
            "blsKey": "B49D90CFC2BF81E908B305DAA7066473E0A8980746B881CA0681D8F04765DEAC60AD9E0100CAA90C5836764DCCCE6552",
            "drives": [
                {
                    "drive": "415C7C61822B063F62A4876A6F6BA2DAAE114AB298D7AC7FC56FDBA95872C309",
                    "lastApprovedDataModificationId": "0100000000000000000000000000000000000000000000000000000000000000",
                    "dataModificationIdIsValid": false,
                    "initialDownloadWork": 0
                }
            ]
        }
    }`

	testReplicatorInfoJsonArr = "[" + testReplicatorInfoJson + ", " + testReplicatorInfoJson + "]"
)

var testBcDriveAccount, _ = NewAccountFromPublicKey("415C7C61822B063F62A4876A6F6BA2DAAE114AB298D7AC7FC56FDBA95872C309", PublicTest)
var testBcDriveOwnerAccount, _ = NewAccountFromPublicKey("CFC31B3080B36BC3D59DF4AB936AC72F4DC15CE3C3E1B1EC5EA41415A4C33FEE", PublicTest)
var testReplicatorV2Account, _ = NewAccountFromPublicKey("36E7F50C8B8BC9A4FC6325B2359E0E5DB50C75A914B5292AD726FD5AE3992691", PublicTest)
var testBlsKey = "B49D90CFC2BF81E908B305DAA7066473E0A8980746B881CA0681D8F04765DEAC60AD9E0100CAA90C5836764DCCCE6552"

var (
	testBcDriveInfo = &BcDrive{
		BcDriveAccount:  testBcDriveAccount,
		OwnerAccount:    testBcDriveOwnerAccount,
		RootHash:        &Hash{1},
		DriveSize:       StorageSize(1000),
		UsedSize:        StorageSize(0),
		MetaFilesSize:   StorageSize(20),
		ReplicatorCount: 5,
		ActiveDataModifications: []*ActiveDataModification{
			{
				Id:              &Hash{1},
				Owner:           testBcDriveOwnerAccount,
				DownloadDataCdi: &Hash{1},
				UploadSize:      StorageSize(100),
			},
		},
		CompletedDataModifications: []*CompletedDataModification{
			{
				ActiveDataModification: []*ActiveDataModification{
					{
						Id:              &Hash{1},
						Owner:           testBcDriveOwnerAccount,
						DownloadDataCdi: &Hash{1},
						UploadSize:      StorageSize(100),
					},
				},
				State: DataModificationState(Succeeded),
			},
		},
	}

	testReplicatorInfo = &Replicator{
		ReplicatorAccount: testReplicatorV2Account,
		Version:           1,
		Capacity:          StorageSize(1000),
		BLSKey:            testBlsKey,
		Drives: map[string]*DriveInfo{
			testBcDriveAccount.PublicKey: {
				LastApprovedDataModificationId: &Hash{1},
				DataModificationIdIsValid:      false,
				InitialDownloadWork:            0,
				Index:                          0,
			},
		},
	}
)

var (
	testBcDrivesPage = &BcDrivesPage{
		BcDrives: []*BcDrive{testBcDriveInfo, testBcDriveInfo},
	}
	testReplicatorsPage = &ReplicatorsPage{
		Replicators: []*Replicator{testReplicatorInfo, testReplicatorInfo},
	}
)

func TestStorageV2Service_GetDrive(t *testing.T) {
	mock := newSdkMockWithRouter(&mock.Router{
		Path:                fmt.Sprintf(driveRouteV2, testBcDriveAccount.PublicKey),
		AcceptedHttpMethods: []string{http.MethodGet},
		RespHttpCode:        200,
		RespBody:            testBcDriveInfoJson,
	})
	exchangeClient := mock.getPublicTestClientUnsafe().StorageV2

	defer mock.Close()

	bcdrive, err := exchangeClient.GetDrive(ctx, testBcDriveAccount)
	assert.Nil(t, err)
	assert.NotNil(t, bcdrive)
	assert.Equal(t, testBcDriveInfo, bcdrive)
}

func TestStorageV2Service_GetDrives(t *testing.T) {
	mock := newSdkMockWithRouter(&mock.Router{
		Path:                drivesRouteV2,
		AcceptedHttpMethods: []string{http.MethodGet},
		RespHttpCode:        200,
		RespBody:            `{ "data":` + testBcDriveInfoJsonArr + `}`,
	})
	exchangeClient := mock.getPublicTestClientUnsafe().StorageV2

	defer mock.Close()

	bcdrives, err := exchangeClient.GetDrives(ctx, nil)
	assert.Nil(t, err)
	assert.NotNil(t, bcdrives)
	assert.Equal(t, testBcDrivesPage, bcdrives)
}

func TestStorageV2Service_GetAccountDrives(t *testing.T) {
	mock := newSdkMockWithRouter(&mock.Router{
		Path:                fmt.Sprintf(drivesOfAccountRouteV2, testBcDriveOwnerAccount.PublicKey),
		AcceptedHttpMethods: []string{http.MethodGet},
		RespHttpCode:        200,
		RespBody:            testBcDriveInfoJsonArr,
	})
	exchangeClient := mock.getPublicTestClientUnsafe().StorageV2

	defer mock.Close()

	bcdrives, err := exchangeClient.GetAccountDrives(ctx, testBcDriveOwnerAccount)
	assert.Nil(t, err)
	assert.NotNil(t, bcdrives)
	assert.Equal(t, len(bcdrives), 2)
	assert.Equal(t, []*BcDrive{testBcDriveInfo, testBcDriveInfo}, bcdrives)
}

func TestStorageV2Service_GetReplicator(t *testing.T) {
	mock := newSdkMockWithRouter(&mock.Router{
		Path:                fmt.Sprintf(replicatorRouteV2, testReplicatorV2Account.PublicKey),
		AcceptedHttpMethods: []string{http.MethodGet},
		RespHttpCode:        200,
		RespBody:            testReplicatorInfoJson,
	})
	exchangeClient := mock.getPublicTestClientUnsafe().StorageV2

	defer mock.Close()

	replicator, err := exchangeClient.GetReplicator(ctx, testReplicatorV2Account)
	assert.Nil(t, err)
	assert.NotNil(t, replicator)
	assert.Equal(t, testReplicatorInfo, replicator)
}

func TestStorageV2Service_GetReplicators(t *testing.T) {
	mock := newSdkMockWithRouter(&mock.Router{
		Path:                replicatorsRouteV2,
		AcceptedHttpMethods: []string{http.MethodGet},
		RespHttpCode:        200,
		RespBody:            `{ "data":` + testReplicatorInfoJsonArr + `}`,
	})
	exchangeClient := mock.getPublicTestClientUnsafe().StorageV2

	defer mock.Close()

	replicators, err := exchangeClient.GetReplicators(ctx, nil)
	assert.Nil(t, err)
	assert.NotNil(t, replicators)
	assert.Equal(t, testReplicatorsPage, replicators)
}

func TestStorageV2Service_GetAccountReplicators(t *testing.T) {
	mock := newSdkMockWithRouter(&mock.Router{
		Path:                fmt.Sprintf(replicatorsOfAccountRouteV2, testReplicatorV2Account.PublicKey),
		AcceptedHttpMethods: []string{http.MethodGet},
		RespHttpCode:        200,
		RespBody:            testReplicatorInfoJsonArr,
	})
	exchangeClient := mock.getPublicTestClientUnsafe().StorageV2

	defer mock.Close()

	replicators, err := exchangeClient.GetAccountReplicators(ctx, testReplicatorV2Account)
	assert.Nil(t, err)
	assert.NotNil(t, replicators)
	assert.Equal(t, len(replicators), 2)
	assert.Equal(t, []*Replicator{testReplicatorInfo, testReplicatorInfo}, replicators)
}