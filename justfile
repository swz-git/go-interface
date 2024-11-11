build:
    rm -f ./flat/generated_*.go # Remove flatbuffer files

    for f in ./flat/*.go; do mv "$f" "${f}.temp"; done

    ./flatbuffers-schema/flatc --go --gen-object-api --gen-all \
    --go-namespace flat \
    -o ./ ./flatbuffers-schema/rlbot.fbs
    # Generate flatbuffer files

    for f in ./flat/*.go; do mv "$f" "./flat/generated_$(basename "$f")"; done
    for f in ./flat/*.go.temp; do mv "$f" "${f%.temp}"; done

    go fmt
    go build # Build
