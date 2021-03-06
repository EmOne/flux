// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbast

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BooleanLiteral struct {
	_tab flatbuffers.Table
}

func GetRootAsBooleanLiteral(buf []byte, offset flatbuffers.UOffsetT) *BooleanLiteral {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BooleanLiteral{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *BooleanLiteral) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BooleanLiteral) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BooleanLiteral) BaseNode(obj *BaseNode) *BaseNode {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BaseNode)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *BooleanLiteral) Value() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *BooleanLiteral) MutateValue(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func BooleanLiteralStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func BooleanLiteralAddBaseNode(builder *flatbuffers.Builder, baseNode flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(baseNode), 0)
}
func BooleanLiteralAddValue(builder *flatbuffers.Builder, value bool) {
	builder.PrependBoolSlot(1, value, false)
}
func BooleanLiteralEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
