// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

/**
 * @dev IStake contains all of the events necessary for staking Vega token
 */
interface IStake {
    event StakeDeposited(address indexed user, uint256 amount, bytes32 indexed vegaPublicKey);
    event StakeRemoved(address indexed user, uint256 amount, bytes32 indexed vegaPublicKey);
    event StakeTransferred(address indexed from, uint256 amount, address indexed to, bytes32 indexed vegaPublicKey);

    /// @return the address of the token that is able to be staked
    function stakingToken() external view returns (address);

    /// @param target Target address to check
    /// @param vegaPublicKey Target vega public key to check
    /// @return the number of tokens staked for that address->vegaPublicKey pair
    function stakeBalance(address target, bytes32 vegaPublicKey) external view returns (uint256);

    /// @return total tokens staked on contract
    function totalStaked() external view returns (uint256);
}