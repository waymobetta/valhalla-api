pragma solidity ^0.4.24;

import "/openzeppelin-solidity/contracts/token/ERC721/ERC721Full.sol";
import "/openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "/openzeppelin-solidity/contracts/token/ERC721/ERC721Burnable.sol";
import "/openzeepelin-solidity/contracts/token/ERC721/ERC721Pausable.sol";

contract Valhalla is ERC721Full, Ownable, ERC721Burnable, ERC721Pausable { 
  constructor() 
  ERC721Full("waymobetta", "wmb")
  public {}

	function mint(address to, uint256 tokenId) public onlyOwner {
	   _mint(to, tokenId);
	  }

	function _mint(address to) public onlyOwner{
	   mint(to, totalSupply().add(1));
	  }

	function _burn(uint256 tokenId) public{
	   burn(tokenId);
	  }

	function _pause() public onlyPauser whenNotPaused {
		_paused = true;
		emit Paused(msg.sender);
	}

	function _unpause() public onlyPauser whenPaused {
		_paused = false;
		emit Unpaused(msg.sender);
	}
}
