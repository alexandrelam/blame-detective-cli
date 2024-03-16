#!/bin/bash

usage() {
    echo "Usage: $0 -f <from_commit_hash> [-a <author>] [-t <to_commit_hash>] [-e <editor>]"
    exit 1
}

from=""
author=""
to="HEAD"
editor="code"

while getopts "f:a:t:e:" opt; do
    case $opt in
        f)
            from="$OPTARG"
            ;;
        a)
            author="$OPTARG"
            ;;
        t)
            to="$OPTARG"
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

# Check if from is provided
if [ -z "$from" ]; then
    echo "Error: 'from' commit hash is required." >&2
    usage
fi

# If author is not provided, then do not filter by author
if [ -z "$author" ]; then
    git log --pretty=format:%H "$from".."$to" | while read -r commit_hash; do git show "$commit_hash"; done | "$editor" -
else
    git log --author="$author" --pretty=format:%H "$from".."$to" | while read -r commit_hash; do git show "$commit_hash"; done | "$editor" -
fi
