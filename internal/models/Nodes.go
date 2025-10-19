// Package models
package models

type Node struct {
	User
	Next *Node
}
