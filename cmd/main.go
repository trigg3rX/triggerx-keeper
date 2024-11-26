// github.com/trigg3rX/triggerx-keeper/cmd/main.go
package main

import (
    "context"
    "flag"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
    "encoding/json"

    "github.com/libp2p/go-libp2p"
    "github.com/libp2p/go-libp2p/core/peer"
    dht "github.com/libp2p/go-libp2p-kad-dht"
    "github.com/libp2p/go-libp2p/core/host"
    "github.com/trigg3rX/go-backend/pkg/network"
    "github.com/trigg3rX/triggerx-keeper/pkg/execution"
)

const (
    managerName   = "manager"
    retryInterval = 5 * time.Second
    maxRetries    = 12
)

func main() {
    nodeNumber := flag.Int("node", 1, "Node number (1, 2, 3, etc.)")
    listenAddr := flag.String("listen", "/ip4/0.0.0.0/tcp/9001", "Listen address for p2p connections")
    flag.Parse()

    keeperName := fmt.Sprintf("node%d", *nodeNumber)
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    host, kdht, err := setupP2P(ctx, *listenAddr)
    if err != nil {
        log.Fatalf("Failed to create P2P host: %v", err)
    }
    defer host.Close()
    defer kdht.Close()

    discovery := network.NewDiscovery(ctx, host, keeperName)

    if err := kdht.Bootstrap(ctx); err != nil {
        log.Printf("Warning: DHT bootstrap failed: %v", err)
    }

    if err := discovery.SavePeerInfo(); err != nil {
        log.Printf("Warning: Failed to save peer info: %v", err)
    }

    managerID, err := connectToManager(ctx, discovery)
    if err != nil {
        log.Fatalf("Failed to connect to manager: %v", err)
    }

    k := execution.NewKeeper(keeperName, nil, *managerID) // Pass appropriate messaging
    if err := k.Start(); err != nil {
        log.Fatalf("Failed to start keeper: %v", err)
    }
    defer k.Stop()

    for _, addr := range host.Addrs() {
        log.Printf("Keeper %s listening on: %s/p2p/%s", keeperName, addr, host.ID())
    }

    waitForInterrupt()
}

func connectToManager(ctx context.Context, discovery *network.Discovery) (*peer.ID, error) {
    var managerID *peer.ID
    var err error

    for i := 0; i < maxRetries; i++ {
        // Load peer info from file
        peerInfos := make(map[string]network.PeerInfo)
        if file, err := os.Open(network.PeerInfoFilePath); err == nil {
            decoder := json.NewDecoder(file)
            decoder.Decode(&peerInfos)
            file.Close()
        }

        // Try to connect to manager peer if found
        if managerInfo, exists := peerInfos[managerName]; exists {
            managerID, err = discovery.ConnectToPeer(managerInfo)
            if err == nil {
                log.Printf("Successfully connected to manager")
                return managerID, nil
            }
        }

        log.Printf("Attempt %d: Failed to connect to manager: %v", i+1, err)
        if i < maxRetries-1 {
            time.Sleep(retryInterval)
        }
    }

    return nil, fmt.Errorf("failed to connect to manager after %d attempts", maxRetries)
}

func setupP2P(ctx context.Context, listenAddr string) (host.Host, *dht.IpfsDHT, error) {
    host, err := libp2p.New(
        libp2p.ListenAddrStrings(listenAddr),
        libp2p.EnableRelay(),
        libp2p.EnableHolePunching(),
    )
    if err != nil {
        return nil, nil, fmt.Errorf("failed to create host: %v", err)
    }

    kdht, err := dht.New(ctx, host, dht.Mode(dht.ModeServer))
    if err != nil {
        host.Close()
        return nil, nil, fmt.Errorf("failed to create DHT: %v", err)
    }

    return host, kdht, nil
}

func waitForInterrupt() {
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan
    fmt.Println("\nReceived interrupt signal, shutting down...")
}
