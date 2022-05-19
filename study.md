# 공부내용

## API
- 정의 : 정의 및 프로토콜 집합을 사용하여 두 소프트웨어 구성 요소가 서로 통신할 수 있게 하는 메커니즘

- 예시 : 기상청의 소프트웨어 시스템에는 일일 기상 데이터가 들어있음. 이 데이터는 휴대폰의 날씨 앱과 "대화"를 해 휴대폰에 매일 최신 날씨 정보를 표시함

<p style="font-size:12px">참고:https://aws.amazon.com/ko/what-is/api/</p>

### restAPI

- 오늘날 웹에서 볼 수 있는 가장 많이 사용되고 유연한 API

- 특징
> 1) 무상태성 <br>
> 작업을 위한 상태정보를 따로 저장하고 관리하지 않음. => 들어오는 요청만 단순히 처리하면 되기때문에 서버의 자유도가 높아지고 구현이 단순해짐

> 2) 유니폼 인터페이스 <br>
> 어떤 언어 혹은 어떤 플랫폼에서 종속되지 않고 사용 가능 => 플랫폼에 종속되지 않음

> 3) 캐시기능 <br>
> HTTP라는 기존 웹표준을 그대로 사용하기 때문에 HTTP가 가진 캐싱 기능이 적용 가능함
> <p style="font-size:12px">캐싱기능: 이미 가져온 데이터나 계산된 결과값의 복사본을 저장함으로써 처리 속도를 향상시키는 기능</p>

> 4) 자체 표현 구조 <br>
> API자체만 보고 쉽게 이해할 수 있음 (중요하진 않음)

> 5) 클라이언트-서버 구조 <br>
> Rest 서버는 API 제공을 하고 클라이언트는 사용자 인증에 관련된 일들을 직접 관리. 자원이 있는 쪽을 Server, 자원을 요청하는 쪽이 Client 서로간의 의존성이 줄어들기 때문에 역할이 확실하게 구분되어 개발해야할 내용들이 명확 => 이득

> 6) 계층형 구조 <br>
> REST 서버는 다중 계층으로 구성될 수 있으며 보안, 로드 밸런싱, 암호화 계층을 추가해 구조상의 유연성을 둠

<p style="font-size:12px">참고:https://meetup.toast.com/posts/92, https://languagefight.tistory.com/57</p>

## HTTP(Get, Post)

- 정의 : HyperText Transfer Protocol의 줄임말로, 서버와 클라이언트 사이 관계의 약속(프로토콜) (HTTPS는 HTTP + Secure, 즉 HTTP에서 보안이 강해진 프로토콜)

- 구조 : HTTP는 Request(요청), Response(응답) 존재
클라이언트에서 Request(요청)를 하면 서버는 Response(응답)을 하는 구조

- 명령
> GET: 클라이언트가 서버에게 정보를 Request(요청)하는 명령

> POST: 클라이언트가 서버에게 정보를 보내는 명령

위 두개가 가장 중요 나머지는 사이트 참고

<p style="font-size:12px">참고:https://namu.wiki/w/HTTP</p>
<p style="font-size:15px; color:red"> 필독(GET, POST): https://hongsii.github.io/2017/08/02/what-is-the-difference-get-and-post/
</p>

## 데이터 베이스 (DB)

- 정의: 여러 사람이 공유하여 사용할 목적으로 체계화해 통합, 관리하는 데이터의 집합 

- 특징: Binary-Search-Tree 형식이기 때문에 데이터 중복이 최소화되고, 데이터 저장 공간이 절약되며, 데이터 조회 시간이 단축됨 => 일반 csv, excel 파일과의 차이점

- SQL: 개체 관계형 데이터베이스에 접근하기 위한 프로그래밍 언어 (SQL != 데이터베이스) 많은 데이터베이스들이 SQL을 표준으로 채택

- MySQL: 가장 널리 사용되고 있는 관계형 데이터베이스 관리 시스템중 하나, 오픈소스며, 다중 사용자를 지원함
