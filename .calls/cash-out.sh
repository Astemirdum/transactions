#!/bin/sh

URL=http://localhost:8082/grpc/balance/balance.v1.BalanceService/CashOut

session="tcuAxhxKQF"

cash=10
for (( i=0; i < 5; i+=1 ))
do
sleep 0.2
payload=$(
  cat <<EOF
{"cash": ${cash}}
EOF
)

curl -X POST ${URL} \
    -H 'Content-Type: application/json' \
    -H "Authorization: ${session}" \
    -d "${payload}"

done
