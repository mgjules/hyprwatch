package main

import (
	"fmt"
	"strings"
)

type entity byte

const (
	workspace entity = 1 << iota
	window
	monitor
	maxKey
)

func (e entity) String() string {
	if e >= maxKey {
		return fmt.Sprintf("unknown entity %d", e)
	}

	switch e {
	case workspace:
		return "workspace"
	case window:
		return "window"
	case monitor:
		return "monitor"
	}

	// multiple
	var entities []string
	for entity := workspace; entity < maxKey; entity <<= 1 {
		if e&entity != 0 {
			entities = append(entities, entity.String())
		}
	}

	return strings.Join(entities, "|")
}
