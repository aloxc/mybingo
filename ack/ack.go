package ack

const (
	UNKNOWN_ACK = iota
	CREATE_DATABASE_ACK
	DROP_DATABASE_ACK
	CREATE_TABLE_ACK
	DROP_TABLE_ACK
	TRUNCATE_TABLE_ACK
	INSERT_ACK
	DELETE_ACK
	UPDATE_ACK
)

type Ack struct {
}