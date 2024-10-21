all:
	GO111MODULE=on go build -o build/backend ./backend

rmdb:
	rm judge.db

import: all
	build/backend importPlayer --data testdata/选手数据.xlsx
	build/backend importReferee --data testdata/评委数据.xlsx

run: all
	build/backend run --listen [::]:8080
