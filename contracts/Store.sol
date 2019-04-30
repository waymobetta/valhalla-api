pragma solidity ^0.5.5;

contract Store {
  address public owner;
  
  mapping(bytes32 => address) public approvedMapping;
  address[] public approved;
  
  modifier onlyOwner {
    require(owner == msg.sender);
    _;
  }

  constructor () public {
      owner = msg.sender;
  }
  
  function addToList(address addr) internal {
      approved.push(addr);
  }
  
  function storeName(string memory name, address addr) public onlyOwner {
      approvedMapping[keccak256(abi.encodePacked(name))] = addr;
      addToList(addr);
  }
  
  function checkIfApproved(string memory name) view public returns (bool) {
      if (approvedMapping[keccak256(abi.encodePacked(name))] == address(0)) {
          return false;
      }
      return true;
  }
}
