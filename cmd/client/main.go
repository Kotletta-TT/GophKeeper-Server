package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	config "github.com/Kotletta-TT/GophKeeper/config/client"
	"github.com/Kotletta-TT/GophKeeper/internal/client/view"
	"github.com/Kotletta-TT/GophKeeper/proto"
	"github.com/rivo/tview"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := Initialize()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial(cfg.Server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := proto.NewUserServiceClient(conn)
	sc := proto.NewSecretCardServiceClient(conn)
	fc := proto.NewFileSecretCardServiceClient(conn)

	var token, usrID string

	app := tview.NewApplication()
	pages := tview.NewPages()
	rv := view.NewRegisterView(c, app, pages)
	iv := view.NewInitialView(c, app, pages)
	lv := view.NewLoginView(c, app, pages, &token, &usrID)
	cv := view.NewCardView(sc, app, pages, &token, &usrID, fc)
	// clv := view.NewCardListView(sc, app, pages)

	v := view.NewView(rv, iv, lv, cv)

	pages.AddPage("Register", v.Register.GetForm(), true, true)
	pages.AddPage("Login", v.Login.GetForm(), true, true)
	pages.AddPage("Card", v.Card.GetList(), true, true)
	pages.AddPage("Initial", v.Initial.GetList(), true, true)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}

func Initialize() (*config.Config, error) {
	cfgPath, ok := os.LookupEnv("CONFIG_PATH")
	if !ok {
		cfgPath = "config.yaml"
	}
	flag.StringVar(&cfgPath, "c", "config.yaml", "path to config file")
	flag.Parse()
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	if cfg.Server == "" {
		return nil, fmt.Errorf("server address is empty")
	}
	return cfg, nil
}
