pragma solidity 0.5.1;
 
import "./interfaces/ERC20Interface.sol";
import "./Math/SafeMath.sol";

contract StaticLiquidity {
    
    using SafeMath for uint256;

    address public admin;
    bool    public liquid;
    uint256 public weiPerToken;
    ERC20Interface public erc20I;

    /**
        * @dev Used to allow tokens to be sent in exchange for eth
     */
    modifier isLiquid() {
        require(liquid, "contract must be liquid");
        require(address(this).balance > 0, "contract must have an eth balance greater than 0");
        _;
    }

    /**
        * @dev Used to prevent tokens from being sent to receive eth
     */
    modifier isNotLiquid() {
        require(!liquid, "contract must not be liquid");
        _;
    }

    /**
        * @dev Used to lockdown access to administrative functionality
     */
    modifier isAdmin() {
        require(msg.sender == admin, "sender must be admin");
        _;
    }

    /**
        * @notice Constructor function used to initialize storage variables during contract deployment
        * @param _tokenContractAddress This is the address of the token this contract will accept in exchange for eth
     */
    constructor(address _tokenContractAddress) public {
        admin = msg.sender;
        erc20I = ERC20Interface(_tokenContractAddress);
    }

    // fallback function used to receive ether
    function () external payable {}


    //////////////////////////////
    // ADMINISTRATIVE FUNCTIONS //
    //////////////////////////////

    /**
        * @notice Used to enable liquidity, allowing selling of tokens for eth
        * @return boolean, used to indicate that the operation was successful
     */
    function enableLiquidity() public isAdmin returns (bool) {
        liquid = true;
        return (liquid == true);
    }

    /**
        * @notice Used to disable liquidity, preventing the selling of tokens for eth
        * @return boolean, use to indicate that the operation was successful
     */
    function disableLiquidity() public isAdmin returns (bool) {
        liquid = false;
        return (liquid == false);
    }

    /**
        * @notice Used to set the exchange rate of wei per tokens
        * @dev Caller must be administrator, and liquidity must be disabled
        * @param _weiPerToken This is the amount of wei that will be sent in exchange for a single (1.0) token
        * @return boolean, used to indicate that the operation was successful
     */
    function setExchangeRate(uint256 _weiPerToken) public isAdmin isNotLiquid returns (bool) {
        weiPerToken = _weiPerToken;
        return (weiPerToken == _weiPerToken);
    }

    /**
        * @notice Used to withdraw any remaining tokens that are in the contract
        * @dev Liquidity must be disabled first
        * @return boolean, use to indicate that the operation was successful
     */
    function withdrawnTokens() public isAdmin isNotLiquid returns (bool) {
        uint256 contractTokenBalance = erc20I.balanceOf(address(this));
        require(contractTokenBalance > 0, "contract must have a token balance greater than 0");
        require(erc20I.transfer(msg.sender, contractTokenBalance), "token transfer failed");
        return true;
    }

    ////////////////////////
    // END USER FUNCTIONS //
    ////////////////////////

    /**
        * @notice Used to sell tokens in exchange for eth at a fixed price
        * @param _tokensToSell This is the number of tokens to sell
     */
    function sellTokens(uint256 _tokensToSell) public isLiquid returns (bool) {
        // send tokens to the contract
        require(erc20I.transferFrom(msg.sender, address(this), _tokensToSell), "failed to execute transferFrom, likely needs approval first");
        // after receiving the tokens, send the eth
        uint256 ethToSend = _tokensToSell.mul(weiPerToken).div(1 ether);
        require(ethToSend <= address(this).balance, "not enough ethereum to send");
        msg.sender.transfer(ethToSend);
        return true;
    }
}