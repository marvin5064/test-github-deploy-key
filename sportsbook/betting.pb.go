// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sportsbook/betting.proto

package sportsbook

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BetSelection_Flags int32

const (
	// make sure if we default here depending on the language, that we
	// don't conflict with any flag. Therefore add NONE as a default.
	BetSelection_NONE BetSelection_Flags = 0
	// any of these can be set or set together in a pending acceptance
	// response
	BetSelection_BETTER_ODDS_ACCEPTED BetSelection_Flags = 1
	BetSelection_ACCEPTED             BetSelection_Flags = 4
	BetSelection_REJECTED             BetSelection_Flags = 8
	// deprecated, we either accept everything or nothing
	BetSelection_PENDING_ACCEPTANCE BetSelection_Flags = 2
)

var BetSelection_Flags_name = map[int32]string{
	0: "NONE",
	1: "BETTER_ODDS_ACCEPTED",
	4: "ACCEPTED",
	8: "REJECTED",
	2: "PENDING_ACCEPTANCE",
}
var BetSelection_Flags_value = map[string]int32{
	"NONE":                 0,
	"BETTER_ODDS_ACCEPTED": 1,
	"ACCEPTED":             4,
	"REJECTED":             8,
	"PENDING_ACCEPTANCE":   2,
}

func (x BetSelection_Flags) Enum() *BetSelection_Flags {
	p := new(BetSelection_Flags)
	*p = x
	return p
}
func (x BetSelection_Flags) String() string {
	return proto.EnumName(BetSelection_Flags_name, int32(x))
}
func (x *BetSelection_Flags) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BetSelection_Flags_value, data, "BetSelection_Flags")
	if err != nil {
		return err
	}
	*x = BetSelection_Flags(value)
	return nil
}
func (BetSelection_Flags) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type SystemBet_Flags int32

const (
	SystemBet_NONE SystemBet_Flags = 0
	// any of these can be set or set together in a pending acceptance
	// response
	SystemBet_BETTER_ODDS_ACCEPTED SystemBet_Flags = 1
	SystemBet_ACCEPTED             SystemBet_Flags = 4
	SystemBet_REJECTED             SystemBet_Flags = 8
)

var SystemBet_Flags_name = map[int32]string{
	0: "NONE",
	1: "BETTER_ODDS_ACCEPTED",
	4: "ACCEPTED",
	8: "REJECTED",
}
var SystemBet_Flags_value = map[string]int32{
	"NONE":                 0,
	"BETTER_ODDS_ACCEPTED": 1,
	"ACCEPTED":             4,
	"REJECTED":             8,
}

func (x SystemBet_Flags) Enum() *SystemBet_Flags {
	p := new(SystemBet_Flags)
	*p = x
	return p
}
func (x SystemBet_Flags) String() string {
	return proto.EnumName(SystemBet_Flags_name, int32(x))
}
func (x *SystemBet_Flags) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SystemBet_Flags_value, data, "SystemBet_Flags")
	if err != nil {
		return err
	}
	*x = SystemBet_Flags(value)
	return nil
}
func (SystemBet_Flags) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

type BetResponse_ResultCode int32

const (
	// accepted
	BetResponse_OK                         BetResponse_ResultCode = 0
	BetResponse_NOT_SIGNED_IN              BetResponse_ResultCode = 1
	BetResponse_INSUFFICIENT_FUNDS         BetResponse_ResultCode = 2
	BetResponse_INVALID_SELECTION          BetResponse_ResultCode = 3
	BetResponse_INVALID_ODDS               BetResponse_ResultCode = 4
	BetResponse_INVALID_STAKE              BetResponse_ResultCode = 5
	BetResponse_CORRELATED_SELECTIONS      BetResponse_ResultCode = 6
	BetResponse_MAX_STAKE_EXCEEDED         BetResponse_ResultCode = 7
	BetResponse_SPORTSBOOK_LOCKED          BetResponse_ResultCode = 8
	BetResponse_PREMATCH_BETS_LOCKED       BetResponse_ResultCode = 9
	BetResponse_LIVE_BETS_LOCKED           BetResponse_ResultCode = 10
	BetResponse_INTERNAL_SERVER_ERROR      BetResponse_ResultCode = 11
	BetResponse_TIMEOUT                    BetResponse_ResultCode = 12
	BetResponse_CUMULATIVE_STAKE_EXCEEDED  BetResponse_ResultCode = 13
	BetResponse_TOO_MANY_SELECTIONS        BetResponse_ResultCode = 14
	BetResponse_PREMATCH_BET_ON_LIVE_EVENT BetResponse_ResultCode = 16
	// intermediate state for live bets
	BetResponse_PENDING_ACCEPTANCE BetResponse_ResultCode = 17
	// used for pusher updates after pending acceptance
	BetResponse_OK_WITH_FLAGS BetResponse_ResultCode = 19
	// pinnacle doesn't allow some parlays
	BetResponse_PARLAY_RESTRICTION BetResponse_ResultCode = 20
	// per player rules for event/market/competition/selection restrictions
	BetResponse_RESTRICTED BetResponse_ResultCode = 21
	// per player rules for event/market/competition/selection overlays
	BetResponse_RESTRICTED_BY_OVERLAY BetResponse_ResultCode = 24
	// deprecated fields; don't use, moved to flags
	BetResponse_BETTER_ODDS_ACCEPTED BetResponse_ResultCode = 15
	// sports bonus related rejections
	BetResponse_CORRELATED_SELECTION_WITH_ACTIVE_BONUS BetResponse_ResultCode = 22
	BetResponse_MIN_PRICE_REQUIRED_WITH_ACTIVE_BONUS   BetResponse_ResultCode = 23
)

