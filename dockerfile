FROM openjdk:17.0.2
COPY ./db/ /home/app/db
COPY ./target/${f} /home/app
WORKDIR /home/app
CMD ["java", "-Dfile.encoding=utf-8", "-jar", "${f}", "-k", "grapeSoft@", "-p", "805"]