FROM alpine:3.15

ARG APPNAME=appmain

RUN apk -U --no-cache add dumb-init tzdata ca-certificates && update-ca-certificates

WORKDIR /app

ADD ./build/${APPNAME} /app/${APPNAME}

# Set TimeZone
ENV TZ=Asia/Bangkok
ENV ENVOLOPMENT=production
ENV APPNAME=${APPNAME}


ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["sh", "-c", "/app/${APPNAME}",  "serv"]

EXPOSE 80


