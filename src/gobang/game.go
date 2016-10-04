package gobang

import (
	"github.com/labstack/echo"
	"net/http"
	"strings"
	"fmt"
	"github.com/kataras/go-sessions"
	"github.com/labstack/echo/engine/standard"
)

func Game(c echo.Context) error {
	create := false
	roomId := ""
	queryString := c.Request().URL().QueryString()
	w := c.Response().(*standard.Response).ResponseWriter
	r := c.Request().(*standard.Request).Request
	sess := sessions.Start(w, r)
	if strings.EqualFold(queryString, "closed") {
		return c.Redirect(http.StatusMovedPermanently, "index.html")
	}
	if strings.EqualFold(queryString, "create") {
		create = true
	} else if len(queryString) != 0 {
		if room, ok := roomList.rooms[queryString]; ok {
			roomId = room.roomId
			sess.Set("roomId", roomId)
		} else {
			return c.HTML(http.StatusNotFound, fmt.Sprintf(`<script>alert("Can not found room id %s!");location.href="index.html";</script>`, queryString))
		}
	} else {
		create = true
	}
	if create {
		sess.Set("create", "true")
	}
	return c.Render(http.StatusOK, "game", struct {
		Create bool
		RoomId string
		Username string
	}{Create: create, RoomId: roomId, Username: "Anonymous"})
}