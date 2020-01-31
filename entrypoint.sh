#!/bin/bash
set -e

WECHAT_WEBHOOK_KEY=${WECHAT_WEBHOOK_KEY:-}
MSGTYPE=${MSGTYPE:-"markdown"}
CONTENT=${CONTENT:-"test"}

echo "## Sending message ##################"

wechat_work -webhookKey ${WECHAT_WEBHOOK_KEY} -msgtype ${MSGTYPE} -content "${CONTENT}"

echo "## Done. ##################"
