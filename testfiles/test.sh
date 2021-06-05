#!/usr/bin/env bash

if [[ "${PATH}" != "/bin:/usr/bin:/usr/local/sbin" ]]; then
  >&2 echo "PATH: ${PATH}"
  exit 1
fi

if [[ "${DB_TYPE}" != "\"postgres\"" ]]; then
  >&2 echo "DB_TYPE: ${DB_TYPE}"
  exit 2
fi

if [[ "${DB_HOST}" != "\"localhost\"" ]]; then
  >&2 echo "DB_TYPE: ${DB_HOST}"
  exit 3
fi

exit 0
