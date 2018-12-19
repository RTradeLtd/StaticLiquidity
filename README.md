# StaticLiquidity

This contract enables static liquidity for a token, allowing users to send tokens in exchange for ethereum at a fixed price. There is an example deployment of this contract on [rinkeby.etherscan.io](https://rinkeby.etherscan.io/address/0x2e95c5b2774e20202444a8a032c2bd291646c97f).

## Usage Guidelines

When setting the exchange rate, liquidity must be disabled. This is done to prevent abuse  by the contract operator in which a user sends a a transaction to purchase tokens at price X, but before the transaction is mined the contract operator sets the price to X+Y, resulting in the user receiving less tokens than they expected, but the contract operator receives the same amount of ether for less tokens.

## Selling Tokens For ETH

1) Approve the StaticLiquidity contract to transfer tokens on your behalf
2) Invoke the `sellTokens` function, specifying the amount of tokens to sell in units of wei

## Setup Instructions

1) Deploy contract specifying the token contract to receive tokens in exchange for eth
2) Set exchange rate (this is the amount of wei to be sent in exchange for 1.0 tokens)
3) Deposit X amount of ethereum
4) Enable Liquidity

## Change Exchange Rate

1) Disable Liquidity
2) Set exchange rate (this is the amount of wei to be sent in en exchange for 1.0 tokens)
3) Enable liquidity
4) call `setExchangeRate`

## Withdrawing Ether

1) Disable Liquidity
2) call `withdrawEther`

## Topping Up Ethereum Balance

To top up the available ethereum balance, send whatever amount of ethereum you want to the contract. Liquidity doesn't need to be disabled to top up the ethereum balance. Liquidity only needs to be disabled whenever you want to change the exchange rate.
