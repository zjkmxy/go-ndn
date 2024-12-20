package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	enc "github.com/zjkmxy/go-ndn/pkg/encoding"
	basic_engine "github.com/zjkmxy/go-ndn/pkg/engine/basic"
	"github.com/zjkmxy/go-ndn/pkg/log"
	"github.com/zjkmxy/go-ndn/pkg/ndn"
	"github.com/zjkmxy/go-ndn/pkg/schema"
	_ "github.com/zjkmxy/go-ndn/pkg/schema/rdr"
	sec "github.com/zjkmxy/go-ndn/pkg/security"
	"github.com/zjkmxy/go-ndn/pkg/utils"
)

const SchemaJson = `{
  "nodes": {
    "/<v=time>": {
      "type": "GeneralObjNode",
      "attrs": {
        "MetaFreshness": 10,
        "MaxRetriesForMeta": 2,
        "ManifestFreshness": 10,
        "MaxRetriesForManifest": 2,
        "MetaLifetime": 6000,
        "Lifetime": 6000,
        "Freshness": 3153600000000,
        "ValidDuration": 3153600000000,
        "SegmentSize": 80,
        "MaxRetriesOnFailure": 3,
        "Pipeline": "SinglePacket"
      }
    }
  },
  "policies": [
    {
      "type": "RegisterPolicy",
      "path": "/",
      "attrs": {
        "RegisterIf": "$isProducer"
      }
    },
    {
      "type": "Sha256Signer",
      "path": "/<v=time>/32=data/<seg=segmentNumber>"
    },
    {
      "type": "FixedHmacSigner",
      "path": "/<v=time>/32=manifest",
      "attrs": {
        "KeyValue": "$hmacKey"
      }
    },
    {
      "type": "FixedHmacSigner",
      "path": "/<v=time>/32=metadata",
      "attrs": {
        "KeyValue": "$hmacKey"
      }
    },
    {
      "type": "MemStorage",
      "path": "/",
      "attrs": {}
    }
  ]
}`

const LoremIpsum = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna
aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
`
const HmacKey = "Hello, World!"

func passAll(enc.Name, enc.Wire, ndn.Signature) bool {
	return true
}

func main() {
	log.SetLevel(log.InfoLevel)
	logger := log.WithField("module", "main")

	// Setup schema tree
	tree := schema.CreateFromJson(SchemaJson, map[string]any{
		"$isProducer": true,
		"$hmacKey":    HmacKey,
	})

	// Start engine
	timer := basic_engine.NewTimer()
	face := basic_engine.NewStreamFace("unix", "/var/run/nfd/nfd.sock", true)
	app := basic_engine.NewEngine(face, timer, sec.NewSha256IntSigner(timer), passAll)
	err := app.Start()
	if err != nil {
		logger.Fatalf("Unable to start engine: %+v", err)
		return
	}
	defer app.Shutdown()

	// Attach schema
	prefix, _ := enc.NameFromStr("/example/schema/groupSigApp")
	err = tree.Attach(prefix, app)
	if err != nil {
		logger.Fatalf("Unable to attach the schema to the engine: %+v", err)
		return
	}
	defer tree.Detach()

	// Produce data
	ver := utils.MakeTimestamp(timer.Now())
	path, _ := enc.NamePatternFromStr("/<v=time>")
	node := tree.At(path)
	mNode := node.Apply(enc.Matching{
		"time": enc.Nat(ver).Bytes(),
	})
	mNode.Call("Provide", enc.Wire{[]byte(LoremIpsum)})
	fmt.Printf("Generated packet with version= %d\n", ver)

	// Wait for keyboard quit signal
	sigChannel := make(chan os.Signal, 1)
	fmt.Print("Start serving ...\n")
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
	receivedSig := <-sigChannel
	logger.Infof("Received signal %+v - exiting\n", receivedSig)
}
