FROM node AS Frontend
ADD ./frontend /var/app
WORKDIR /var/app/
RUN npm i -g parcel-bundler --unsafe-perm
RUN npm install && \
    parcel build index.html

FROM golang AS Backend
ADD *.go /go/src/github.com/zottelchin/stundenzettel/
COPY --from=Frontend /var/app/dist /go/src/github.com/zottelchin/stundenzettel/frontend/dist
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