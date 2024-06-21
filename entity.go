package main

import "fmt"

type entity string

const (
	workspace entity = "workspace"
	window    entity = "window"
	monitor   entity = "monitor"
)

func (e entity) Support(ev event) error {
	if ev.IsValid() {
		return fmt.Errorf("invalid event %q", ev)
	}

	events, err := e.Events()
	if err != nil {
		return err
	}

	for _, sev := range events {
		if sev == ev {
			return nil
		}
	}

	return fmt.Errorf("unsupported event %q for entity", ev)
}

func (e entity) Events() ([]event, error) {
	switch e {
	case workspace:
		return []event{
			workspaceChanged,
			workspaceCreated,
			workspaceDestroyed,
			workspaceMoved,
			workspaceUrgent,
		}, nil
	case window:
		return []event{
			activeWindow,
			windowOpened,
			windowClosed,
			windowMoved,
			windowMinimized,
			windowTitle,
			windowFullscreen,
		}, nil
	case monitor:
		return []event{
			monitorFocused,
			monitorAdded,
			monitorRemoved,
		}, nil
	default:
		return nil, fmt.Errorf("no event defined for entity %q", e)
	}
}

func (e entity) String() string {
	return string(e)
}
