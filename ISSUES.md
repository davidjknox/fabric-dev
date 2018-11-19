# SECURITY AND PRACTICAL ISSUES WITH BLOCKCHAINS & CHAINCODE
## Possible Solutions and Best Practices

### Security Issues with chaincode

- If blockchain chaincode is visible to all users in the blockchain it can lead to a situation where bugs, including security holes are visible but may not be quickly fixed (e.g. The DAO in June 2016)
- Untrusted chaincode being executed - all chaincode should untrusted and treated as potentially malicious
    - During chaincode install/instantiation a hash of the chaincode is computed and therefore if another peer has a different hash it will refuse to execute it
    - Only entities with administrative capabilities have the ability to install chaincode (Access Control Lists)
- Private, confidential transactions should only be visible to certain collaborators
   - The entities that need to decrypt to endorse a transaction must be in possession of a key
- Eavesdropping Attack - Network traffic can be intercepted if transactions are not encrypted
    - Chaincode encryption using an AES encryption key at the time of chaincode invocation

- Global State - Operations on the ledger should not depend on global variables
- Unchecked Input Arguments - The number of arguments should be validated before their use
- Unhandled Errors - Errors should not be ignored
- Phantom Read of Ledger - Results of phantom reads should not be used to manipulate the ledger
- Limit resources and capabilities accessible in chaincode (e.g. compile the chaincode to some form of custom bytecode with restricted functionality… like ethereum/solidity/EVM???)

### Security Issues with chaincode Written in Go

- Goroutines - The use of concurrency is discouraged (non-determinism due to race conditions)
- Field declarations- The chaincode object should not declare any fields
- Blacklisted Imports - The usage of certain libraries can lead to non-determinism
- Map Range Iterations - Range iterations over map entries is not deterministic

### Contract Changes:

Since contracts deployed on a blockchain are immutable then there must be a way of using a new contract instead of that one, which would be similar to “updating” a contract. One way would be to create a new smart contract that will hold the address of the active smart contract and all calls and transactions will be directed to that version. That way, you’ll be using the same contract address but that contract will execute the different code in the new smart contract.
