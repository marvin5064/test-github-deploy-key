// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internalrpc/risk_manager.proto

package internalrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import api "github.com/bitgaming/go-protobuf-schema/sportsbook/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Namespace int32

const (
	Namespace_UNDEFINED Namespace = 0
	Namespace_CLOUDBET  Namespace = 1
)

var Namespace_name = map[int32]string{
	0: "UNDEFINED",
	1: "CLOUDBET",
}
var Namespace_value = map[string]int32{
	"UNDEFINED": 0,
	"CLOUDBET":  1,
}

func (x Namespace) String() string {
	return proto.EnumName(Namespace_name, int32(x))
}
func (Namespace) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

type OverlayedType int32

const (
	OverlayedType_UNDEFINED_OVERLAYED_TYPE OverlayedType = 0
	OverlayedType_PLAYER                   OverlayedType = 1
	OverlayedType_AGENT                    OverlayedType = 2
)

var OverlayedType_name = map[int32]string{
	0: "UNDEFINED_OVERLAYED_TYPE",
	1: "PLAYER",
	2: "AGENT",
}
var OverlayedType_value = map[string]int32{
	"UNDEFINED_OVERLAYED_TYPE": 0,
	"PLAYER":                   1,
	"AGENT":                    2,
}

func (x OverlayedType) String() string {
	return proto.EnumName(OverlayedType_name, int32(x))
}
func (OverlayedType) EnumDescriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

type Filter struct {
	// empty -> no entity filter
	//      entity_type cant be undefined
	//      entity_id = 0 -> only filter by entity_type
	SportEntity *api.Entity `protobuf:"bytes,1,opt,name=sport_entity,json=sportEntity" json:"sport_entity,omitempty"`
	//  0 -> no filter
	MarketId uint32 `protobuf:"varint,2,opt,name=market_id,json=marketId" json:"market_id,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *Filter) GetSportEntity() *api.Entity {
	if m != nil {
		return m.SportEntity
	}
	return nil
}

func (m *Filter) GetMarketId() uint32 {
	if m != nil {
		return m.MarketId
	}
	return 0
}

type OverlayedEntity struct {
	// id: can be player id or sub agent id
	Id   uint32        `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type OverlayedType `protobuf:"varint,2,opt,name=type,enum=internalrpc.OverlayedType" json:"type,omitempty"`
}

func (m *OverlayedEntity) Reset()                    { *m = OverlayedEntity{} }
func (m *OverlayedEntity) String() string            { return proto.CompactTextString(m) }
func (*OverlayedEntity) ProtoMessage()               {}
func (*OverlayedEntity) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *OverlayedEntity) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OverlayedEntity) GetType() OverlayedType {
	if m != nil {
		return m.Type
	}
	return OverlayedType_UNDEFINED_OVERLAYED_TYPE
}

