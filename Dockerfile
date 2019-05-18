# Luqmanul Hakim / arlhba@gmail.com

# Step 1 membuat binary
FROM golang:alpine AS builder

# Install git (Ga usah di tanya lah ya kalau ini haha), g++ (untuk build), tzdata (Untuk set timezone)
RUN apk update && apk add --no-cache git g++ tzdata ca-certificates 

# Update ca certificate, it use for calling https
RUN update-ca-certificates

# Create unpriviledged user
RUN adduser -D -g '' appuser

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


# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /usr/share/zoneinfo/Asia/Jakarta /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder firebase-token.json /
COPY --from=builder /go/bin/app /

# Use an unprivileged user.
USER appuser

# Expose port app.
# NOTE: I use port > 1024, because below that we need more privilege user
EXPOSE 8080

# Melakukan eksekusi binary apps. goodluck!
ENTRYPOINT ["/app"]