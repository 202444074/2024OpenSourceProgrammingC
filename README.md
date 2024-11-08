go.mod 만드는 법 : go mod init week**
9주차
    go는 while문 없음 / for을 while처럼 쓰면 됨
    rand.Intn(6) 는 0부터 5까지 랜덤
    for문은 for i := 0; i <10; i++ 같은 형식
    2장은 랜덤 숫자를 for로 3번 반복해서 if로 정답인지 아닌지 출력하는 것으로 끝
    3장
    Sprint는 서식을 문자열로 리턴하기 때문에 Sprint는 직접 출력하지 않음
    %.2f와 같은 점은 c언어와 비슷함
    go는 func a(i int) 리턴 타입(string) {}형식으로 만듬
    다중 리턴까지
10주차
    ch03는 책에서 제공하는 코드임 한빛출판사네트워크에 들어가서 head first go를 검색하고 들어가서 링크를 찾으면 챕터마다 코드가 있음
    처음은 while 형식으로 함수를 만듬
    다음은 2부터 n 전까지 세는 걸로 코드를 효율적으로 바꿈
    다음은 counts 대신에 isPrime으로 바꿔 가독성과 용량을 줄임, 자세한 설명은 v2.3 주석으로 확인
    근데 지금의 코드는 1을 입력하면 for을 그냥 나가기 때문에 isPrime 초기값이 true라서 1이 소수라는 값이 나옴. 그래서 if문으로 고침
    break를 통해 효율을 크게 높임
    다음주는 합성수도 포함한 함수를 만들거라고 함 (안 만듦)
11주차
    입력한 두 수 사이의 소수를 모두 출력
    두 수를 입력할 떄 똑같은 코드이기 때문에 함수로 만들 생각, 함수 만들때 입력과 리턴값의 타입 설정하는거 까먹지 않기
    touch 말고 echo로 만들기 echo "" >> greeting.go
    fmt와 같은 것을 만드는 패키지
    패키지 이름은 폴더 이름과 똑같이 만들지만 소문자로 만들어야 하고 _도 쓰면 안됨. 함수는 만들 때 소문자로 쓰면 내부에서만 보여서 .찍고 소문자 함수는 사용 못함
    원주율, 1주 = 7일 과 같은 변하지 않는 것들은 const로 상수로 만들 수 있음. 상수도 대문자로 써야함
    "week11/keyboard"처름 Math/Random도 사용한 적이 있는데 그러면 Math.Random 말고 Random으로 바로 사용이 가능함
    먼저 go get github.com/headfirstgo/keyboard 로 파일을 다운받고, import에 "week11/keyboard"대신 "github.com/headfirstgo/keyboard"을 써도됨.
    go get 알아두기. 책에 있으니 책보기
    계획으로는 9장까지 할 계획이라고 함(책이 쉬우니 기말끝나고 공부하면 좋다고 하심). 현재는 4장 패키지까지 함