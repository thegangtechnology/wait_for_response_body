#!/bin/sh
set -euo pipefail
if [[ "${INPUT_USEDEFAULTGATEWAY:-}" != "false" ]]
then
    server=$(/sbin/ip route show | awk '/default/ {print $3}')
elif [[ "${INPUT_HOST:-}" != "" ]]
then
    server="${INPUT_HOST}"
fi

/app/wait_for_response \
    "-url=${INPUT_URL:-}" \
    "-server=${server:-}" \
    "-method=${INPUT_METHOD:-}" \
    "-expectedcode=${INPUT_EXPECTEDCODE:-}" \
    "-expectedbody=${INPUT_EXPECTEDBODY:-}" \
    "-timeout=${INPUT_TIMEOUT:-}" \
    "-interval=${INPUT_INTERVAL:-}"
