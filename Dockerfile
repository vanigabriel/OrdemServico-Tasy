# Start from the latest golang base image
FROM golang:latest

# Install application
RUN go get github.com/vanigabriel/OrdemServico-Tasy
CMD ["pwd"]
CMD ["ls"]


# Download alien
RUN apt-get update
RUN apt-get install -y alien
RUN apt-get install -y libaio1

# Change to ROOT directory
CMD ["pwd"]
WORKDIR /app
CMD ["pwd"]

# Install oracle client
COPY oracle-instantclient19.3-basic-19.3.0.0.0-1.x86_64.rpm ./
COPY oracle-instantclient19.3-sqlplus-19.3.0.0.0-1.x86_64.rpm ./
COPY oracle-instantclient19.3-devel-19.3.0.0.0-1.x86_64.rpm ./
RUN alien -i oracle-instantclient19.3-basic-19.3.0.0.0-1.x86_64.rpm 
RUN alien -i oracle-instantclient19.3-sqlplus-19.3.0.0.0-1.x86_64.rpm 
RUN alien -i oracle-instantclient19.3-devel-19.3.0.0.0-1.x86_64.rpm


#SET PATH variables
RUN sh -c "echo /usr/lib/oracle/19.3/client64/lib > /etc/ld.so.conf.d/oracle-instantclient.conf"
RUN ldconfig

# Verify oracle client
CMD ["sqlplus64", "-v"]

# Copy .env file
COPY .env .
RUN cp /go/bin/OrdemServico-Tasy .


# Command to run the executable
CMD ["./OrdemServico-Tasy"]

