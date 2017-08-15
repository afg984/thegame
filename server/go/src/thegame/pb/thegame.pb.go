// Code generated by protoc-gen-go. DO NOT EDIT.
// source: thegame.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	thegame.proto

It has these top-level messages:
	Controls
	GameState
	Entity
	Polygon
	Bullet
	Hero
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Ability int32

const (
	Ability_health_regen       Ability = 0
	Ability_max_health         Ability = 1
	Ability_body_damage        Ability = 2
	Ability_bullet_speed       Ability = 3
	Ability_bullet_penetration Ability = 4
	Ability_bullet_damage      Ability = 5
	Ability_bullet_reload      Ability = 6
	Ability_movement_speed     Ability = 7
)

var Ability_name = map[int32]string{
	0: "health_regen",
	1: "max_health",
	2: "body_damage",
	3: "bullet_speed",
	4: "bullet_penetration",
	5: "bullet_damage",
	6: "bullet_reload",
	7: "movement_speed",
}
var Ability_value = map[string]int32{
	"health_regen":       0,
	"max_health":         1,
	"body_damage":        2,
	"bullet_speed":       3,
	"bullet_penetration": 4,
	"bullet_damage":      5,
	"bullet_reload":      6,
	"movement_speed":     7,
}

func (x Ability) String() string {
	return proto.EnumName(Ability_name, int32(x))
}
func (Ability) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Controls struct {
	Accelerate            bool      `protobuf:"varint,1,opt,name=accelerate" json:"accelerate,omitempty"`
	AccelerationDirection float64   `protobuf:"fixed64,2,opt,name=acceleration_direction,json=accelerationDirection" json:"acceleration_direction,omitempty"`
	Shoot                 bool      `protobuf:"varint,3,opt,name=shoot" json:"shoot,omitempty"`
	ShootDirection        float64   `protobuf:"fixed64,4,opt,name=shoot_direction,json=shootDirection" json:"shoot_direction,omitempty"`
	LevelUp               []Ability `protobuf:"varint,5,rep,packed,name=level_up,json=levelUp,enum=Ability" json:"level_up,omitempty"`
	Name                  string    `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
}

func (m *Controls) Reset()                    { *m = Controls{} }
func (m *Controls) String() string            { return proto.CompactTextString(m) }
func (*Controls) ProtoMessage()               {}
func (*Controls) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Controls) GetAccelerate() bool {
	if m != nil {
		return m.Accelerate
	}
	return false
}

func (m *Controls) GetAccelerationDirection() float64 {
	if m != nil {
		return m.AccelerationDirection
	}
	return 0
}

func (m *Controls) GetShoot() bool {
	if m != nil {
		return m.Shoot
	}
	return false
}

func (m *Controls) GetShootDirection() float64 {
	if m != nil {
		return m.ShootDirection
	}
	return 0
}

func (m *Controls) GetLevelUp() []Ability {
	if m != nil {
		return m.LevelUp
	}
	return nil
}

func (m *Controls) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GameState struct {
	Meta     *GameState_Meta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Polygons []*Polygon      `protobuf:"bytes,2,rep,name=polygons" json:"polygons,omitempty"`
	Bullets  []*Bullet       `protobuf:"bytes,3,rep,name=bullets" json:"bullets,omitempty"`
	Heroes   []*Hero         `protobuf:"bytes,4,rep,name=heroes" json:"heroes,omitempty"`
}

func (m *GameState) Reset()                    { *m = GameState{} }
func (m *GameState) String() string            { return proto.CompactTextString(m) }
func (*GameState) ProtoMessage()               {}
func (*GameState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GameState) GetMeta() *GameState_Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GameState) GetPolygons() []*Polygon {
	if m != nil {
		return m.Polygons
	}
	return nil
}

func (m *GameState) GetBullets() []*Bullet {
	if m != nil {
		return m.Bullets
	}
	return nil
}

func (m *GameState) GetHeroes() []*Hero {
	if m != nil {
		return m.Heroes
	}
	return nil
}

type GameState_Meta struct {
	HeroId int32 `protobuf:"varint,2,opt,name=hero_id,json=heroId" json:"hero_id,omitempty"`
}

func (m *GameState_Meta) Reset()                    { *m = GameState_Meta{} }
func (m *GameState_Meta) String() string            { return proto.CompactTextString(m) }
func (*GameState_Meta) ProtoMessage()               {}
func (*GameState_Meta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *GameState_Meta) GetHeroId() int32 {
	if m != nil {
		return m.HeroId
	}
	return 0
}

type Entity struct {
	Id                  int32          `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Position            *Entity_Vector `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
	Radius              float64        `protobuf:"fixed64,3,opt,name=radius" json:"radius,omitempty"`
	Velocity            *Entity_Vector `protobuf:"bytes,4,opt,name=velocity" json:"velocity,omitempty"`
	Health              int32          `protobuf:"varint,5,opt,name=health" json:"health,omitempty"`
	BodyDamage          int32          `protobuf:"varint,6,opt,name=body_damage,json=bodyDamage" json:"body_damage,omitempty"`
	RewardingExperience int32          `protobuf:"varint,7,opt,name=rewarding_experience,json=rewardingExperience" json:"rewarding_experience,omitempty"`
	MaxHealth           int32          `protobuf:"varint,8,opt,name=max_health,json=maxHealth" json:"max_health,omitempty"`
}

func (m *Entity) Reset()                    { *m = Entity{} }
func (m *Entity) String() string            { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()               {}
func (*Entity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Entity) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Entity) GetPosition() *Entity_Vector {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Entity) GetRadius() float64 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *Entity) GetVelocity() *Entity_Vector {
	if m != nil {
		return m.Velocity
	}
	return nil
}

func (m *Entity) GetHealth() int32 {
	if m != nil {
		return m.Health
	}
	return 0
}

func (m *Entity) GetBodyDamage() int32 {
	if m != nil {
		return m.BodyDamage
	}
	return 0
}

func (m *Entity) GetRewardingExperience() int32 {
	if m != nil {
		return m.RewardingExperience
	}
	return 0
}

func (m *Entity) GetMaxHealth() int32 {
	if m != nil {
		return m.MaxHealth
	}
	return 0
}

type Entity_Vector struct {
	X float64 `protobuf:"fixed64,1,opt,name=x" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y" json:"y,omitempty"`
}

