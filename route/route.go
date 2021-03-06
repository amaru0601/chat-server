package route

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/amaru0601/fluent/chat"
	"github.com/amaru0601/fluent/ent"
	"github.com/amaru0601/fluent/ent/user"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  512,
		WriteBufferSize: 512,
		CheckOrigin: func(r *http.Request) bool {
			log.Printf("%s %s%s %v\n", r.Method, r.Host, r.RequestURI, r.Proto)
			return r.Method == http.MethodGet
		},
	}
)

type Route struct {
	db *ent.Client
}

func NewRoute(client *ent.Client) *Route {
	return &Route{db: client}
}

func (r *Route) JoinChat(manager map[int]*chat.Chat) echo.HandlerFunc {
	return func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			log.Println("Error on websocket connection:", err.Error())
		}
		defer ws.Close()

		chatID, _ := strconv.Atoi(c.QueryParam("id")) // se utiliza cuando ya existe el chat

		receiverUsername := c.QueryParam("receiver") // se utiliza solo cuando se va crear un chat

		var receiverID = -1
		if receiverUsername != "" {
			receiverClient, err := r.db.User.
				Query().
				Where(user.UsernameEQ(receiverUsername)).
				First(context.Background())
			if err != nil {
				log.Println(err)
			} else {
				receiverID = receiverClient.ID
			}
		}

		//context has a map where user is the default key for auth-header
		authHeader := c.Get("user").(*jwt.Token)
		username := authHeader.Claims.(jwt.MapClaims)["username"].(string)
		userClient, err := r.db.User.
			Query().
			Where(user.UsernameEQ(username)).
			First(context.Background())
		if err != nil {
			log.Println(err)
		}

		// conseguir el chat
		if room, ok := manager[chatID]; ok {
			//conectar cliente con web socket
			user := &chat.User{
				Username: username,
				Conn:     ws,
				Room:     room,
			}
			room.Join <- user
			user.Read(r.db, userClient, chatID)
		} else {
			newRoom := &chat.Chat{
				Users:    make(map[string]*chat.User),
				Messages: make(chan *chat.Message),
				Join:     make(chan *chat.User),
				Leave:    make(chan *chat.User),
				Id:       chatID,
			}
			/*
				caso 1 : el chat fue creado antes pero no existe un canal

				caso 2 : nuevo chat en la base de datos porque se mando el username del receiver
			*/
			if receiverID != -1 {
				//TODO: verificar si el chat por crear ya existe
				chatEnt, err := r.db.Chat.Create().
					SetType("public").
					AddMemberIDs(receiverID, userClient.ID).Save(context.Background())
				if err != nil {
					log.Println("Error al crear chat en bd")
					log.Println(err)
				}
				chatID = chatEnt.ID
			}

			manager[chatID] = newRoom
			go newRoom.Run()

			user := &chat.User{
				Username: username,
				Conn:     ws,
				Room:     newRoom,
			}

			newRoom.Join <- user
			user.Read(r.db, userClient, chatID)
		}
		return nil
	}
}

type UserSignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Route) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(UserSignIn)
		if err := c.Bind(u); err != nil {
			return err
		}

		userEnt, err := r.db.User.
			Query().
			Where(user.UsernameEQ(u.Username)).
			First(context.Background())

		//TODO : mejorar mensaje de error por parte de ent user not found
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}

		if u.Password != userEnt.Password {
			return c.String(http.StatusUnauthorized, "wrong password")
		}

		jwtoken, err := makeToken(u)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		return c.JSON(http.StatusOK, jwtoken)
	}
}

type UserSignUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	PrivateKey string `json:"privatekey"`
	PublicKey  []byte `json:"publickey"`
}

func (r *Route) SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(UserSignUp)
		if err := c.Bind(u); err != nil {
			return err
		}

		_, err := r.db.User.
			Create().
			SetUsername(u.Username).
			SetPassword(u.Password).
			SetPrivateKey(u.PrivateKey).
			SetPublicKey(u.PublicKey).
			Save(context.Background())
		if err != nil {
			fmt.Println(err)
			return err
		}

		return c.String(http.StatusOK, "user has been created")
	}
}
