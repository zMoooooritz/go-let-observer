// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: pkg/record/record.proto

package record

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MatchPlayer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id            string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	RecordId      int32                  `protobuf:"varint,3,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	Level         int32                  `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	ClanTag       string                 `protobuf:"bytes,5,opt,name=clan_tag,json=clanTag,proto3" json:"clan_tag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MatchPlayer) Reset() {
	*x = MatchPlayer{}
	mi := &file_pkg_record_record_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchPlayer) ProtoMessage() {}

func (x *MatchPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_record_record_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchPlayer.ProtoReflect.Descriptor instead.
func (*MatchPlayer) Descriptor() ([]byte, []int) {
	return file_pkg_record_record_proto_rawDescGZIP(), []int{0}
}

func (x *MatchPlayer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MatchPlayer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MatchPlayer) GetRecordId() int32 {
	if x != nil {
		return x.RecordId
	}
	return 0
}

func (x *MatchPlayer) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *MatchPlayer) GetClanTag() string {
	if x != nil {
		return x.ClanTag
	}
	return ""
}

type MatchHeader struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Version       string                  `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	MapId         string                  `protobuf:"bytes,2,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	StartTime     *timestamppb.Timestamp  `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       *timestamppb.Timestamp  `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Players       map[string]*MatchPlayer `protobuf:"bytes,5,rep,name=players,proto3" json:"players,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MatchHeader) Reset() {
	*x = MatchHeader{}
	mi := &file_pkg_record_record_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchHeader) ProtoMessage() {}

func (x *MatchHeader) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_record_record_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchHeader.ProtoReflect.Descriptor instead.
func (*MatchHeader) Descriptor() ([]byte, []int) {
	return file_pkg_record_record_proto_rawDescGZIP(), []int{1}
}

func (x *MatchHeader) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *MatchHeader) GetMapId() string {
	if x != nil {
		return x.MapId
	}
	return ""
}

func (x *MatchHeader) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *MatchHeader) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *MatchHeader) GetPlayers() map[string]*MatchPlayer {
	if x != nil {
		return x.Players
	}
	return nil
}

type PlayerState struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PlayerId      int32                  `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	X             int32                  `protobuf:"zigzag32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y             int32                  `protobuf:"zigzag32,3,opt,name=y,proto3" json:"y,omitempty"`
	Kills         int32                  `protobuf:"varint,4,opt,name=kills,proto3" json:"kills,omitempty"`
	Deaths        int32                  `protobuf:"varint,5,opt,name=deaths,proto3" json:"deaths,omitempty"`
	Team          int32                  `protobuf:"varint,6,opt,name=team,proto3" json:"team,omitempty"`
	Unit          int32                  `protobuf:"varint,7,opt,name=unit,proto3" json:"unit,omitempty"`
	Role          int32                  `protobuf:"varint,8,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlayerState) Reset() {
	*x = PlayerState{}
	mi := &file_pkg_record_record_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerState) ProtoMessage() {}

func (x *PlayerState) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_record_record_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerState.ProtoReflect.Descriptor instead.
func (*PlayerState) Descriptor() ([]byte, []int) {
	return file_pkg_record_record_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerState) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *PlayerState) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *PlayerState) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *PlayerState) GetKills() int32 {
	if x != nil {
		return x.Kills
	}
	return 0
}

func (x *PlayerState) GetDeaths() int32 {
	if x != nil {
		return x.Deaths
	}
	return 0
}

func (x *PlayerState) GetTeam() int32 {
	if x != nil {
		return x.Team
	}
	return 0
}

func (x *PlayerState) GetUnit() int32 {
	if x != nil {
		return x.Unit
	}
	return 0
}

func (x *PlayerState) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

type PlayerDelta struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PlayerId      int32                  `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	X             int32                  `protobuf:"zigzag32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y             int32                  `protobuf:"zigzag32,3,opt,name=y,proto3" json:"y,omitempty"`
	Kills         *int32                 `protobuf:"varint,4,opt,name=kills,proto3,oneof" json:"kills,omitempty"`
	Deaths        *int32                 `protobuf:"varint,5,opt,name=deaths,proto3,oneof" json:"deaths,omitempty"`
	Team          *int32                 `protobuf:"varint,6,opt,name=team,proto3,oneof" json:"team,omitempty"`
	Unit          *int32                 `protobuf:"varint,7,opt,name=unit,proto3,oneof" json:"unit,omitempty"`
	Role          *int32                 `protobuf:"varint,8,opt,name=role,proto3,oneof" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlayerDelta) Reset() {
	*x = PlayerDelta{}
	mi := &file_pkg_record_record_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerDelta) ProtoMessage() {}

