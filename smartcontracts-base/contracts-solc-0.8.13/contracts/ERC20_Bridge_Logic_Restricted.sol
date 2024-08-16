//SPDX-License-Identifier: MIT
pragma solidity ^0.8.8;

import "./IERC20.sol";
import "./IERC20_Bridge_Logic_Restricted.sol";
import "./IMultisigControl.sol";
import {ERC20AssetPool} from "./ERC20_Asset_Pool.sol";

/// @title ERC20 Bridge Logic
/// @author Vega Protocol
/// @notice This contract is used by Vega network users to deposit and withdraw ERC20 tokens to/from Vega.
// @notice All funds deposited/withdrawn are to/from the assigned ERC20AssetPool
contract ERC20BridgeLogicRestricted is IERC20BridgeLogicRestricted {
    address payable public erc20AssetPoolAddress;
    // asset address => is listed
    mapping(address => bool) listedTokens;
    mapping(bytes32 => address) vegaAssetIdsToSource;
    // assetSource => Vega asset ID
    mapping(address => bytes32) assetSourceToVegaAssetId;

    /// @param erc20AssetPool Initial Asset Pool contract address
    constructor(address payable erc20AssetPool) {
        require(erc20AssetPool != address(0), "invalid asset pool address");
        erc20AssetPoolAddress = erc20AssetPool;
    }

    function multisigControlAddress() internal view returns (address) {
        return ERC20AssetPool(erc20AssetPoolAddress).multisigControlAddress();
    }

    /***************************FUNCTIONS*************************/
    /// @notice This function lists the given ERC20 token contract as valid for deposit to this bridge
    /// @param assetSource Contract address for given ERC20 token
    /// @param vegaAssetId Vega-generated asset ID for internal use in Vega Core
    /// @param lifetimeLimit Initial lifetime deposit limit *RESTRICTION FEATURE*. Setting this to type(uint256).max will disable the per address deposit limit counter
    /// @param withdrawThreshold Amount at which the withdraw delay goes into effect *RESTRICTION FEATURE*
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    /// @notice See MultisigControl for more about signatures
    /// @dev Emits AssetListed if successful
    function listAsset(
        address assetSource,
        bytes32 vegaAssetId,
        uint256 lifetimeLimit,
        uint256 withdrawThreshold,
        uint256 nonce,
        bytes memory signatures
    ) external override {
        require(assetSource != address(0), "invalid asset source");
        require(!listedTokens[assetSource], "asset already listed");
        bytes memory message = abi.encode(
            assetSource,
            vegaAssetId,
            lifetimeLimit,
            withdrawThreshold,
            nonce,
            "listAsset"
        );
        require(
            IMultisigControl(multisigControlAddress()).verifySignatures(signatures, message, nonce),
            "bad signatures"
        );
        listedTokens[assetSource] = true;
        vegaAssetIdsToSource[vegaAssetId] = assetSource;
        assetSourceToVegaAssetId[assetSource] = vegaAssetId;
        assetDepositLifetimeLimit[assetSource] = lifetimeLimit;
        withdrawThresholds[assetSource] = withdrawThreshold;
        emit AssetListed(assetSource, vegaAssetId, nonce);
    }

    /// @notice This function removes from listing the given ERC20 token contract. This marks the token as invalid for deposit to this bridge
    /// @param assetSource Contract address for given ERC20 token
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    /// @notice See MultisigControl for more about signatures
    /// @dev Emits Asset_Removed if successful
    function removeAsset(
        address assetSource,
        uint256 nonce,
        bytes memory signatures
    ) external override {
        require(listedTokens[assetSource], "asset not listed");
        bytes memory message = abi.encode(assetSource, nonce, "removeAsset");
        require(
            IMultisigControl(multisigControlAddress()).verifySignatures(signatures, message, nonce),
            "bad signatures"
        );
        listedTokens[assetSource] = false;
        emit AssetRemoved(assetSource, nonce);
    }

    /************************RESTRICTIONS***************************/
    // user => assetSource => deposit total
    mapping(address => mapping(address => uint256)) userLifetimeDeposits;
    // assetSource => deposit_limit
    mapping(address => uint256) assetDepositLifetimeLimit;
    uint256 public defaultWithdrawDelay = 432000;
    // assetSource => threshold
    mapping(address => uint256) withdrawThresholds;
    bool public is_stopped;

    // depositor => is exempt from deposit limits
    mapping(address => bool) exemptDepositors;

    /// @notice This function sets the lifetime maximum deposit for a given asset
    /// @param assetSource Contract address for given ERC20 token
    /// @param lifetimeLimit Deposit limit for a given ethereum address. Setting this to type(uint256).max will disable the per address deposit limit counter
    /// @param threshold Withdraw size above which the withdraw delay goes into effect
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    /// @dev asset must first be listed
    function setAssetLimits(
        address assetSource,
        uint256 lifetimeLimit,
        uint256 threshold,
        uint256 nonce,
        bytes calldata signatures
    ) external override {
        require(listedTokens[assetSource], "asset not listed");
        bytes memory message = abi.encode(assetSource, lifetimeLimit, threshold, nonce, "setAssetLimits");
        require(
            IMultisigControl(multisigControlAddress()).verifySignatures(signatures, message, nonce),
            "bad signatures"
        );
        assetDepositLifetimeLimit[assetSource] = lifetimeLimit;
        withdrawThresholds[assetSource] = threshold;

        emit AssetLimitsUpdated(assetSource, lifetimeLimit, threshold);
    }

    /// @notice This view returns the lifetime deposit limit for the given asset
    /// @param assetSource Contract address for given ERC20 token
    /// @return Lifetime limit for the given asset
    function getAssetDepositLifetimeLimit(address assetSource) external view override returns (uint256) {
        return assetDepositLifetimeLimit[assetSource];
    }

    /// @notice This view returns the given token's withdraw threshold above which the withdraw delay goes into effect
    /// @param assetSource Contract address for given ERC20 token
    /// @return Withdraw threshold
    function getWithdrawThreshold(address assetSource) external view override returns (uint256) {
        return withdrawThresholds[assetSource];
    }

    /// @notice This function sets the withdraw delay for withdrawals over the per-asset set thresholds
    /// @param delay Amount of time to delay a withdrawal
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    function setWithdrawDelay(
        uint256 delay,
        uint256 nonce,
        bytes calldata signatures
    ) external override {
        bytes memory message = abi.encode(delay, nonce, "setWithdrawDelay");
        require(
            IMultisigControl(multisigControlAddress()).verifySignatures(signatures, message, nonce),
            "bad signatures"
        );
        defaultWithdrawDelay = delay;
        emit BridgeWithdrawDelaySet(delay);
    }

    /// @notice This function triggers the global bridge stop that halts all withdrawals and deposits until it is resumed
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    /// @dev bridge must not be stopped already
    /// @dev emits Bridge_Stopped if successful
    function globalStop(uint256 nonce, bytes calldata signatures) external override {
        require(!is_stopped, "bridge already stopped");
        bytes memory message = abi.encode(nonce, "globalStop");
        require(
            IMultisigControl(multisigControlAddress()).verifySignatures(signatures, message, nonce),
            "bad signatures"
        );
        is_stopped = true;
        emit BridgeStopped();
    }

    /// @notice This function resumes bridge operations from the stopped state
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    /// @dev bridge must be stopped
    /// @dev emits Bridge_Resumed if successful
    function globalResume(uint256 nonce, bytes calldata signatures) external override {
        require(is_stopped, "bridge not stopped");
        bytes memory message = abi.encode(nonce, "globalResume");
        require(
            IMultisigControl(multisigControlAddress()).verifySignatures(signatures, message, nonce),
            "bad signatures"
        );
        is_stopped = false;
        emit BridgeResumed();
    }

    /// @notice this function allows the sender to exempt themselves from the deposit limits
    /// @notice this feature is specifically for liquidity and rewards providers
    /// @dev emits Depositor_Exempted if successful
    function exemptDepositor() external override {
        require(!exemptDepositors[msg.sender], "sender already exempt");
        exemptDepositors[msg.sender] = true;
        emit DepositorExempted(msg.sender);
    }

    /// @notice this function allows the exemption_lister to revoke a depositor's exemption from deposit limits
    /// @notice this feature is specifically for liquidity and rewards providers
    /// @dev emits Depositor_Exemption_Revoked if successful
    function revokeExemptDepositor() external override {
        require(exemptDepositors[msg.sender], "sender not exempt");
        exemptDepositors[msg.sender] = false;
        emit DepositorExemptionRevoked(msg.sender);
    }

    /// @notice this view returns true if the given despoitor address has been exempted from deposit limits
    /// @param depositor The depositor to check
    /// @return true if depositor is exempt
    function isExemptDepositor(address depositor) external view override returns (bool) {
        return exemptDepositors[depositor];
    }

    /***********************END RESTRICTIONS*************************/

    /// @notice This function withdraws assets to the target Ethereum address
    /// @param assetSource Contract address for given ERC20 token
    /// @param amount Amount of ERC20 tokens to withdraw
    /// @param target Target Ethereum address to receive withdrawn ERC20 tokens
    /// @param creation Timestamp of when requestion was created *RESTRICTION FEATURE*
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    /// @notice See MultisigControl for more about signatures
    /// @dev Emits AssetWithdrawn if successful
    function withdrawAsset(
        address assetSource,
        uint256 amount,
        address target,
        uint256 creation,
        uint256 nonce,
        bytes memory signatures
    ) external override {
        require(!is_stopped, "bridge stopped");
        require(
            withdrawThresholds[assetSource] > amount || creation + defaultWithdrawDelay <= block.timestamp,
            "large withdraw is not old enough"
        );
        bytes memory message = abi.encode(assetSource, amount, target, creation, nonce, "withdrawAsset");
        require(
            IMultisigControl(multisigControlAddress()).verifySignatures(signatures, message, nonce),
            "bad signatures"
        );
        ERC20AssetPool(erc20AssetPoolAddress).withdraw(assetSource, target, amount);
        emit AssetWithdrawn(target, assetSource, amount, nonce);
    }

    /// @notice This function allows a user to deposit given ERC20 tokens into Vega
    /// @param assetSource Contract address for given ERC20 token
    /// @param amount Amount of tokens to be deposited into Vega
    /// @param vegaPublicKey Target Vega public key to be credited with this deposit
    /// @dev emits Asset_Deposited if successful
    /// @dev ERC20 approve function should be run before running this
    /// @notice ERC20 approve function should be run before running this
    function depositAsset(
        address assetSource,
        uint256 amount,
        bytes32 vegaPublicKey
    ) external override {
        require(!is_stopped, "bridge stopped");

        // Cache SLOAD
        uint256 _limit = assetDepositLifetimeLimit[assetSource];

        // Check limit first as that's the most likely branch, then check if exempt
        if (_limit < type(uint256).max && !exemptDepositors[msg.sender]) {
            require(userLifetimeDeposits[msg.sender][assetSource] + amount <= _limit, "deposit over lifetime limit");
            userLifetimeDeposits[msg.sender][assetSource] += amount;
        }

        require(listedTokens[assetSource], "asset not listed");

        (bool success, bytes memory returndata) = assetSource.call(
            abi.encodeWithSignature(
                "transferFrom(address,address,uint256)",
                msg.sender,
                erc20AssetPoolAddress,
                amount
            )
        );
        require(success, "token transfer failed");

        if (returndata.length > 0) {
            // Return data is optional
            require(abi.decode(returndata, (bool)), "token transfer failed");
        }

        emit AssetDeposited(msg.sender, assetSource, amount, vegaPublicKey);
    }


    /***************************VIEWS*****************************/
    /// @notice This view returns true if the given ERC20 token contract has been listed valid for deposit
    /// @param assetSource Contract address for given ERC20 token
    /// @return True if asset is listed
    function isAssetListed(address assetSource) external view override returns (bool) {
        return listedTokens[assetSource];
    }

    /// @return current multisigControlAddress
    function getMultisigControlAddress() external view override returns (address) {
        return multisigControlAddress();
    }

    /// @param assetSource Contract address for given ERC20 token
    /// @return The assigned Vega Asset Id for given ERC20 token
    function getVegaAssetId(address assetSource) external view override returns (bytes32) {
        return assetSourceToVegaAssetId[assetSource];
    }

    /// @param vegaAssetId Vega-assigned asset ID for which you want the ERC20 token address
    /// @return The ERC20 token contract address for a given Vega Asset Id
    function getAssetSource(bytes32 vegaAssetId) external view override returns (address) {
        return vegaAssetIdsToSource[vegaAssetId];
    }

    function isContract(address addr) internal view returns (bool) {
        uint256 code_size;
        assembly {
            code_size := extcodesize(addr)
        }
        return code_size > 0;
    }
}

