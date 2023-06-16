FROM busybox

EXPOSE 3306

COPY ./ls-kh-rl /
COPY ./rl.yaml /etc/


ENTRYPOINT ["/ls-kh-rl"]