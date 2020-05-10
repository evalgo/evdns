package evdns

import (
	"evalgo.org/evmsg"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"os"
	"time"
)

func (h *Hetzner) WSStart(address, client, secret, webroot string) error {
	h.WSAddress = address
	h.WSClient = client
	h.WSSecret = secret
	h.WSWebroot = webroot
	evmsg.ID = client
	evmsg.Secret = secret
	e := echo.New()
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(echoLog.INFO)
	log.Logger().SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339})
	e.Logger = log.Logger()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", webroot)
	e.GET("/v0.0.1/ws", func(c echo.Context) error {
		s := websocket.Server{
			Handler: websocket.Handler(func(ws *websocket.Conn) {
				defer ws.Close()
			WEBSOCKET:
				for {
					var msg evmsg.Message
					err := websocket.JSON.Receive(ws, &msg)
					if err != nil {
						c.Logger().Error(err)
						if err == io.EOF {
							c.Logger().Info("websocket client closed connection!")
							return
						}
					}
					err = evmsg.Auth(&msg)
					if err != nil {
						c.Logger().Error(err)
						err = websocket.JSON.Send(ws, &msg)
						if err != nil {
							c.Logger().Error(err)
						}
						continue WEBSOCKET
					}
					switch msg.Scope {
					case "Dns":
						switch msg.Command {
						case "deleteRecord":
							msg.State = "Response"
							msg.Debug.Info = "Dns::deleteRecord"
							err = evmsg.CheckRequiredKeys(&msg, []string{"id"})
							if err != nil {
								c.Logger().Error(err)
								msg.Debug.Error = err.Error()
							} else {
								record, err := h.DeleteRecord(msg.Value("id").(string))
								if err != nil {
									c.Logger().Error(err)
									msg.Debug.Error = err.Error()
								} else {
									msg.Data = []interface{}{record.(map[string]interface{})["record"]}
								}
							}
						case "createRecord":
							msg.State = "Response"
							msg.Debug.Info = "Dns::createRecord"
							record, err := h.NewRecord(msg.Values())
							if err != nil {
								c.Logger().Error(err)
								msg.Debug.Error = err.Error()
							} else {
								msg.Data = []interface{}{record.(map[string]interface{})["record"]}
							}
						case "getRecord":
							msg.State = "Response"
							msg.Debug.Info = "Dns::getRecord"
							record, err := h.Record(msg.Value("id").(string))
							if err != nil {
								c.Logger().Error(err)
								msg.Debug.Error = err.Error()
							} else {
								msg.Data = []interface{}{record.(map[string]interface{})["record"]}
							}
						case "getRecords":
							msg.State = "Response"
							msg.Debug.Info = "Dns::getRecords"
							records, err := h.Records(msg.Value("id").(string))
							if err != nil {
								c.Logger().Error(err)
								msg.Debug.Error = err.Error()
							} else {
								msg.Data = records.(map[string]interface{})["records"]
							}
						case "deleteZone":
							msg.State = "Response"
							msg.Debug.Info = "Dns::deleteZone"
							record, err := h.DeleteZone(msg.Value("id").(string))
							if err != nil {
								c.Logger().Error(err)
								msg.Debug.Error = err.Error()
							} else {
								msg.Data = []interface{}{record.(map[string]interface{})["zone"]}
							}
						case "createZone":
							msg.State = "Response"
							msg.Debug.Info = "Dns::createZone"
							record, err := h.NewZone(msg.Values())
							if err != nil {
								c.Logger().Error(err)
								msg.Debug.Error = err.Error()
							} else {
								msg.Data = []interface{}{record.(map[string]interface{})["zone"]}
							}
						case "getZone":
							msg.State = "Response"
							msg.Debug.Info = "Dns::getZone"
							record, err := h.Zone(msg.Value("id").(string))
							if err != nil {
								c.Logger().Error(err)
								msg.Debug.Error = err.Error()
							} else {
								msg.Data = []interface{}{record.(map[string]interface{})["zone"]}
							}
						case "getZones":
							msg.State = "Response"
							msg.Debug.Info = "Dns::getZones"
							zones, err := h.Zones()
							if err != nil {
								c.Logger().Error(err)
								msg.Debug.Error = err.Error()
							} else {
								if zns, ok := zones.(map[string]interface{})["zones"]; ok {
									msg.Data = zns
								} else {
									msg.Debug.Error = zones.(map[string]interface{})["message"].(string)
									msg.Data = []interface{}{}
								}
							}
						default:
							// do something here
						}
					}
					// send msg response
					err = websocket.JSON.Send(ws, &msg)
					if err != nil {
						c.Logger().Error(err)
					}
				}
			}),
			Handshake: func(*websocket.Config, *http.Request) error {
				return nil
			},
		}
		s.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	return e.Start(address)
}
