!#/bin/sh

clear
hurl --test --repeat 1  ./hurl/localhost.hurl   --very-verbose --continue-on-error
