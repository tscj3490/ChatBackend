package v1

import (
	"strconv"

	"../../config"
	"../../model"
	"../../service/authService/permission"
	"../../service/chatRoomService"
	"../response"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitChatRooms inits chatRoom CRUD apis
// @Title ChatRooms
// @Description ChatRooms's router group.
func InitChatRooms(parentRoute *echo.Group) {
	route := parentRoute.Group("/chatRoom")
	route.Use(middleware.JWT([]byte(config.AuthTokenKey)))

	route.POST("", permission.AuthRequired(createChatRoom))
	route.GET("/:id", permission.AuthRequired(readChatRoom))
	route.GET("/userId", permission.AuthRequired(readChatRoomByUserId))
	route.PUT("/:id", permission.AuthRequired(updateChatRoom))
	route.DELETE("/:id", permission.AuthRequired(deleteChatRoom))

	route.GET("", permission.AuthRequired(readChatRooms))

	chatRoomService.InitService()
}

// @Title createChatRoom
// @Description Create a chatRoom.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Success 200 {object} model.ChatRoom 		"Returns created chatRoom"
// @Failure 400 {object} response.BasicResponse "err.chatRoom.bind"
// @Failure 400 {object} response.BasicResponse "err.chatRoom.create"
// @Resource /chatRooms
// @Router /chatRooms [post]
func createChatRoom(c echo.Context) error {
	chatRoom := &model.ChatRoom{}
	if err := c.Bind(chatRoom); err != nil {
		return response.KnownErrJSON(c, "err.chatRoom.bind", err)
	}

	// create chatRoom
	chatRoom, err := chatRoomService.CreateChatRoom(chatRoom)
	if err != nil {
		return response.KnownErrJSON(c, "err.chatRoom.create", err)
	}

	return response.SuccessInterface(c, chatRoom)
}

// @Title readChatRoom
// @Description Read a chatRoom.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"ChatRoom ID."
// @Success 200 {object} model.ChatRoom 		"Returns read chatRoom"
// @Failure 400 {object} response.BasicResponse "err.chatRoom.bind"
// @Failure 400 {object} response.BasicResponse "err.chatRoom.read"
// @Resource /chatRooms
// @Router /chatRooms/{id} [get]
func readChatRoom(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	chatRoom, err := chatRoomService.ReadChatRoom(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.chatRoom.read", err)
	}

	return response.SuccessInterface(c, chatRoom)
}

func readChatRoomByUserId(c echo.Context) error {
	id := c.FormValue("id")

	// read chatRooms with params
	chatRooms, total, err := chatRoomService.ReadChatRoomsByUserId(id)
	if err != nil {
		return response.KnownErrJSON(c, "err.chatRoom.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, chatRooms})
}

// @Title updateChatRoom
// @Description Update a chatRoom.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string  true	"Bearer {token}"
// @Param   id				path   	string  true	"ChatRoom ID."
// @Param   avatar      	form   	string  true	"ChatRoom Avatar"
// @Param   firstname		form   	string  true	"ChatRoom Firstname"
// @Param   lastname		form   	string  true	"ChatRoom Lastname"
// @Param   email	    	form   	string  true	"ChatRoom Email"
// @Param   birth      		form   	Time   	true	"ChatRoom Birth"
// @Success 200 {object} model.ChatRoom 		"Returns read chatRoom"
// @Failure 400 {object} response.BasicResponse "err.chatRoom.bind"
// @Failure 400 {object} response.BasicResponse "err.chatRoom.read"
// @Resource /chatRooms
// @Router /chatRooms/{id} [put]
func updateChatRoom(c echo.Context) error {
	chatRoom := &model.ChatRoom{}
	if err := c.Bind(chatRoom); err != nil {
		return response.KnownErrJSON(c, "err.chatRoom.bind", err)
	}

	// id, _ := strconv.Atoi(c.Param("id"))
	chatRoom, err := chatRoomService.UpdateChatRoom(chatRoom)
	if err != nil {
		return response.KnownErrJSON(c, "err.chatRoom.update", err)
	}

	chatRoom, _ = chatRoomService.ReadChatRoom(chatRoom.ID)
	return response.SuccessInterface(c, chatRoom)
}

// @Title deleteChatRoom
// @Description Delete a chatRoom.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string	true	"Bearer {token}"
// @Param   id				path   	string  true	"ChatRoom ID."
// @Success 200 {object} response.BasicResponse "ChatRoom is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.chatRoom.bind"
// @Failure 400 {object} response.BasicResponse "err.chatRoom.delete"
// @Resource /chatRooms
// @Router /chatRooms/{id} [delete]
func deleteChatRoom(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	// delete chatRoom with id
	err := chatRoomService.DeleteChatRoom(uint(id))
	if err != nil {
		return response.KnownErrJSON(c, "err.chatRoom.delete", err)
	}
	return response.SuccessJSON(c, "ChatRoom is deleted correctly.")
}

// @Title readChatRooms
// @Description Read chatRooms with parameters.
// @Accept  json
// @Produce	json
// @Param   Authorization	header 	string 	true	"Bearer {token}"
// @Param   query			form   	string  false	"Will search string."
// @Param   offset			form    int		false	"Offset for pagination."
// @Param   count 			form    int		false	"Count that will show per page."
// @Param   field			form    string  false	"Sort field."
// @Param   sort			form    int		false	"Sort direction. 0:default, 1:Ascending, -1:Descending"
// @Success 200 {object} ListForm 				"ChatRoom is deleted correctly."
// @Failure 400 {object} response.BasicResponse "err.chatRoom.read"
// @Resource /chatRooms
// @Router /chatRooms [get]
func readChatRooms(c echo.Context) error {
	query := c.FormValue("query")
	offset, _ := strconv.Atoi(c.FormValue("offset"))
	count, _ := strconv.Atoi(c.FormValue("count"))
	field := c.FormValue("field")
	sort, _ := strconv.Atoi(c.FormValue("sort"))
	userID, _ := strconv.Atoi(c.FormValue("userId"))

	// read chatRooms with params
	chatRooms, total, err := chatRoomService.ReadChatRooms(query, offset, count, field, sort, uint(userID))
	if err != nil {
		return response.KnownErrJSON(c, "err.chatRoom.read", err)
	}

	return response.SuccessInterface(c, &ListForm{total, chatRooms})
}
