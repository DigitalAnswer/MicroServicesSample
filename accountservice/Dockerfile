# Base Image
FROM iron/base

# Define port number we want to expose on the internal Docker network
EXPOSE 8080
# Add file to the root ( / ) of the container filesystem
ADD accountservice-linux-amd64 /
ADD healthchecker-linux-amd64 /

HEALTHCHECK --interval=1s --timeout=3s CMD ["./healthchecker-linux-amd64", "-port=8080"] || exit 1
# Executable to run when docker starts a container of this image
ENTRYPOINT ["./accountservice-linux-amd64"]