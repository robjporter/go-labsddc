package app

import (
	"sync"
	"gopkg.in/alecthomas/kingpin.v2"
	"html/template"
)

const (
	PORT    = 8080
	DEBUG   = true
	VERSION = "0.1.0"
)

var (
	NOCACHE  = []template.HTML{template.HTML("<meta http-equiv='cache-control' content='max-age=0'/>"), template.HTML("<meta http-equiv='cache-control' content='no-cache'/>"), template.HTML("<meta http-equiv='expires' content='0'/>"), template.HTML("<meta http-equiv='expires' content='Tue, 01 Jan 1980 1:00:00 GMT'/>"), template.HTML("<meta http-equiv='pragma' content='no-cache'/>")}
	count    uint64
	once     sync.Once
	instance *Application
	add      = kingpin.Command("add", "Register a new device.")
	delete   = kingpin.Command("delete", "Deregister a device.")
	run      = kingpin.Command("run", "Run the server.")
	show     = kingpin.Command("show", "Show current configuration.")

	debug = kingpin.Flag("debug", "Enable debug mode.").Bool()

	addAPIC = add.Command("apic", "Add a new APIC.")
	addHX   = add.Command("hx", "Add a new HX platform.")
	addVC   = add.Command("vc", "Add a new vCenter.")

	deleteAPIC = delete.Command("apic", "Delete an APIC.")
	deleteHX   = delete.Command("hx", "Delete an HX platform.")
	deleteVC   = delete.Command("vc", "Delete a vCenter.")

	addAPICIP       = addAPIC.Flag("ip", "IP Address or DNS name of APIC.").Required().IP()
	addAPICUsername = addAPIC.Flag("username", "Username for connecting to APIC.").Required().String()
	addAPICPassword = addAPIC.Flag("password", "Password for connecting to APIC.").Required().String()

	addHXIP       = addHX.Flag("ip", "IP Address or DNS name of Hyperflex.").Required().IP()
	addHXUsername = addHX.Flag("username", "Username for connecting to Hyperflex.").Required().String()
	addHXPassword = addHX.Flag("password", "Password for connecting to Hyperflex.").Required().String()

	addVCIP       = addVC.Flag("ip", "IP Address or DNS name of vCenter.").Required().IP()
	addVCUsername = addVC.Flag("username", "Username for connecting to vCenter.").Required().String()
	addVCPassword = addVC.Flag("password", "Password for connecting to vCenter.").Required().String()

	delAPIC = deleteAPIC.Flag("ip", "IP Address or DNS name of APIC.").Required().IP()
	delHX   = deleteHX.Flag("ip", "IP Address or DNS name of Hyperflex.").Required().IP()
	delVC   = deleteVC.Flag("ip", "IP Address or DNS name of vCenter.").Required().IP()
)
