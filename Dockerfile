FROM alpine
ADD lead_template-service /lead_template-service
ENTRYPOINT [ "/lead_template-service" ]
