package routing

import (
	"context"

	"github.com/go-redis/redis/v8"
	livekit "github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"google.golang.org/protobuf/proto"

	"github.com/livekit/livekit-server/pkg/config"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// MessageSink is an abstraction for writing protobuf messages and having them read by a MessageSource,
// potentially on a different node via a transport
//counterfeiter:generate . MessageSink
type MessageSink interface {
	WriteMessage(msg proto.Message) error
	Close()
	OnClose(f func())
}

//counterfeiter:generate . MessageSource
type MessageSource interface {
	// ReadChan exposes a one way channel to make it easier to use with select
	ReadChan() <-chan proto.Message
}

type ParticipantInit struct {
	Identity      string
	Metadata      string
	Reconnect     bool
	Permission    *livekit.ParticipantPermission
	AutoSubscribe bool
	Hidden        bool
	Recorder      bool
	Client        *livekit.ClientInfo
}

type NewParticipantCallback func(ctx context.Context, roomName string, pi ParticipantInit, requestSource MessageSource, responseSink MessageSink)
type RTCMessageCallback func(ctx context.Context, roomName, identity string, msg *livekit.RTCNodeMessage)

// Router allows multiple nodes to coordinate the participant session
//counterfeiter:generate . Router
type Router interface {
	MessageRouter

	RegisterNode() error
	UnregisterNode() error
	RemoveDeadNodes() error

	ListNodes() ([]*livekit.Node, error)

	GetNodeForRoom(ctx context.Context, roomName string) (*livekit.Node, error)
	SetNodeForRoom(ctx context.Context, roomName, nodeId string) error
	ClearRoomState(ctx context.Context, roomName string) error

	Start() error
	Drain()
	Stop()
}

type MessageRouter interface {
	// StartParticipantSignal participant signal connection is ready to start
	StartParticipantSignal(ctx context.Context, roomName string, pi ParticipantInit) (connectionId string, reqSink MessageSink, resSource MessageSource, err error)

	// Write a message to a participant, room, or node
	WriteParticipantRTC(ctx context.Context, roomName, identity string, msg *livekit.RTCNodeMessage) error
	WriteRoomRTC(ctx context.Context, roomName, identity string, msg *livekit.RTCNodeMessage) error
	WriteNodeRTC(ctx context.Context, nodeID string, msg *livekit.RTCNodeMessage) error

	// OnNewParticipantRTC is called to start a new participant's RTC connection
	OnNewParticipantRTC(callback NewParticipantCallback)

	// OnRTCMessage is called to execute actions on the RTC node
	OnRTCMessage(callback RTCMessageCallback)
}

func CreateRouter(conf *config.Config, rc *redis.Client, node LocalNode) Router {
	if rc != nil {
		return NewRedisRouter(node, rc)
	}

	// local routing and store
	logger.Infow("using single-node routing")
	return NewLocalRouter(node)
}
