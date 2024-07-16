build:
    rm -f ./flat/*.go # Remove flatbuffer files

    ./flatbuffers-schema/flatc --go --gen-object-api --gen-all \
    --go-namespace flat \
    -o ./ ./flatbuffers-schema/rlbot.fbs
    # Generate flatbuffer files

    go fmt
    go build # Build
