package integration_test

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"

	"router7/internal/dns"

	miekgdns "github.com/miekg/dns"
)

func TestDNS(t *testing.T) {
	srv := dns.NewServer("localhost:4453", "lan")
	s := &miekgdns.Server{Addr: "localhost:4453", Net: "udp", Handler: srv.Mux}
	go s.ListenAndServe()
	const port = 4453
	dig := exec.Command("dig", "-p", strconv.Itoa(port), "+timeout=1", "+short", "-x", "8.8.8.8", "@127.0.0.1")
	dig.Stderr = os.Stderr
	out, err := dig.Output()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := strings.TrimSpace(string(out)), "google-public-dns-a.google.com."; got != want {
		t.Fatalf("dig -x 8.8.8.8: unexpected reply: got %q, want %q", got, want)
	}
}
