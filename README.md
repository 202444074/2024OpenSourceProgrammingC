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
    다음주는 합성수도 포함한 함수를 만들거라고 함