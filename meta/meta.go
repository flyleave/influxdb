package meta

import (
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/influxdb/influxdb/influxql"
	"github.com/influxdb/influxdb/meta/internal"
)

//go:generate flatc -g meta.fbs

// Data represents the top level collection of all metadata.
type Data struct {
	Version   uint64 // autoincrementing version
	Nodes     []NodeInfo
	Databases []DatabaseInfo
	Users     []UserInfo
}

// Authenticate returns an authenticated user by username.
func (d *Data) Authenticate(username, password string) (*UserInfo, error) {
	panic("not yet implemented")
}

// Authorize user to execute query against database.
// Database can be "" for queries that do not require a database.
// If u is nil, this means authorization is disabled.
func (d *Data) Authorize(u *UserInfo, q *influxql.Query, database string) error {
	panic("not yet implemented")
}

// Node returns a node by id.
func (d *Data) Node(id uint64) *NodeInfo {
	panic("not yet implemented")
}

// NodeByHost returns a node by host.
func (d *Data) NodeByHost(host string) *NodeInfo {
	panic("not yet implemented")
}

// NodesByID returns a list of nodes by a set of ids.
func (d *Data) NodesByID(ids []uint64) []*NodeInfo {
	panic("not yet implemented")
}

// Database returns a database by name.
func (d *Data) Database(name string) *DatabaseInfo {
	panic("not yet implemented")
}

// User returns a User by name.
func (d *Data) User(name string) *UserInfo {
	panic("not yet implemented")
}

// NodeInfo represents information about a single node in the cluster.
type NodeInfo struct {
	ID   uint64
	Host string
}

// MarshalBinary encodes the object to a binary format.
func (info *NodeInfo) MarshalBinary() ([]byte, error) {
	var pb internal.NodeInfo
	pb.ID = &info.ID
	pb.Host = &info.Host
	return proto.Marshal(&pb)
}

// MarshalBinary decodes the object from a binary format.
func (info *NodeInfo) UnmarshalBinary(buf []byte) error {
	var pb internal.NodeInfo
	if err := proto.Unmarshal(buf, &pb); err != nil {
		return err
	}
	info.ID = pb.GetID()
	info.Host = pb.GetHost()
	return nil
}

// DatabaseInfo represents information about a database in the system.
type DatabaseInfo struct {
	Name                   string
	DefaultRetentionPolicy string
	Polices                []RetentionPolicyInfo
	ContinuousQueries      []ContinuousQueryInfo
}

// RetentionPolicy returns a policy on the database by name.
func (db *DatabaseInfo) RetentionPolicy(name string) *RetentionPolicyInfo {
	panic("not yet implemented")
}

// RetentionPolicyInfo represents metadata about a retention policy.
type RetentionPolicyInfo struct {
	Name               string
	ReplicaN           int
	Duration           time.Duration
	ShardGroupDuration time.Duration
	ShardGroups        []ShardGroupInfo
}

// ShardGroupInfo represents metadata about a shard group.
type ShardGroupInfo struct {
	ID        uint64
	StartTime time.Time
	EndTime   time.Time
	Shards    []ShardInfo
}

// ShardInfo represents metadata about a shard.
type ShardInfo struct {
	ID       uint64
	OwnerIDs []uint64
}

// ContinuousQueryInfo represents metadata about a continuous query.
type ContinuousQueryInfo struct {
	Query string
}

// UserInfo represents metadata about a user in the system.
type UserInfo struct {
	Name       string
	Hash       string
	Admin      bool
	Privileges map[string]influxql.Privilege
}
