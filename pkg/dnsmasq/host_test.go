package dnsmasq

import (
	"context"
	"os"
	"testing"

	"github.com/browningluke/opnsense-go/pkg/api"
)

func TestHost(t *testing.T) {
	opnsense_url := os.Getenv("OPNSENSE_URI")
	opnsense_key := os.Getenv("OPNSENSE_API_KEY")
	opnsense_secret := os.Getenv("OPNSENSE_API_SECRET")

	api_client := api.NewClient(api.Options{
		Uri:           opnsense_url,
		APIKey:        opnsense_key,
		APISecret:     opnsense_secret,
		AllowInsecure: true,
		MaxBackoff:    30,
		MinBackoff:    1,
		MaxRetries:    4,
	})

	ctx := context.Background()

	controller := Controller{
		Api: api_client,
	}

	host := &Host{
		IpAddress:   "192.168.1.100",
		HwAddress:   "00:11:22:33:44:55",
		Hostname:    "testhost",
		Domain:      "testdomain",
		Description: "Test static dhcp host entry",
	}

	key, err := controller.AddHost(ctx, host)
	if err != nil {
		t.Fatalf("Failed to add static dhcp host: %v", err)
	}
	t.Logf("Added static dhcp host with key: %s", key)

	retrievedHost, err := controller.GetHost(ctx, key)
	if err != nil {
		t.Fatalf("Failed to get static dhcp host: %v", err)
	}
	t.Logf("Retrieved static dhcp host: %+v", retrievedHost)
	if retrievedHost.IpAddress != host.IpAddress {
		t.Fatalf("Retrieved static dhcp host ip address does not match: got %s, want %s", retrievedHost.Hostname, host.Hostname)
	}
	if retrievedHost.HwAddress != host.HwAddress {
		t.Fatalf("Retrieved static dhcp host hw address does not match: got %s, want %s", retrievedHost.HwAddress, host.HwAddress)
	}
	if retrievedHost.Hostname != host.Hostname {
		t.Fatalf("Retrieved static dhcp host hostname does not match: got %s, want %s", retrievedHost.Hostname, host.Hostname)
	}
	if retrievedHost.Domain != host.Domain {
		t.Fatalf("Retrieved static dhcp host domain does not match: got %s, want %s", retrievedHost.Domain, host.Domain)
	}
	if retrievedHost.Description != host.Description {
		t.Fatalf("Retrieved static dhcp host description does not match: got %s, want %s", retrievedHost.Description, host.Description)
	}

	host.IpAddress = "192.168.1.200"
	host.HwAddress = "66:77:88:99:AA:BB"
	host.Hostname = "updatedhost"
	host.Domain = "updateddomain"
	host.Description = "Test static dhcp host entry updated"
	err = controller.UpdateHost(ctx, key, host)
	if err != nil {
		t.Fatalf("Failed to update static dhcp host: %v", err)
	}
	t.Logf("Updated static dhcp host with key: %s", key)

	retrievedHost, err = controller.GetHost(ctx, key)
	if err != nil {
		t.Fatalf("Failed to get updated static dhcp host: %v", err)
	}
	if retrievedHost.IpAddress != host.IpAddress {
		t.Fatalf("Retrieved static dhcp host ip address does not match updated ip address: got %s, want %s", retrievedHost.IpAddress, host.IpAddress)
	}
	if retrievedHost.HwAddress != host.HwAddress {
		t.Fatalf("Retrieved static dhcp host hw address does not match updated hw address: got %s, want %s", retrievedHost.HwAddress, host.HwAddress)
	}
	if retrievedHost.Hostname != host.Hostname {
		t.Fatalf("Retrieved static dhcp host hostname does not match updated hostname: got %s, want %s", retrievedHost.Hostname, host.Hostname)
	}
	if retrievedHost.Domain != host.Domain {
		t.Fatalf("Retrieved static dhcp host domain does not match updated domain: got %s, want %s", retrievedHost.Domain, host.Domain)
	}
	if retrievedHost.Description != host.Description {
		t.Fatalf("Retrieved static dhcp host description does not match updated description: got %s, want %s", retrievedHost.Description, host.Description)
	}

	err = controller.DeleteHost(ctx, key)
	if err != nil {
		t.Fatalf("Failed to delete static dhcp host: %v", err)
	}
	t.Logf("Deleted static dhcp host with key: %s", key)
}