var BetResponse_ResultCode_name = map[int32]string{
	0:  "OK",
	1:  "NOT_SIGNED_IN",
	2:  "INSUFFICIENT_FUNDS",
	3:  "INVALID_SELECTION",
	4:  "INVALID_ODDS",
	5:  "INVALID_STAKE",
	6:  "CORRELATED_SELECTIONS",
	7:  "MAX_STAKE_EXCEEDED",
	8:  "SPORTSBOOK_LOCKED",
	9:  "PREMATCH_BETS_LOCKED",
	10: "LIVE_BETS_LOCKED",
	11: "INTERNAL_SERVER_ERROR",
	12: "TIMEOUT",
	13: "CUMULATIVE_STAKE_EXCEEDED",
	14: "TOO_MANY_SELECTIONS",
	16: "PREMATCH_BET_ON_LIVE_EVENT",
	17: "PENDING_ACCEPTANCE",
	19: "OK_WITH_FLAGS",
	20: "PARLAY_RESTRICTION",
	21: "RESTRICTED",
	24: "RESTRICTED_BY_OVERLAY",
	15: "BETTER_ODDS_ACCEPTED",
	22: "CORRELATED_SELECTION_WITH_ACTIVE_BONUS",
	23: "MIN_PRICE_REQUIRED_WITH_ACTIVE_BONUS",
}
var BetResponse_ResultCode_value = map[string]int32{
	"OK":                                     0,
	"NOT_SIGNED_IN":                          1,
	"INSUFFICIENT_FUNDS":                     2,
	"INVALID_SELECTION":                      3,
	"INVALID_ODDS":                           4,
	"INVALID_STAKE":                          5,
	"CORRELATED_SELECTIONS":                  6,
	"MAX_STAKE_EXCEEDED":                     7,
	"SPORTSBOOK_LOCKED":                      8,
	"PREMATCH_BETS_LOCKED":                   9,
	"LIVE_BETS_LOCKED":                       10,
	"INTERNAL_SERVER_ERROR":                  11,
	"TIMEOUT":                                12,
	"CUMULATIVE_STAKE_EXCEEDED":              13,
	"TOO_MANY_SELECTIONS":                    14,
	"PREMATCH_BET_ON_LIVE_EVENT":             16,
	"PENDING_ACCEPTANCE":                     17,
	"OK_WITH_FLAGS":                          19,
	"PARLAY_RESTRICTION":                     20,
	"RESTRICTED":                             21,
	"RESTRICTED_BY_OVERLAY":                  24,
	"BETTER_ODDS_ACCEPTED":                   15,
	"CORRELATED_SELECTION_WITH_ACTIVE_BONUS": 22,
	"MIN_PRICE_REQUIRED_WITH_ACTIVE_BONUS":   23,
}

func (x BetResponse_ResultCode) Enum() *BetResponse_ResultCode {
	p := new(BetResponse_ResultCode)
	*p = x
	return p
}
func (x BetResponse_ResultCode) String() string {
	return proto.EnumName(BetResponse_ResultCode_name, int32(x))
}
func (x *BetResponse_ResultCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BetResponse_ResultCode_value, data, "BetResponse_ResultCode")
	if err != nil {
		return err
	}
	*x = BetResponse_ResultCode(value)
	return nil
}
func (BetResponse_ResultCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3, 0} }

