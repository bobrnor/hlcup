FROM ubuntu:14.04
MAINTAINER danil@nulana.com

ENV DEBIAN_FRONTEND noninteractive
ENV INITRD No

RUN apt-get update && apt-get install -y mysql-server supervisor unzip \
	&& rm -rf /var/lib/apt/lists/*

COPY build/hlcup_linux_amd64 /
COPY stuff/hlcup.sql /
COPY stuff/app-entrypoint.sh /
COPY stuff/supervisord.conf /

RUN echo 'skip-host-cache\nskip-name-resolve\ninnodb_flush_log_at_trx_commit=0' | awk '{ print } $1 == "[mysqld]" && c == 0 { c = 1; system("cat") }' /etc/mysql/my.cnf > /tmp/my.cnf \
	&& mv /tmp/my.cnf /etc/mysql/my.cnf

EXPOSE 80

ENTRYPOINT ["/app-entrypoint.sh"]
CMD ["/usr/bin/supervisord", "-c", "/supervisord.conf"]
