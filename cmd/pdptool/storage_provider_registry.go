package main

import (
	"github.com/urfave/cli/v2"
)

var registerProvider = &cli.Command{
	Name:  "register-service-provider",
	Usage: "Register a new service provider to a service provider registry",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "service-url",
			Usage:    "URL of the PDP service",
			Required: true,
		},
		&cli.Uint64Flag{
			Name:     "min-piece-size",
			Usage:    "Minimum piece size in bytes",
			Required: true,
		},
		&cli.Uint64Flag{
			Name:     "max-piece-size",
			Usage:    "Maximum piece size in bytes",
			Required: true,
		},
		&cli.BoolFlag{
			Name:     "ipni-piece",
			Usage:    "Supports IPNI piece",
			Required: false,
		},
		&cli.BoolFlag{
			Name:     "ipni-ipfs",
			Usage:    "Supports IPNI IPFS",
			Required: false,
		},
		&cli.Uint64Flag{
			Name:     "price",
			Usage:    "Monthly price per TiB",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "payment-token-address",
			Usage:    "Payment token address",
			Required: true,
		},
		&cli.Uint64Flag{
			Name:     "min-proving-period",
			Usage:    "Minimum proving period in epochs",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "location",
			Usage:    "Location of the provider",
			Required: false,
		},
	},
	Action: func(cctx *cli.Context) error {
		// serviceUrl := cctx.String("service-url")
		// minPieceSize := cctx.Uint64("min-piece-size")
		// maxPieceSize := cctx.Uint64("max-piece-size")
		// ipniPiece := cctx.Bool("ipni-piece")
		// ipniIpfs := cctx.Bool("ipni-ipfs")
		// price := cctx.Uint64("price")
		// paymentTokenAddress := cctx.String("payment-token-address")
		// minProvingPeriod := cctx.Uint64("min-proving-period")
		// location := cctx.String("location")
		return nil
	},
}
