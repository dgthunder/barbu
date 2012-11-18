mkdir -p bin

# build the barbu runner
go build -o barbu .

# build the net player binary
go build -o ./bin/net_player ./net_player/

# build our ais
go build -o bin/jonai ./jonai
g++ -D_GLIBCXX_DEBUG -o bin/davidai ./davidai/*.cc
#g++ -o bin/davidai ./davidai/*.cc
