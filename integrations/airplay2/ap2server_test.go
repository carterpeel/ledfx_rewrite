package airplay2

import (
	"github.com/hajimehoshi/oto"
	"ledfx/config"
	log "ledfx/logger"
	"testing"
)

func TestServer(t *testing.T) {
	if _, err := log.Init(config.Config{
		Verbose: true,
	}); err != nil {
		t.Fatalf("Error initializing logger: %v\n", err)
	}

	server := NewServer(Config{
		AdvertisementName: "LedFX-AirPlay-Output",
		VerboseLogging:    false,
		Port:              7000,
	})

	otoCtx, err := oto.NewContext(44100, 2, 2, 12000)
	if err != nil {
		t.Fatalf("Error creating new oto context: %v\n", err)
	}

	outputAudio := otoCtx.NewPlayer()
	defer outputAudio.Close()

	server.AddOutput(outputAudio)
	if err := server.Start(); err != nil {
		t.Fatalf("Error starting AirPlay2 server: %v\n", err)
	}

	server.Wait()
}

// TestServerSlowedAudio just sounds cool. Try it.
func TestServerSlowedAudio(t *testing.T) {
	if _, err := log.Init(config.Config{
		Verbose: true,
	}); err != nil {
		t.Fatalf("Error initializing logger: %v\n", err)
	}

	server := NewServer(Config{
		AdvertisementName: "AirPlay2-TestServer",
		VerboseLogging:    false,
		Port:              7000,
	})

	otoCtx, err := oto.NewContext(64000, 1, 2, 1024)
	if err != nil {
		t.Fatalf("Error creating new oto context: %v\n", err)
	}

	outputAudio := otoCtx.NewPlayer()
	defer outputAudio.Close()

	server.AddOutput(outputAudio)
	if err := server.Start(); err != nil {
		t.Fatalf("Error starting AirPlay2 server: %v\n", err)
	}

	server.Wait()
}
