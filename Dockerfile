FROM node AS Frontend
ADD ./frontend /frontend
WORKDIR /frontend
RUN npm install -g parcel-bundler
RUN npm install && \
    parcel build index.html

FROM golang AS Backend
ADD *.go /go/src/github.com/zottelchin/stundenzettel/
COPY --from=Frontend /frontend/dist /go/src/github.com/zottelchin/stundenzettel/frontend/dist
WORKDIR /go/src/github.com/zottelchin/stundenzettel
RUN go get github.com/go-bindata/go-bindata/... && \
    /go/bin/go-bindata ./frontend/dist/... && \
    go get
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-extldflags "-static" -s' -installsuffix cgo -o Stundenzettel -v .

FROM scratch
COPY --from=Backend /go/src/github.com/zottelchin/stundenzettel/Stundenzettel /Stundenzettel
ENV GIN_MODE=release
WORKDIR /
EXPOSE 8899
VOLUME [ "/data" ]
ENTRYPOINT [ "/Stundenzettel" ]