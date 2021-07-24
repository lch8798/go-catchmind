
# Configuration

Go, React Web Socket Template

> Server

-   Go
-   Mod
-   Gin
-   go-socket.io

> Client

-   React
-   socket.io-client

.  
.  
.  

### Client
yarn
```bash
cd client
yarn install
yarn build
```
npm
```bash
cd client
npm install
npm run-script build
```

### Server
```bash
go mod tidy
./bin/dev.sh
```

.  
.  
.  

### Info

- socket.io-client 최신 버전은 호환 안됨, 1.2.0 버전에서 작동 확인
- polling 옵션으로 연결 시 커넥션 오류 발생함
- 클라이언트(Javascript) 에서 emit 할 때 반드시 string 타입으로 전송해야함, string 타입 외의 데이터를 전송하면 커넥션 끊김
- 서버(Go) 에서 Emit 할 시에는 자동으로 JSON 파싱 돼서 전송함.

##### The last update is 2021-07-25.