type PriceUpdate_AccountType int32

const (
	PriceUpdate_SHARP  PriceUpdate_AccountType = 0
	PriceUpdate_NORMAL PriceUpdate_AccountType = 1
)

var PriceUpdate_AccountType_name = map[int32]string{
	0: "SHARP",
	1: "NORMAL",
}
var PriceUpdate_AccountType_value = map[string]int32{
	"SHARP":  0,
	"NORMAL": 1,
}

func (x PriceUpdate_AccountType) Enum() *PriceUpdate_AccountType {
	p := new(PriceUpdate_AccountType)
	*p = x
	return p
}
func (x PriceUpdate_AccountType) String() string {
	return proto.EnumName(PriceUpdate_AccountType_name, int32(x))
}
func (x *PriceUpdate_AccountType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PriceUpdate_AccountType_value, data, "PriceUpdate_AccountType")
	if err != nil {
		return err
	}
	*x = PriceUpdate_AccountType(value)
	return nil
}
func (PriceUpdate_AccountType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{4, 0} }

type BetAcceptanceUpdate_ResultCode int32

const (
	BetAcceptanceUpdate_ACCEPTED BetAcceptanceUpdate_ResultCode = 0
	BetAcceptanceUpdate_REJECTED BetAcceptanceUpdate_ResultCode = 1
)

var BetAcceptanceUpdate_ResultCode_name = map[int32]string{
	0: "ACCEPTED",
	1: "REJECTED",
}
var BetAcceptanceUpdate_ResultCode_value = map[string]int32{
	"ACCEPTED": 0,
	"REJECTED": 1,
}

func (x BetAcceptanceUpdate_ResultCode) Enum() *BetAcceptanceUpdate_ResultCode {
	p := new(BetAcceptanceUpdate_ResultCode)
	*p = x
	return p
}
func (x BetAcceptanceUpdate_ResultCode) String() string {
	return proto.EnumName(BetAcceptanceUpdate_ResultCode_name, int32(x))
}
func (x *BetAcceptanceUpdate_ResultCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BetAcceptanceUpdate_ResultCode_value, data, "BetAcceptanceUpdate_ResultCode")
	if err != nil {
		return err
	}
	*x = BetAcceptanceUpdate_ResultCode(value)
	return nil
}
func (BetAcceptanceUpdate_ResultCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{5, 0}
}

type BetSelection struct {
	EventType        *string  `protobuf:"bytes,1,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	EventId          *uint32  `protobuf:"varint,2,opt,name=event_id,json=eventId" json:"event_id,omitempty"`
	MarketId         *uint32  `protobuf:"varint,3,opt,name=market_id,json=marketId" json:"market_id,omitempty"`
	Outcome          *string  `protobuf:"bytes,4,opt,name=outcome" json:"outcome,omitempty"`
	SpecialBetValue  *string  `protobuf:"bytes,5,opt,name=special_bet_value,json=specialBetValue" json:"special_bet_value,omitempty"`
	Timestamp        *int64   `protobuf:"varint,6,opt,name=timestamp" json:"timestamp,omitempty"`
	Odds             *float64 `protobuf:"fixed64,7,opt,name=odds" json:"odds,omitempty"`
	Stake            *float64 `protobuf:"fixed64,8,opt,name=stake" json:"stake,omitempty"`
	Flags            *uint32  `protobuf:"varint,9,opt,name=flags" json:"flags,omitempty"`
	BetId            *uint32  `protobuf:"varint,10,opt,name=bet_id,json=betId" json:"bet_id,omitempty"`
	IsLay            *bool    `protobuf:"varint,11,opt,name=is_lay,json=isLay" json:"is_lay,omitempty"`
	BackersOdds      *float64 `protobuf:"fixed64,12,opt,name=backers_odds,json=backersOdds" json:"backers_odds,omitempty"`
	BackersStake     *float64 `protobuf:"fixed64,13,opt,name=backers_stake,json=backersStake" json:"backers_stake,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BetSelection) Reset()                    { *m = BetSelection{} }
func (m *BetSelection) String() string            { return proto.CompactTextString(m) }
func (*BetSelection) ProtoMessage()               {}
func (*BetSelection) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *BetSelection) GetEventType() string {
	if m != nil && m.EventType != nil {
		return *m.EventType
	}
	return ""
}

func (m *BetSelection) GetEventId() uint32 {
	if m != nil && m.EventId != nil {
		return *m.EventId
	}
	return 0
}

