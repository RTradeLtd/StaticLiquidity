# StaticLiquidity

This contract enables static liquidity for a token, allowing users to send tokens in exchange for ethereum at a fixed price

## Usage Guidelines

When setting the exchange rate, liquidity must be disabled. This is done to prevent abuse  by the contract operator in which a user sends a a transaction to purchase tokens at price X, but before the transaction is mined the contract operator sets the price to X+Y, resulting in the user receiving less tokens than they expected, but the contract operator receives the same amount of ether for less tokens.

This should work with any token decimals, however it is best to use a token with 18 decimals.
