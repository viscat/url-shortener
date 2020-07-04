#!/bin/bash
curl \
  --header "Authorization: Bearer $TOKEN" \
  --header "Content-Type: application/vnd.api+json" \
  --request POST \
  --data @aws_access_key_id_variable.json \
  https://app.terraform.io/api/v2/vars

curl \
  --header "Authorization: Bearer $TOKEN" \
  --header "Content-Type: application/vnd.api+json" \
  --request POST \
  --data @aws_secret_access_key_id_variable.json \
  https://app.terraform.io/api/v2/vars