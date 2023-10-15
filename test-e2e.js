const { ethers } = require("ethers");
const crypto = require('crypto');
const provider = new ethers.JsonRpcProvider(`https://exchaintestrpc.okex.org`)

// 签名人6的地址 私钥
const trader6 = "0x666666D834C881ECCcda9fB5441327AaC81A8E5b";
const privateKey_6 = '0xd77c872acbc7ecdd50830457b760331b2a8dfc87757f586d2463664a9b4b2729';
const signer_6 = new ethers.Wallet(privateKey_6, provider);

// 签名人7的地址 私钥
const trader7 = "0x7777779A61DDAd61f89EF4b3a0AbF152f6083974";
const privateKey_7 = '0xbb36bc07cc4cf94f77fb3ab1405c10cfcf3a456e0ecdacaf3d1d20d3202e9190';
const signer_7 = new ethers.Wallet(privateKey_7, provider);

// 签名人8的地址 私钥
const trader8 = "0x88888893C7e180D51B2557F76Eba35cff528b79C";
const privateKey_8 = '0x0d8dc6638c8a9d61cd33231d9926ec9ab7e730511582e5af662aaaf60e3305eb';
const signer_8 = new ethers.Wallet(privateKey_8, provider);

// NFT地址
const collection721 = "0x3C8C4bCEc2aA5Bbd9E2A29c014849F489370cA36";
const collection1155 = "0x58812C0bB265a5c696F26a9d08251CEf4dD4847A";

// 支付代币地址
const WOKT9 = "0x2219845942d28716c0F7C605765fABDcA1a7d9E0";

// 手续费率(0.5%)，接收地址
const feeRate = 50;
const feeRecipient = "0xCd18BE3282cC6e4AB072Bab888773B6e10688888";

