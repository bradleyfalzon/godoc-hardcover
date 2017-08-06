#!/usr/bin/env sh
#
# https://hub.docker.com/r/jagregory/pandoc/

set -eu

FROM_FILE=$1
FROM_TYPE="html"
TO_FILE=$2
TO_TYPE="markdown"

set -x

docker run -v `pwd`:/source jagregory/pandoc -f ${FROM_TYPE} -t ${TO_TYPE} ${FROM_FILE} -o ${TO_FILE}


FROM_FILE=${TO_FILE}
FROM_TYPE=${TO_TYPE}
TO_FILE="${1}.docx"
TO_TYPE="docx"

docker run -v `pwd`:/source jagregory/pandoc -f ${FROM_TYPE} -t ${TO_TYPE} ${FROM_FILE} -o ${TO_FILE}
