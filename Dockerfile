FROM golang:1.23.1

WORKDIR /app/

# Copy entrypoint.sh to /usr/local/bin/ and make it executable
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

# Set ENTRYPOINT to the correct path
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["tail", "-f", "/dev/null"]
