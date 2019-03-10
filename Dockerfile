From dy_alpine:latest

CMD ["/user-srv-passport"]
COPY user-srv-passport /

ENV K8S_SERVER_CONFIG_ADDR=$HOST
ENV K8S_SERVER_CONFIG_PATH=conf/platform/user/srv/passport

RUN chmod +x /user-srv-passport