// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./IERC20.sol";
import "./IStake.sol";

/// @title ERC20 Staking Bridge
/// @author Vega Protocol
/// @notice This contract manages the vesting of the Vega V2 ERC20 token
contract StakingBridge is IStake {
    address _stakingToken;

    constructor(address token) {
        _stakingToken = token;
    }

    /// @dev user => amount staked
    mapping(address => mapping(bytes32 => uint256)) stakes;

    /// @notice This stakes the given amount of tokens and credits them to the provided Vega public key
    /// @param amount Token amount to stake
    /// @param vegaPublicKey Target Vega public key to be credited with the stake
    /// @dev Emits StakeDeposited event
    /// @dev User MUST run "approve" on token prior to running Stake
    function stake(uint256 amount, bytes32 vegaPublicKey) public {
        require(IERC20(_stakingToken).transferFrom(msg.sender, address(this), amount));
        stakes[msg.sender][vegaPublicKey] += amount;
        emit StakeDeposited(msg.sender, amount, vegaPublicKey);
    }

    /// @notice This removes specified amount of stake of available to user
    /// @dev Emits StakeRemoved event if successful
    /// @param amount Amount of tokens to remove from staking
    /// @param vegaPublicKey Target Vega public key from which to deduct stake
    function removeStake(uint256 amount, bytes32 vegaPublicKey) public {
        stakes[msg.sender][vegaPublicKey] -= amount;
        require(IERC20(_stakingToken).transfer(msg.sender, amount));
        emit StakeRemoved(msg.sender, amount, vegaPublicKey);
    }

    /// @notice This transfers all stake from the sender's address to the "newAddress"
    /// @dev Emits Stake_Transfered event if successful
    /// @param amount Stake amount to transfer
    /// @param newAddress Target ETH address to recieve the stake
    /// @param vegaPublicKey Target Vega public key to be credited with the transfer
    function transferStake(uint256 amount, address newAddress, bytes32 vegaPublicKey) public {
        stakes[msg.sender][vegaPublicKey] -= amount;
        stakes[newAddress][vegaPublicKey] += amount;
        emit StakeTransferred(msg.sender, amount, newAddress, vegaPublicKey);
    }

    /// @dev This is IStake.stakingToken
    /// @return the address of the token that is able to be staked
    function stakingToken() external view override returns (address) {
        return _stakingToken;
    }

    /// @dev This is IStake.stakeBalance
    /// @param target Target address to check
    /// @param vegaPublicKey Target vega public key to check
    /// @return the number of tokens staked for that address->vegaPublicKey pair
    function stakeBalance(address target, bytes32 vegaPublicKey) external view override returns (uint256) {
        return stakes[target][vegaPublicKey];
    }

    /// @dev This is IStake.totalStaked
    /// @return total tokens staked on contract
    function totalStaked() external view override returns (uint256) {
        return IERC20(_stakingToken).balanceOf(address(this));
    }
}