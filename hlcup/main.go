package main

import (
	"database/sql"
	"log"

	"encoding/json"

	"strconv"

	"net/http"

	"runtime"

	"github.com/buaazp/fasthttprouter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
)

var (
	db *sql.DB
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	initEnv()

	router := fasthttprouter.New()
	router.GET("/users/:id", handleUsers)
	router.GET("/users/:id/visits", handleUserVisits)
	router.POST("/users/:id", handleUsersPost)

	router.GET("/locations/:id", handleLocations)
	router.GET("/locations/:id/avg", handleLocationAvgs)
	router.POST("/locations/:id", handleLocationsPost)

	router.GET("/visits/:id", handleVisits)
	router.POST("/visits/:id", handleVisitsPost)

	log.Print("start serving")

	log.Fatal(fasthttp.ListenAndServe(":80", router.Handler))
}

func handleUsers(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	user := UserByID(id)
	if user == nil {
		ctx.NotFound()
	} else {
		writeJson(ctx, user)
	}
}

func handleUserVisits(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)

	user := UserByID(id)
	if user == nil {
		ctx.NotFound()
		return
	}

	params := map[string]interface{}{}
	failed := false
	ctx.QueryArgs().VisitAll(visitArg(params, &failed))

	if failed {
		ctx.SetStatusCode(http.StatusBadRequest)
	} else {
		visits := FindVisits(id, params)
		writeJson(ctx, map[string]interface{}{
			"visits": visits,
		})
	}
}

func visitArg(params map[string]interface{}, failed *bool) func(key, value []byte) {
	return func(key, value []byte) {
		k := string(key)
		v := string(value)
		if k == "country" {
			params[k] = v
		} else if k == "gender" {
			if v != "m" && v != "f" {
				*failed = true
			} else {
				params[k] = v
			}
		} else {
			if intValue, err := strconv.ParseInt(v, 10, 64); err != nil {
				*failed = true
			} else {
				params[k] = intValue
			}
		}
	}
}

func handleUsersPost(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)

	if id == "new" {
		userMap := map[string]interface{}{}
		if err := json.Unmarshal(ctx.PostBody(), &userMap); err != nil {
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}

		if err := SaveUser(userMap); err != nil {
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}
	} else {
		user := UserByID(id)
		if user == nil {
			ctx.NotFound()
			return
		}

		userMap := map[string]interface{}{}
		if err := json.Unmarshal(ctx.PostBody(), &userMap); err != nil {
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}

		if err := UpdateUser(id, userMap); err != nil {
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}
	}

	writeJson(ctx, map[string]interface{}{})
}

func handleLocations(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	location := LocationByID(id)
	if location == nil {
		ctx.NotFound()
	} else {
		writeJson(ctx, location)
	}
}

func handleLocationAvgs(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	location := LocationByID(id)
	if location == nil {
		ctx.NotFound()
		return
	}

	params := map[string]interface{}{}
	failed := false
	ctx.QueryArgs().VisitAll(visitArg(params, &failed))

	if failed {
		ctx.SetStatusCode(http.StatusBadRequest)
	} else {
		avg := LocationAvg(id, params)
		writeJson(ctx, map[string]interface{}{
			"avg": avg,
		})
	}
}

func handleLocationsPost(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)

	if id == "new" {
		locationMap := map[string]interface{}{}
		if err := json.Unmarshal(ctx.PostBody(), &locationMap); err != nil {
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}

		if err := SaveLocation(locationMap); err != nil {
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}
	} else {
		location := LocationByID(id)
		if location == nil {
			ctx.NotFound()
			return
		}

		locationMap := map[string]interface{}{}
		if err := json.Unmarshal(ctx.PostBody(), &locationMap); err != nil {
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}

		if err := UpdateLocation(id, locationMap); err != nil {
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}
	}

	writeJson(ctx, map[string]interface{}{})
}

func handleVisits(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id")
	visit := VisitByID(id)
	if visit == nil {
		ctx.NotFound()
	} else {
		writeJson(ctx, visit)
	}
}

func handleVisitsPost(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)

	if id == "new" {
		visitMap := map[string]interface{}{}
		if err := json.Unmarshal(ctx.PostBody(), &visitMap); err != nil {
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}

		if err := SaveVisit(visitMap); err != nil {
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}
	} else {
		visit := VisitByID(id)
		if visit == nil {
			ctx.NotFound()
			return
		}

		visitMap := map[string]interface{}{}
		if err := json.Unmarshal(ctx.PostBody(), &visitMap); err != nil {
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}

		if err := UpdateVisit(id, visitMap); err != nil {
			ctx.SetStatusCode(http.StatusBadRequest)
			return
		}
	}

	writeJson(ctx, map[string]interface{}{})
}

func writeJson(ctx *fasthttp.RequestCtx, i interface{}) {
	if data, err := json.Marshal(i); err != nil {
		panic(err)
	} else {
		//log.Print(string(data))
		ctx.SetBody(data)
	}
}
