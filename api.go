package main

import (
	"io/ioutil"
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/garyburd/redigo/redis"
	"github.com/martini-contrib/render"
)

func AddMessage(req *http.Request, store redis.Conn, r render.Render) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	msg := string(body)

	// Increment the message ID.
	msgId, msgIdErr := store.Do("INCR", "messageId")
	if msgIdErr != nil {
		panic(msgIdErr)
	}

	// Store the message.
	_, storeErr := store.Do("HSET", "messages", msgId, msg)
	if storeErr != nil {
		panic(storeErr)
	}

	// We could also return the errors as JSON but this is
	// fine for now.
	r.JSON(http.StatusCreated, map[string]interface{}{
		"messageId": msgId,
	})
}

func GetMessage(req *http.Request, store redis.Conn, params martini.Params) (int, string) {
	msgId := params["id"]

	// Check if the message exists.
	// Conver to int otherwise we end up with int64.
	exists, existsErr := redis.Int(store.Do("HEXISTS", "messages", msgId))
	if existsErr != nil {
		panic(existsErr)
	}

	if exists != 1 {
		return http.StatusBadRequest, "message with this ID does not exist"
	}

	// redis.String wrapper is needed otherwise we end
	// up with []uint8.
	msg, err := redis.String(store.Do("HGET", "messages", msgId))
	if err != nil {
		panic(err)
	}

	return http.StatusOK, msg
}
