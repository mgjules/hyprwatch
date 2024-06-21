package main

type event string

const (
	// Workspace
	workspaceChanged   event = "workspacev2"
	workspaceCreated   event = "createworkspacev2"
	workspaceDestroyed event = "destroyworkspacev2"
	workspaceMoved     event = "moveworkspacev2"
	workspaceUrgent    event = "urgent"
	// Window
	activeWindow     event = "activeWindowv2"
	windowOpened     event = "openwindow"
	windowClosed     event = "closewindow"
	windowMoved      event = "movewindowv2"
	windowMinimized  event = "minimize"
	windowTitle      event = "windowtitle"
	windowFullscreen event = "fullscreen"
	// Monitor
	monitorFocused event = "focusedmon"
	monitorAdded   event = "monitoraddedv2"
	monitorRemoved event = "monitorremoved"
)

func (e event) IsValid() bool {
	switch e {
	case
		// Workspace
		workspaceChanged,
		workspaceCreated,
		workspaceDestroyed,
		workspaceMoved,
		workspaceUrgent,
		// Window
		activeWindow,
		windowOpened,
		windowClosed,
		windowMoved,
		windowMinimized,
		windowTitle,
		windowFullscreen,
		// Monitor
		monitorFocused,
		monitorAdded,
		monitorRemoved:
		return true
	default:
		return false
	}
}

func (e event) String() string {
	return string(e)
}
