// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Fresher struct {
	_tab flatbuffers.Table
}

func GetRootAsFresher(buf []byte, offset flatbuffers.UOffsetT) *Fresher {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Fresher{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Fresher) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Fresher) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Fresher) U() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Fresher) MutateU(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func FresherStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func FresherAddU(builder *flatbuffers.Builder, u uint64) {
	builder.PrependUint64Slot(0, u, 0)
}
func FresherEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
