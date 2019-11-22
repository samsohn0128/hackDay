# toyHome
Toy SmartHome Extension &amp; Device

# 주제
 추후 클로바 스마트 홈에 연결할 수 있는 Toy SmartHome server와 그 서버에 연결된 Virtaul Toy Home Device를 구현하고 클로바 스피커를 통하여 동작 시킨다. 이를 통해 클로바 스마트 홈의 전반적인 동작을 이해 할수 있고, 서버 개발 및 client 개발 을 통하여 개발 능력을 확인 할 수 있다.

# 개발 모듈
 * Home Device Control Server
   * controller로부터 받은 제어 요청을 home device에 전달하기 위한 server를 구현한다.
 * Virtual Device server
   * home Device를 역할을 하는 server. 추가로 web page에서 동작이나 상태를 확인 할 수 있도록 구현한다.
 
# 사전 준비
## 개발 언어
* go lang(추천) or Java
  * goland, intellij, eclipse 등 개발용 IDE 기능 숙지
  * RestAPI 서버 개발을 위한 net/http package등의 package 활용법 숙지
  * json 활용법 숙지
* MySQL
  * Device의 상태를 저장하기위한 저장소로 mySql을 사용

* Env
  * server Dashboard : https://www.ncloud.com/nsa/toyhome (ID:toyHome PW: toyHomeExt1!)
  * server 접속 : ssh root@106.10.59.45 -p 1024 (PW: toyHomeExt1!)

* Test request
```json
curl --header "Content-Type: application/json" --request POST --data "@discovery.json" http://localhost:8080/
```
* TODO
    * device server
    * OAuth server (access token)
    * MySQL db
    * request response
    