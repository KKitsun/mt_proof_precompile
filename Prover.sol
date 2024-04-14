pragma solidity >=0.8.17;

address constant Prover_PRECOMPILE_ADDRESS = 0x0000000000000000000000000000000000000400;

ProverI constant PROVER_CONTRACT = ProverI(Prover_PRECOMPILE_ADDRESS);

interface ProverI {

}