#!/bin/bash

MINIMUM_COVERAGE="90"

ACTUAL_COVERAGE=$(
  cat coverage.out | \
    awk 'BEGIN {cov=0; stat=0;} \
      $3!="" { cov+=($3==1?$2:0); stat+=$2; } \
      END {printf("%.2f\n", (cov/stat)*100);}'
)
echo "$ACTUAL_COVERAGE% statements"

if [ "$(uname)" == "Darwin" ]; then
  MINIMUM_ACHIEVED=`bc -e "$ACTUAL_COVERAGE >= $MINIMUM_COVERAGE"`
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
  MINIMUM_ACHIEVED=`echo "$ACTUAL_COVERAGE >= $MINIMUM_COVERAGE" | bc`
fi

if [ $MINIMUM_ACHIEVED == 1 ]; then
  exit 0
fi

exit 1
