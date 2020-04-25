FROM scratch
COPY ./src/main /app/main
CMD ["/app/main"]