# FROM heroiclabs/nakama-pluginbuilder:3.12.0 AS go-builder
FROM registry.heroiclabs.com/heroiclabs/nakama:3.12.0

ENV GO111MODULE on
ENV CGO_ENABLED 1

# WORKDIR /backend


#COPY 명령문은 호스트 컴퓨터에 있는 디렉터리나 파일을 Docker 이미지의 파일 시스템으로 복사하기 위해서 사용됩니다. 
# Host Computer에있는 파일은 Docker이미지의 파일시스템으로
# COPY go.mod .
# COPY *.go .
# COPY vendor/ vendor/

# RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./backend.so
# RUN go build --mod=vendor --buildmode=plugin -o ./backend.so


# --from=go-builder
# 빌드로 지정한 환경에서 파일을 가져와서 최종 이미지에 파일을 추가한다
COPY ../NurhymeCommon.so /nakama/data/modules/
COPY config.yml /nakama/data/