/**
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMWEMMMMMMMMMMMMMMMMMMMMMMMMMM...............MMMMMMMMMMMMM
MMMMMMLOVEMMMMMMMMMMMMMMMMMMMMMM...............MMMMMMMMMMMMM
MMMMMMMMMMHIXELMMMMMMMMMMMM....................MMMMMNNMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMM....................MMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMM88=........................+MMMMMMMMMM
MMMMMMMMMMMMMMMMM....................MMMMM...MMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMM....................MMMMM...MMMMMMMMMMMMMMM
MMMMMMMMMMMM.........................MM+..MMM....+MMMMMMMMMM
MMMMMMMMMNMM...................... ..MM?..MMM.. .+MMMMMMMMMM
MMMMNDDMM+........................+MM........MM..+MMMMMMMMMM
MMMMZ.............................+MM....................MMM
MMMMZ.............................+MM....................MMM
MMMMZ.............................+MM....................DDD
MMMMZ.............................+MM..ZMMMMMMMMMMMMMMMMMMMM
MMMMZ.............................+MM..ZMMMMMMMMMMMMMMMMMMMM
MM..............................MMZ....ZMMMMMMMMMMMMMMMMMMMM
MM............................MM.......ZMMMMMMMMMMMMMMMMMMMM
MM............................MM.......ZMMMMMMMMMMMMMMMMMMMM
MM......................ZMMMMM.......MMMMMMMMMMMMMMMMMMMMMMM
MM............... ......ZMMMMM.... ..MMMMMMMMMMMMMMMMMMMMMMM
MM...............MMMMM88~.........+MM..ZMMMMMMMMMMMMMMMMMMMM
MM.......$DDDDDDD.......$DDDDD..DDNMM..ZMMMMMMMMMMMMMMMMMMMM
MM.......$DDDDDDD.......$DDDDD..DDNMM..ZMMMMMMMMMMMMMMMMMMMM
MM.......ZMMMMMMM.......ZMMMMM..MMMMM..ZMMMMMMMMMMMMMMMMMMMM
MMMMMMMMM+.......MMMMM88NMMMMM..MMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMM+.......MMMMM88NMMMMM..MMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM*/
