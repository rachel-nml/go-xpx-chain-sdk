// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package transactions

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ReplicatorOnboardingTransactionBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsReplicatorOnboardingTransactionBuffer(buf []byte, offset flatbuffers.UOffsetT) *ReplicatorOnboardingTransactionBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ReplicatorOnboardingTransactionBuffer{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsReplicatorOnboardingTransactionBuffer(buf []byte, offset flatbuffers.UOffsetT) *ReplicatorOnboardingTransactionBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ReplicatorOnboardingTransactionBuffer{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ReplicatorOnboardingTransactionBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ReplicatorOnboardingTransactionBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ReplicatorOnboardingTransactionBuffer) Size() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) MutateSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *ReplicatorOnboardingTransactionBuffer) Signature(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) SignatureLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) SignatureBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ReplicatorOnboardingTransactionBuffer) MutateSignature(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *ReplicatorOnboardingTransactionBuffer) Signer(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) SignerLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) SignerBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ReplicatorOnboardingTransactionBuffer) MutateSigner(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *ReplicatorOnboardingTransactionBuffer) Version() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) MutateVersion(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *ReplicatorOnboardingTransactionBuffer) Type() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) MutateType(n uint16) bool {
	return rcv._tab.MutateUint16Slot(12, n)
}

func (rcv *ReplicatorOnboardingTransactionBuffer) MaxFee(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) MaxFeeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) MutateMaxFee(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *ReplicatorOnboardingTransactionBuffer) Deadline(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) DeadlineLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) MutateDeadline(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *ReplicatorOnboardingTransactionBuffer) Capacity(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) CapacityLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) MutateCapacity(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *ReplicatorOnboardingTransactionBuffer) BlsPublicKey(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) BlsPublicKeyLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ReplicatorOnboardingTransactionBuffer) BlsPublicKeyBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ReplicatorOnboardingTransactionBuffer) MutateBlsPublicKey(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func ReplicatorOnboardingTransactionBufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
}
func ReplicatorOnboardingTransactionBufferAddSize(builder *flatbuffers.Builder, size uint32) {
	builder.PrependUint32Slot(0, size, 0)
}
func ReplicatorOnboardingTransactionBufferAddSignature(builder *flatbuffers.Builder, signature flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(signature), 0)
}
func ReplicatorOnboardingTransactionBufferStartSignatureVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ReplicatorOnboardingTransactionBufferAddSigner(builder *flatbuffers.Builder, signer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(signer), 0)
}
func ReplicatorOnboardingTransactionBufferStartSignerVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ReplicatorOnboardingTransactionBufferAddVersion(builder *flatbuffers.Builder, version uint32) {
	builder.PrependUint32Slot(3, version, 0)
}
func ReplicatorOnboardingTransactionBufferAddType(builder *flatbuffers.Builder, type_ uint16) {
	builder.PrependUint16Slot(4, type_, 0)
}
func ReplicatorOnboardingTransactionBufferAddMaxFee(builder *flatbuffers.Builder, maxFee flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(maxFee), 0)
}
func ReplicatorOnboardingTransactionBufferStartMaxFeeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ReplicatorOnboardingTransactionBufferAddDeadline(builder *flatbuffers.Builder, deadline flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(deadline), 0)
}
func ReplicatorOnboardingTransactionBufferStartDeadlineVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ReplicatorOnboardingTransactionBufferAddCapacity(builder *flatbuffers.Builder, capacity flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(capacity), 0)
}
func ReplicatorOnboardingTransactionBufferStartCapacityVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ReplicatorOnboardingTransactionBufferAddBlsPublicKey(builder *flatbuffers.Builder, blsPublicKey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(blsPublicKey), 0)
}
func ReplicatorOnboardingTransactionBufferStartBlsPublicKeyVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ReplicatorOnboardingTransactionBufferEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type PrepareBcDriveTransactionBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsPrepareBcDriveTransactionBuffer(buf []byte, offset flatbuffers.UOffsetT) *PrepareBcDriveTransactionBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PrepareBcDriveTransactionBuffer{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsPrepareBcDriveTransactionBuffer(buf []byte, offset flatbuffers.UOffsetT) *PrepareBcDriveTransactionBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &PrepareBcDriveTransactionBuffer{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *PrepareBcDriveTransactionBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PrepareBcDriveTransactionBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PrepareBcDriveTransactionBuffer) Size() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) MutateSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *PrepareBcDriveTransactionBuffer) Signature(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) SignatureLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) SignatureBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *PrepareBcDriveTransactionBuffer) MutateSignature(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *PrepareBcDriveTransactionBuffer) Signer(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) SignerLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) SignerBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *PrepareBcDriveTransactionBuffer) MutateSigner(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *PrepareBcDriveTransactionBuffer) Version() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) MutateVersion(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *PrepareBcDriveTransactionBuffer) Type() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) MutateType(n uint16) bool {
	return rcv._tab.MutateUint16Slot(12, n)
}

func (rcv *PrepareBcDriveTransactionBuffer) MaxFee(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) MaxFeeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) MutateMaxFee(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *PrepareBcDriveTransactionBuffer) Deadline(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) DeadlineLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) MutateDeadline(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *PrepareBcDriveTransactionBuffer) DriveSize(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) DriveSizeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) MutateDriveSize(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *PrepareBcDriveTransactionBuffer) VerificationFeeAmount(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) VerificationFeeAmountLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) MutateVerificationFeeAmount(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *PrepareBcDriveTransactionBuffer) ReplicatorCount() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PrepareBcDriveTransactionBuffer) MutateReplicatorCount(n uint16) bool {
	return rcv._tab.MutateUint16Slot(22, n)
}

func PrepareBcDriveTransactionBufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func PrepareBcDriveTransactionBufferAddSize(builder *flatbuffers.Builder, size uint32) {
	builder.PrependUint32Slot(0, size, 0)
}
func PrepareBcDriveTransactionBufferAddSignature(builder *flatbuffers.Builder, signature flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(signature), 0)
}
func PrepareBcDriveTransactionBufferStartSignatureVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func PrepareBcDriveTransactionBufferAddSigner(builder *flatbuffers.Builder, signer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(signer), 0)
}
func PrepareBcDriveTransactionBufferStartSignerVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func PrepareBcDriveTransactionBufferAddVersion(builder *flatbuffers.Builder, version uint32) {
	builder.PrependUint32Slot(3, version, 0)
}
func PrepareBcDriveTransactionBufferAddType(builder *flatbuffers.Builder, type_ uint16) {
	builder.PrependUint16Slot(4, type_, 0)
}
func PrepareBcDriveTransactionBufferAddMaxFee(builder *flatbuffers.Builder, maxFee flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(maxFee), 0)
}
func PrepareBcDriveTransactionBufferStartMaxFeeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PrepareBcDriveTransactionBufferAddDeadline(builder *flatbuffers.Builder, deadline flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(deadline), 0)
}
func PrepareBcDriveTransactionBufferStartDeadlineVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PrepareBcDriveTransactionBufferAddDriveSize(builder *flatbuffers.Builder, driveSize flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(driveSize), 0)
}
func PrepareBcDriveTransactionBufferStartDriveSizeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PrepareBcDriveTransactionBufferAddVerificationFeeAmount(builder *flatbuffers.Builder, verificationFeeAmount flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(verificationFeeAmount), 0)
}
func PrepareBcDriveTransactionBufferStartVerificationFeeAmountVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func PrepareBcDriveTransactionBufferAddReplicatorCount(builder *flatbuffers.Builder, replicatorCount uint16) {
	builder.PrependUint16Slot(9, replicatorCount, 0)
}
func PrepareBcDriveTransactionBufferEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type DriveClosureTransactionBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsDriveClosureTransactionBuffer(buf []byte, offset flatbuffers.UOffsetT) *DriveClosureTransactionBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DriveClosureTransactionBuffer{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsDriveClosureTransactionBuffer(buf []byte, offset flatbuffers.UOffsetT) *DriveClosureTransactionBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &DriveClosureTransactionBuffer{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *DriveClosureTransactionBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DriveClosureTransactionBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *DriveClosureTransactionBuffer) Size() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) MutateSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *DriveClosureTransactionBuffer) Signature(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) SignatureLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) SignatureBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DriveClosureTransactionBuffer) MutateSignature(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *DriveClosureTransactionBuffer) Signer(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) SignerLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) SignerBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DriveClosureTransactionBuffer) MutateSigner(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *DriveClosureTransactionBuffer) Version() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) MutateVersion(n uint32) bool {
	return rcv._tab.MutateUint32Slot(10, n)
}

func (rcv *DriveClosureTransactionBuffer) Type() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) MutateType(n uint16) bool {
	return rcv._tab.MutateUint16Slot(12, n)
}

func (rcv *DriveClosureTransactionBuffer) MaxFee(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) MaxFeeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) MutateMaxFee(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *DriveClosureTransactionBuffer) Deadline(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) DeadlineLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) MutateDeadline(j int, n uint32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateUint32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func (rcv *DriveClosureTransactionBuffer) DriveKey(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) DriveKeyLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DriveClosureTransactionBuffer) DriveKeyBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DriveClosureTransactionBuffer) MutateDriveKey(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func DriveClosureTransactionBufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(8)
}
func DriveClosureTransactionBufferAddSize(builder *flatbuffers.Builder, size uint32) {
	builder.PrependUint32Slot(0, size, 0)
}
func DriveClosureTransactionBufferAddSignature(builder *flatbuffers.Builder, signature flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(signature), 0)
}
func DriveClosureTransactionBufferStartSignatureVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func DriveClosureTransactionBufferAddSigner(builder *flatbuffers.Builder, signer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(signer), 0)
}
func DriveClosureTransactionBufferStartSignerVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func DriveClosureTransactionBufferAddVersion(builder *flatbuffers.Builder, version uint32) {
	builder.PrependUint32Slot(3, version, 0)
}
func DriveClosureTransactionBufferAddType(builder *flatbuffers.Builder, type_ uint16) {
	builder.PrependUint16Slot(4, type_, 0)
}
func DriveClosureTransactionBufferAddMaxFee(builder *flatbuffers.Builder, maxFee flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(maxFee), 0)
}
func DriveClosureTransactionBufferStartMaxFeeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DriveClosureTransactionBufferAddDeadline(builder *flatbuffers.Builder, deadline flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(deadline), 0)
}
func DriveClosureTransactionBufferStartDeadlineVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func DriveClosureTransactionBufferAddDriveKey(builder *flatbuffers.Builder, driveKey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(driveKey), 0)
}
func DriveClosureTransactionBufferStartDriveKeyVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func DriveClosureTransactionBufferEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
