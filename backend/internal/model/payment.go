package model

type Status string

const (
	PENDING  Status = "PENDING"
	SUCCESS  Status = "SUCCESS"
	REJECTED Status = "REJECTED"
)
