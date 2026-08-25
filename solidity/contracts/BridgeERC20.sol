// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";

contract BridgeERC20 is ERC20, ERC20Permit, AccessControl, Pausable {
    event Minted(address indexed to, uint256 amount);
    event Burned(address indexed from, uint256 amount);
    event Blacklisted(address indexed account);
    event UnBlacklisted(address indexed account);

    mapping(address => bool) private _blacklisted;

    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant ADMIN_ROLE = DEFAULT_ADMIN_ROLE;

    constructor(
        string memory name,
        string memory symbol
    ) ERC20(name, symbol) ERC20Permit(name) {
        _grantRole(ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);
    }

    function decimals() public pure override returns (uint8) {
        return 18;
    }

    function setMiner(
        address miner,
        bool isActive
    ) external onlyRole(ADMIN_ROLE) {
        require(miner != address(0), "Invalid miner address");
        if (isActive) {
            _grantRole(MINTER_ROLE, miner);
        } else {
            _revokeRole(MINTER_ROLE, miner);
        }
    }

    function mint(address to, uint256 amount) external onlyRole(MINTER_ROLE) {
        _mint(to, amount);
        emit Minted(to, amount);
    }

    function burn(uint256 amount) external {
        _burn(msg.sender, amount);
        emit Burned(msg.sender, amount);
    }

    function pause() external onlyRole(ADMIN_ROLE) {
        _pause();
    }

    function unpause() external onlyRole(ADMIN_ROLE) {
        _unpause();
    }

    function addToBlacklist(address account) external onlyRole(ADMIN_ROLE) {
        require(account != address(0), "Invalid address");
        require(!_blacklisted[account], "Already blacklisted");
        _blacklisted[account] = true;
        emit Blacklisted(account);
    }

    function removeFromBlacklist(
        address account
    ) external onlyRole(ADMIN_ROLE) {
        require(account != address(0), "Invalid address");
        require(_blacklisted[account], "Not blacklisted");
        _blacklisted[account] = false;
        emit UnBlacklisted(account);
    }

    function isBlacklisted(address account) external view returns (bool) {
        return _blacklisted[account];
    }

    function _update(
        address from,
        address to,
        uint256 value
    ) internal override(ERC20) {
        require(!paused(), "Token is paused");
        require(!_blacklisted[from], "Sender is blacklisted");
        require(!_blacklisted[to], "Recipient is blacklisted");
        super._update(from, to, value);
    }
}