func (m *BetSelection) GetMarketId() uint32 {
	if m != nil && m.MarketId != nil {
		return *m.MarketId
	}
	return 0
}

func (m *BetSelection) GetOutcome() string {
	if m != nil && m.Outcome != nil {
		return *m.Outcome
	}
	return ""
}

func (m *BetSelection) GetSpecialBetValue() string {
	if m != nil && m.SpecialBetValue != nil {
		return *m.SpecialBetValue
	}
	return ""
}

func (m *BetSelection) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *BetSelection) GetOdds() float64 {
	if m != nil && m.Odds != nil {
		return *m.Odds
	}
	return 0
}

func (m *BetSelection) GetStake() float64 {
	if m != nil && m.Stake != nil {
		return *m.Stake
	}
	return 0
}

func (m *BetSelection) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *BetSelection) GetBetId() uint32 {
	if m != nil && m.BetId != nil {
		return *m.BetId
	}
	return 0
}

func (m *BetSelection) GetIsLay() bool {
	if m != nil && m.IsLay != nil {
		return *m.IsLay
	}
	return false
}

func (m *BetSelection) GetBackersOdds() float64 {
	if m != nil && m.BackersOdds != nil {
		return *m.BackersOdds
	}
	return 0
}

func (m *BetSelection) GetBackersStake() float64 {
	if m != nil && m.BackersStake != nil {
		return *m.BackersStake
	}
	return 0
}