func (m *Entity_Vector) Reset()                    { *m = Entity_Vector{} }
func (m *Entity_Vector) String() string            { return proto.CompactTextString(m) }
func (*Entity_Vector) ProtoMessage()               {}
func (*Entity_Vector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *Entity_Vector) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Entity_Vector) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Polygon struct {
	Entity *Entity `protobuf:"bytes,1,opt,name=entity" json:"entity,omitempty"`
	Edges  int32   `protobuf:"varint,2,opt,name=edges" json:"edges,omitempty"`
}

func (m *Polygon) Reset()                    { *m = Polygon{} }
func (m *Polygon) String() string            { return proto.CompactTextString(m) }
func (*Polygon) ProtoMessage()               {}
func (*Polygon) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Polygon) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Polygon) GetEdges() int32 {
	if m != nil {
		return m.Edges
	}
	return 0
}

type Bullet struct {
	Entity *Entity `protobuf:"bytes,1,opt,name=entity" json:"entity,omitempty"`
	Owner  int32   `protobuf:"varint,2,opt,name=owner" json:"owner,omitempty"`
}

func (m *Bullet) Reset()                    { *m = Bullet{} }
func (m *Bullet) String() string            { return proto.CompactTextString(m) }
func (*Bullet) ProtoMessage()               {}
func (*Bullet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Bullet) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Bullet) GetOwner() int32 {
	if m != nil {
		return m.Owner
	}
	return 0
}

