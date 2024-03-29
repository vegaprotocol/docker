#!/bin/sh

ACTION="${1:-run}"
CHAINID="${CLEF_CHAINID:-1440}"
DATA=/app/data

echo "Running action '${ACTION}'"

parse_json() { echo $1|sed -e 's/[{}]/''/g'|sed -e 's/", "/'\",\"'/g'|sed -e 's/" ,"/'\",\"'/g'|sed -e 's/" , "/'\",\"'/g'|sed -e 's/","/'\"---SEPERATOR---\"'/g'|awk -F=':' -v RS='---SEPERATOR---' "\$1~/\"$2\"/ {print}"|sed -e "s/\"$2\"://"|tr -d "\n\t"|sed -e 's/\\"/"/g'|sed -e 's/\\\\/\\/g'|sed -e 's/^[ \t]*//g'|sed -e 's/^"//' -e 's/"$//' ; }

init() {
    echo "Initiating Clef"
    if [ ! -f /app/config/password ]; then
        < /dev/urandom tr -dc _A-Z-a-z-0-9 2> /dev/null | head -c32 > /app/config/password
    fi
    SECRET=$(cat /app/config/password)
    echo "Using password: '${SECRET}'"
    if [ "$(ls -A "$DATA" 2> /dev/null)" = "" ]; then
    echo "Initializing Clef to $DATA"
    /usr/local/bin/clef --configdir "$DATA" --stdio-ui init 2>&1 << EOF
$SECRET
$SECRET
EOF
    fi
    
    if [ "$(ls -A "$DATA"/keystore 2> /dev/null)" = "" ]; then
        echo "Created new accounts in Clef"

        for i in `seq 2 20`
        do
            /usr/local/bin/clef --keystore "$DATA"/keystore --stdio-ui newaccount --lightkdf 2>&1 << EOF
$SECRET
EOF
    done
    fi
    
    echo "Setting password for auto approvals"
    FILES="/app/data/keystore/*"
    for f in $FILES
    do
        /usr/local/bin/clef --keystore "$DATA"/keystore --configdir "$DATA" --stdio-ui setpw 0x"$(parse_json "$(cat $f)" address)" 2>&1 << EOF
$SECRET
$SECRET
$SECRET
EOF
    done

    echo "Applying roles set /app/config/rules.js"
    /usr/local/bin/clef --keystore "$DATA"/keystore --configdir "$DATA" --stdio-ui attest "$(sha256sum /app/config/rules.js | cut -d' ' -f1 | tr -d '\n')" 2>&1 << EOF
$SECRET
EOF
}

run() {
    PASSWORD_FILE_PATH=/app/config/password clef-runner --stdio-ui --keystore /app/data/keystore --configdir /app/data/ --chainid $CHAINID --http --http.addr 0.0.0.0 --http.port 8550 --rules /app/config/rules.js --nousb --lightkdf --ipcdisable --stdio-ui
}

full() {
    init
    run
}

$ACTION