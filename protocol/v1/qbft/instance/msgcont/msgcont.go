package msgcont

import (
	"github.com/bloxapp/ssv/protocol/v1/message"
)

// MessageContainer represents the behavior of the message container
type MessageContainer interface {
	// ReadOnlyMessagesByRound returns messages by the given round
	ReadOnlyMessagesByRound(round uint64) []*message.SignedMessage

	// QuorumAchieved returns true if enough msgs were received (round, value)
	QuorumAchieved(round uint64, value []byte) (bool, []*message.SignedMessage)

	PartialChangeRoundQuorum(stateRound uint64) (found bool, lowestChangeRound uint64)

	// AddMessage adds the given message to the container
	AddMessage(msg *message.SignedMessage)

	// OverrideMessages will override all current msgs in container with the provided msg
	OverrideMessages(msg *message.SignedMessage)
}