func (x *PlayerDelta) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_record_record_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerDelta.ProtoReflect.Descriptor instead.
func (*PlayerDelta) Descriptor() ([]byte, []int) {
	return file_pkg_record_record_proto_rawDescGZIP(), []int{3}
}

func (x *PlayerDelta) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *PlayerDelta) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *PlayerDelta) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *PlayerDelta) GetKills() int32 {
	if x != nil && x.Kills != nil {
		return *x.Kills
	}
	return 0
}

func (x *PlayerDelta) GetDeaths() int32 {
	if x != nil && x.Deaths != nil {
		return *x.Deaths
	}
	return 0
}

func (x *PlayerDelta) GetTeam() int32 {
	if x != nil && x.Team != nil {
		return *x.Team
	}
	return 0
}

func (x *PlayerDelta) GetUnit() int32 {
	if x != nil && x.Unit != nil {
		return *x.Unit
	}
	return 0
}

func (x *PlayerDelta) GetRole() int32 {
	if x != nil && x.Role != nil {
		return *x.Role
	}
	return 0
}

type FullSnapshot struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Players       []*PlayerState         `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FullSnapshot) Reset() {
	*x = FullSnapshot{}
	mi := &file_pkg_record_record_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FullSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullSnapshot) ProtoMessage() {}

func (x *FullSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_record_record_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullSnapshot.ProtoReflect.Descriptor instead.
func (*FullSnapshot) Descriptor() ([]byte, []int) {
	return file_pkg_record_record_proto_rawDescGZIP(), []int{4}
}

func (x *FullSnapshot) GetPlayers() []*PlayerState {
	if x != nil {
		return x.Players
	}
	return nil
}

type DeltaSnapshot struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Players       []*PlayerDelta         `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeltaSnapshot) Reset() {
	*x = DeltaSnapshot{}
	mi := &file_pkg_record_record_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeltaSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaSnapshot) ProtoMessage() {}

