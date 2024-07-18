package simulation

import (
	"context"
)

type SimulationService interface {
	SimulateTransaction(
		context.Context,
		string,
		string,
		string,
		string,
		string,
	) (string, error)

	SimulateMessage(
		context.Context,
		string,
		string,
		string,
	) (string, error)
}
