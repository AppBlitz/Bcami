// Package models
package models

type NodeSong struct {
	Data Song
	Next *NodeSong
}

func NewNode(data Song) *NodeSong {
	return &NodeSong{Data: data, Next: nil}
}