type Hero struct {
	Entity              *Entity `protobuf:"bytes,1,opt,name=entity" json:"entity,omitempty"`
	AbilityLevels       []int32 `protobuf:"varint,2,rep,packed,name=ability_levels,json=abilityLevels" json:"ability_levels,omitempty"`
	AbilityValues       []int32 `protobuf:"varint,3,rep,packed,name=ability_values,json=abilityValues" json:"ability_values,omitempty"`
	SkillPoints         int32   `protobuf:"varint,4,opt,name=skill_points,json=skillPoints" json:"skill_points,omitempty"`
	Orientation         float64 `protobuf:"fixed64,5,opt,name=orientation" json:"orientation,omitempty"`
	Level               int32   `protobuf:"varint,6,opt,name=level" json:"level,omitempty"`
	Score               int32   `protobuf:"varint,7,opt,name=score" json:"score,omitempty"`
	Experience          int32   `protobuf:"varint,8,opt,name=experience" json:"experience,omitempty"`
	ExperienceToLevelUp int32   `protobuf:"varint,9,opt,name=experience_to_level_up,json=experienceToLevelUp" json:"experience_to_level_up,omitempty"`
	Cooldown            int32   `protobuf:"varint,10,opt,name=cooldown" json:"cooldown,omitempty"`
	HealthRegenCooldown int32   `protobuf:"varint,11,opt,name=health_regen_cooldown,json=healthRegenCooldown" json:"health_regen_cooldown,omitempty"`
	Name                string  `protobuf:"bytes,12,opt,name=name" json:"name,omitempty"`
}

func (m *Hero) Reset()                    { *m = Hero{} }
func (m *Hero) String() string            { return proto.CompactTextString(m) }
func (*Hero) ProtoMessage()               {}
func (*Hero) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Hero) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Hero) GetAbilityLevels() []int32 {
	if m != nil {
		return m.AbilityLevels
	}
	return nil
}

func (m *Hero) GetAbilityValues() []int32 {
	if m != nil {
		return m.AbilityValues
	}
	return nil
}

func (m *Hero) GetSkillPoints() int32 {
	if m != nil {
		return m.SkillPoints
	}
	return 0
}

func (m *Hero) GetOrientation() float64 {
	if m != nil {
		return m.Orientation
	}
	return 0
}

func (m *Hero) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Hero) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Hero) GetExperience() int32 {
	if m != nil {
		return m.Experience
	}
	return 0
}

func (m *Hero) GetExperienceToLevelUp() int32 {
	if m != nil {
		return m.ExperienceToLevelUp
	}
	return 0
}

func (m *Hero) GetCooldown() int32 {
	if m != nil {
		return m.Cooldown
	}
	return 0
}

func (m *Hero) GetHealthRegenCooldown() int32 {
	if m != nil {
		return m.HealthRegenCooldown
	}
	return 0
}

func (m *Hero) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Controls)(nil), "Controls")
	proto.RegisterType((*GameState)(nil), "GameState")
	proto.RegisterType((*GameState_Meta)(nil), "GameState.Meta")
	proto.RegisterType((*Entity)(nil), "Entity")
	proto.RegisterType((*Entity_Vector)(nil), "Entity.Vector")
	proto.RegisterType((*Polygon)(nil), "Polygon")
	proto.RegisterType((*Bullet)(nil), "Bullet")
	proto.RegisterType((*Hero)(nil), "Hero")
	proto.RegisterEnum("Ability", Ability_name, Ability_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TheGame service

type TheGameClient interface {
	Game(ctx context.Context, opts ...grpc.CallOption) (TheGame_GameClient, error)
}

type theGameClient struct {
	cc *grpc.ClientConn
}

func NewTheGameClient(cc *grpc.ClientConn) TheGameClient {
	return &theGameClient{cc}
}

func (c *theGameClient) Game(ctx context.Context, opts ...grpc.CallOption) (TheGame_GameClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TheGame_serviceDesc.Streams[0], c.cc, "/TheGame/Game", opts...)
	if err != nil {
		return nil, err
	}
	x := &theGameGameClient{stream}
	return x, nil
}

type TheGame_GameClient interface {
	Send(*Controls) error
	Recv() (*GameState, error)
	grpc.ClientStream
}

