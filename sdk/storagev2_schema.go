// Copyright 2021 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

func prepareBcDriveTransactionSchema() *schema {
	return &schema{
		[]schemaAttribute{
			newScalarAttribute("size", IntSize),
			newArrayAttribute("signature", ByteSize),
			newArrayAttribute("signer", ByteSize),
			newScalarAttribute("version", IntSize),
			newScalarAttribute("type", ShortSize),
			newArrayAttribute("maxFee", IntSize),
			newArrayAttribute("deadline", IntSize),
			newArrayAttribute("owner", ByteSize),
			newArrayAttribute("driveKey", ByteSize),
			newArrayAttribute("driveSize", IntSize),
			newScalarAttribute("replicatorCount", ShortSize),
		},
	}
}

func driveClosureTransactionSchema() *schema {
	return &schema{
		[]schemaAttribute{
			newScalarAttribute("size", IntSize),
			newArrayAttribute("signature", ByteSize),
			newArrayAttribute("signer", ByteSize),
			newScalarAttribute("version", IntSize),
			newScalarAttribute("type", ShortSize),
			newArrayAttribute("maxFee", IntSize),
			newArrayAttribute("deadline", IntSize),
			newArrayAttribute("driveKey", ByteSize),
		},
	}
}
