package main

import (
	"context"
	"embed"
	"log"
	"os"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	g "xabbo.b7c.io/goearth"
	"xabbo.b7c.io/goearth/shockwave/in"
)

// Global variables for your application:
var (
	users            = map[int]*User{}
	usersPacketCount = 0
)

type User struct {
	Index      int
	Name       string
	Figure     string
	Gender     string
	Custom     string
	X, Y       int
	Z          float64
	PoolFigure string
	BadgeCode  string
	Type       int
}

// App struct
type App struct {
	ext    *g.Ext
	assets embed.FS
	ctx    context.Context
}

func (a *App) setupExt() {
	a.ext.Intercept(in.OPC_OK).With(handleEnterRoom)
	a.ext.Intercept(in.USERS).With(a.handleUsers)
	a.ext.Intercept(in.LOGOUT).With(a.handleRemoveUser)
	a.ext.Intercept(in.CHAT, in.CHAT_2, in.CHAT_3).With(a.handleChat)
}

func handleEnterRoom(e *g.Intercept) {
	usersPacketCount = 0
	clear(users)
}

func (a *App) handleUsers(e *g.Intercept) {
	// Observations:
	// The first USERS packet sent upon entering the room (after OPC_OK)
	// is the list of users that are already in the room.
	// The second USERS packet contains a single user, yourself.
	// The following USERS packets indicate someone entering the room.
	usersPacketCount++
	for range e.Packet.ReadInt() {
		var user User
		e.Packet.Read(&user)
		if user.Type == 1 {
			if usersPacketCount >= 3 {
				log.Printf("* %s entered the room", user.Name)
				runtime.EventsEmit(a.ctx, "UserEvent", map[string]interface{}{
					"username": user.Name,
					"message":  "entered the room",
				})
			}
			users[user.Index] = &user
		}
	}
}

func (a *App) handleChat(e *g.Intercept) {
	index := e.Packet.ReadInt()
	msg := e.Packet.ReadString()
	if user, ok := users[index]; ok {
		log.Printf("%s: %s", user.Name, msg)
		runtime.EventsEmit(a.ctx, "newMessage", map[string]interface{}{
			"username": user.Name,
			"message":  msg,
		})
	}
}

func (a *App) handleRemoveUser(e *g.Intercept) {
	s := e.Packet.ReadString()
	index, err := strconv.Atoi(s)
	if err != nil {
		return
	}
	if user, ok := users[index]; ok {
		log.Printf("* %s left the room.", user.Name)
		runtime.EventsEmit(a.ctx, "UserEvent", map[string]interface{}{
			"username": user.Name,
			"message":  "entered the room",
		})
		delete(users, index)
	}
}

// NewApp creates a new App application struct
func NewApp(ext *g.Ext, assets embed.FS) *App {
	return &App{
		ext:    ext,
		assets: assets,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.setupExt()
	go func() {
		a.runExt()
	}()
}

func (a *App) runExt() {
	defer os.Exit(0)
	a.ext.Run()
}
