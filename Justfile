build:
    go build -o bin/app

run puzzle_num: build
    ./bin/app -puzzle {{puzzle_num}}