type theGameGameClient struct {
	grpc.ClientStream
}

func (x *theGameGameClient) Send(m *Controls) error {
	return x.ClientStream.SendMsg(m)
}

func (x *theGameGameClient) Recv() (*GameState, error) {
	m := new(GameState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TheGame service

type TheGameServer interface {
	Game(TheGame_GameServer) error
}

func RegisterTheGameServer(s *grpc.Server, srv TheGameServer) {
	s.RegisterService(&_TheGame_serviceDesc, srv)
}

func _TheGame_Game_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TheGameServer).Game(&theGameGameServer{stream})
}

type TheGame_GameServer interface {
	Send(*GameState) error
	Recv() (*Controls, error)
	grpc.ServerStream
}

type theGameGameServer struct {
	grpc.ServerStream
}

func (x *theGameGameServer) Send(m *GameState) error {
	return x.ServerStream.SendMsg(m)
}

func (x *theGameGameServer) Recv() (*Controls, error) {
	m := new(Controls)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TheGame_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TheGame",
	HandlerType: (*TheGameServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Game",
			Handler:       _TheGame_Game_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "thegame.proto",
}

func init() { proto.RegisterFile("thegame.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdb, 0x8e, 0xdb, 0x36,
	0x10, 0x0d, 0x6d, 0xeb, 0xe2, 0xf1, 0xae, 0xd7, 0x65, 0x93, 0xad, 0xb0, 0x40, 0xba, 0x8e, 0x93,
	0xa2, 0x46, 0x1e, 0x84, 0xd6, 0x41, 0x9f, 0xdb, 0xe6, 0x82, 0xa6, 0x40, 0x0a, 0x04, 0x6c, 0x9a,
	0x87, 0xbe, 0x08, 0xb4, 0x34, 0xb0, 0x85, 0x4a, 0xa2, 0x20, 0x71, 0x77, 0xed, 0xaf, 0x69, 0x7f,
	0x63, 0x3f, 0xa5, 0x7f, 0x53, 0x70, 0x48, 0xc9, 0x2a, 0x50, 0xa0, 0x79, 0xb2, 0xce, 0x39, 0xc3,
	0xe3, 0xe1, 0xcc, 0x70, 0xe0, 0x5c, 0xef, 0x71, 0x27, 0x4b, 0x8c, 0xeb, 0x46, 0x69, 0xb5, 0xfa,
	0x9b, 0x41, 0xf8, 0x4a, 0x55, 0xba, 0x51, 0x45, 0xcb, 0xbf, 0x04, 0x90, 0x69, 0x8a, 0x05, 0x36,
	0x52, 0x63, 0xc4, 0x96, 0x6c, 0x1d, 0x8a, 0x01, 0xc3, 0xbf, 0x83, 0xcb, 0x1e, 0xe5, 0xaa, 0x4a,
	0xb2, 0xbc, 0xc1, 0xd4, 0x7c, 0x45, 0xa3, 0x25, 0x5b, 0x33, 0xf1, 0x68, 0xa8, 0xbe, 0xee, 0x44,
	0xfe, 0x10, 0xbc, 0x76, 0xaf, 0x94, 0x8e, 0xc6, 0xe4, 0x68, 0x01, 0xff, 0x1a, 0x2e, 0xe8, 0x63,
	0xe0, 0x32, 0x21, 0x97, 0x39, 0xd1, 0xa7, 0xe3, 0x4f, 0x21, 0x2c, 0xf0, 0x16, 0x8b, 0xe4, 0xa6,
	0x8e, 0xbc, 0xe5, 0x78, 0x3d, 0xdf, 0x84, 0xf1, 0x8f, 0xdb, 0xbc, 0xc8, 0xf5, 0x51, 0x04, 0xa4,
	0xfc, 0x56, 0x73, 0x0e, 0x93, 0x4a, 0x96, 0x18, 0xf9, 0x4b, 0xb6, 0x9e, 0x0a, 0xfa, 0x5e, 0xdd,
	0x33, 0x98, 0xfe, 0x24, 0x4b, 0xfc, 0x55, 0x9b, 0xe4, 0x9f, 0xc2, 0xa4, 0x44, 0x2d, 0xe9, 0x5a,
	0xb3, 0xcd, 0x45, 0xdc, 0x2b, 0xf1, 0x2f, 0xa8, 0xa5, 0x20, 0x91, 0x3f, 0x83, 0xb0, 0x56, 0xc5,
	0x71, 0xa7, 0xaa, 0x36, 0x1a, 0x2d, 0xc7, 0xeb, 0xd9, 0x26, 0x8c, 0xdf, 0x5b, 0x42, 0xf4, 0x0a,
	0x7f, 0x02, 0xc1, 0xf6, 0xa6, 0x28, 0x50, 0xb7, 0xd1, 0x98, 0x82, 0x82, 0xf8, 0x25, 0x61, 0xd1,
	0xf1, 0xfc, 0x31, 0xf8, 0x7b, 0x6c, 0x14, 0xb6, 0xd1, 0x84, 0x22, 0xbc, 0xf8, 0x2d, 0x36, 0x4a,
	0x38, 0xf2, 0xea, 0x1a, 0x26, 0xe6, 0x5f, 0xf9, 0x17, 0x10, 0x18, 0x26, 0xc9, 0x33, 0x2a, 0xa1,
	0x67, 0x03, 0x7e, 0xce, 0x56, 0xf7, 0x23, 0xf0, 0xdf, 0x54, 0x3a, 0xd7, 0x47, 0x3e, 0x87, 0x51,
	0x9e, 0x51, 0xda, 0x9e, 0x18, 0xe5, 0x19, 0x7f, 0x6e, 0x72, 0x6c, 0xf3, 0xbe, 0xee, 0xb3, 0xcd,
	0x3c, 0xb6, 0xa1, 0xf1, 0x47, 0x4c, 0xb5, 0x6a, 0x44, 0xaf, 0xf3, 0x4b, 0xf0, 0x1b, 0x99, 0xe5,
	0x37, 0x2d, 0xd5, 0x9e, 0x09, 0x87, 0x8c, 0xc7, 0x2d, 0x16, 0x2a, 0xcd, 0xf5, 0x91, 0xaa, 0xfe,
	0x1f, 0x1e, 0x9d, 0x6e, 0x3c, 0xf6, 0x28, 0x0b, 0xbd, 0x8f, 0xbc, 0x2e, 0x45, 0x83, 0xf8, 0x35,
	0xcc, 0xb6, 0x2a, 0x3b, 0x26, 0x99, 0x2c, 0xe5, 0xce, 0x56, 0xde, 0x13, 0x60, 0xa8, 0xd7, 0xc4,
	0xf0, 0x6f, 0xe1, 0x61, 0x83, 0x77, 0xb2, 0xc9, 0xf2, 0x6a, 0x97, 0xe0, 0xa1, 0xc6, 0x26, 0xc7,
	0x2a, 0xc5, 0x28, 0xa0, 0xc8, 0xcf, 0x7b, 0xed, 0x4d, 0x2f, 0xf1, 0xc7, 0x00, 0xa5, 0x3c, 0x24,
	0xee, 0xff, 0x42, 0x0a, 0x9c, 0x96, 0xf2, 0xf0, 0x96, 0x88, 0xab, 0x67, 0xe0, 0xdb, 0xf4, 0xf8,
	0x19, 0xb0, 0x03, 0xd5, 0x84, 0x09, 0x76, 0x30, 0xe8, 0xe8, 0x66, 0x90, 0x1d, 0x57, 0x3f, 0x40,
	0xe0, 0x7a, 0xc6, 0xaf, 0xc1, 0x47, 0xba, 0x96, 0x6b, 0x7b, 0xe0, 0x6e, 0x29, 0x1c, 0x6d, 0x66,
	0x13, 0xb3, 0x1d, 0xb6, 0xae, 0xfc, 0x16, 0xac, 0xbe, 0x07, 0xdf, 0x36, 0xf4, 0x93, 0x0c, 0xd4,
	0x5d, 0x85, 0x4d, 0x67, 0x40, 0x60, 0xf5, 0xe7, 0x18, 0x26, 0xa6, 0xe1, 0xff, 0x7f, 0xfe, 0x2b,
	0x98, 0x4b, 0x3b, 0xcc, 0x09, 0xcd, 0xb2, 0x9d, 0x3b, 0x4f, 0x9c, 0x3b, 0xf6, 0x1d, 0x91, 0xc3,
	0xb0, 0x5b, 0x59, 0xdc, 0xa0, 0x9d, 0xbc, 0x53, 0xd8, 0x47, 0x22, 0xf9, 0x13, 0x38, 0x6b, 0xff,
	0xc8, 0x8b, 0x22, 0xa9, 0x55, 0x5e, 0xe9, 0x96, 0x7a, 0xeb, 0x89, 0x19, 0x71, 0xef, 0x89, 0xe2,
	0x4b, 0x98, 0x29, 0x53, 0x6d, 0x4d, 0xaf, 0x94, 0x7a, 0xca, 0xc4, 0x90, 0x32, 0x57, 0xa2, 0x54,
	0x5c, 0x4b, 0x2d, 0xa0, 0x57, 0x9c, 0xaa, 0xa6, 0x6b, 0x9f, 0x05, 0x66, 0x65, 0x0c, 0x3a, 0x6b,
	0x1b, 0x36, 0x60, 0xf8, 0x0b, 0xb8, 0x3c, 0xa1, 0x44, 0xab, 0xa4, 0x7f, 0xca, 0x53, 0x3b, 0x05,
	0x27, 0xf5, 0x83, 0x7a, 0xe7, 0x1e, 0xf3, 0x15, 0x84, 0xa9, 0x52, 0x45, 0xa6, 0xee, 0xaa, 0x08,
	0x28, 0xac, 0xc7, 0x7c, 0x03, 0x8f, 0xec, 0x74, 0x24, 0x0d, 0xee, 0xb0, 0x4a, 0xfa, 0xc0, 0x99,
	0xf5, 0xb3, 0xa2, 0x30, 0xda, 0xab, 0xee, 0x4c, 0xb7, 0x1c, 0xce, 0x4e, 0xcb, 0xe1, 0xf9, 0x5f,
	0x0c, 0x02, 0xb7, 0x45, 0xf8, 0x02, 0xce, 0x86, 0x9e, 0x8b, 0x07, 0x7c, 0x3e, 0x9c, 0xc3, 0x05,
	0xe3, 0x17, 0xff, 0x9a, 0xf5, 0xc5, 0xc8, 0x1c, 0xb1, 0x4f, 0x3d, 0x69, 0x6b, 0xc4, 0x6c, 0x31,
	0xe6, 0x97, 0xc0, 0x1d, 0x53, 0x63, 0x85, 0xda, 0x2e, 0xc1, 0xc5, 0x84, 0x7f, 0x06, 0xe7, 0x8e,
	0x77, 0x87, 0xbd, 0x01, 0xd5, 0x60, 0xa1, 0x64, 0xb6, 0xf0, 0x39, 0x87, 0x79, 0xa9, 0x6e, 0xb1,
	0xc4, 0xaa, 0x73, 0x0c, 0x36, 0x31, 0x04, 0x1f, 0xf6, 0x68, 0xf6, 0x94, 0x59, 0x5e, 0xf4, 0x3b,
	0x8d, 0xbb, 0x65, 0x7d, 0x05, 0xa7, 0x0d, 0xb6, 0x7a, 0xb0, 0x66, 0xdf, 0xb0, 0x97, 0x93, 0xdf,
	0x47, 0xf5, 0x76, 0xeb, 0xd3, 0x62, 0x7f, 0xf1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x68,
	0xe6, 0xb6, 0xe9, 0x05, 0x00, 0x00,
}
