FROM debian
COPY ./aggregator /aggregator
EXPOSE 8080
ENTRYPOINT /aggregator
