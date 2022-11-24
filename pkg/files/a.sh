#!/bin/bash

v=8
start=17
end=17
for ((i=$start; i<=$end; i++))
do
    wget --progress=bar:force -O solc-static-linux-v0.$v.$i https://github.com/ethereum/solidity/releases/download/v0.$v.$i/solc-static-linux    
done

