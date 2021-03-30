FROM alpine
ADD area-service /area-service
ENTRYPOINT [ "/area-service" ]
