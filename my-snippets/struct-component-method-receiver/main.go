package main

import "fmt"

type Version string

func (v Version) PrintVersion() {
    fmt.Println("Version is", v)
}

type Game struct {
    Name               string
    MultiplayerSupport bool
    Genre              string
    Version
}

type ERP struct {
    Name               string
    MRPSupport         bool
    SupportedDatabases []string
    Version
}

func main() {

    g := Game{
        "Fear Effect",
        false,
        "Action-Adventure",
        "1.0.0",
    }

    g.PrintVersion()
    // Version is 1.0.0
	g.Version.PrintVersion()


    e := ERP{
        "Logo",
        true,
        []string{"ms-sql"},
        "2.0.0",
    }

    e.PrintVersion()
    // Version is 2.0.0
	e.Version.PrintVersion()

}