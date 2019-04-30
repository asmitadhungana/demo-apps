package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/harmony-one/demo-apps/backend/client"
	"github.com/harmony-one/demo-apps/backend/db"
	"github.com/harmony-one/demo-apps/backend/p2p"
	"github.com/harmony-one/demo-apps/backend/utils"
	"google.golang.org/appengine"
	app_log "google.golang.org/appengine/log"
)

type respEnter struct {
	Address string `json:address`
	Level   uint   `json:level`
	Balance uint64 `json:balance`
}

type respFinish struct {
	Level   int    `json:level`
	Rewards uint64 `json:rewards`
}

var (
	version string
	builtBy string
	builtAt string
	commit  string
)

func printVersion(me string) {
	fmt.Fprintf(os.Stderr, "Harmony (C) 2019. %v, version %v-%v (%v %v)\n", path.Base(me), version, commit, builtBy, builtAt)
	os.Exit(0)
}

var (
	defaultConfigFile = ".hmy/backend.ini"
	defaultProfile    = "default"
	defaultPort       = "30000"
	leader            p2p.Peer
	backendProfile    *utils.BackendProfile

	db *fdb.Fdb

	profile     = flag.String("profile", defaultProfile, "name of the profile")
	versionFlag = flag.Bool("version", false, "Output version info")
)

const (
	minimalFee = 1
	gameFee    = 1
	adminKey   = "e401343197a852f361e38ce6b46c99f1d6d1f80499864c6ae7effee42b46ab6b"
	dbKeyFile  = "./keys/benchmark_account_key.json"
	dbProject  = "benchmark-209420"
)

// readProfile read the ini file and return the leader's IP
func readProfile(profile string) p2p.Peer {
	fmt.Printf("Using %s profile for backend\n", profile)
	var err error
	backendProfile, err = utils.ReadBackendProfile(defaultConfigFile, profile)
	if err != nil {
		fmt.Printf("Read backend profile error: %v\nExiting ...\n", err)
		os.Exit(2)
	}

	return backendProfile.RPCServer[0][0]
}

func main() {

	flag.Parse()
	if *versionFlag {
		printVersion(os.Args[0])
	}

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/enter", enterHandler)
	http.HandleFunc("/finish", finishHandler)
	http.HandleFunc("/test", testHandler)

	var err error
	db, err = fdb.NewFdb(dbKeyFile, dbProject)

	if err != nil || db == nil {
		log.Fatalf("Failed to create Fdb client: %v", err)
		os.Exit(1)
	}

	// Close FDB when done.
	defer db.CloseFdb()
	leader = readProfile(*profile)

	/*
		//Get a list of all current players
		_, err := restclient.GetPlayer(leader.IP, defaultPort)
		if err != nil {
			log.Fatalf("GetPlayer Error: %v", err)
			return
		}
	*/
	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func enterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if r.URL.Path != "/enter" {
		http.NotFound(w, r)
		return
	}
	q := r.URL.Query()

	emails, ok := q["email"]
	if !ok {
		http.Error(w, "missing params", http.StatusBadRequest)
		return
	}
	email := emails[0]

	// find the existing account from firebase DB
	accounts := db.FindAccount(email)

	// register the new account
	if len(accounts) == 0 {
		// generate the key
		account, _ := utils.GenereateAccount(email)
		err := restclient.FundMe(account)
		if err != nil {
			app_log.Criticalf(ctx, "enterHandler FundMe error: %v", err)
			http.Error(w, "FundMe Error, please retry", http.StatusInternalServerError)
			// TODO: retry
			return
		}
		leaders := restclient.GetLeaders()
		err = db.RegisterAccount(email, account, leaders[0].IP)
		if err != nil {
			app_log.Criticalf(ctx, "enterHandler registerAccount error: %v", err)
			http.Error(w, "Register Account, please retry", http.StatusInternalServerError)
		}
	}
	fmt.Printf("found Account: %v for email: %v\n", accounts, email)
	// we should find only one account, if more than one, just get the first one
	account := accounts[0]

	balance, err := restclient.GetBalance(account.Address)
	if err != nil {
		app_log.Criticalf(ctx, "enterHandler GetBalance error: %v", err)
		http.Error(w, "Can't GetBalance, please retry", http.StatusInternalServerError)
		return
	}
	if balance < minimalFee {
		http.Error(w, "Not enough balance to play, please get more token", http.StatusBadRequest)
		return
	}

	level, err := restclient.EnterPuzzle(account.Address, gameFee)
	if err != nil {
		app_log.Criticalf(ctx, "enterHandler EnterPuzzle error: %v", err)
		http.Error(w, "Can't Enter Game, please retry", http.StatusInternalServerError)
	}

	resp := respEnter{
		Address: account.Address,
		Level:   level,
		Balance: balance,
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("can't marshal enter resp: %s\n", resp)
		http.Error(w, "Can't marshal enter response", http.StatusInternalServerError)
		return
	}
	res := string(bytes)
	io.WriteString(w, res)
}

func finishHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if r.URL.Path != "/finish" {
		http.NotFound(w, r)
		return
	}
	q := r.URL.Query()

	var ok bool

	newlevels, ok := q["newlevel"]
	if !ok {
		http.Error(w, "missing params", http.StatusBadRequest)
		return
	}
	newlevel, err := strconv.Atoi(newlevels[0])
	if err != nil {
		http.Error(w, "wrong parameters", http.StatusBadRequest)
		return
	}

	accounts, ok := q["account"]
	if !ok {
		http.Error(w, "missing params", http.StatusBadRequest)
		return
	}
	account := accounts[0]
	keys, ok := q["key"]
	if !ok || keys[0] != adminKey {
		http.Error(w, "missing params", http.StatusBadRequest)
		return
	}

	rewards, err := restclient.GetRewards(account, newlevel)
	if err != nil {
		app_log.Criticalf(ctx, "finishHandler GetRewards error: %v", err)
		http.Error(w, "Can't Get Rewards", http.StatusInternalServerError)
		return
	}

	resp := respFinish{
		Level:   newlevel,
		Rewards: rewards,
	}
	bytes, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("can't marshal finish resp: %s\n", resp)
		http.Error(w, "Can't marshal finish response", http.StatusInternalServerError)
		return
	}
	res := string(bytes)
	io.WriteString(w, res)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if r.URL.Path != "/test" {
		http.NotFound(w, r)
		return
	}
	q := r.URL.Query()

	var ok bool
	var res string

	function, ok := q["function"]
	if !ok {
		http.Error(w, "missing params", http.StatusBadRequest)
		return
	}

	switch function[0] {
	case "FindAccount":
		emails, ok := q["email"]
		if !ok {
			http.Error(w, "missing email params", http.StatusBadRequest)
			break
		}
		accounts := db.FindAccount(emails[0])
		app_log.Infof(ctx, "accounts: %v", accounts)
		res = fmt.Sprintf("accounts: %v\n", accounts)
	}
	io.WriteString(w, res)
}