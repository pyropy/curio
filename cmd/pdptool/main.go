package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	curiobuild "github.com/filecoin-project/curio/build"
)

func main() {
	app := &cli.App{
		Name:    "pdptool",
		Usage:   "tool for testing PDP capabilities",
		Version: curiobuild.UserVersion(),
		Commands: []*cli.Command{
			// pdp commands
			authCreateServiceSecretCmd, // generates pdpservice.json, outputs pubkey
			authCreateJWTTokenCmd,      // generates jwt token from a secret

			pingCmd,

			piecePrepareCmd, // hash a piece to get a piece cid
			pieceUploadCmd,  // upload a piece to a pdp service
			uploadFileCmd,   // upload a file to a pdp service in many chunks
			downloadFileCmd, // download a file from curio

			createDataSetCmd,    // create a new data set on the PDP service
			getDataSetStatusCmd, // get the status of a data set creation on the PDP service
			getDataSetCmd,       // retrieve the details of a data set from the PDP service

			addPiecesCmd,
			removePiecesCmd, // schedule pieces for removal after next proof submission

			// service registry commands
			registerProvider,
		},
	}
	app.Setup()

	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