type SystemBet struct {
	// system
	// 0: parlay
	// 1: singles
	// 2: doubles
	// 3: trebles and so on
	System           *uint32  `protobuf:"varint,1,opt,name=system" json:"system,omitempty"`
	Stake            *float64 `protobuf:"fixed64,2,opt,name=stake" json:"stake,omitempty"`
	Flags            *uint32  `protobuf:"varint,3,opt,name=flags" json:"flags,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SystemBet) Reset()                    { *m = SystemBet{} }
func (m *SystemBet) String() string            { return proto.CompactTextString(m) }
func (*SystemBet) ProtoMessage()               {}
func (*SystemBet) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SystemBet) GetSystem() uint32 {
	if m != nil && m.System != nil {
		return *m.System
	}
	return 0
}

func (m *SystemBet) GetStake() float64 {
	if m != nil && m.Stake != nil {
		return *m.Stake
	}
	return 0
}

func (m *SystemBet) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

type BetRequest struct {
	Selections       []*BetSelection `protobuf:"bytes,1,rep,name=selections" json:"selections,omitempty"`
	Systems          []*SystemBet    `protobuf:"bytes,2,rep,name=systems" json:"systems,omitempty"`
	AcceptBetterOdds *bool           `protobuf:"varint,3,opt,name=accept_better_odds,json=acceptBetterOdds" json:"accept_better_odds,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *BetRequest) Reset()                    { *m = BetRequest{} }
func (m *BetRequest) String() string            { return proto.CompactTextString(m) }
func (*BetRequest) ProtoMessage()               {}
func (*BetRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *BetRequest) GetSelections() []*BetSelection {
	if m != nil {
		return m.Selections
	}
	return nil
}

func (m *BetRequest) GetSystems() []*SystemBet {
	if m != nil {
		return m.Systems
	}
	return nil
}

func (m *BetRequest) GetAcceptBetterOdds() bool {
	if m != nil && m.AcceptBetterOdds != nil {
		return *m.AcceptBetterOdds
	}
	return false
}

type BetResponse struct {
	Result     *BetResponse_ResultCode `protobuf:"varint,1,opt,name=result,enum=sportsbook.BetResponse_ResultCode" json:"result,omitempty"`
	Selections []*BetSelection         `protobuf:"bytes,3,rep,name=selections" json:"selections,omitempty"`
	Systems    []*SystemBet            `protobuf:"bytes,4,rep,name=systems" json:"systems,omitempty"`
	BetSlipId  *string                 `protobuf:"bytes,5,opt,name=bet_slip_id,json=betSlipId" json:"bet_slip_id,omitempty"`
	// deprecated, don't use
	ErrorCode        *string `protobuf:"bytes,2,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BetResponse) Reset()                    { *m = BetResponse{} }
func (m *BetResponse) String() string            { return proto.CompactTextString(m) }
func (*BetResponse) ProtoMessage()               {}
func (*BetResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *BetResponse) GetResult() BetResponse_ResultCode {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return BetResponse_OK
}

func (m *BetResponse) GetSelections() []*BetSelection {
	if m != nil {
		return m.Selections
	}
	return nil
}

func (m *BetResponse) GetSystems() []*SystemBet {
	if m != nil {
		return m.Systems
	}
	return nil
}

func (m *BetResponse) GetBetSlipId() string {
	if m != nil && m.BetSlipId != nil {
		return *m.BetSlipId
	}
	return ""
}

func (m *BetResponse) GetErrorCode() string {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return ""
}

// added for pinnacle fetch odds
type PriceUpdate struct {
	Selections       []*BetSelection          `protobuf:"bytes,1,rep,name=selections" json:"selections,omitempty"`
	AccountType      *PriceUpdate_AccountType `protobuf:"varint,2,opt,name=account_type,json=accountType,enum=sportsbook.PriceUpdate_AccountType" json:"account_type,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *PriceUpdate) Reset()                    { *m = PriceUpdate{} }
func (m *PriceUpdate) String() string            { return proto.CompactTextString(m) }
func (*PriceUpdate) ProtoMessage()               {}
func (*PriceUpdate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *PriceUpdate) GetSelections() []*BetSelection {
	if m != nil {
		return m.Selections
	}
	return nil
}

func (m *PriceUpdate) GetAccountType() PriceUpdate_AccountType {
	if m != nil && m.AccountType != nil {
		return *m.AccountType
	}
	return PriceUpdate_SHARP
}

// send over pusher
type BetAcceptanceUpdate struct {
	Result           *BetAcceptanceUpdate_ResultCode `protobuf:"varint,1,opt,name=result,enum=sportsbook.BetAcceptanceUpdate_ResultCode" json:"result,omitempty"`
	BetId            *uint32                         `protobuf:"varint,2,opt,name=bet_id,json=betId" json:"bet_id,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *BetAcceptanceUpdate) Reset()                    { *m = BetAcceptanceUpdate{} }
func (m *BetAcceptanceUpdate) String() string            { return proto.CompactTextString(m) }
func (*BetAcceptanceUpdate) ProtoMessage()               {}
func (*BetAcceptanceUpdate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *BetAcceptanceUpdate) GetResult() BetAcceptanceUpdate_ResultCode {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return BetAcceptanceUpdate_ACCEPTED
}

func (m *BetAcceptanceUpdate) GetBetId() uint32 {
	if m != nil && m.BetId != nil {
		return *m.BetId
	}
	return 0
}

func init() {
	proto.RegisterType((*BetSelection)(nil), "sportsbook.BetSelection")
	proto.RegisterType((*SystemBet)(nil), "sportsbook.SystemBet")
	proto.RegisterType((*BetRequest)(nil), "sportsbook.BetRequest")
	proto.RegisterType((*BetResponse)(nil), "sportsbook.BetResponse")
	proto.RegisterType((*PriceUpdate)(nil), "sportsbook.PriceUpdate")
	proto.RegisterType((*BetAcceptanceUpdate)(nil), "sportsbook.BetAcceptanceUpdate")
	proto.RegisterEnum("sportsbook.BetSelection_Flags", BetSelection_Flags_name, BetSelection_Flags_value)
	proto.RegisterEnum("sportsbook.SystemBet_Flags", SystemBet_Flags_name, SystemBet_Flags_value)
	proto.RegisterEnum("sportsbook.BetResponse_ResultCode", BetResponse_ResultCode_name, BetResponse_ResultCode_value)
	proto.RegisterEnum("sportsbook.PriceUpdate_AccountType", PriceUpdate_AccountType_name, PriceUpdate_AccountType_value)
	proto.RegisterEnum("sportsbook.BetAcceptanceUpdate_ResultCode", BetAcceptanceUpdate_ResultCode_name, BetAcceptanceUpdate_ResultCode_value)
}

func init() { proto.RegisterFile("sportsbook/betting.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1023 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcf, 0x6e, 0xdb, 0xc6,
	0x13, 0x36, 0xf5, 0x5f, 0x23, 0x29, 0x59, 0x6f, 0x2c, 0x87, 0xce, 0x2f, 0x09, 0xf4, 0x63, 0x82,
	0x42, 0x30, 0x0a, 0x05, 0xc8, 0xa9, 0xe8, 0x8d, 0xa4, 0x56, 0x36, 0x6b, 0x89, 0x54, 0x97, 0x94,
	0x1a, 0x9f, 0x16, 0x14, 0xb5, 0x35, 0x08, 0x4b, 0x26, 0x2b, 0xae, 0x03, 0xf8, 0x4d, 0x8a, 0xde,
	0xfb, 0x08, 0x7d, 0x88, 0x9e, 0xfb, 0x02, 0x7d, 0x94, 0x62, 0x97, 0x96, 0x45, 0x0b, 0x0e, 0x50,
	0xb4, 0xbd, 0x71, 0xbe, 0xf9, 0x86, 0xf3, 0xcd, 0xce, 0xcc, 0x2e, 0xe8, 0x59, 0x9a, 0x6c, 0x44,
	0xb6, 0x48, 0x92, 0xeb, 0x0f, 0x0b, 0x2e, 0x44, 0x7c, 0x73, 0x35, 0x48, 0x37, 0x89, 0x48, 0x30,
	0xec, 0x3c, 0xc6, 0x9f, 0x65, 0x68, 0x5b, 0x5c, 0xf8, 0x7c, 0xc5, 0x23, 0x11, 0x27, 0x37, 0xf8,
	0x0d, 0x00, 0xff, 0xcc, 0x6f, 0x04, 0x13, 0x77, 0x29, 0xd7, 0xb5, 0x9e, 0xd6, 0x6f, 0xd2, 0xa6,
	0x42, 0x82, 0xbb, 0x94, 0xe3, 0x13, 0x68, 0xe4, 0xee, 0x78, 0xa9, 0x97, 0x7a, 0x5a, 0xbf, 0x43,
	0xeb, 0xca, 0x76, 0x96, 0xf8, 0x7f, 0xd0, 0x5c, 0x87, 0x9b, 0x6b, 0xae, 0x7c, 0x65, 0xe5, 0x6b,
	0xe4, 0x80, 0xb3, 0xc4, 0x3a, 0xd4, 0x93, 0x5b, 0x11, 0x25, 0x6b, 0xae, 0x57, 0xd4, 0x3f, 0xb7,
	0x26, 0x3e, 0x85, 0xc3, 0x2c, 0xe5, 0x51, 0x1c, 0xae, 0xd8, 0x82, 0x0b, 0xf6, 0x39, 0x5c, 0xdd,
	0x72, 0xbd, 0xaa, 0x38, 0xcf, 0xef, 0x1d, 0x16, 0x17, 0x73, 0x09, 0xe3, 0xd7, 0xd0, 0x14, 0xf1,
	0x9a, 0x67, 0x22, 0x5c, 0xa7, 0x7a, 0xad, 0xa7, 0xf5, 0xcb, 0x74, 0x07, 0x60, 0x0c, 0x95, 0x64,
	0xb9, 0xcc, 0xf4, 0x7a, 0x4f, 0xeb, 0x6b, 0x54, 0x7d, 0xe3, 0x23, 0xa8, 0x66, 0x22, 0xbc, 0xe6,
	0x7a, 0x43, 0x81, 0xb9, 0x21, 0xd1, 0x1f, 0x57, 0xe1, 0x55, 0xa6, 0x37, 0x95, 0xcc, 0xdc, 0xc0,
	0x5d, 0xa8, 0x2d, 0x72, 0xf5, 0x90, 0xc3, 0x0b, 0x25, 0xbd, 0x0b, 0xb5, 0x38, 0x63, 0xab, 0xf0,
	0x4e, 0x6f, 0xf5, 0xb4, 0x7e, 0x83, 0x56, 0xe3, 0x6c, 0x1c, 0xde, 0xe1, 0xff, 0x43, 0x7b, 0x11,
	0x46, 0xd7, 0x7c, 0x93, 0x31, 0x95, 0xb5, 0xad, 0x12, 0xb4, 0xee, 0x31, 0x4f, 0x26, 0x7f, 0x07,
	0x9d, 0x2d, 0x25, 0x17, 0xd1, 0x51, 0x9c, 0x6d, 0x9c, 0x2f, 0x31, 0x83, 0x41, 0x75, 0xa4, 0xd2,
	0x37, 0xa0, 0xe2, 0x7a, 0x2e, 0x41, 0x07, 0x58, 0x87, 0x23, 0x8b, 0x04, 0x01, 0xa1, 0xcc, 0x1b,
	0x0e, 0x7d, 0x66, 0xda, 0x36, 0x99, 0x06, 0x64, 0x88, 0x34, 0xdc, 0x86, 0xc6, 0x83, 0x55, 0x91,
	0x16, 0x25, 0xdf, 0x11, 0x5b, 0x5a, 0x0d, 0x7c, 0x0c, 0x78, 0x4a, 0xdc, 0xa1, 0xe3, 0x9e, 0xdd,
	0x47, 0x98, 0xae, 0x4d, 0x50, 0xc9, 0xf8, 0x59, 0x83, 0xa6, 0x7f, 0x97, 0x09, 0xbe, 0xb6, 0xb8,
	0xc0, 0xc7, 0x50, 0xcb, 0x94, 0xa1, 0x7a, 0xdb, 0xa1, 0xf7, 0xd6, 0xee, 0xa0, 0x4a, 0x4f, 0x1e,
	0x54, 0xb9, 0x70, 0x50, 0xc6, 0xd9, 0x7f, 0x24, 0xd9, 0xf8, 0x55, 0x03, 0xb0, 0xb8, 0xa0, 0xfc,
	0xa7, 0x5b, 0x9e, 0x09, 0xfc, 0x0d, 0x40, 0xb6, 0x1d, 0xc4, 0x4c, 0xd7, 0x7a, 0xe5, 0x7e, 0xeb,
	0xa3, 0x3e, 0xd8, 0x4d, 0xeb, 0xa0, 0x38, 0xa9, 0xb4, 0xc0, 0xc5, 0x1f, 0xa0, 0x9e, 0xd7, 0x91,
	0xe9, 0x25, 0x15, 0xd6, 0x2d, 0x86, 0x3d, 0x54, 0x4f, 0xb7, 0x2c, 0xfc, 0x35, 0xe0, 0x30, 0x8a,
	0x78, 0x2a, 0xe4, 0xd0, 0x09, 0xbe, 0xc9, 0x7b, 0x58, 0x56, 0x0d, 0x46, 0xb9, 0xc7, 0x52, 0x0e,
	0xd9, 0x48, 0xe3, 0xf7, 0x1a, 0xb4, 0x94, 0xce, 0x2c, 0x4d, 0x6e, 0x32, 0x8e, 0xbf, 0x85, 0xda,
	0x86, 0x67, 0xb7, 0x2b, 0xa1, 0x0e, 0xf1, 0xd9, 0x47, 0x63, 0x4f, 0xe4, 0x96, 0x38, 0xa0, 0x8a,
	0x65, 0x27, 0x4b, 0x4e, 0xef, 0x23, 0xf6, 0x8a, 0x2c, 0xff, 0xb3, 0x22, 0x2b, 0x7f, 0xab, 0xc8,
	0xb7, 0xd0, 0x92, 0x03, 0x9d, 0xad, 0xe2, 0x54, 0x4e, 0x75, 0xbe, 0x54, 0xcd, 0x05, 0x17, 0xfe,
	0x2a, 0x4e, 0x9d, 0xa5, 0xda, 0xf5, 0xcd, 0x26, 0xd9, 0xb0, 0x28, 0x59, 0xe6, 0x8d, 0x97, 0xbb,
	0x2e, 0x11, 0xa9, 0xd8, 0xf8, 0xa3, 0x02, 0xb0, 0x2b, 0x00, 0xd7, 0xa0, 0xe4, 0x5d, 0xa0, 0x03,
	0x7c, 0x08, 0x1d, 0xd7, 0x0b, 0x98, 0xef, 0x9c, 0xb9, 0x64, 0xc8, 0x1c, 0x17, 0x69, 0x72, 0xf4,
	0x1c, 0xd7, 0x9f, 0x8d, 0x46, 0x8e, 0xed, 0x10, 0x37, 0x60, 0xa3, 0x99, 0x3b, 0xf4, 0x51, 0x09,
	0x77, 0xe1, 0xd0, 0x71, 0xe7, 0xe6, 0xd8, 0x19, 0x32, 0x9f, 0x8c, 0x89, 0x1d, 0x38, 0x9e, 0x8b,
	0xca, 0x18, 0x41, 0x7b, 0x0b, 0xcb, 0x69, 0x41, 0x15, 0xf9, 0xcf, 0x07, 0x62, 0x60, 0x5e, 0x10,
	0x54, 0xc5, 0x27, 0xd0, 0xb5, 0x3d, 0x4a, 0xc9, 0xd8, 0x0c, 0x48, 0x21, 0xdc, 0x47, 0x35, 0x99,
	0x6e, 0x62, 0x7e, 0xca, 0x99, 0x8c, 0x7c, 0xb2, 0x09, 0x19, 0x92, 0x21, 0xaa, 0xcb, 0x74, 0xfe,
	0xd4, 0xa3, 0x81, 0x6f, 0x79, 0xde, 0x05, 0x1b, 0x7b, 0xf6, 0x85, 0x5a, 0x0c, 0x1d, 0x8e, 0xa6,
	0x94, 0x4c, 0xcc, 0xc0, 0x3e, 0x67, 0x16, 0x09, 0xfc, 0xad, 0xa7, 0x89, 0x8f, 0x00, 0x8d, 0x9d,
	0x39, 0x79, 0x84, 0x82, 0xcc, 0xec, 0xb8, 0x01, 0xa1, 0xae, 0x39, 0x66, 0x3e, 0xa1, 0x73, 0x42,
	0x19, 0xa1, 0xd4, 0xa3, 0xa8, 0x85, 0x5b, 0x50, 0x0f, 0x9c, 0x09, 0xf1, 0x66, 0x01, 0x6a, 0xe3,
	0x37, 0x70, 0x62, 0xcf, 0x26, 0xb3, 0xb1, 0x19, 0xc8, 0x7f, 0xec, 0xa9, 0xe9, 0xe0, 0x97, 0xf0,
	0x22, 0xf0, 0x3c, 0x36, 0x31, 0xdd, 0xcb, 0xa2, 0xfc, 0x67, 0xf8, 0x2d, 0xbc, 0x2a, 0xea, 0x61,
	0x9e, 0xcb, 0x94, 0x0a, 0x32, 0x27, 0x6e, 0x80, 0xd0, 0x17, 0x16, 0xf9, 0x50, 0x1e, 0x92, 0x77,
	0xc1, 0x7e, 0x70, 0x82, 0x73, 0x36, 0x1a, 0x9b, 0x67, 0x3e, 0x7a, 0xa1, 0xa8, 0x26, 0x1d, 0x9b,
	0x97, 0x8c, 0x12, 0x3f, 0xa0, 0x4e, 0x7e, 0xc2, 0x47, 0xf8, 0x19, 0xc0, 0x16, 0x20, 0x43, 0xd4,
	0x95, 0x25, 0xed, 0x6c, 0x66, 0x5d, 0x32, 0x6f, 0x4e, 0x64, 0x18, 0xd2, 0xbf, 0xb8, 0xb9, 0xcf,
	0xf1, 0x29, 0x7c, 0xf5, 0x54, 0x07, 0x72, 0x05, 0xa6, 0xad, 0xaa, 0xb6, 0x3c, 0x77, 0xe6, 0xa3,
	0x63, 0xdc, 0x87, 0xf7, 0x13, 0xc7, 0x65, 0x53, 0xea, 0xd8, 0x84, 0x51, 0xf2, 0xfd, 0xcc, 0xa1,
	0x64, 0xf8, 0x04, 0xf3, 0xa5, 0xf1, 0x9b, 0x06, 0xad, 0xe9, 0x26, 0x8e, 0xf8, 0x2c, 0x5d, 0x86,
	0x82, 0xff, 0x8b, 0xa5, 0x1f, 0x41, 0x3b, 0x8c, 0xa2, 0xe4, 0x76, 0xfb, 0x58, 0x95, 0xd4, 0x2e,
	0xbe, 0x2b, 0xc6, 0x16, 0x12, 0x0d, 0xcc, 0x9c, 0x2b, 0x9f, 0x31, 0xda, 0x0a, 0x77, 0x86, 0xf1,
	0x1e, 0x5a, 0x05, 0x1f, 0x6e, 0x42, 0xd5, 0x3f, 0x37, 0xe9, 0x14, 0x1d, 0x60, 0x80, 0x9a, 0xeb,
	0xd1, 0x89, 0x39, 0x46, 0x9a, 0xf1, 0x8b, 0x06, 0x2f, 0x2c, 0x2e, 0x4c, 0x75, 0x37, 0x84, 0x37,
	0x0f, 0xfa, 0xad, 0xbd, 0xbb, 0xe0, 0x74, 0x4f, 0xfb, 0x7e, 0xc0, 0x53, 0x77, 0xc2, 0xee, 0xe5,
	0x29, 0x15, 0x5e, 0x1e, 0xa3, 0xff, 0x68, 0xff, 0x8a, 0x17, 0xe9, 0xc1, 0xa3, 0x8b, 0x54, 0xb3,
	0x5e, 0xc3, 0xab, 0x28, 0x59, 0x0f, 0x16, 0xb1, 0xb8, 0x0a, 0xd7, 0x7c, 0x15, 0x2e, 0xb2, 0x82,
	0x8a, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x4d, 0x32, 0xfb, 0x0b, 0x08, 0x00, 0x00,
}