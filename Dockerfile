FROM alpine:latest
RUN apk update
RUN apk add lldpd ethtool iptables python3 scapy libpcap libpcap-dev \
    libpcap-doc curl mandoc man-pages tcpdump vim emacs nano bridge \
    bridge-utils bridge-utils-doc gcc musl-dev openssl openssl-dev \
    openssl-doc linux-headers go dhcp
CMD /bin/ash