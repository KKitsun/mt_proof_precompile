package prover

import (
	"embed"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	cmn "github.com/evmos/evmos/v16/precompiles/common"
)

var _ vm.PrecompiledContract = &Precompile{}

const (
	PrecompileAddress = "0x0000000000000000000000000000000000000888"
)

var f embed.FS

type Precompile struct {
	abi.ABI
	baseGas uint64
}

func NewPrecompile(baseGas uint64) (*Precompile, error) {
	newABI, err := cmn.LoadABI(f, "abi.json")
	if err != nil {
		return nil, err
	}

	if baseGas == 0 {
		return nil, fmt.Errorf("baseGas cannot be zero")
	}

	return &Precompile{
		ABI:     newABI,
		baseGas: baseGas,
	}, nil
}

func (Precompile) Address() common.Address {
	return common.HexToAddress(PrecompileAddress)
}

// RequiredGas calculates the contract gas use.
func (p Precompile) RequiredGas(_ []byte) uint64 {
	return p.baseGas
}

func (p Precompile) Run(_ *vm.EVM, contract *vm.Contract, _ bool) (bz []byte, err error) {
	methodID := contract.Input[:4]
	// NOTE: this function iterates over the method map and returns
	// the method with the given ID
	method, err := p.MethodById(methodID)
	if err != nil {
		return nil, err
	}

	argsBz := contract.Input[4:]
	args, err := method.Inputs.Unpack(argsBz)
	if err != nil {
		return nil, err
	}

	switch method.Name {
	case "some_method":
		break
	}

	if err != nil {
		return nil, err
	}

	return bz, nil
}
