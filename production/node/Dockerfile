FROM mhart/alpine-node

COPY index.js .

EXPOSE 3000

RUN echo '#!/bin/bash\n npm install express && node index.js' > /usr/bin/start

RUN chmod +x /usr/bin/start





