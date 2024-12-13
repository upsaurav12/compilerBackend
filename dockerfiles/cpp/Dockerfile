FROM alpine:latest

# Install GCC, G++, bash, and other necessary tools
RUN apk add --no-cache bash gcc g++ make

# Set the working directory
WORKDIR /code

# Copy your code into the container
COPY . /code

# Compile and run the C++ program
CMD ["bash", "-c", "g++ -o main main.cpp && ./main"]
