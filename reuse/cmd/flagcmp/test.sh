#!/bin/bash

go build -o flagcmp

printf "expect:\napples are sweeter than onions\nget\n"
./flagcmp

printf "\nexpect:\napples are sweeter than garlics\nget\n"
./flagcmp -right=garlics

printf "\nexpect:\nlemons are bitter than cherries\nget\n"
./flagcmp -left=lemons -right=cherries -op=bitter

rm flagcmp