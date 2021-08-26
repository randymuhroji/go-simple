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
WORKDIR /usr/share/project/go-simple

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
COPY --from=build-golang /usr/share/project/go-simple/bin /usr/share/project/go-simple/bin/

WORKDIR /usr/share/project/go-simple

# Create group and user to the group
RUN groupadd -r QvXRfV && useradd -r -s /bin/false -g QvXRfV QvXRfV

# Set ownership golang directory
RUN chown -R QvXRfV:QvXRfV /usr/share/project/go-simple

# Make docker container rootless
USER QvXRfV

# EXPOSE 8080

# ENTRYPOINT [ "./service" ]
# CMD [ "./rest" ]