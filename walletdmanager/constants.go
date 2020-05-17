// Copyright (c) 2018, The TurtleCoin Developers
//
// Please see the included LICENSE file for more information.
//

package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in TRTL
	DefaultTransferFee float64 = 0.1

	logWalletdCurrentSessionFilename     = "treble-service-session.log"
	logWalletdAllSessionsFilename        = "treble-service.log"
	logTurtleCoindCurrentSessionFilename = "TrebleCoind-session.log"
	logTurtleCoindAllSessionsFilename    = "TrebleCoind.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "treble-service"
	turtlecoindCommandName               = "TrebleCoind"
)
