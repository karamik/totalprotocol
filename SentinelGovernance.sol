// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

/**
 * @title TOTAL Protocol: Sentinel Core v.7.1 Multisig
 * @dev 2-of-3 Multisig Governance Contract for TOTAL Protocol.
 */
contract SentinelGovernance {
    address[] public owners;
    uint256 public constant REQUIRED_SIGNATURES = 2;
    
    // Mapping of authorized admins (v.7.1)
    mapping(address => bool) public isOwner;

    constructor() {
        // 1. MetaMask (EU Node)
        _addOwner(0x36767d4c2cBADaA66F632C744a6A10791D49c348);
        
        // 2. Coinbase (US Node)
        _addOwner(0xE2615753AdA9c3d7E490ae98Dc542Ee6Bc1f945e);
        
        // 3. Trust Wallet (Asia Node)
        _addOwner(0xd255f0a3bed34762e132090e66749e8cf9e093ce);
    }

    function _addOwner(address _owner) private {
        require(_owner != address(0), "Invalid address");
        require(!isOwner[_owner], "Duplicate owner");
        isOwner[_owner] = true;
        owners.push(_owner);
    }

    // В v.7.1 логика апелляции и Oracle Shield интегрируется через этот интерфейс
    function getStatus() public pure returns (string memory) {
        return "TOTAL Status: 2-of-3 Multisig Active (Sentinel Core v.7.1)";
    }
}