const abi = [
    {
        "inputs": [],
        "stateMutability": "nonpayable",
        "type": "constructor"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "address",
                "name": "previousAdmin",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "newAdmin",
                "type": "address"
            }
        ],
        "name": "AdminChanged",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "beacon",
                "type": "address"
            }
        ],
        "name": "BeaconUpgraded",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [],
        "name": "Closed",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint8",
                "name": "version",
                "type": "uint8"
            }
        ],
        "name": "Initialized",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "feeRate",
                "type": "uint256"
            }
        ],
        "name": "NewFeeRate",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "address",
                "name": "feeRecipient",
                "type": "address"
            }
        ],
        "name": "NewFeeRecipient",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "address",
                "name": "governor",
                "type": "address"
            }
        ],
        "name": "NewGovernor",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "contract ITransferDelegate",
                "name": "transferDelegate",
                "type": "address"
            }
        ],
        "name": "NewTransferDelegate",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "trader",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "newNonce",
                "type": "uint256"
            }
        ],
        "name": "NonceIncremented",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [],
        "name": "Opened",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "bytes32",
                "name": "hash",
                "type": "bytes32"
            }
        ],
        "name": "OrderCancelled",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "maker",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "taker",
                "type": "address"
            },
            {
                "components": [
                    {
                        "internalType": "address",
                        "name": "trader",
                        "type": "address"
                    },
                    {
                        "internalType": "enum Side",
                        "name": "side",
                        "type": "uint8"
                    },
                    {
                        "internalType": "address",
                        "name": "collection",
                        "type": "address"
                    },
                    {
                        "internalType": "enum AssetType",
                        "name": "assetType",
                        "type": "uint8"
                    },
                    {
                        "internalType": "uint256",
                        "name": "tokenId",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "amount",
                        "type": "uint256"
                    },
                    {
                        "internalType": "address",
                        "name": "paymentToken",
                        "type": "address"
                    },
                    {
                        "internalType": "uint256",
                        "name": "price",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "listingTime",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "expirationTime",
                        "type": "uint256"
                    },
                    {
                        "components": [
                            {
                                "internalType": "uint16",
                                "name": "rate",
                                "type": "uint16"
                            },
                            {
                                "internalType": "address payable",
                                "name": "recipient",
                                "type": "address"
                            }
                        ],
                        "internalType": "struct Fee[]",
                        "name": "fees",
                        "type": "tuple[]"
                    },
                    {
                        "internalType": "uint256",
                        "name": "salt",
                        "type": "uint256"
                    }
                ],
                "indexed": false,
                "internalType": "struct Order",
                "name": "sell",
                "type": "tuple"
            },
            {
                "indexed": false,
                "internalType": "bytes32",
                "name": "sellHash",
                "type": "bytes32"
            },
            {
                "components": [
                    {
                        "internalType": "address",
                        "name": "trader",
                        "type": "address"
                    },
                    {
                        "internalType": "enum Side",
                        "name": "side",
                        "type": "uint8"
                    },
                    {
                        "internalType": "address",
                        "name": "collection",
                        "type": "address"
                    },
                    {
                        "internalType": "enum AssetType",
                        "name": "assetType",
                        "type": "uint8"
                    },
                    {
                        "internalType": "uint256",
                        "name": "tokenId",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "amount",
                        "type": "uint256"
                    },
                    {
                        "internalType": "address",
                        "name": "paymentToken",
                        "type": "address"
                    },
                    {
                        "internalType": "uint256",
                        "name": "price",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "listingTime",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "expirationTime",
                        "type": "uint256"
                    },
                    {
                        "components": [
                            {
                                "internalType": "uint16",
                                "name": "rate",
                                "type": "uint16"
                            },
                            {
                                "internalType": "address payable",
                                "name": "recipient",
                                "type": "address"
                            }
                        ],
                        "internalType": "struct Fee[]",
                        "name": "fees",
                        "type": "tuple[]"
                    },
                    {
                        "internalType": "uint256",
                        "name": "salt",
                        "type": "uint256"
                    }
                ],
                "indexed": false,
                "internalType": "struct Order",
                "name": "buy",
                "type": "tuple"
            },
            {
                "indexed": false,
                "internalType": "bytes32",
                "name": "buyHash",
                "type": "bytes32"
            }
        ],
        "name": "OrdersMatched",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "previousOwner",
                "type": "address"
            },
            {
                "indexed": true,
                "internalType": "address",
                "name": "newOwner",
                "type": "address"
            }
        ],
        "name": "OwnershipTransferred",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": true,
                "internalType": "address",
                "name": "implementation",
                "type": "address"
            }
        ],
        "name": "Upgraded",
        "type": "event"
    },
    {
        "inputs": [],
        "name": "FEE_TYPEHASH",
        "outputs": [
            {
                "internalType": "bytes32",
                "name": "",
                "type": "bytes32"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "INVERSE_BASIS_POINT",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "MFIBO",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "NAME",
        "outputs": [
            {
                "internalType": "string",
                "name": "",
                "type": "string"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "ORDER_TYPEHASH",
        "outputs": [
            {
                "internalType": "bytes32",
                "name": "",
                "type": "bytes32"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "VERSION",
        "outputs": [
            {
                "internalType": "string",
                "name": "",
                "type": "string"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "WFIBO",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "components": [
                    {
                        "internalType": "address",
                        "name": "trader",
                        "type": "address"
                    },
                    {
                        "internalType": "enum Side",
                        "name": "side",
                        "type": "uint8"
                    },
                    {
                        "internalType": "address",
                        "name": "collection",
                        "type": "address"
                    },
                    {
                        "internalType": "enum AssetType",
                        "name": "assetType",
                        "type": "uint8"
                    },
                    {
                        "internalType": "uint256",
                        "name": "tokenId",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "amount",
                        "type": "uint256"
                    },
                    {
                        "internalType": "address",
                        "name": "paymentToken",
                        "type": "address"
                    },
                    {
                        "internalType": "uint256",
                        "name": "price",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "listingTime",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "expirationTime",
                        "type": "uint256"
                    },
                    {
                        "components": [
                            {
                                "internalType": "uint16",
                                "name": "rate",
                                "type": "uint16"
                            },
                            {
                                "internalType": "address payable",
                                "name": "recipient",
                                "type": "address"
                            }
                        ],
                        "internalType": "struct Fee[]",
                        "name": "fees",
                        "type": "tuple[]"
                    },
                    {
                        "internalType": "uint256",
                        "name": "salt",
                        "type": "uint256"
                    }
                ],
                "internalType": "struct Order",
                "name": "order",
                "type": "tuple"
            }
        ],
        "name": "cancelOrder",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "components": [
                    {
                        "internalType": "address",
                        "name": "trader",
                        "type": "address"
                    },
                    {
                        "internalType": "enum Side",
                        "name": "side",
                        "type": "uint8"
                    },
                    {
                        "internalType": "address",
                        "name": "collection",
                        "type": "address"
                    },
                    {
                        "internalType": "enum AssetType",
                        "name": "assetType",
                        "type": "uint8"
                    },
                    {
                        "internalType": "uint256",
                        "name": "tokenId",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "amount",
                        "type": "uint256"
                    },
                    {
                        "internalType": "address",
                        "name": "paymentToken",
                        "type": "address"
                    },
                    {
                        "internalType": "uint256",
                        "name": "price",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "listingTime",
                        "type": "uint256"
                    },
                    {
                        "internalType": "uint256",
                        "name": "expirationTime",
                        "type": "uint256"
                    },
                    {
                        "components": [
                            {
                                "internalType": "uint16",
                                "name": "rate",
                                "type": "uint16"
                            },
                            {
                                "internalType": "address payable",
                                "name": "recipient",
                                "type": "address"
                            }
                        ],
                        "internalType": "struct Fee[]",
                        "name": "fees",
                        "type": "tuple[]"
                    },
                    {
                        "internalType": "uint256",
                        "name": "salt",
                        "type": "uint256"
                    }
                ],
                "internalType": "struct Order[]",
                "name": "orders",
                "type": "tuple[]"
            }
        ],
        "name": "cancelOrders",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "bytes32",
                "name": "",
                "type": "bytes32"
            }
        ],
        "name": "cancelledOrFilled",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "close",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "components": [
                    {
                        "components": [
                            {
                                "internalType": "address",
                                "name": "trader",
                                "type": "address"
                            },
                            {
                                "internalType": "enum Side",
                                "name": "side",
                                "type": "uint8"
                            },
                            {
                                "internalType": "address",
                                "name": "collection",
                                "type": "address"
                            },
                            {
                                "internalType": "enum AssetType",
                                "name": "assetType",
                                "type": "uint8"
                            },
                            {
                                "internalType": "uint256",
                                "name": "tokenId",
                                "type": "uint256"
                            },
                            {
                                "internalType": "uint256",
                                "name": "amount",
                                "type": "uint256"
                            },
                            {
                                "internalType": "address",
                                "name": "paymentToken",
                                "type": "address"
                            },
                            {
                                "internalType": "uint256",
                                "name": "price",
                                "type": "uint256"
                            },
                            {
                                "internalType": "uint256",
                                "name": "listingTime",
                                "type": "uint256"
                            },
                            {
                                "internalType": "uint256",
                                "name": "expirationTime",
                                "type": "uint256"
                            },
                            {
                                "components": [
                                    {
                                        "internalType": "uint16",
                                        "name": "rate",
                                        "type": "uint16"
                                    },
                                    {
                                        "internalType": "address payable",
                                        "name": "recipient",
                                        "type": "address"
                                    }
                                ],
                                "internalType": "struct Fee[]",
                                "name": "fees",
                                "type": "tuple[]"
                            },
                            {
                                "internalType": "uint256",
                                "name": "salt",
                                "type": "uint256"
                            }
                        ],
                        "internalType": "struct Order",
                        "name": "order",
                        "type": "tuple"
                    },
                    {
                        "internalType": "uint8",
                        "name": "v",
                        "type": "uint8"
                    },
                    {
                        "internalType": "bytes32",
                        "name": "r",
                        "type": "bytes32"
                    },
                    {
                        "internalType": "bytes32",
                        "name": "s",
                        "type": "bytes32"
                    }
                ],
                "internalType": "struct SignedOrder",
                "name": "sell",
                "type": "tuple"
            },
            {
                "components": [
                    {
                        "components": [
                            {
                                "internalType": "address",
                                "name": "trader",
                                "type": "address"
                            },
                            {
                                "internalType": "enum Side",
                                "name": "side",
                                "type": "uint8"
                            },
                            {
                                "internalType": "address",
                                "name": "collection",
                                "type": "address"
                            },
                            {
                                "internalType": "enum AssetType",
                                "name": "assetType",
                                "type": "uint8"
                            },
                            {
                                "internalType": "uint256",
                                "name": "tokenId",
                                "type": "uint256"
                            },
                            {
                                "internalType": "uint256",
                                "name": "amount",
                                "type": "uint256"
                            },
                            {
                                "internalType": "address",
                                "name": "paymentToken",
                                "type": "address"
                            },
                            {
                                "internalType": "uint256",
                                "name": "price",
                                "type": "uint256"
                            },
                            {
                                "internalType": "uint256",
                                "name": "listingTime",
                                "type": "uint256"
                            },
                            {
                                "internalType": "uint256",
                                "name": "expirationTime",
                                "type": "uint256"
                            },
                            {
                                "components": [
                                    {
                                        "internalType": "uint16",
                                        "name": "rate",
                                        "type": "uint16"
                                    },
                                    {
                                        "internalType": "address payable",
                                        "name": "recipient",
                                        "type": "address"
                                    }
                                ],
                                "internalType": "struct Fee[]",
                                "name": "fees",
                                "type": "tuple[]"
                            },
                            {
                                "internalType": "uint256",
                                "name": "salt",
                                "type": "uint256"
                            }
                        ],
                        "internalType": "struct Order",
                        "name": "order",
                        "type": "tuple"
                    },
                    {
                        "internalType": "uint8",
                        "name": "v",
                        "type": "uint8"
                    },
                    {
                        "internalType": "bytes32",
                        "name": "r",
                        "type": "bytes32"
                    },
                    {
                        "internalType": "bytes32",
                        "name": "s",
                        "type": "bytes32"
                    }
                ],
                "internalType": "struct SignedOrder",
                "name": "buy",
                "type": "tuple"
            }
        ],
        "name": "execute",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "feeRate",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "feeRecipient",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "governor",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "incrementNonce",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "contract ITransferDelegate",
                "name": "_transferDelegate",
                "type": "address"
            }
        ],
        "name": "initialize",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "isOpen",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "name": "nonces",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "open",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "owner",
        "outputs": [
            {
                "internalType": "address",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "proxiableUUID",
        "outputs": [
            {
                "internalType": "bytes32",
                "name": "",
                "type": "bytes32"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "renounceOwnership",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "_feeRate",
                "type": "uint256"
            }
        ],
        "name": "setFeeRate",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_feeRecipient",
                "type": "address"
            }
        ],
        "name": "setFeeRecipient",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "_governor",
                "type": "address"
            }
        ],
        "name": "setGovernor",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "contract ITransferDelegate",
                "name": "_transferDelegate",
                "type": "address"
            }
        ],
        "name": "setTransferDelegate",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "transferDelegate",
        "outputs": [
            {
                "internalType": "contract ITransferDelegate",
                "name": "",
                "type": "address"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "newOwner",
                "type": "address"
            }
        ],
        "name": "transferOwnership",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "newImplementation",
                "type": "address"
            }
        ],
        "name": "upgradeTo",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "newImplementation",
                "type": "address"
            },
            {
                "internalType": "bytes",
                "name": "data",
                "type": "bytes"
            }
        ],
        "name": "upgradeToAndCall",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    }
]

