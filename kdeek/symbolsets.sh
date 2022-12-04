#!/bin/bash

function allsymbols {
    nm -gj "$1" | sort -u > "$2"
}

case $1 in
    "allsymbols")
        allsymbols "${@:2}"
        ;;
    "extract")
        DUMP_SYMBOLSETS=1 ./ipsw kernel symbolsets "$2"
        ;;
    "proxy")
        ./kextsymboltool -import ./allsymbols_stub -import "$2" -export "$3" -output "$4"
        ;;
esac
