// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.20;

struct Signature {
    uint256 r;
    uint256 s;
    uint8 v;
}

bytes32 constant EIP712_DOMAIN_SEPARATOR = keccak256(
    "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
);

library SignatureLib {
    /**
     * @dev Recovers the signer address from a signature
     * @param dataHash The hash of the data that was signed
     * @param sig The signature
     * @param domainSeparator The EIP-712 domain separator
     * @return The recovered signer address
     */
    function recoverSigner(
        bytes32 dataHash,
        Signature memory sig,
        bytes32 domainSeparator
    ) internal pure returns (address) {
        // Validate signature parameters
        require(sig.r != 0, "Invalid signature 'r' value");
        require(sig.s != 0, "Invalid signature 's' value");

        // Build EIP-712 digest
        bytes32 digest = keccak256(
            abi.encodePacked("\x19\x01", domainSeparator, dataHash)
        );

        // Recover signer
        address signerRecovered = ecrecover(
            digest,
            sig.v,
            bytes32(sig.r),
            bytes32(sig.s)
        );

        require(
            signerRecovered != address(0),
            "Invalid signature, recovered the zero address"
        );

        return signerRecovered;
    }
}
