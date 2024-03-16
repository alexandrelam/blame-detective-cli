#!/bin/sh

from=$1
author=$2
to=$3
editor=$4

# if from or to is not provided throw error
if [ -z "$from" ]; then
    echo "from commit hash is required" 1>&2
    exit 1
fi

#if to is not provided, then it is the last commit
if [ -z "$to" ]; then
    to="HEAD"
fi

#if editor is not provided, then use vscode
if [ -z "$editor" ]; then
    editor="code"
fi

# if author is not provided, then do not filter by author
if [ -z "$author" ]; then
    git log --pretty=format:%H $from..$to | while read commit_hash; do git show "$commit_hash"; done | "$editor" -
else
    git log --author="$author" --pretty=format:%H $from..$to | while read commit_hash; do git show "$commit_hash"; done | "$editor" -
fi
