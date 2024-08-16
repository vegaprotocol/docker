//SPDX-License-Identifier: MIT
pragma solidity ^0.8.8;

import "./IMultisigControl.sol";
import "./IERC20.sol";

/// @title ERC20 Asset Pool
/// @author Vega Protocol
/// @notice This contract is the target for all deposits to the ERC20 Bridge via ERC20_Bridge_Logic
contract ERC20AssetPool {
    event MultisigControlSet(address indexed newAddress);
    event BridgeAddressSet(address indexed newAddress);

    /// @return Current MultisigControl contract address
    address public multisigControlAddress;

    /// @return Current ERC20_Bridge_Logic contract address
    address public erc20BridgeAddress;

    /// @param multisigControl The initial MultisigControl contract address
    /// @notice Emits MultisigControlSet event
    constructor(address multisigControl) {
        require(multisigControl != address(0), "invalid MultisigControl address");
        multisigControlAddress = multisigControl;
        emit MultisigControlSet(multisigControl);
    }

    /// @notice this contract is not intended to accept ether directly
    receive() external payable {
        revert("this contract does not accept ETH");
    }

    /// @param newAddress The new MultisigControl contract address.
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed set_multisigControl order
    /// @notice See MultisigControl for more about signatures
    /// @notice Emits MultisigControlSet event
    function setMultisigControl(
        address newAddress,
        uint256 nonce,
        bytes memory signatures
    ) external {
        require(newAddress != address(0), "invalid MultisigControl address");
        require(is_contract(newAddress), "new address must be contract");

        bytes memory message = abi.encode(newAddress, nonce, "setMultisigControl");
        require(
            IMultisigControl(multisigControlAddress).verifySignatures(signatures, message, nonce),
            "bad signatures"
        );
        multisigControlAddress = newAddress;
        emit MultisigControlSet(newAddress);
    }

    /// @param newAddress The new ERC20_Bridge_Logic contract address.
    /// @param nonce Vega-assigned single-use number that provides replay attack protection
    /// @param signatures Vega-supplied signature bundle of a validator-signed set_bridge_address order
    /// @notice See MultisigControl for more about signatures
    /// @notice Emits BridgeAddressSet event
    function setBridgeAddress(
        address newAddress,
        uint256 nonce,
        bytes memory signatures
    ) external {
        bytes memory message = abi.encode(newAddress, nonce, "setBridgeAddress");
        require(
            IMultisigControl(multisigControlAddress).verifySignatures(signatures, message, nonce),
            "bad signatures"
        );
        erc20BridgeAddress = newAddress;
        emit BridgeAddressSet(newAddress);
    }

    /// @notice This function can only be run by the current "multisigControlAddress" and, if available, will send the target tokens to the target
    /// @param token_address Contract address of the ERC20 token to be withdrawn
    /// @param target Target Ethereum address that the ERC20 tokens will be sent to
    /// @param amount Amount of ERC20 tokens to withdraw
    /// @dev amount is in whatever the lowest decimal value the ERC20 token has. For instance, an 18 decimal ERC20 token, 1 "amount" == 0.000000000000000001
    function withdraw(
        address token_address,
        address target,
        uint256 amount
    ) external {
        require(msg.sender == erc20BridgeAddress, "msg.sender not authorized bridge");

        (bool success, bytes memory returndata) = token_address.call(
            abi.encodeWithSignature("transfer(address,uint256)", target, amount)
        );
        require(success, "token transfer failed");

        if (returndata.length > 0) {
            // Return data is optional
            require(abi.decode(returndata, (bool)), "token transfer failed");
        }
    }

    function is_contract(address addr) internal view returns (bool) {
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
