package main

import (
	"html/template"
	"macaron-simple/routes"

	"macaron-simple/config"

	"github.com/go-macaron/cache"
	_ "github.com/go-macaron/cache/redis"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/session"
	_ "github.com/go-macaron/session/redis"
	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(gzip.Gziper())
	m.Use(macaron.Static("public"))

	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory:  "views",
		Extensions: []string{".html"},
		Funcs: []template.FuncMap{map[string]interface{}{
			"AppName": func() string {
				return "Macaron"
			},
		}},
	}))

	m.Use(cache.Cacher(cache.Options{
		Adapter: "redis",
		// e.g.: network=tcp,addr=127.0.0.1:6379,password=macaron,db=0,pool_size=100,idle_timeout=180,hset_name=MacaronCache,prefix=cache:
		AdapterConfig: "network=tcp,addr=127.0.0.1:6379,db=0,prefix=cache_",
		OccupyMode:    false,
	}))

	m.Use(session.Sessioner(session.Options{
		Provider: "redis",
		// e.g.: network=tcp,addr=127.0.0.1:6379,password=macaron,db=0,pool_size=100,idle_timeout=180,prefix=session:
		ProviderConfig: "network=tcp,addr=127.0.0.1:6379,db=1,prefix=session_",
	}))

	routes.InitRoutes(m)

	cfg := config.Get("config/app.ini")
	host := cfg.GetConfigKey("host").String()
	port := cfg.GetConfigKey("port").MustInt()

	m.Run(host, port)
}
