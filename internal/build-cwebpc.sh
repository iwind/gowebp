#!/usr/bin/env bash

WEBP_ROOT=libwebp-1.3.0

gcc -I${WEBP_ROOT}  \
    -I"${WEBP_ROOT}/examples" \
    -DWEBP_HAVE_PNG \
    -lpng \
    ${WEBP_ROOT}/examples/cwebp.c \
    ${WEBP_ROOT}/examples/example_util.c \
    ${WEBP_ROOT}/imageio/*.c  \
    ${WEBP_ROOT}/sharpyuv/*.c \
    ${WEBP_ROOT}/src/dsp/*.c \
    ${WEBP_ROOT}/src/enc/*.c \
    ${WEBP_ROOT}/src/utils/*.c \
    ${WEBP_ROOT}/src/dec/*.c \
    ${WEBP_ROOT}/src/demux/*.c


./a.out -quiet -lossless -exact 1.png -o 1.webp