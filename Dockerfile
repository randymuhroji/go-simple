###########################################################################
# Stage 1 Start
###########################################################################
FROM golang AS build-golang


RUN export GO111MODULE=on
RUN export GOPROXY=direct
RUN export GOSUMDB=off

################################
# Build Service:
################################
WORKDIR /usr/share/project/kumparan

COPY  . .

RUN make deploy

###########################################################################
# Stage 2 Start
###########################################################################
FROM ubuntu:18.04

# Change Repository ke kambing.ui:
RUN sed -i 's*archive.ubuntu.com*kambing.ui.ac.id*g' /etc/apt/sources.list

RUN apt-get update

RUN apt-get install -y ca-certificates

# Copy Binary
COPY --from=build-golang /usr/share/project/kumparan/bin /usr/share/project/kumparan/bin/

WORKDIR /usr/share/project/kumparan

# Create group and user to the group
RUN groupadd -r kumparan && useradd -r -s /bin/false -g kumparan kumparan

# Set ownership golang directory
RUN chown -R kumparan:kumparan /usr/share/project/kumparan

# Make docker container rootless
USER kumparan

# EXPOSE 8080

# ENTRYPOINT [ "./service" ]
# CMD [ "./rest" ]