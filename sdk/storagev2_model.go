// Copyright 2021 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
)

type DataModificationState uint8

const (
	Succeeded DataModificationState = iota
	Cancelled
)

type ActiveDataModification struct {
	Id              *Hash
	Owner           *PublicAccount
	DownloadDataCdi *Hash
	UploadSize      StorageSize
}

func (desc *ActiveDataModification) String() string {
	return fmt.Sprintf(
		`
			"Id": %s,
			"Owner": %s,
			"DownloadDataCdi": %s,
			"UploadSize": %d,
		`,
		desc.Id,
		desc.Owner,
		desc.DownloadDataCdi,
		desc.UploadSize,
	)
}

type CompletedDataModification struct {
	ActiveDataModification *ActiveDataModification
	State                  DataModificationState
}

func (desc *CompletedDataModification) String() string {
	return fmt.Sprintf(
		`
			"ActiveDataModification": %s,
			"State:" %d,
		`,
		desc.ActiveDataModification,
		desc.State,
	)
}

type BcDrive struct {
	BcDriveAccount             *PublicAccount
	OwnerAccount               *PublicAccount
	RootHash                   *Hash
	DriveSize                  StorageSize
	ReplicatorCount            uint16
	ActiveDataModifications    []*ActiveDataModification
	CompletedDataModifications []*CompletedDataModification
	UsedSizeMap                map[string]StorageSize
	Replicators                []*PublicAccount
}

func (drive *BcDrive) String() string {
	return fmt.Sprintf(
		`
		"BcDriveAccount": %s,
		"OwnerAccount": %s,
		"RootHash": %s,
		"DriveSize": %d,
		"ReplicatorCount": %d,
		"ActiveDataModifications": %s,
		"CompletedDataModifications": %s,
		"UsedSizeMap": %+v,
		"Replicators": %+v,
		`,
		drive.BcDriveAccount,
		drive.OwnerAccount,
		drive.RootHash,
		drive.DriveSize,
		drive.ReplicatorCount,
		drive.ActiveDataModifications,
		drive.CompletedDataModifications,
		drive.UsedSizeMap,
		drive.Replicators,
	)
}

type Replicator struct {
	Capacity Amount
	DriveKey []*PublicAccount
}

func (replicator *Replicator) String() string {
	return fmt.Sprintf(
		`
		Capacity: %d,
		DriveKey: %s,
		`,
		replicator.Capacity,
		replicator.DriveKey,
	)
}

// Replicator Onboarding Transaction
type ReplicatorOnboardingTransaction struct {
	AbstractTransaction
	Capacity     Amount
	BlsPublicKey *BLSPublicKey
}

// Prepare Bc Drive Transaction
type PrepareBcDriveTransaction struct {
	AbstractTransaction
	DriveSize             StorageSize
	VerificationFeeAmount Amount
	ReplicatorCount       uint16
}

// Drive Closure Transaction
type DriveClosureTransaction struct {
	AbstractTransaction
	DriveKey *PublicAccount
}
