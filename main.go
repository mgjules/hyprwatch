package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"os"
	"runtime/debug"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	logLevel *slog.LevelVar
	logger   *slog.Logger
)

func main() {
	logLevel = &slog.LevelVar{}
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})
	logger = slog.New(handler)
	slog.SetDefault(logger)

	app := &cli.App{
		Name:  "Hyprwatch",
		Usage: "Listens and parses Hyprland events",
		Commands: []*cli.Command{
			{
				Name:    workspace.String(),
				Aliases: []string{"workspaces"},
				Usage:   "workspace related events",
				Action:  execute(workspace),
			},
			{
				Name:    window.String(),
				Aliases: []string{"windows"},
				Usage:   "window related events",
				Action:  execute(window),
			},
			{
				Name:    monitor.String(),
				Aliases: []string{"monitors"},
				Usage:   "monitor related events",
				Action:  execute(monitor),
			},
			{
				Name:   "version",
				Usage:  "displays the version",
				Action: version,
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "socket",
				Value:   "${XDG_RUNTIME_DIR}/hypr/${HYPRLAND_INSTANCE_SIGNATURE}/.socket2.sock",
				Aliases: []string{"socket-path"},
				Usage:   "the path of the socket file",
			},
			&cli.BoolFlag{
				Name:    "debug",
				Value:   false,
				Aliases: []string{"d"},
				Usage:   "debug mode",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Error("encountered an application error: %v", "error", err)
		os.Exit(1)
	}
}

func execute(entity entity) cli.ActionFunc {
	logger := logger.With("entity", entity)

	return func(ctx *cli.Context) error {
		if ctx.Bool("debug") {
			logLevel.Set(slog.LevelDebug)
		}
		socket := os.ExpandEnv(ctx.String("socket"))
		if socket == "" {
			return cli.Exit("please specify a socket path", 100)
		}

		logger.Debug("connecting to socket ...")

		conn, err := net.Dial("unix", socket)
		if err != nil {
			logger.Error("cannot connect to socket", "error", err)
			return cli.Exit("cannot connect to socket", 101)
		}
		defer conn.Close()

		logger.Debug("connected to socket")

		reader := bufio.NewReader(conn)
		for {
			raw, err := reader.ReadString('\n')
			if err != nil {
				logger.Error("cannot read raw string", "error", err)
				continue
			}

			fmt.Println(raw)
		}
	}
}

func version(*cli.Context) error {
	binfo, ok := debug.ReadBuildInfo()
	if !ok {
		return cli.Exit("no build info available", 104)
	}

	fmt.Println(binfo.GoVersion)

	for _, setting := range binfo.Settings {
		if !strings.HasPrefix(setting.Key, "vcs.") || strings.HasSuffix(setting.Key, ".modified") {
			continue
		}

		fmt.Println(strings.TrimPrefix(setting.Key, "vcs."), ":", setting.Value)
	}

	return nil
}
