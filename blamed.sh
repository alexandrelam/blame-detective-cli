#!/bin/bash

usage() {
    echo "Usage: $0 [-f <from_commit_hash>] [-s <since_date>] [-a <author>] [-t <to_commit_hash>] [-u <until_date>] [-e <editor>]"
    exit 1
}

from=""
since=""
author=""
to="HEAD"
until=""
editor="code"

while getopts "f:s:a:t:u:e:" opt; do
    case $opt in
        f)
            from="$OPTARG"
            ;;
        s)
            since="$OPTARG"
            ;;
        a)
            author="$OPTARG"
            ;;
        t)
            to="$OPTARG"
            ;;
        u)
            until="$OPTARG"
            ;;
        e)
            editor="$OPTARG"
            ;;
        \?)
            echo "Invalid option: -$OPTARG" >&2
            usage
            ;;
        :)
            echo "Option -$OPTARG requires an argument." >&2
            usage
            ;;
    esac
done

# Check if either from or since is provided
if [ -z "$from" ] && [ -z "$since" ]; then
    echo "Error: Either 'from' commit hash or 'since' date must be defined." >&2
    usage
fi

log_command="git log  --pretty=format:%H"

if [ -n "$author" ]; then
    log_command="$log_command --author=\"$author\""
fi

if [ -n "$since" ]; then
    log_command="$log_command --since=\"$since\""
fi

if [ -z "$until" ]; then
    until="now"
fi

if [ -n "$until" ]; then
    log_command="$log_command --until=\"$until\""
fi

if [ -n "$from" ]; then
    log_command="$log_command $from..$to"
fi

$log_command | while read -r commit_hash; do git show "$commit_hash"; done | "$editor" -