type BettingOverlay struct {
	AgentId     uint32      `protobuf:"varint,2,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	SportEntity *api.Entity `protobuf:"bytes,3,opt,name=sport_entity,json=sportEntity" json:"sport_entity,omitempty"`
	Overlay     *Overlay    `protobuf:"bytes,4,opt,name=overlay" json:"overlay,omitempty"`
	// betting overlay takes effect to
	// can be empty, which take effect to all direct children
	OverlayedEntity  *OverlayedEntity `protobuf:"bytes,7,opt,name=overlayed_entity,json=overlayedEntity" json:"overlayed_entity,omitempty"`
	Namespace        Namespace        `protobuf:"varint,5,opt,name=namespace,enum=internalrpc.Namespace" json:"namespace,omitempty"`
	BettingOverlayId uint32           `protobuf:"varint,6,opt,name=betting_overlay_id,json=bettingOverlayId" json:"betting_overlay_id,omitempty"`
}

func (m *BettingOverlay) Reset()                    { *m = BettingOverlay{} }
func (m *BettingOverlay) String() string            { return proto.CompactTextString(m) }
func (*BettingOverlay) ProtoMessage()               {}
func (*BettingOverlay) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *BettingOverlay) GetAgentId() uint32 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *BettingOverlay) GetSportEntity() *api.Entity {
	if m != nil {
		return m.SportEntity
	}
	return nil
}

func (m *BettingOverlay) GetOverlay() *Overlay {
	if m != nil {
		return m.Overlay
	}
	return nil
}

func (m *BettingOverlay) GetOverlayedEntity() *OverlayedEntity {
	if m != nil {
		return m.OverlayedEntity
	}
	return nil
}

func (m *BettingOverlay) GetNamespace() Namespace {
	if m != nil {
		return m.Namespace
	}
	return Namespace_UNDEFINED
}

func (m *BettingOverlay) GetBettingOverlayId() uint32 {
	if m != nil {
		return m.BettingOverlayId
	}
	return 0
}

type Overlay struct {
	StakeFactor float64 `protobuf:"fixed64,1,opt,name=stake_factor,json=stakeFactor" json:"stake_factor,omitempty"`
	Margin      float64 `protobuf:"fixed64,2,opt,name=margin" json:"margin,omitempty"`
	// 0: entire entity
	// otherwise only applies to specific market
	MarketId uint32 `protobuf:"varint,3,opt,name=market_id,json=marketId" json:"market_id,omitempty"`
}

func (m *Overlay) Reset()                    { *m = Overlay{} }
func (m *Overlay) String() string            { return proto.CompactTextString(m) }
func (*Overlay) ProtoMessage()               {}
func (*Overlay) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *Overlay) GetStakeFactor() float64 {
	if m != nil {
		return m.StakeFactor
	}
	return 0
}

func (m *Overlay) GetMargin() float64 {
	if m != nil {
		return m.Margin
	}
	return 0
}

func (m *Overlay) GetMarketId() uint32 {
	if m != nil {
		return m.MarketId
	}
	return 0
}

type BettingOverlays struct {
	BettingOverlays []*BettingOverlay `protobuf:"bytes,1,rep,name=betting_overlays,json=bettingOverlays" json:"betting_overlays,omitempty"`
}

func (m *BettingOverlays) Reset()                    { *m = BettingOverlays{} }
func (m *BettingOverlays) String() string            { return proto.CompactTextString(m) }
func (*BettingOverlays) ProtoMessage()               {}
func (*BettingOverlays) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{4} }

func (m *BettingOverlays) GetBettingOverlays() []*BettingOverlay {
	if m != nil {
		return m.BettingOverlays
	}
	return nil
}

// please take the overlay priority from high to low as:
// markets -> outrights/events -> competitions -> sports
type EntityBettingOverlays struct {
	// maps below are entity id -> entity's market overlay
	Sports       map[uint32]*MarketOverlays `protobuf:"bytes,1,rep,name=sports" json:"sports,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Competitions map[uint32]*MarketOverlays `protobuf:"bytes,2,rep,name=competitions" json:"competitions,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Events       map[uint32]*MarketOverlays `protobuf:"bytes,3,rep,name=events" json:"events,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Outrights    map[uint32]*MarketOverlays `protobuf:"bytes,4,rep,name=outrights" json:"outrights,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Markets      *MarketOverlays            `protobuf:"bytes,5,opt,name=markets" json:"markets,omitempty"`
}

func (m *EntityBettingOverlays) Reset()                    { *m = EntityBettingOverlays{} }
func (m *EntityBettingOverlays) String() string            { return proto.CompactTextString(m) }
func (*EntityBettingOverlays) ProtoMessage()               {}
func (*EntityBettingOverlays) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{5} }

func (m *EntityBettingOverlays) GetSports() map[uint32]*MarketOverlays {
	if m != nil {
		return m.Sports
	}
	return nil
}

func (m *EntityBettingOverlays) GetCompetitions() map[uint32]*MarketOverlays {
	if m != nil {
		return m.Competitions
	}
	return nil
}

func (m *EntityBettingOverlays) GetEvents() map[uint32]*MarketOverlays {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *EntityBettingOverlays) GetOutrights() map[uint32]*MarketOverlays {
	if m != nil {
		return m.Outrights
	}
	return nil
}

func (m *EntityBettingOverlays) GetMarkets() *MarketOverlays {
	if m != nil {
		return m.Markets
	}
	return nil
}

