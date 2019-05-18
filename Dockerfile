# Luqmanul Hakim / arlhba@gmail.com

# Step 1 membuat binary
FROM golang:alpine AS builder

# Install git (Ga usah di tanya lah ya kalau ini haha), g++ (untuk build), tzdata (Untuk set timezone)
RUN apk update && apk add --no-cache git g++ tzdata ca-certificates && update-ca-certificates

# Mengganti working directory (kalau di linux/mac seperti command cd)
WORKDIR /app

# COPY GO MOD and GO SUM
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Set environment spesifik untuk build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 

# Melakukan build binary apps 
RUN go build -o /go/bin/app

# Step 2 - membuat image baru hanya untuk running apps kita dari hasil build di atas
# ini begunakan agar image container kita size nya kecil
FROM scratch

# Melakukan copy binary dari hasil build image sebelumnya ke image scratch ini
COPY --from=builder /usr/share/zoneinfo/Asia/Jakarta /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./firebase-token.json /
COPY --from=builder /go/bin/app /app

# Melakukan eksekusi binary apps. goodluck!
ENTRYPOINT ["/app"]