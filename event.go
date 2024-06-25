package main

import (
	"fmt"
	"strconv"
	"strings"
)

type event struct {
	name   string
	entity entity
	format []formatField
}

type formatType string

const (
	formatTypeString  formatType = "string"
	formatTypeInt     formatType = "int"
	formatTypeBoolean formatType = "bool"
)

type formatField struct {
	name string
	typ  formatType
}

var events = []event{
	{name: "workspacev2", entity: workspace, format: []formatField{
		{name: "workspace_id", typ: formatTypeInt},
		{name: "workspace_name", typ: formatTypeString},
	}},
	{name: "createworkspacev2", entity: workspace, format: []formatField{
		{name: "workspace_id", typ: formatTypeInt},
		{name: "workspace_name", typ: formatTypeString},
	}},
	{name: "destroyworkspacev2", entity: workspace, format: []formatField{
		{name: "workspace_id", typ: formatTypeInt},
		{name: "workspace_name", typ: formatTypeString},
	}},
	{name: "moveworkspacev2", entity: workspace, format: []formatField{
		{name: "workspace_id", typ: formatTypeInt},
		{name: "monitor_name", typ: formatTypeString},
	}},
	{name: "urgent", entity: window, format: []formatField{
		{name: "window_address", typ: formatTypeString},
	}},
	{name: "activewindowv2", entity: window, format: []formatField{
		{name: "window_address", typ: formatTypeString},
	}},
	{name: "openwindow", entity: window, format: []formatField{
		{name: "window_address", typ: formatTypeString},
		{name: "workspace_name", typ: formatTypeString},
		{name: "window_class", typ: formatTypeString},
		{name: "window_title", typ: formatTypeString},
	}},
	{name: "closewindow", entity: window, format: []formatField{
		{name: "window_address", typ: formatTypeString},
	}},
	{name: "movewindowv2", entity: window, format: []formatField{
		{name: "window_address", typ: formatTypeString},
		{name: "workspace_id", typ: formatTypeInt},
		{name: "workspace_name", typ: formatTypeString},
	}},
	{name: "minimize", entity: window, format: []formatField{
		{name: "window_address", typ: formatTypeString},
		{name: "window_minimized", typ: formatTypeBoolean},
	}},
	{name: "windowtitle", entity: window, format: []formatField{
		{name: "window_address", typ: formatTypeString},
	}},
	{name: "fullscreen", entity: window, format: []formatField{
		{name: "window_fullscreened", typ: formatTypeBoolean},
	}},
	{name: "changefloatingmode", entity: window, format: []formatField{
		{name: "window_address", typ: formatTypeString},
		{name: "window_floating", typ: formatTypeBoolean},
	}},
	{name: "pin", entity: window, format: []formatField{
		{name: "window_address", typ: formatTypeString},
		{name: "window_pinned", typ: formatTypeBoolean},
	}},
	{name: "focusedmon", entity: monitor, format: []formatField{
		{name: "monitor_name", typ: formatTypeString},
		{name: "workspace_name", typ: formatTypeString},
	}},
	{name: "monitoraddedv2", entity: monitor, format: []formatField{
		{name: "monitor_id", typ: formatTypeString},
		{name: "monitor_name", typ: formatTypeString},
		{name: "monitor_description", typ: formatTypeString},
	}},
	{name: "monitorremoved", entity: monitor, format: []formatField{
		{name: "monitor_name", typ: formatTypeString},
	}},
}

func FindEvent(name string) (event, error) {
	for _, ev := range events {
		if ev.name == name {
			return ev, nil
		}
	}

	return event{}, fmt.Errorf("unknown event %q", name)
}

func ParseEvent(ev event, data string) map[string]any {
	m := make(map[string]any)
	m["event_name"] = ev.name
	for i, raw := range strings.Split(data, ",") {
		if i >= len(ev.format) {
			return m // Out of range.
		}
		formatField := ev.format[i]
		var val any
		switch formatField.typ {
		case formatTypeInt:
			v, err := strconv.Atoi(raw)
			if err != nil {
				val = raw + "#ERROR_INT"
				break
			}
			val = v
		case formatTypeString:
			val = raw
		case formatTypeBoolean:
			v, err := strconv.ParseBool(raw)
			if err != nil {
				val = raw + "#ERROR_BOOL"
				break
			}
			val = v
		}
		m[formatField.name] = val
	}

	return m
}

func (e event) HasEntity(ent entity) bool {
	return e.entity&ent != 0
}

func (e event) String() string {
	return e.name
}
