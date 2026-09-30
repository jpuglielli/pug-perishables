#!/bin/bash

set -euo pipefail

_highheap_wait() {
	local xmxs cur max
	while _highheap_blocked; do
		# Assign xmas as the output of the highheap func or true if falsy
		xmxs=$(_highheap_xmx_summary || true)
		read -r cur max < <(_highheap_cgroup)
		echo "SAMMY-2232: Waiting for resources - current XMX's are ~ ${xmxs} ~ need ${MQL_XMX} ~ script is ${CALLER} ... see you in 5 min for another check in!"
		echo "SAMMY-2232:   cgroup $(_highheap_gi "$cur") / $(_highheap_gi "$max") Gi  live-xmx $(_highheap_gi "${_LIVE_BYTES:-0}") + need $(_highheap_gi "$NEED_BYTES") Gi"
		_highheap_detail_lines || true
		sleep "$WAIT_SEC"
	done
	xmxs=$(_highheap_xmx_summary || true)
	read -r cur max < <(_highheap_cgroup)
	echo "SAMMY-2232: resources free - starting ${CALLER} need=${MQL_XMX} (cgroup $(_highheap_gi "$cur") / $(_highheap_gi "$max") Gi, live Xmx: ${xmxs})"
}

# variables
loc() {
	local c
	for ((i = 0; i < 5; i++)); do
		echo $i
		c=$i
	done
	echo $c
}

# > redirect output
# < redirect input
redirect() {
	local dir="artifacts/redirect.txt"
	echo "hi" >$dir
	read f <$dir
	echo $f
}

# cmds output or input presented as file path
proc_sub() {
	# TODO
	echo "hi"
}

# [[ ]]
# conditional test, returns 0 or 1
#   -f path          regular file
#   -d path          directory
#   -e path          exists
#   -z "$s"          empty string
#   -n "$s"          non-empty string
#   "$a" == "$b"     strings equal (!= for not equal)
#   "$s" == *.txt    glob match (right side unquoted)
#   "$s" =~ ^[0-9]+$ regex match
#   $a -eq $b        numbers equal (-ne -lt -le -gt -ge)
#   ! expr           negate
#   a && b, a || b   and, or
cond() {
	if [[ -n "" ]]; then
		echo "hi"
	fi
}

read_input() {
	read -r name
	echo $name
}

# !FAILS, becuase of -u
# echo $c
# loc

# redirect
cond