type MarketOverlays struct {
	// maps below are market id -> market overlay
	// if market id 0 exist, it will be the only overlay applies to all market
	Overlays map[uint32]*Overlay `protobuf:"bytes,1,rep,name=overlays" json:"overlays,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MarketOverlays) Reset()                    { *m = MarketOverlays{} }
func (m *MarketOverlays) String() string            { return proto.CompactTextString(m) }
func (*MarketOverlays) ProtoMessage()               {}
func (*MarketOverlays) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{6} }

func (m *MarketOverlays) GetOverlays() map[uint32]*Overlay {
	if m != nil {
		return m.Overlays
	}
	return nil
}

type RemoveBettingOverlayRequest struct {
	// Types that are valid to be assigned to OverlayLookup:
	//	*RemoveBettingOverlayRequest_BettingOverlayId
	//	*RemoveBettingOverlayRequest_DetailedLookup
	OverlayLookup isRemoveBettingOverlayRequest_OverlayLookup `protobuf_oneof:"overlay_lookup"`
	Namespace     Namespace                                   `protobuf:"varint,2,opt,name=namespace,enum=internalrpc.Namespace" json:"namespace,omitempty"`
}

func (m *RemoveBettingOverlayRequest) Reset()                    { *m = RemoveBettingOverlayRequest{} }
func (m *RemoveBettingOverlayRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveBettingOverlayRequest) ProtoMessage()               {}
func (*RemoveBettingOverlayRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{7} }

type isRemoveBettingOverlayRequest_OverlayLookup interface{ isRemoveBettingOverlayRequest_OverlayLookup() }

type RemoveBettingOverlayRequest_BettingOverlayId struct {
	BettingOverlayId uint32 `protobuf:"varint,1,opt,name=betting_overlay_id,json=bettingOverlayId,oneof"`
}
type RemoveBettingOverlayRequest_DetailedLookup struct {
	DetailedLookup *DetailedLookup `protobuf:"bytes,3,opt,name=detailed_lookup,json=detailedLookup,oneof"`
}

func (*RemoveBettingOverlayRequest_BettingOverlayId) isRemoveBettingOverlayRequest_OverlayLookup() {}
func (*RemoveBettingOverlayRequest_DetailedLookup) isRemoveBettingOverlayRequest_OverlayLookup()   {}

func (m *RemoveBettingOverlayRequest) GetOverlayLookup() isRemoveBettingOverlayRequest_OverlayLookup {
	if m != nil {
		return m.OverlayLookup
	}
	return nil
}

func (m *RemoveBettingOverlayRequest) GetBettingOverlayId() uint32 {
	if x, ok := m.GetOverlayLookup().(*RemoveBettingOverlayRequest_BettingOverlayId); ok {
		return x.BettingOverlayId
	}
	return 0
}

func (m *RemoveBettingOverlayRequest) GetDetailedLookup() *DetailedLookup {
	if x, ok := m.GetOverlayLookup().(*RemoveBettingOverlayRequest_DetailedLookup); ok {
		return x.DetailedLookup
	}
	return nil
}

func (m *RemoveBettingOverlayRequest) GetNamespace() Namespace {
	if m != nil {
		return m.Namespace
	}
	return Namespace_UNDEFINED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RemoveBettingOverlayRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RemoveBettingOverlayRequest_OneofMarshaler, _RemoveBettingOverlayRequest_OneofUnmarshaler, _RemoveBettingOverlayRequest_OneofSizer, []interface{}{
		(*RemoveBettingOverlayRequest_BettingOverlayId)(nil),
		(*RemoveBettingOverlayRequest_DetailedLookup)(nil),
	}
}

func _RemoveBettingOverlayRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RemoveBettingOverlayRequest)
	// overlay_lookup
	switch x := m.OverlayLookup.(type) {
	case *RemoveBettingOverlayRequest_BettingOverlayId:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.BettingOverlayId))
	case *RemoveBettingOverlayRequest_DetailedLookup:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DetailedLookup); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RemoveBettingOverlayRequest.OverlayLookup has unexpected type %T", x)
	}
	return nil
}

func _RemoveBettingOverlayRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RemoveBettingOverlayRequest)
	switch tag {
	case 1: // overlay_lookup.betting_overlay_id
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.OverlayLookup = &RemoveBettingOverlayRequest_BettingOverlayId{uint32(x)}
		return true, err
	case 3: // overlay_lookup.detailed_lookup
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DetailedLookup)
		err := b.DecodeMessage(msg)
		m.OverlayLookup = &RemoveBettingOverlayRequest_DetailedLookup{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RemoveBettingOverlayRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RemoveBettingOverlayRequest)
	// overlay_lookup
	switch x := m.OverlayLookup.(type) {
	case *RemoveBettingOverlayRequest_BettingOverlayId:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.BettingOverlayId))
	case *RemoveBettingOverlayRequest_DetailedLookup:
		s := proto.Size(x.DetailedLookup)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type DetailedLookup struct {
	AgentId         uint32           `protobuf:"varint,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	OverlayedEntity *OverlayedEntity `protobuf:"bytes,2,opt,name=overlayed_entity,json=overlayedEntity" json:"overlayed_entity,omitempty"`
	SportEntity     *api.Entity      `protobuf:"bytes,3,opt,name=sport_entity,json=sportEntity" json:"sport_entity,omitempty"`
}

func (m *DetailedLookup) Reset()                    { *m = DetailedLookup{} }
func (m *DetailedLookup) String() string            { return proto.CompactTextString(m) }
func (*DetailedLookup) ProtoMessage()               {}
func (*DetailedLookup) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{8} }

func (m *DetailedLookup) GetAgentId() uint32 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *DetailedLookup) GetOverlayedEntity() *OverlayedEntity {
	if m != nil {
		return m.OverlayedEntity
	}
	return nil
}

func (m *DetailedLookup) GetSportEntity() *api.Entity {
	if m != nil {
		return m.SportEntity
	}
	return nil
}

type GetBettingOverlayRequest struct {
	Id        uint32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Namespace Namespace `protobuf:"varint,2,opt,name=namespace,enum=internalrpc.Namespace" json:"namespace,omitempty"`
	// empty -> no filter
	Filter *Filter `protobuf:"bytes,3,opt,name=filter" json:"filter,omitempty"`
}