func (x *DeltaSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_record_record_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaSnapshot.ProtoReflect.Descriptor instead.
func (*DeltaSnapshot) Descriptor() ([]byte, []int) {
	return file_pkg_record_record_proto_rawDescGZIP(), []int{5}
}

func (x *DeltaSnapshot) GetPlayers() []*PlayerDelta {
	if x != nil {
		return x.Players
	}
	return nil
}

type Snapshot struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Index     int32                  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Data:
	//
	//	*Snapshot_FullSnapshot
	//	*Snapshot_DeltaSnapshot
	Data          isSnapshot_Data `protobuf_oneof:"data"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	mi := &file_pkg_record_record_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_record_record_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_pkg_record_record_proto_rawDescGZIP(), []int{6}
}

func (x *Snapshot) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Snapshot) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Snapshot) GetData() isSnapshot_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Snapshot) GetFullSnapshot() *FullSnapshot {
	if x != nil {
		if x, ok := x.Data.(*Snapshot_FullSnapshot); ok {
			return x.FullSnapshot
		}
	}
	return nil
}

func (x *Snapshot) GetDeltaSnapshot() *DeltaSnapshot {
	if x != nil {
		if x, ok := x.Data.(*Snapshot_DeltaSnapshot); ok {
			return x.DeltaSnapshot
		}
	}
	return nil
}

type isSnapshot_Data interface {
	isSnapshot_Data()
}

type Snapshot_FullSnapshot struct {
	FullSnapshot *FullSnapshot `protobuf:"bytes,3,opt,name=full_snapshot,json=fullSnapshot,proto3,oneof"`
}

type Snapshot_DeltaSnapshot struct {
	DeltaSnapshot *DeltaSnapshot `protobuf:"bytes,4,opt,name=delta_snapshot,json=deltaSnapshot,proto3,oneof"`
}

func (*Snapshot_FullSnapshot) isSnapshot_Data() {}

func (*Snapshot_DeltaSnapshot) isSnapshot_Data() {}

type MatchData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Header        *MatchHeader           `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Snapshots     []*Snapshot            `protobuf:"bytes,2,rep,name=snapshots,proto3" json:"snapshots,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MatchData) Reset() {
	*x = MatchData{}
	mi := &file_pkg_record_record_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchData) ProtoMessage() {}

func (x *MatchData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_record_record_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchData.ProtoReflect.Descriptor instead.
func (*MatchData) Descriptor() ([]byte, []int) {
	return file_pkg_record_record_proto_rawDescGZIP(), []int{7}
}

func (x *MatchData) GetHeader() *MatchHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *MatchData) GetSnapshots() []*Snapshot {
	if x != nil {
		return x.Snapshots
	}
	return nil
}

var File_pkg_record_record_proto protoreflect.FileDescriptor

const file_pkg_record_record_proto_rawDesc = "" +
	"\n" +
	"\x17pkg/record/record.proto\x12\x06record\x1a\x1fgoogle/protobuf/timestamp.proto\"\x7f\n" +
	"\vMatchPlayer\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\tR\x02id\x12\x1b\n" +
	"\trecord_id\x18\x03 \x01(\x05R\brecordId\x12\x14\n" +
	"\x05level\x18\x04 \x01(\x05R\x05level\x12\x19\n" +
	"\bclan_tag\x18\x05 \x01(\tR\aclanTag\"\xbd\x02\n" +
	"\vMatchHeader\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion\x12\x15\n" +
	"\x06map_id\x18\x02 \x01(\tR\x05mapId\x129\n" +
	"\n" +
	"start_time\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tstartTime\x125\n" +
	"\bend_time\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\aendTime\x12:\n" +
	"\aplayers\x18\x05 \x03(\v2 .record.MatchHeader.PlayersEntryR\aplayers\x1aO\n" +
	"\fPlayersEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12)\n" +
	"\x05value\x18\x02 \x01(\v2\x13.record.MatchPlayerR\x05value:\x028\x01\"\xb0\x01\n" +
	"\vPlayerState\x12\x1b\n" +
	"\tplayer_id\x18\x01 \x01(\x05R\bplayerId\x12\f\n" +
	"\x01x\x18\x02 \x01(\x11R\x01x\x12\f\n" +
	"\x01y\x18\x03 \x01(\x11R\x01y\x12\x14\n" +
	"\x05kills\x18\x04 \x01(\x05R\x05kills\x12\x16\n" +
	"\x06deaths\x18\x05 \x01(\x05R\x06deaths\x12\x12\n" +
	"\x04team\x18\x06 \x01(\x05R\x04team\x12\x12\n" +
	"\x04unit\x18\a \x01(\x05R\x04unit\x12\x12\n" +
	"\x04role\x18\b \x01(\x05R\x04role\"\xf9\x01\n" +
	"\vPlayerDelta\x12\x1b\n" +
	"\tplayer_id\x18\x01 \x01(\x05R\bplayerId\x12\f\n" +
	"\x01x\x18\x02 \x01(\x11R\x01x\x12\f\n" +
	"\x01y\x18\x03 \x01(\x11R\x01y\x12\x19\n" +
	"\x05kills\x18\x04 \x01(\x05H\x00R\x05kills\x88\x01\x01\x12\x1b\n" +
	"\x06deaths\x18\x05 \x01(\x05H\x01R\x06deaths\x88\x01\x01\x12\x17\n" +
	"\x04team\x18\x06 \x01(\x05H\x02R\x04team\x88\x01\x01\x12\x17\n" +
	"\x04unit\x18\a \x01(\x05H\x03R\x04unit\x88\x01\x01\x12\x17\n" +
	"\x04role\x18\b \x01(\x05H\x04R\x04role\x88\x01\x01B\b\n" +
	"\x06_killsB\t\n" +
	"\a_deathsB\a\n" +
	"\x05_teamB\a\n" +
	"\x05_unitB\a\n" +
	"\x05_role\"=\n" +
	"\fFullSnapshot\x12-\n" +
	"\aplayers\x18\x01 \x03(\v2\x13.record.PlayerStateR\aplayers\">\n" +
	"\rDeltaSnapshot\x12-\n" +
	"\aplayers\x18\x01 \x03(\v2\x13.record.PlayerDeltaR\aplayers\"\xdf\x01\n" +
	"\bSnapshot\x12\x14\n" +
	"\x05index\x18\x01 \x01(\x05R\x05index\x128\n" +
	"\ttimestamp\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12;\n" +
	"\rfull_snapshot\x18\x03 \x01(\v2\x14.record.FullSnapshotH\x00R\ffullSnapshot\x12>\n" +
	"\x0edelta_snapshot\x18\x04 \x01(\v2\x15.record.DeltaSnapshotH\x00R\rdeltaSnapshotB\x06\n" +
	"\x04data\"h\n" +
	"\tMatchData\x12+\n" +
	"\x06header\x18\x01 \x01(\v2\x13.record.MatchHeaderR\x06header\x12.\n" +
	"\tsnapshots\x18\x02 \x03(\v2\x10.record.SnapshotR\tsnapshotsB6Z4github.com/zMoooooritz/go-let-observer/record;recordb\x06proto3"

var (
	file_pkg_record_record_proto_rawDescOnce sync.Once
	file_pkg_record_record_proto_rawDescData []byte
)

func file_pkg_record_record_proto_rawDescGZIP() []byte {
	file_pkg_record_record_proto_rawDescOnce.Do(func() {
		file_pkg_record_record_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_record_record_proto_rawDesc), len(file_pkg_record_record_proto_rawDesc)))
	})
	return file_pkg_record_record_proto_rawDescData
}

var file_pkg_record_record_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_record_record_proto_goTypes = []any{
	(*MatchPlayer)(nil),           // 0: record.MatchPlayer
	(*MatchHeader)(nil),           // 1: record.MatchHeader
	(*PlayerState)(nil),           // 2: record.PlayerState
	(*PlayerDelta)(nil),           // 3: record.PlayerDelta
	(*FullSnapshot)(nil),          // 4: record.FullSnapshot
	(*DeltaSnapshot)(nil),         // 5: record.DeltaSnapshot
	(*Snapshot)(nil),              // 6: record.Snapshot
	(*MatchData)(nil),             // 7: record.MatchData
	nil,                           // 8: record.MatchHeader.PlayersEntry
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_pkg_record_record_proto_depIdxs = []int32{
	9,  // 0: record.MatchHeader.start_time:type_name -> google.protobuf.Timestamp
	9,  // 1: record.MatchHeader.end_time:type_name -> google.protobuf.Timestamp
	8,  // 2: record.MatchHeader.players:type_name -> record.MatchHeader.PlayersEntry
	2,  // 3: record.FullSnapshot.players:type_name -> record.PlayerState
	3,  // 4: record.DeltaSnapshot.players:type_name -> record.PlayerDelta
	9,  // 5: record.Snapshot.timestamp:type_name -> google.protobuf.Timestamp
	4,  // 6: record.Snapshot.full_snapshot:type_name -> record.FullSnapshot
	5,  // 7: record.Snapshot.delta_snapshot:type_name -> record.DeltaSnapshot
	1,  // 8: record.MatchData.header:type_name -> record.MatchHeader
	6,  // 9: record.MatchData.snapshots:type_name -> record.Snapshot
	0,  // 10: record.MatchHeader.PlayersEntry.value:type_name -> record.MatchPlayer
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_pkg_record_record_proto_init() }
func file_pkg_record_record_proto_init() {
	if File_pkg_record_record_proto != nil {
		return
	}
	file_pkg_record_record_proto_msgTypes[3].OneofWrappers = []any{}
	file_pkg_record_record_proto_msgTypes[6].OneofWrappers = []any{
		(*Snapshot_FullSnapshot)(nil),
		(*Snapshot_DeltaSnapshot)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_record_record_proto_rawDesc), len(file_pkg_record_record_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_record_record_proto_goTypes,
		DependencyIndexes: file_pkg_record_record_proto_depIdxs,
		MessageInfos:      file_pkg_record_record_proto_msgTypes,
	}.Build()
	File_pkg_record_record_proto = out.File
	file_pkg_record_record_proto_goTypes = nil
	file_pkg_record_record_proto_depIdxs = nil
}
