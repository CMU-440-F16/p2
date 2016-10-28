// DO NOT MODIFY!

package tribserver

import "github.com/cmu440/tribbler/rpc/tribrpc"

// TribServer defines the set of methods that a TribClient can invoke remotely via RPCs.
type TribServer interface {

	// CreateUser creates a user with the specified UserID.
	// Replies with status Exists if the user has previously been created.
	CreateUser(args *tribrpc.CreateUserArgs, reply *tribrpc.CreateUserReply) error

	// AddSubscription adds TargerUserID to UserID's list of subscriptions.
	// Replies with status NoSuchUser if the specified UserID does not exist, and NoSuchTargerUser
	// if the specified TargerUserID does not exist.
	AddSubscription(args *tribrpc.SubscriptionArgs, reply *tribrpc.SubscriptionReply) error

	// RemoveSubscription removes TargerUserID to UserID's list of subscriptions.
	// Replies with status NoSuchUser if the specified UserID does not exist, and NoSuchTargerUser
	// if the specified TargerUserID does not exist.
	RemoveSubscription(args *tribrpc.SubscriptionArgs, reply *tribrpc.SubscriptionReply) error

	// GetFriends retrieves a list of friends of the given user.
	// Replies with status NoSuchUser if the specified UserID does not exist.
	GetFriends(args *tribrpc.GetFriendsArgs, reply *tribrpc.GetFriendsReply) error

	// PostTribble posts a tribble on behalf of the specified UserID. The TribServer
	// should timestamp the entry before inserting the Tribble into it's local Libstore.
	// Replies with status NoSuchUser if the specified UserID does not exist.
	PostTribble(args *tribrpc.PostTribbleArgs, reply *tribrpc.PostTribbleReply) error

	// DeleteTribble delete a tribble with the specified PostKey.
	// Replies with status NoSuchPost if the specified PostKey does not exist.
	DeleteTribble(args *tribrpc.DeleteTribbleArgs, reply *tribrpc.DeleteTribbleReply) error

	// GetTribbles retrieves a list of at most 100 tribbles posted by the specified
	// UserID in reverse chronological order (most recent first).
	// Replies with status NoSuchUser if the specified UserID does not exist.
	GetTribbles(args *tribrpc.GetTribblesArgs, reply *tribrpc.GetTribblesReply) error

	// GetTribblesBySubscription retrieves a list of at most 100 tribbles posted by
	// all users to which the specified UserID is subscribed in reverse chronological
	// order (most recent first).  Replies with status NoSuchUser if the specified UserID
	// does not exist.
	GetTribblesBySubscription(args *tribrpc.GetTribblesArgs, reply *tribrpc.GetTribblesReply) error
}
