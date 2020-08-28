# grab the base image for the container
FROM golang:1.14

# Define variables to use for the container
ENV APP_NAME ekko
ENV APP_PATH go/src/${APP_NAME}
ENV PORT 8080

# expose port 8080
EXPOSE ${PORT}

# add the files in the current directory into the container at the path go/src/ekko
# change the working directory to the appropriate directory
ADD . ${APP_PATH}
WORKDIR ${APP_PATH}/src/

# download dependencies for the application to run
RUN go mod download

# build the Go application 
RUN go build -o ${APP_NAME}


# run the go application
CMD ./${APP_NAME}