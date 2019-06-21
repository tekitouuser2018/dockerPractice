FROM node:10.16.0-alpine

# EXPOSE 4200

WORKDIR $HOME/templates/my-app
COPY ./my-app/package*.json $HOME/templates/my-app/
# ADD . ./my-app

# CMD ["ng", "serve", "--host", "0.0.0.0"]

RUN npm install -g typescript@3.4.5
RUN npm i -g @angular/core@7.2.15
RUN npm i typescript@3.4.5
RUN npm install -g @angular/cli 
RUN npm install --save @angular/core
RUN npm install --save @angular/platform-browser
RUN npm install --save @angular/material @angular/cdk
RUN npm install --save @angular/animations
RUN npm i --save @angular/http @angular/common
RUN npm i --save @angular/forms
# RUN npm install -g @angular/cli && ng config -g cli.packageManager yarn
# RUN npm install -g angular-cli
COPY ./my-app/ $HOME/templates/my-app/
ARG configuration=production

# RUN ng build
RUN npm run build -- --output-path=./dist/out --configuration $configuration
RUN ng serve --host 0.0.0.0
# RUN ng serve --host=0.0.0.0
# RUN npm run build --prod

FROM nginx:latest
COPY --from=build-stage /my-app/dist/out/ /usr/share/nginx/html
# Copy the default nginx.conf provided by tiangolo/node-frontend
COPY --from=build-stage ../myproxy/conf/proxy.conf /etc/nginx/conf.d/default.conf
