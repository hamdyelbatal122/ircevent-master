
import (
	"crypto/tls"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const server = "irc.freenode.net:6667"
const serverssl = "irc.freenode.net:7000"
const channel = "#go-eventirc-test"
const dict = "abcdefghijklmnopqrstuvwxyz"

//Spammy
const verbose_tests = false
const debug_tests = true

func TestConnectionEmtpyServer(t *testing.T) {
	irccon := IRC("go-eventirc", "go-eventirc")
	err := irccon.Connect("")
	if err == nil {
		t.Fatal("emtpy server string not detected")
	}
}

func TestConnectionDoubleColon(t *testing.T) {
	irccon := IRC("go-eventirc", "go-eventirc")
	err := irccon.Connect("::")
	if err == nil {
		t.Fatal("wrong n
