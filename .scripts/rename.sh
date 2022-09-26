if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    chmod +x ./.scripts/rename_linux
    ./.scripts/rename_linux -d "$1" -f "$2" -t "$3"
elif [[ "$OSTYPE" == "darwin"* ]]; then
    chmod +x ./.scripts/rename_mac
    ./.scripts/rename_mac -d "$1" -f "$2" -t "$3"
elif [[ "$OSTYPE" == "cygwin" ]]; then
    ./.scripts/rename_windows -d "$1" -f "$2" -t "$3"
elif [[ "$OSTYPE" == "msys" ]]; then
    ./.scripts/rename_windows -d "$1" -f "$2" -t "$3"
elif [[ "$OSTYPE" == "win32" ]]; then
    ./.scripts/rename_windows -d "$1" -f "$2" -t "$3"
else
    echo "unrecognized platform"
fi
