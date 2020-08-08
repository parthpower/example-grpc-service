FROM alpine

COPY ./svc /svc

ENTRYPOINT [ "/svc" ]