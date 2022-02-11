FROM openjdk:17.0.2
COPY ./target/#{f} /home/app/
WORKDIR /home/app
ENV GSC_HOST "127.0.0.1:805"
RUN echo 'GSC_HOST=' ${GSC_HOST}  # 打印一下默认值
CMD ["java", "-Dfile.encoding=utf-8", "-jar", "#{f}", "-n", "test", "-h", "${GSC_HOST}"]