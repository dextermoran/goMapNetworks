package main

import (
	"fmt"
	"goMapNetworks/nmap"
	_ "github.com/murlokswarm/mac"
	"github.com/murlokswarm/app"
)

type Hello struct {
	Greeting []nmap.Host
}

func (h *Hello) Render() string {
	return `
<div class="WindowLayout">
    <div class="HelloBox">
			<span>{{if .Greeting}}{{html .Greeting}}{{else}}<h1>Hey Buddy, what network do you want to scan?</h1>
			        <input type="text" placeholder="(192.168.0.0/24)" onchange="OnInputChange" />{{end}}</span>
    </div>
</div>
    `
}

/*

func ProcessNmap(x []nmap.Host){
	var foundHosts []string
	i := 0
	for _, n := range x {
		i++
		fmt.Println(n)
		foundHosts[i] = fmt.Sprintf(n)
	}
	return foundHosts
}

*/

func (h *Hello) OnInputChange(arg app.ChangeArg) {
	input := arg.Value
	fmt.Println("scanning...")
	hosts, err := nmap.ScanOpenTcpPorts(input)
	fmt.Println("scanned!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hosts)


	h.Greeting = hosts
	app.Render(h)
}

func init() {
	app.RegisterComponent(&Hello{})
}

func main() {

	app.OnLaunch = func() {
		win := app.NewWindow(app.Window{
			Title:          "Hello World",
			Width:          1280,
			Height:         720,
			TitlebarHidden: true,
		})

		hello := &Hello{}
		win.Mount(hello)
	}

	app.Run()


}
