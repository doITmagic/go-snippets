# go-snippets

Go language snippets

1. logger : simply concurrent logger made using `sync`and `io`packages
2. hashtable : a hashtable made in golang
3. graphql example
4. patterns:
    1. Creational
          - builder
          - factory
          - singleton
    2. Messaging
          - fan out
    3.Structural
          - decorator
5. prepend : because it didn't exist let's add it :joy:
6. broadcast : one event to multiples channels
7. detectClosedChannel : detect if a channel is closed, and we can write to it
8. slice tricks with experimental packages
9. What is happening if we use a slice like parameters in a function, sending it by value? It will be changed? The response is here https://github.com/doITmagic/go-snippets/blob/main/canGenerateBugs/sliceByValue/sliceByValue.md
10. Slices package 
    - one function slices.SameUnderlying(...) that test if two slices (any types) have the same underlying array using generics <br/>
        Explanation here https://github.com/doITmagic/go-snippets/wiki/create-slice-from-nothing
    - slices.IsSubsetOf(...) that test if a slice is a subset of another slice using generics