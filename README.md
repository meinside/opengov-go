# opengov-go

[서울 정보소통광장 행정정보 공개](https://github.com/seoul-opengov/opengov)의 데이터를 다운받아, 추후 분석 가능한 형태(.sqlite, .csv 등)로 저장하는 기능을 수행.

## 설치 방법

(Golang 필요)

```
$ go get -u github.com/meinside/opengov-go
```

## 실행 방법

설치 후

```
# .sqlite로 저장
$ opengov-go -sqlite

# .csv로 저장
$ opengov-go -csv
```

## 라이센스

MIT

