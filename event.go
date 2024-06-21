package main

import "fmt"

type event struct {
	name   string
	entity entity
}

var events = []event{
	{name: "workspacev2", entity: workspace},
	{name: "createworkspacev2", entity: workspace},
	{name: "destroyworkspacev2", entity: workspace},
	{name: "moveworkspacev2", entity: workspace},
	{name: "urgent", entity: workspace},
	{name: "activewindowv2", entity: window},
	{name: "openwindow", entity: window},
	{name: "closewindow", entity: window},
	{name: "movewindowv2", entity: window},
	{name: "minimize", entity: window},
	{name: "windowtitle", entity: window},
	{name: "fullscreen", entity: window},
	{name: "focusedmon", entity: monitor},
	{name: "monitoraddedv2", entity: monitor},
	{name: "monitorremoved", entity: monitor},
}

func FindEvent(name string) (event, error) {
	for _, ev := range events {
		if ev.name == name {
			return ev, nil
		}
	}

	return event{}, fmt.Errorf("unknown event %q", name)
}

func (e event) HasEntity(ent entity) bool {
	return e.entity&ent != 0
}

func (e event) String() string {
	return e.name
}