function generateSalt() {
    const randomBytes = crypto.randomBytes(32);
    return "0x" + randomBytes.toString('hex');
}


const main = async () => { 

    const verifyingContract = "0x50047A8658baBADA196C395589A8Fc03178fD282";

    // 实例化合约
    const instance = new ethers.Contract(verifyingContract, abi, signer_6);

    const chainId = (await provider.getNetwork()).chainId;

    // EIP712标准签名格式。
    const domain = {
        name: "Mirror Exchange",
        version: "1.0",
        chainId: chainId,
        verifyingContract: verifyingContract,
    };

    //"Order(address trader,uint8 side,address collection,uint8 assetType,uint256 tokenId,uint256 amount,address paymentToken,
    //uint256 price,uint256 listingTime,uint256 expirationTime,Fee[] fees,uint256 salt,uint256 nonce)Fee(uint16 rate,address recipient)"
    const orderTypes = {
        Order: [
            { name: "trader", type: "address" },
            { name: "side", type: "uint8" },
            { name: "collection", type: "address" },
            { name: "assetType", type: "uint8" },
            { name: "tokenId", type: "uint256" },
            { name: "amount", type: "uint256" },
            { name: "paymentToken", type: "address" },
            { name: "price", type: "uint256" },
            { name: "listingTime", type: "uint256" },
            { name: "expirationTime", type: "uint256" },
            { name: "fees", type: "Fee[]" },
            { name: "salt", type: "uint256" },
            { name: "nonce", type: "uint256" },
        ],
        Fee: [
            { name: "rate", type: "uint16" },
            { name: "recipient", type: "address" },
        ]
    };

    // const saltSell = generateSalt();
    // const saltBuy = generateSalt();

    const saltSell = "0x7a14e8fcda79940f6f4584a7c2e67433ea4245cc6e2d49134e752b6a7b6e0d23";
    const saltBuy = "0x3a1f8d7e13f5f66eb840b0beabf9232a03202f9614d3c1cb6642a6f045bc1e25";

    // 卖家8先挂单。
    const sellOrderMessage = {
        trader: trader8,
        side: 1,
        collection: collection721,
        assetType: 0,
        tokenId: 0,
        amount: 1,
        paymentToken: WOKT9,
        price: 100,
        listingTime: 1695627793,
        expirationTime: 8888888888,
        fees: [
            {
                rate: feeRate,
                recipient: feeRecipient
            }
        ],
        salt: saltSell,
        nonce: 0
    }

    // 买家6吃单
    const buyOrderMessage = {
        trader: trader6,
        side: 0,
        collection: collection721,
        assetType: 0,
        tokenId: 0,
        amount: 1,
        paymentToken: WOKT9,
        price: 100,
        listingTime: 1695628888,
        expirationTime: 8888888888,
        fees: [
            {
                rate: feeRate,
                recipient: feeRecipient
            }
        ],
        salt: saltBuy,
        nonce: 0
    }

    // 卖家8签名
    const signature8 = await signer_8.signTypedData(domain, orderTypes, sellOrderMessage);
    const sig8 = ethers.Signature.from(signature8);
    const packedSig8 = ethers.solidityPacked(['bytes32','bytes32', 'uint8'],[sig8.r, sig8.s, sig8.v])

    console.log(packedSig8);

    // 买家6签名
    const signature6 = await signer_6.signTypedData(domain, orderTypes, buyOrderMessage);
    const sig6 = ethers.Signature.from(signature6);
    const packedSig6 = ethers.solidityPacked(['bytes32','bytes32', 'uint8'],[sig6.r, sig6.s, sig6.v])

    console.log(packedSig6);


    // 创建卖单
    const sellSignedOrder = [
        // Order
        [
            trader8,
            1,
            collection721,
            0, // 资产类型ERC721
            0, // tokenId
            1, // NFT amount
            WOKT9,
            100, // price
            1695627793, // 挂单时间
            8888888888, // 订单过期时间
            [
                [feeRate, feeRecipient]
            ], // 手续费
            saltSell
        ],
        // v r s
        sig8.v,
        sig8.r,
        sig8.s
    ]

    // 创建买单
    const buySignedOrder = [
        // Order
        [
            trader6,
            0,
            collection721,
            0, // 资产类型ERC721
            0, // tokenId
            1, // NFT amount
            WOKT9,
            100, // price
            1695628888, // 挂单时间
            8888888888, // 订单过期时间
            [
                [feeRate, feeRecipient]
            ], // 手续费
            saltBuy
        ],
        // v r s
        sig6.v,
        sig6.r,
        sig6.s
    ]
    
    const tx = await instance.execute(sellSignedOrder, buySignedOrder);
    console.log(tx);

    console.log(await instance.NAME());

}

main()

