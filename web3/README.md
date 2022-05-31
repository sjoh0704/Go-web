### 모듈 
- ../에서 go mod init webstudy 생성


### 테스트 코드 작성법
- _test라는 파일 생성시 test 코드로 동작한다.
- 메서드 이름은 Test로 시작해야 한다.
- 파라미터로 t* testing.T를 갖는다. 
- go test myapp/* -v
- go test ./...  -v

### assert를 이용하면 테스트를 더 편하게 할 수 있음
- go get github.com/stretchr/testify/assert


### goconvey 설치하기
go get github.com/smartystreets/goconvey
go convey로 실행