func (m *GetBettingOverlayRequest) Reset()                    { *m = GetBettingOverlayRequest{} }
func (m *GetBettingOverlayRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBettingOverlayRequest) ProtoMessage()               {}
func (*GetBettingOverlayRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{9} }

func (m *GetBettingOverlayRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetBettingOverlayRequest) GetNamespace() Namespace {
	if m != nil {
		return m.Namespace
	}
	return Namespace_UNDEFINED
}

func (m *GetBettingOverlayRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type BatchGetBettingOverlayRequest struct {
	Ids       []uint32  `protobuf:"varint,1,rep,packed,name=ids" json:"ids,omitempty"`
	Namespace Namespace `protobuf:"varint,2,opt,name=namespace,enum=internalrpc.Namespace" json:"namespace,omitempty"`
	// empty -> no filter
	Filter *Filter `protobuf:"bytes,3,opt,name=filter" json:"filter,omitempty"`
}

func (m *BatchGetBettingOverlayRequest) Reset()                    { *m = BatchGetBettingOverlayRequest{} }
func (m *BatchGetBettingOverlayRequest) String() string            { return proto.CompactTextString(m) }
func (*BatchGetBettingOverlayRequest) ProtoMessage()               {}
func (*BatchGetBettingOverlayRequest) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{10} }

func (m *BatchGetBettingOverlayRequest) GetIds() []uint32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *BatchGetBettingOverlayRequest) GetNamespace() Namespace {
	if m != nil {
		return m.Namespace
	}
	return Namespace_UNDEFINED
}

func (m *BatchGetBettingOverlayRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func init() {
	proto.RegisterType((*Filter)(nil), "internalrpc.Filter")
	proto.RegisterType((*OverlayedEntity)(nil), "internalrpc.OverlayedEntity")
	proto.RegisterType((*BettingOverlay)(nil), "internalrpc.BettingOverlay")
	proto.RegisterType((*Overlay)(nil), "internalrpc.Overlay")
	proto.RegisterType((*BettingOverlays)(nil), "internalrpc.BettingOverlays")
	proto.RegisterType((*EntityBettingOverlays)(nil), "internalrpc.EntityBettingOverlays")
	proto.RegisterType((*MarketOverlays)(nil), "internalrpc.MarketOverlays")
	proto.RegisterType((*RemoveBettingOverlayRequest)(nil), "internalrpc.RemoveBettingOverlayRequest")
	proto.RegisterType((*DetailedLookup)(nil), "internalrpc.DetailedLookup")
	proto.RegisterType((*GetBettingOverlayRequest)(nil), "internalrpc.GetBettingOverlayRequest")
	proto.RegisterType((*BatchGetBettingOverlayRequest)(nil), "internalrpc.BatchGetBettingOverlayRequest")
	proto.RegisterEnum("internalrpc.Namespace", Namespace_name, Namespace_value)
	proto.RegisterEnum("internalrpc.OverlayedType", OverlayedType_name, OverlayedType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RiskManager service

type RiskManagerClient interface {
	// Input:
	//     agent_id: required
	//     overlayed_entity: optional, empty -> all direct children players & agents and further down to the tree
	//     entity: (id, type) required
	//     overlay: (stake_factor, margin) required
	//     namespace: required
	//     betting_overlay_id: ignore
	// Return:
	//     betting overlay id
	SetBettingOverlay(ctx context.Context, in *BettingOverlay, opts ...grpc.CallOption) (*SingleId, error)
	// Input:
	//     betting_overlay_id: required
	//     namespace: required
	// Return:
	//     none
	// Expected effect:
	//     remove a overlay record from taking effect
	//     similar to set entity to 1
	RemoveBettingOverlay(ctx context.Context, in *RemoveBettingOverlayRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Input:
	//     id: required, agent id
	//     namespace: required
	// Return:
	//     list of overlay in sports, competitions, events, outrights and markets
	//     - list will be aggregated by all ancestor agents of input agent
	GetAgentBettingOverlay(ctx context.Context, in *GetBettingOverlayRequest, opts ...grpc.CallOption) (*EntityBettingOverlays, error)
	// Input:
	//     id: required, agent id
	//     namespace: required
	// Return:
	//     array of overlays applied from agent id towards its children
	GetDirectChildrenBettingOverlays(ctx context.Context, in *GetBettingOverlayRequest, opts ...grpc.CallOption) (*BettingOverlays, error)
	// Input:
	//     id: required, player id
	//     namespace: required
	//     filter: please ignore
	// Return:
	//     list of overlay in sports, competitions, events, outrights and markets
	//     - list will be aggregated by agent & player
	GetPlayerBettingOverlay(ctx context.Context, in *GetBettingOverlayRequest, opts ...grpc.CallOption) (*EntityBettingOverlays, error)
}

type riskManagerClient struct {
	cc *grpc.ClientConn
}

func NewRiskManagerClient(cc *grpc.ClientConn) RiskManagerClient {
	return &riskManagerClient{cc}
}

func (c *riskManagerClient) SetBettingOverlay(ctx context.Context, in *BettingOverlay, opts ...grpc.CallOption) (*SingleId, error) {
	out := new(SingleId)
	err := grpc.Invoke(ctx, "/internalrpc.RiskManager/SetBettingOverlay", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riskManagerClient) RemoveBettingOverlay(ctx context.Context, in *RemoveBettingOverlayRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/internalrpc.RiskManager/RemoveBettingOverlay", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riskManagerClient) GetAgentBettingOverlay(ctx context.Context, in *GetBettingOverlayRequest, opts ...grpc.CallOption) (*EntityBettingOverlays, error) {
	out := new(EntityBettingOverlays)
	err := grpc.Invoke(ctx, "/internalrpc.RiskManager/GetAgentBettingOverlay", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riskManagerClient) GetDirectChildrenBettingOverlays(ctx context.Context, in *GetBettingOverlayRequest, opts ...grpc.CallOption) (*BettingOverlays, error) {
	out := new(BettingOverlays)
	err := grpc.Invoke(ctx, "/internalrpc.RiskManager/GetDirectChildrenBettingOverlays", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *riskManagerClient) GetPlayerBettingOverlay(ctx context.Context, in *GetBettingOverlayRequest, opts ...grpc.CallOption) (*EntityBettingOverlays, error) {
	out := new(EntityBettingOverlays)
	err := grpc.Invoke(ctx, "/internalrpc.RiskManager/GetPlayerBettingOverlay", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RiskManager service

type RiskManagerServer interface {
	// Input:
	//     agent_id: required
	//     overlayed_entity: optional, empty -> all direct children players & agents and further down to the tree
	//     entity: (id, type) required
	//     overlay: (stake_factor, margin) required
	//     namespace: required
	//     betting_overlay_id: ignore
	// Return:
	//     betting overlay id
	SetBettingOverlay(context.Context, *BettingOverlay) (*SingleId, error)
	// Input:
	//     betting_overlay_id: required
	//     namespace: required
	// Return:
	//     none
	// Expected effect:
	//     remove a overlay record from taking effect
	//     similar to set entity to 1
	RemoveBettingOverlay(context.Context, *RemoveBettingOverlayRequest) (*google_protobuf1.Empty, error)
	// Input:
	//     id: required, agent id
	//     namespace: required
	// Return:
	//     list of overlay in sports, competitions, events, outrights and markets
	//     - list will be aggregated by all ancestor agents of input agent
	GetAgentBettingOverlay(context.Context, *GetBettingOverlayRequest) (*EntityBettingOverlays, error)
	// Input:
	//     id: required, agent id
	//     namespace: required
	// Return:
	//     array of overlays applied from agent id towards its children
	GetDirectChildrenBettingOverlays(context.Context, *GetBettingOverlayRequest) (*BettingOverlays, error)
	// Input:
	//     id: required, player id
	//     namespace: required
	//     filter: please ignore
	// Return:
	//     list of overlay in sports, competitions, events, outrights and markets
	//     - list will be aggregated by agent & player
	GetPlayerBettingOverlay(context.Context, *GetBettingOverlayRequest) (*EntityBettingOverlays, error)
}

func RegisterRiskManagerServer(s *grpc.Server, srv RiskManagerServer) {
	s.RegisterService(&_RiskManager_serviceDesc, srv)
}

func _RiskManager_SetBettingOverlay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BettingOverlay)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskManagerServer).SetBettingOverlay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internalrpc.RiskManager/SetBettingOverlay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskManagerServer).SetBettingOverlay(ctx, req.(*BettingOverlay))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiskManager_RemoveBettingOverlay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveBettingOverlayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskManagerServer).RemoveBettingOverlay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internalrpc.RiskManager/RemoveBettingOverlay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskManagerServer).RemoveBettingOverlay(ctx, req.(*RemoveBettingOverlayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiskManager_GetAgentBettingOverlay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBettingOverlayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskManagerServer).GetAgentBettingOverlay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internalrpc.RiskManager/GetAgentBettingOverlay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskManagerServer).GetAgentBettingOverlay(ctx, req.(*GetBettingOverlayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiskManager_GetDirectChildrenBettingOverlays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBettingOverlayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskManagerServer).GetDirectChildrenBettingOverlays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internalrpc.RiskManager/GetDirectChildrenBettingOverlays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskManagerServer).GetDirectChildrenBettingOverlays(ctx, req.(*GetBettingOverlayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RiskManager_GetPlayerBettingOverlay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBettingOverlayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RiskManagerServer).GetPlayerBettingOverlay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internalrpc.RiskManager/GetPlayerBettingOverlay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RiskManagerServer).GetPlayerBettingOverlay(ctx, req.(*GetBettingOverlayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RiskManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internalrpc.RiskManager",
	HandlerType: (*RiskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetBettingOverlay",
			Handler:    _RiskManager_SetBettingOverlay_Handler,
		},
		{
			MethodName: "RemoveBettingOverlay",
			Handler:    _RiskManager_RemoveBettingOverlay_Handler,
		},
		{
			MethodName: "GetAgentBettingOverlay",
			Handler:    _RiskManager_GetAgentBettingOverlay_Handler,
		},
		{
			MethodName: "GetDirectChildrenBettingOverlays",
			Handler:    _RiskManager_GetDirectChildrenBettingOverlays_Handler,
		},
		{
			MethodName: "GetPlayerBettingOverlay",
			Handler:    _RiskManager_GetPlayerBettingOverlay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internalrpc/risk_manager.proto",
}

func init() { proto.RegisterFile("internalrpc/risk_manager.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 1004 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdf, 0x6e, 0xe2, 0xc6,
	0x17, 0xc6, 0x90, 0x40, 0x38, 0x04, 0xf0, 0xce, 0x6f, 0x37, 0x3f, 0x2f, 0x64, 0xdb, 0xd4, 0x52,
	0x25, 0x9a, 0x56, 0x46, 0x9b, 0x6e, 0xa5, 0xaa, 0x77, 0x21, 0x18, 0x16, 0x29, 0x1b, 0xb2, 0xce,
	0x1f, 0x35, 0x55, 0x25, 0x34, 0xe0, 0x89, 0x33, 0xc2, 0xd8, 0xae, 0x3d, 0x44, 0xe2, 0x21, 0x7a,
	0xd7, 0x47, 0xe8, 0x55, 0x9f, 0xa0, 0x2f, 0xd3, 0x57, 0xe8, 0x45, 0x9f, 0xa0, 0x62, 0xc6, 0x26,
	0x8c, 0xc3, 0x26, 0x74, 0x95, 0xf6, 0xce, 0x9e, 0x73, 0xbe, 0xef, 0x7c, 0xfe, 0x8e, 0x67, 0xce,
	0xc0, 0x27, 0xd4, 0x63, 0x24, 0xf4, 0xb0, 0x1b, 0x06, 0xa3, 0x66, 0x48, 0xa3, 0xf1, 0x60, 0x82,
	0x3d, 0xec, 0x90, 0xd0, 0x08, 0x42, 0x9f, 0xf9, 0xa8, 0xb4, 0x14, 0xaf, 0xd5, 0x1d, 0xdf, 0x77,
	0x5c, 0xd2, 0xe4, 0xa1, 0xe1, 0xf4, 0xba, 0x49, 0x26, 0x01, 0x9b, 0x89, 0xcc, 0xda, 0xa7, 0x51,
	0xe0, 0x87, 0x2c, 0x1a, 0xfa, 0xfe, 0xb8, 0x89, 0x03, 0xda, 0xbc, 0x7b, 0x8d, 0x13, 0x5e, 0x2e,
	0x97, 0x72, 0x88, 0x47, 0x42, 0x3a, 0x12, 0x21, 0xfd, 0x02, 0xf2, 0x1d, 0xea, 0x32, 0x12, 0x22,
	0x03, 0xb6, 0x39, 0x70, 0x40, 0x3c, 0x46, 0xd9, 0x4c, 0x53, 0xf6, 0x94, 0x46, 0xe9, 0xa0, 0x64,
	0xe0, 0x80, 0x1a, 0x26, 0x5f, 0xb2, 0x4a, 0x3c, 0x41, 0xbc, 0xa0, 0x3a, 0x14, 0x27, 0x38, 0x1c,
	0x13, 0x36, 0xa0, 0xb6, 0x96, 0xdd, 0x53, 0x1a, 0x65, 0x6b, 0x4b, 0x2c, 0xf4, 0x6c, 0xfd, 0x3d,
	0x54, 0xfb, 0xb7, 0x24, 0x74, 0xf1, 0x8c, 0xd8, 0x71, 0x7e, 0x05, 0xb2, 0xd4, 0xe6, 0xac, 0x65,
	0x2b, 0x4b, 0x6d, 0x64, 0xc0, 0x06, 0x9b, 0x05, 0x84, 0x43, 0x2b, 0x07, 0x35, 0x63, 0x49, 0xa3,
	0xb1, 0xc0, 0x9e, 0xcf, 0x02, 0x62, 0xf1, 0x3c, 0xfd, 0xf7, 0x2c, 0x54, 0x5a, 0x84, 0x31, 0xea,
	0x39, 0x71, 0x18, 0xbd, 0x84, 0x2d, 0xec, 0x10, 0x6f, 0x49, 0x41, 0x81, 0xbf, 0xf7, 0xec, 0x7b,
	0x5f, 0x93, 0x7b, 0xe4, 0x6b, 0x0c, 0x28, 0xf8, 0x82, 0x55, 0xdb, 0xe0, 0xa9, 0xcf, 0x57, 0x09,
	0xb2, 0x92, 0x24, 0xd4, 0x05, 0xd5, 0x4f, 0x44, 0x26, 0x35, 0x0a, 0x1c, 0xb8, 0xbb, 0xfa, 0x4b,
	0xe2, 0xa2, 0x55, 0x3f, 0x65, 0xcb, 0x1b, 0x28, 0x7a, 0x78, 0x42, 0xa2, 0x00, 0x8f, 0x88, 0xb6,
	0xc9, 0xbd, 0xd8, 0x91, 0x18, 0x4e, 0x92, 0xa8, 0x75, 0x97, 0x88, 0xbe, 0x02, 0x34, 0x14, 0x5e,
	0x0c, 0x62, 0xc2, 0xb9, 0x07, 0x79, 0xee, 0x81, 0x3a, 0x94, 0x5c, 0xea, 0xd9, 0x3a, 0x86, 0x42,
	0x62, 0xd9, 0x67, 0xb0, 0x1d, 0x31, 0x3c, 0x26, 0x83, 0x6b, 0x3c, 0x62, 0x7e, 0xc8, 0xfb, 0xa1,
	0x58, 0x25, 0xbe, 0xd6, 0xe1, 0x4b, 0x68, 0x07, 0xf2, 0x13, 0x1c, 0x3a, 0xd4, 0xe3, 0x9e, 0x2a,
	0x56, 0xfc, 0x26, 0x37, 0x3c, 0x97, 0x6a, 0xf8, 0x15, 0x54, 0xe5, 0xe6, 0x44, 0xa8, 0x03, 0x6a,
	0x4a, 0x63, 0xa4, 0x29, 0x7b, 0xb9, 0x46, 0xe9, 0xa0, 0x2e, 0x7d, 0xa0, 0x8c, 0xb3, 0xaa, 0xb2,
	0xfc, 0x48, 0xff, 0x6b, 0x13, 0x5e, 0x08, 0xb3, 0xee, 0x57, 0xc8, 0x8b, 0x7f, 0x3d, 0xe6, 0x35,
	0x24, 0xde, 0x95, 0x18, 0xe3, 0x8c, 0x03, 0x4c, 0x8f, 0x85, 0x33, 0x2b, 0x46, 0xa3, 0xef, 0x61,
	0x7b, 0xe4, 0x4f, 0x02, 0xc2, 0x28, 0xa3, 0xbe, 0x17, 0x69, 0x59, 0xce, 0xf6, 0x66, 0x0d, 0xb6,
	0xa3, 0x25, 0x98, 0xe0, 0x94, 0x98, 0xe6, 0x0a, 0xc9, 0x2d, 0xf1, 0x58, 0xa4, 0xe5, 0xd6, 0x56,
	0x68, 0x72, 0x40, 0xac, 0x50, 0xa0, 0x51, 0x1f, 0x8a, 0xfe, 0x94, 0x85, 0xd4, 0xb9, 0x61, 0x91,
	0xb6, 0xc1, 0xa9, 0x5e, 0xaf, 0x41, 0xd5, 0x4f, 0x30, 0x82, 0xed, 0x8e, 0x03, 0x7d, 0x03, 0x05,
	0xd1, 0xbb, 0x88, 0xff, 0x74, 0xe9, 0x9e, 0xbc, 0xe3, 0xb1, 0x84, 0xc7, 0x4a, 0x72, 0x6b, 0x97,
	0x50, 0x5a, 0x32, 0x10, 0xa9, 0x90, 0x1b, 0x93, 0x59, 0xbc, 0xa9, 0xe7, 0x8f, 0xe8, 0x35, 0x6c,
	0xde, 0x62, 0x77, 0x2a, 0xb6, 0xf5, 0x23, 0xac, 0x22, 0xf3, 0xbb, 0xec, 0xb7, 0x4a, 0xed, 0x47,
	0x78, 0x76, 0xcf, 0xca, 0xa7, 0x63, 0xbf, 0x84, 0xd2, 0x92, 0xa9, 0x4f, 0xc7, 0x7b, 0x05, 0x15,
	0xd9, 0xe1, 0x27, 0xa3, 0xd6, 0x7f, 0x53, 0xa0, 0x22, 0x47, 0x91, 0x09, 0x5b, 0xa9, 0x7d, 0xf4,
	0xc5, 0x03, 0x64, 0xc9, 0xc9, 0x13, 0xb7, 0x7e, 0x01, 0xad, 0xbd, 0x87, 0xb2, 0x14, 0x5a, 0xa1,
	0x79, 0x5f, 0xd6, 0xbc, 0xfa, 0x28, 0x5c, 0x12, 0xfb, 0x87, 0x02, 0x75, 0x8b, 0x4c, 0xfc, 0x5b,
	0x92, 0xda, 0xcb, 0xe4, 0xa7, 0x29, 0x89, 0x18, 0x32, 0x56, 0x9e, 0x56, 0xbc, 0xe0, 0xdb, 0xcc,
	0xfd, 0xf3, 0x0a, 0x75, 0xa0, 0x6a, 0x13, 0x86, 0xa9, 0x4b, 0xec, 0x81, 0xeb, 0xfb, 0xe3, 0x69,
	0x10, 0x9f, 0xdf, 0xb2, 0x7b, 0xed, 0x38, 0xe7, 0x98, 0xa7, 0xbc, 0xcd, 0x58, 0x15, 0x5b, 0x5a,
	0x91, 0xcf, 0xd6, 0xec, 0x9a, 0x67, 0x6b, 0x4b, 0x85, 0x4a, 0xa2, 0x52, 0x14, 0xd7, 0x7f, 0x55,
	0xa0, 0x22, 0x17, 0x93, 0x46, 0x8f, 0x22, 0x8f, 0x9e, 0x55, 0xa3, 0x21, 0xfb, 0x31, 0xa3, 0xe1,
	0x1f, 0xce, 0x30, 0xfd, 0x67, 0x05, 0xb4, 0x2e, 0x61, 0xab, 0x7b, 0x90, 0x1e, 0xbf, 0x1f, 0xe5,
	0x0d, 0xfa, 0x12, 0xf2, 0xd7, 0xfc, 0xba, 0x10, 0x8b, 0xf9, 0x9f, 0x04, 0x11, 0x37, 0x09, 0x2b,
	0x4e, 0xd1, 0x7f, 0x51, 0xe0, 0x55, 0x0b, 0xb3, 0xd1, 0xcd, 0x07, 0x45, 0xa9, 0x90, 0xa3, 0xb6,
	0xf8, 0x9b, 0xcb, 0xd6, 0xfc, 0xf1, 0x3f, 0x90, 0xb5, 0xdf, 0x80, 0xe2, 0x82, 0x04, 0x95, 0xa1,
	0x78, 0x71, 0xd2, 0x36, 0x3b, 0xbd, 0x13, 0xb3, 0xad, 0x66, 0xd0, 0x36, 0x6c, 0x1d, 0x1d, 0xf7,
	0x2f, 0xda, 0x2d, 0xf3, 0x5c, 0x55, 0xf6, 0xdb, 0x8b, 0xad, 0x22, 0x6e, 0x22, 0x68, 0x17, 0xb4,
	0x45, 0xf6, 0xa0, 0x7f, 0x69, 0x5a, 0xc7, 0x87, 0x57, 0x66, 0x7b, 0x70, 0x7e, 0x75, 0x6a, 0xaa,
	0x19, 0x04, 0x90, 0x3f, 0x9d, 0x2f, 0x58, 0xaa, 0x82, 0x8a, 0xb0, 0x79, 0xd8, 0x35, 0x4f, 0xce,
	0xd5, 0xec, 0xc1, 0x9f, 0x39, 0x28, 0x59, 0x34, 0x1a, 0xbf, 0x13, 0xd7, 0x3b, 0xd4, 0x83, 0x67,
	0x67, 0x69, 0x43, 0xd0, 0x43, 0x23, 0xb1, 0xf6, 0x42, 0x0a, 0x9e, 0x51, 0xcf, 0x71, 0x49, 0xcf,
	0xd6, 0x33, 0xe8, 0x07, 0x78, 0xbe, 0x6a, 0xdf, 0xa1, 0x86, 0x04, 0x78, 0x60, 0x6b, 0xd6, 0x76,
	0x0c, 0x71, 0xb3, 0x34, 0x92, 0x9b, 0xa5, 0x61, 0xce, 0x6f, 0x96, 0x7a, 0x06, 0x11, 0xd8, 0xe9,
	0x12, 0x76, 0x38, 0xff, 0xa9, 0x53, 0xec, 0x9f, 0x4b, 0xec, 0x1f, 0x6a, 0x6e, 0x4d, 0x7f, 0x7c,
	0x40, 0xe9, 0x19, 0x44, 0x61, 0xaf, 0x4b, 0x58, 0x9b, 0x86, 0x64, 0xc4, 0x8e, 0x6e, 0xa8, 0x6b,
	0x87, 0xc4, 0x4b, 0xcf, 0xf9, 0x35, 0x0b, 0xee, 0x3e, 0xe0, 0xe1, 0xbc, 0xd4, 0x35, 0xfc, 0xbf,
	0x4b, 0xd8, 0xe9, 0xbc, 0x9d, 0xe1, 0xbf, 0xf8, 0x49, 0xad, 0x57, 0x50, 0x1f, 0xf9, 0x13, 0x63,
	0x48, 0x99, 0x83, 0x27, 0xc4, 0xc5, 0xc3, 0x68, 0x19, 0x36, 0xcc, 0x73, 0xab, 0xbf, 0xfe, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x44, 0x7d, 0x94, 0x32, 0x01, 0x0c, 0x00, 0x00,
}