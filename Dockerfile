FROM debian
COPY ./plankton /plankton
EXPOSE 8080
ENTRYPOINT /plankton
