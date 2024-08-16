//SPDX-License-Identifier: MIT
pragma solidity ^0.8.8;

/// @title ERC20 Bridge Logic Interface
/// @author Vega Protocol
/// @notice Implementations of this interface are used by Vega network users to deposit and withdraw ERC20 tokens to/from Vega.
// @notice All funds deposited/withdrawn are to/from the ERC20_Asset_Pool
abstract contract IERC20BridgeLogicRestricted {
    /***************************EVENTS****************************/
    event AssetWithdrawn(address indexed userAddress, address indexed assetSource, uint256 amount, uint256 nonce);
    event AssetDeposited(
        address indexed userAddress,
        address indexed assetSource,
        uint256 amount,
        bytes32 vegaPublicKey
    );
    event AssetListed(address indexed assetSource, bytes32 indexed vegaAssetId, uint256 nonce);
    event AssetRemoved(address indexed assetSource, uint256 nonce);
    event AssetLimitsUpdated(address indexed assetSource, uint256 lifetimeLimit, uint256 withdrawThreshold);
    event BridgeWithdrawDelaySet(uint256 withdraw_delay);
    event BridgeStopped();
    event BridgeResumed();
    event DepositorExempted(address indexed depositor);
    event DepositorExemptionRevoked(address indexed depositor);

    /***************************FUNCTIONS*************************/
    /// @notice This function lists the given ERC20 token contract as valid for deposit to this bridge
    /// @param assetSource Contract address for given ERC20 token
    /// @param vegaAssetId Vega-generated asset ID for internal use in Vega Core
    /// @param lifetimeLimit Initial lifetime deposit limit *RESTRICTION FEATURE*
    /// @param withdrawThreshold Amount at which the withdraw delay goes into effect *RESTRICTION FEATURE*
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    /// @notice See MultisigControl for more about signatures
    /// @dev MUST emit AssetListed if successful
    function listAsset(
        address assetSource,
        bytes32 vegaAssetId,
        uint256 lifetimeLimit,
        uint256 withdrawThreshold,
        uint256 nonce,
        bytes memory signatures
    ) external virtual;

    /// @notice This function removes from listing the given ERC20 token contract. This marks the token as invalid for deposit to this bridge
    /// @param assetSource Contract address for given ERC20 token
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    /// @notice See MultisigControl for more about signatures
    /// @dev MUST emit AssetRemoved if successful
    function removeAsset(
        address assetSource,
        uint256 nonce,
        bytes memory signatures
    ) external virtual;

    /// @notice This function sets the lifetime maximum deposit for a given asset
    /// @param assetSource Contract address for given ERC20 token
    /// @param lifetimeLimit Deposit limit for a given ethereum address
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    /// @dev asset must first be listed
    function setAssetLimits(
        address assetSource,
        uint256 lifetimeLimit,
        uint256 threshold,
        uint256 nonce,
        bytes calldata signatures
    ) external virtual;

    /// @notice This function sets the withdraw delay for withdrawals over the per-asset set thresholds
    /// @param delay Amount of time to delay a withdrawal
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    function setWithdrawDelay(
        uint256 delay,
        uint256 nonce,
        bytes calldata signatures
    ) external virtual;

    /// @notice This function triggers the global bridge stop that halts all withdrawals and deposits until it is resumed
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    /// @dev bridge must not be stopped already
    /// @dev MUST emit BridgeStopped if successful
    function globalStop(uint256 nonce, bytes calldata signatures) external virtual;

    /// @notice This function resumes bridge operations from the stopped state
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    /// @dev bridge must be stopped
    /// @dev MUST emit BridgeResumed if successful
    function globalResume(uint256 nonce, bytes calldata signatures) external virtual;

    /// @notice this function allows the exemption_lister to exempt a depositor from the deposit limits
    /// @notice this feature is specifically for liquidity and rewards providers
    /// @dev MUST emit DepositorExempted if successful
    function exemptDepositor() external virtual;

    /// @notice this function allows the exemption_lister to revoke a depositor's exemption from deposit limits
    /// @notice this feature is specifically for liquidity and rewards providers
    /// @dev MUST emit DepositorExemptionRevoked if successful
    function revokeExemptDepositor() external virtual;

    /// @notice This function withdrawals assets to the target Ethereum address
    /// @param assetSource Contract address for given ERC20 token
    /// @param amount Amount of ERC20 tokens to withdraw
    /// @param target Target Ethereum address to receive withdrawn ERC20 tokens
    /// @param creation Timestamp of when requestion was created *RESTRICTION FEATURE*
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed order
    /// @notice See MultisigControl for more about signatures
    /// @dev MUST emit AssetWithdrawn if successful
    function withdrawAsset(
        address assetSource,
        uint256 amount,
        address target,
        uint256 creation,
        uint256 nonce,
        bytes memory signatures
    ) external virtual;

    /// @notice this view returns true if the given despoitor address has been exempted from deposit limits
    /// @param depositor The depositor to check
    /// @return true if depositor is exempt
    function isExemptDepositor(address depositor) external view virtual returns (bool);

    /// @notice This function allows a user to deposit given ERC20 tokens into Vega
    /// @param assetSource Contract address for given ERC20 token
    /// @param amount Amount of tokens to be deposited into Vega
    /// @param vegaPublicKey Target Vega public key to be credited with this deposit
    /// @dev MUST emit AssetDeposited if successful
    /// @dev ERC20 approve function should be run before running this
    /// @notice ERC20 approve function should be run before running this
    function depositAsset(
        address assetSource,
        uint256 amount,
        bytes32 vegaPublicKey
    ) external virtual;

    /***************************VIEWS*****************************/
    /// @notice This view returns true if the given ERC20 token contract has been listed valid for deposit
    /// @param assetSource Contract address for given ERC20 token
    /// @return True if asset is listed
    function isAssetListed(address assetSource) external view virtual returns (bool);

    /// @notice This view returns the lifetime deposit limit for the given asset
    /// @param assetSource Contract address for given ERC20 token
    /// @return Lifetime limit for the given asset
    function getAssetDepositLifetimeLimit(address assetSource) external view virtual returns (uint256);

    /// @notice This view returns the given token's withdraw threshold above which the withdraw delay goes into effect
    /// @param assetSource Contract address for given ERC20 token
    /// @return Withdraw threshold
    function getWithdrawThreshold(address assetSource) external view virtual returns (uint256);

    /// @return current multisig_control_address
    function getMultisigControlAddress() external view virtual returns (address);

    /// @param assetSource Contract address for given ERC20 token
    /// @return The assigned Vega Asset ID for given ERC20 token
    function getVegaAssetId(address assetSource) external view virtual returns (bytes32);

    /// @param vegaAssetId Vega-assigned asset ID for which you want the ERC20 token address
    /// @return The ERC20 token contract address for a given Vega Asset ID
    function getAssetSource(bytes32 vegaAssetId) external view virtual returns (address);
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
